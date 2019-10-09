package mollie

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Mollie  constants holding values to initialize the client and create requests.
const (
	BaseURL            string = "https://api.mollie.com/"
	AuthHeader         string = "Authorization"
	TokenType          string = "Bearer"
	APITokenEnv        string = "MOLLIE_API_TOKEN"
	OrgTokenEnv        string = "MOLLIE_ORG_TOKEN"
	RequestContentType string = "application/json"
)

var (
	errEmptyAPIKey = errors.New("you must provide a non-empty API key")
	errBadBaseURL  = errors.New("malformed base url, it must contain a trailing slash")
)

// Client manages communication with Mollie's API.
type Client struct {
	BaseURL        *url.URL
	authentication string
	client         *http.Client
	common         service // Reuse a single struct instead of allocating one for each service on the heap.
	config         *Config
	// Services
	Payments      *PaymentsService
	Chargebacks   *ChargebacksService
	Methods       *MethodsService
	Invoices      *InvoicesService
	Organizations *OrganizationsService
<<<<<<< HEAD
	Profiles    *ProfilesService
=======
	Refunds       *RefundsService
>>>>>>> Implements the Get method and add Refund Services into Client Struct
}

type service struct {
	client *Client
}

// WithAuthenticationValue offers a convenient setter for any of the valid authentication
// tokens provided by Mollie.
//
// Ideally your API key will be provided from and environment variable or
// a secret management engine.
// This should only be used when environment variables are "impossible" to be used.
func (c *Client) WithAuthenticationValue(k string) error {
	if k == "" {
		return errEmptyAPIKey
	}

	c.authentication = strings.TrimSpace(k)

	return nil
}

// NewAPIRequest is a wrapper around the http.NewRequest function.
//
// It will setup the authentication headers/parameters according to the client config.
func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	if c.config.testing == true {
		u.Query().Add("testmode", "true")
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(AuthHeader, strings.Join([]string{TokenType, c.authentication}, " "))

	if body != nil {
		req.Header.Set("Content-Type", RequestContentType)
	}
	req.Header.Set("Accept", RequestContentType)

	return
}

// Do sends an API request and returns the API response or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return nil, err
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
func NewClient(baseClient *http.Client, c *Config) (mollie *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURL)

	mollie = &Client{
		BaseURL: u,
		client:  baseClient,
		config:  c,
	}

	mollie.common.client = mollie

	// services for resources
	mollie.Payments = (*PaymentsService)(&mollie.common)
	mollie.Chargebacks = (*ChargebacksService)(&mollie.common)
	mollie.Methods = (*MethodsService)(&mollie.common)
	mollie.Invoices = (*InvoicesService)(&mollie.common)
	mollie.Organizations = (*OrganizationsService)(&mollie.common)
	mollie.Profiles = (*ProfilesService)(&mollie.common)

	// Parse authorization from environment
	if tkn, ok := os.LookupEnv(APITokenEnv); ok {
		mollie.authentication = tkn
	}
	return
}

/*
Error reports details on a failed API request.
The success or failure of each HTTP request is shown in the status field of the HTTP response header,
which contains standard HTTP status codes:
- a 2xx code for success
- a 4xx or 5xx code for failure
*/
type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Response *http.Response `json:"response"` // the full response that produced the error
}

// Error functions implement the Error interface on the zuora.Error struct.
func (e *Error) Error() string {
	return fmt.Sprintf("response failed with status %v", e.Message)
}

/*
Constructor for Error
*/
func newError(r *http.Response) *Error {
	var e Error
	e.Response = r
	e.Code = r.StatusCode
	e.Message = r.Status
	return &e
}

// Response is a Mollie API response. This wraps the standard http.Response
// returned from Mollie and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response
	content []byte
}

func newResponse(r *http.Response) *Response {
	var res Response
	if c, err := ioutil.ReadAll(r.Body); err == nil {
		res.content = c
	}
	json.NewDecoder(r.Body).Decode(&res)
	res.Response = r
	return &res
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body.
func CheckResponse(r *http.Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}
	return nil
}
