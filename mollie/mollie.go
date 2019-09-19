package mollie

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Mollie  constants holding values to initialize the client and create requests.
const (
	BaseURLV2          string = "https://api.mollie.com/v2/"
	AuthHeader         string = "authorization"
	TokenType          string = "Bearer"
	APITokenEnv        string = "MOLLIE_API_TOKEN"
	OrgTokenEnv        string = "MOLLIE_ORG_TOKEN"
	RequestContentType string = "application/json"
)

var (
	errEmptyAPIKey = errors.New("you must provide a non-empty API key")
	errBadBaseURL  = errors.New("malformed base url, it must contain a trailing slash")
)

// Client for Mollie API plus information related to the
// different authentication methods provided by the API.
type Client struct {
	client            *http.Client
	BaseURL           *url.URL
	APIKey            string
	OrganizationToken string
}

//WithAPIKey offers a convenient setter with some base validation to attach
//an API key to an Client.
//
//Ideally your API key will be provided from and environment variable or
//a secret management engine.
func (c *Client) WithAPIKey(k string) error {
	if k == "" {
		return errEmptyAPIKey
	}

	c.APIKey = strings.TrimSpace(k)

	return nil
}

//WithOrganizationToken offers a convenient token with some base validation to
//attach a Mollie Organization Token to an Client.
//
//Ideally your API key will be provided from and environment variable or
//a secret management engine.
func (c *Client) WithOrganizationToken(t string) error {
	if t == "" {
		return errEmptyAPIKey
	}
	c.OrganizationToken = strings.TrimSpace(t)
	return nil
}

//NewAPIRequest is a wrapper around the http.NewRequest function.
//
//For setting up the headers it takes a hierarchical approach, this meaning that
//if set the Client.OrganizationToken will be used, if this value is empty then
//it will attempt to use the Client.APIKey, and if this value is also empty it
//will return an error.
func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
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
	if body != nil {
		req.Header.Set("Content-Type", RequestContentType)
	}
	req.Header.Set("Accept", RequestContentType)

	return
}

// Do sends an API request and returns the API response or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//NewClient returns a new Mollie HTTP API client.
// You can pass a previously build http client, if none is provided then
// http.DefaultClient will be used.
//
//By default NewClient will lookup the environment for values to assign to the
//API token (`MOLLIE_API_TOKEN`) and the Organization token (`MOLLIE_ORG_TOKEN`).
//
//You can also set the token values programmatically by using the Client
//WithAPIKey and WithOrganizationKey functions.
func NewClient(baseClient *http.Client) (mollie *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURLV2)

	mollie = &Client{
		client:  baseClient,
		BaseURL: u,
	}

	// Try to parse tokens from environment
	if tkn, ok := os.LookupEnv(APITokenEnv); ok {
		mollie.APIKey = tkn
	}
	if orgTkn, ok := os.LookupEnv(OrgTokenEnv); ok {
		mollie.OrganizationToken = orgTkn
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
