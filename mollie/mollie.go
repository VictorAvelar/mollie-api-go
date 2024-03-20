package mollie

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/VictorAvelar/mollie-api-go/v4/pkg/idempotency"
	"github.com/google/go-querystring/query"
)

// Constants holding values for client initialization and request instantiation.
const (
	BaseURL              string = "https://api.mollie.com/"
	AuthHeader           string = "Authorization"
	TokenType            string = "Bearer"
	APITokenEnv          string = "MOLLIE_API_TOKEN"
	OrgTokenEnv          string = "MOLLIE_ORG_TOKEN"
	RequestContentType   string = "application/json"
	IdempotencyKeyHeader string = "Idempotency-Key"
	ClientName           string = "MollieGoClient"
	Version              string = "3.14.0"
)

var (
	accessTokenExpr = regexp.MustCompile(`(?m)^access_`)
	errEmptyAuthKey = errors.New("you must provide a non-empty authentication key")
	errBadBaseURL   = errors.New("malformed base url, it must contain a trailing slash")

	goData = strings.Join([]string{runtime.GOOS, runtime.GOARCH, runtime.Version()}, "/")
)

// Client manages communication with Mollie's API.
type Client struct {
	BaseURL        *url.URL
	authentication string
	userAgent      string
	client         *http.Client
	common         service // Reuse a single struct instead of allocating one for each service on the heap.
	config         *Config
	// Tools
	idempotencyKeyProvider idempotency.KeyGenerator
	// Services
	Payments       *PaymentsService
	Chargebacks    *ChargebacksService
	PaymentMethods *PaymentMethodsService
	Invoices       *InvoicesService
	Organizations  *OrganizationsService
	Profiles       *ProfilesService
	Refunds        *RefundsService
	Shipments      *ShipmentsService
	Orders         *OrdersService
	Settlements    *SettlementsService
	Captures       *CapturesService
	Subscriptions  *SubscriptionsService
	Customers      *CustomersService
	Wallets        *WalletsService
	Mandates       *MandatesService
	Permissions    *PermissionsService
	Onboarding     *OnboardingService
	PaymentLinks   *PaymentLinksService
	Clients        *ClientsService
	Balances       *BalancesService
	ClientLinks    *ClientLinksService
	Terminals      *TerminalsService
}

type service struct {
	client *Client
}

func (c *Client) get(ctx context.Context, uri string, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewAPIRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) post(ctx context.Context, uri string, body interface{}, options interface{}) (
	res *Response,
	err error,
) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewAPIRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) patch(ctx context.Context, uri string, body interface{}) (
	res *Response,
	err error,
) {
	req, err := c.NewAPIRequest(ctx, http.MethodPatch, uri, body)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) delete(ctx context.Context, uri string) (res *Response, err error) {
	req, err := c.NewAPIRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return
	}

	return c.Do(req)
}

// WithAuthenticationValue offers a convenient setter for any of the valid authentication
// tokens provided by Mollie.
//
// Ideally your API key will be provided from and environment variable or
// a secret management engine.
// This should only be used when environment variables are "impossible" to be used.
func (c *Client) WithAuthenticationValue(k string) error {
	if k == "" {
		return errEmptyAuthKey
	}

	c.authentication = strings.TrimSpace(k)

	return nil
}

// HasAccessToken will return true when the provided authentication token
// complies with the access token REGEXP match check.
// This will enable TestMode inside the request body.
//
// See: https://github.com/VictorAvelar/mollie-api-go/issues/123
func (c *Client) HasAccessToken() bool {
	return accessTokenExpr.Match([]byte(c.authentication))
}

// SetIdempotencyKeyGenerator allows you to pass your own idempotency
// key generator.
func (c *Client) SetIdempotencyKeyGenerator(kg idempotency.KeyGenerator) {
	c.idempotencyKeyProvider = kg
}

// NewAPIRequest is a wrapper around the http.NewRequest function.
//
// It will setup the authentication headers/parameters according to the client config.
func (c *Client) NewAPIRequest(ctx context.Context, method string, uri string, body interface{}) (
	req *http.Request,
	err error,
) {
	//nolint: contextcheck
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	url, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("url_parsing_error: %w", err)
	}

	if c.config.testing && c.HasAccessToken() {
		qp := url.Query()
		qp.Add("testmode", "true")
		url.RawQuery = qp.Encode()
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)

		err := enc.Encode(body)
		if err != nil {
			return nil, fmt.Errorf("encoding_error: %w", err)
		}
	}

	if ctx == nil {
		ctx = context.Background()
	}

	req, err = http.NewRequestWithContext(ctx, method, url.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("new_request: %w", err)
	}

	c.addRequestHeaders(req)

	return req, nil
}

func (c *Client) addRequestHeaders(req *http.Request) {
	req.Header.Add(AuthHeader, strings.Join([]string{TokenType, c.authentication}, " "))
	req.Header.Set("Content-Type", RequestContentType)
	req.Header.Set("Accept", RequestContentType)
	req.Header.Set("User-Agent", c.userAgent)

	if c.config.reqIdempotency &&
		c.idempotencyKeyProvider != nil &&
		req.Method == http.MethodPost {
		req.Header.Set(IdempotencyKeyHeader, c.idempotencyKeyProvider.Generate())
	}
}

// Do sends an API request and returns the API response or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http_error: %w", err)
	}
	defer resp.Body.Close()

	response, err := newResponse(resp)
	if err != nil {
		return response, err
	}

	err = CheckResponse(response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewClient returns a new Mollie HTTP API client.
// You can pass a previously build http client, if none is provided then
// http.DefaultClient will be used.
//
// NewClient will lookup the environment for values to assign to the
// API token (`MOLLIE_API_TOKEN`) and the Organization token (`MOLLIE_ORG_TOKEN`)
// according to the provided Config object.
//
// You can also set the token values programmatically by using the Client
// WithAPIKey and WithOrganizationKey functions.
func NewClient(baseClient *http.Client, conf *Config) (mollie *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	uri, _ := url.Parse(BaseURL)

	mollie = &Client{
		BaseURL:                uri,
		client:                 baseClient,
		config:                 conf,
		idempotencyKeyProvider: nil,
	}

	mollie.common.client = mollie

	if mollie.config.reqIdempotency {
		mollie.common.client.idempotencyKeyProvider = idempotency.NewStdGenerator()
	}

	// services for resources
	mollie.Payments = (*PaymentsService)(&mollie.common)
	mollie.Chargebacks = (*ChargebacksService)(&mollie.common)
	mollie.PaymentMethods = (*PaymentMethodsService)(&mollie.common)
	mollie.Invoices = (*InvoicesService)(&mollie.common)
	mollie.Organizations = (*OrganizationsService)(&mollie.common)
	mollie.Profiles = (*ProfilesService)(&mollie.common)
	mollie.Refunds = (*RefundsService)(&mollie.common)
	mollie.Shipments = (*ShipmentsService)(&mollie.common)
	mollie.Orders = (*OrdersService)(&mollie.common)
	mollie.Captures = (*CapturesService)(&mollie.common)
	mollie.Settlements = (*SettlementsService)(&mollie.common)
	mollie.Subscriptions = (*SubscriptionsService)(&mollie.common)
	mollie.Customers = (*CustomersService)(&mollie.common)
	mollie.Wallets = (*WalletsService)(&mollie.common)
	mollie.Mandates = (*MandatesService)(&mollie.common)
	mollie.Permissions = (*PermissionsService)(&mollie.common)
	mollie.Onboarding = (*OnboardingService)(&mollie.common)
	mollie.PaymentLinks = (*PaymentLinksService)(&mollie.common)
	mollie.Clients = (*ClientsService)(&mollie.common)
	mollie.Balances = (*BalancesService)(&mollie.common)
	mollie.ClientLinks = (*ClientLinksService)(&mollie.common)
	mollie.Terminals = (*TerminalsService)(&mollie.common)

	mollie.userAgent = strings.Join([]string{
		ClientName,
		Version,
		goData,
	}, "/")

	// Parse authorization from specified environment variable
	tkn, ok := os.LookupEnv(mollie.config.auth)
	if ok {
		mollie.authentication = tkn
	}

	return mollie, nil
}

/*
Constructor for Error.
*/
func newError(rsp *Response) error {
	baseErr := &BaseError{}

	if rsp.ContentLength > 0 {
		err := json.Unmarshal(rsp.content, baseErr)
		if err != nil {
			return err
		}
	} else {
		baseErr.Status = rsp.StatusCode
		baseErr.Title = rsp.Status
		baseErr.Detail = string(rsp.content)
	}

	return baseErr
}

// Response is a Mollie API response. This wraps the standard http.Response
// returned from Mollie and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response
	content []byte
}

func newResponse(rsp *http.Response) (*Response, error) {
	res := Response{Response: rsp}

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return &res, err
	}

	res.content = data

	rsp.Body = io.NopCloser(bytes.NewBuffer(data))
	res.Response = rsp

	return &res, nil
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body.
func CheckResponse(r *Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}

	return nil
}
