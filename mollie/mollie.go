package mollie

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Mollie  constanst holding values to initialize the client and create requests.
const (
	BaseURL     string = "https://api.mollie.com"
	APIVersion  string = "v2"
	AuthHeader  string = "authorization"
	TokenType   string = "Bearer"
	APITokenEnv        = "MOLLIE_API_TOKEN"
	OrgTokenEnv        = "MOLLIE_ORG_TOKEN"
)

var (
	errEmptyAPIKey = errors.New("you must provide a non-empty API key")
)

type httpClient interface {
	Do(req *http.Request) (res *http.Response, err error)
}

//APIClient struct contains the httpClient plus information related to the
// different authentication methods provided by the API.
type APIClient struct {
	ClientContext     context.Context
	HTTPClient        httpClient
	BaseURL           *url.URL
	APIKey            string
	OrganizationToken string
}

//WithAPIKey offers a convenient setter with some base validation to attach
//an API key to an APIClient.
//Ideally your API key will be provided from and environment variable or
//a secret management engine.
func (c *APIClient) WithAPIKey(k string) error {
	if k == "" {
		return errEmptyAPIKey
	}

	c.APIKey = strings.Trim(k, " ")

	return nil
}

//WithOrganizationToken offers a convenient token with some base validation to
//attach a Mollie Organization Token to an APIClient.
//Ideally your API key will be provided from and environment variable or
//a secret management engine.
func (c *APIClient) WithOrganizationToken(t string) error {
	if t == "" {
		return errEmptyAPIKey
	}

	c.OrganizationToken = t

	return nil
}

//NewAPIRequest is a wrapper arount the http.NewRequest function.
//It takes the same parameters plus a flag to indicate if the request needs
//to have the authorization headers.
//The http request method defaults to GET if no method is set.
//For setting up the headers it takes a hierarchical approach, this meaning that
//if set the APIClient.OrganizationToken will be used, if this value is empty then
//it will attemp to use the APIClient.APIKey, and if this value is also empty it
//will return an error.
func (c *APIClient) NewAPIRequest(m string, uri string, body io.Reader, auth bool) (req *http.Request, err error) {
	uri = strings.Trim(uri, "/")
	uri = strings.Join([]string{
		APIVersion,
		uri,
	}, "/")
	rel, _ := url.Parse(uri)

	uri = c.BaseURL.ResolveReference(rel).String()
	req, bErr := http.NewRequest(m, uri, body)
	if err != nil {
		return nil, bErr
	}

	req.Header.Add("accept", "application/json")

	if auth {
		var parts = []string{TokenType}
		if c.OrganizationToken != "" {
			parts = append(parts, c.OrganizationToken)
		} else if c.APIKey != "" {
			parts = append(parts, c.APIKey)
		} else {
			return nil, errEmptyAPIKey
		}

		v := strings.Join(parts, " ")
		req.Header.Add(AuthHeader, v)
	}
	return
}

//NewClient returns a fully qualified Mollie HTTP API client with context.
//It receives a context, a httpClient and a uri to initialize the client
//but also accepts nil in some cases.
//If nil is passed to context, context.Background will be initialized.
//if nil is passed as httpClient then the http.DefaultClient will be initialized.
//The uri will be parsed with url.Parse() function.
//By default NewClient will lookup the environment for values to assign to the
//API token (MOLLIE_API_TOKEN) and the Organization token (MOLLIE_ORG_TOKEN).
//You can also set the token values programmatically by using the APIClient
//WithAPIKey and WithOrganizationKey functions.
func NewClient(ctx context.Context, baseClient httpClient, uri string) (mollie *APIClient, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, err := url.Parse(uri)
	if err != nil {
		return
	}

	c := &APIClient{
		ClientContext: ctx,
		HTTPClient:    baseClient,
		BaseURL:       u,
	}

	if tkn, ok := os.LookupEnv(APITokenEnv); ok {
		c.APIKey = tkn
	}

	if orgTkn, ok := os.LookupEnv(OrgTokenEnv); ok {
		c.OrganizationToken = orgTkn
	}

	return c, nil
}
