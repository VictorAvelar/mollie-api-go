# mollie
--
    import "github.com/VictorAvelar/mollie-api-go/mollie"

Package mollie is a wrapper around Mollie's REST API.

See: https://www.mollie.com/en/developers

The Mollie API implements a Representational state transfer (REST) architecture.

Sounds technical, but it’s really quite easy. It mainly breaks down to
HTTP-methods GET, PATCH, POST and DELETE matching the operations to read,
update, create and delete.

REST also implies a nice and clean structure for URLs or endpoints. This means
you can reach any part of the Mollie API on https://api.mollie.com/v2/ adding
the name of the resource you want to interact with.

## Usage

```go
const (
	BaseURL     string = "https://api.mollie.com"
	APIVersion  string = "v2"
	AuthHeader  string = "authorization"
	TokenType   string = "Bearer"
	APITokenEnv string = "MOLLIE_API_TOKEN"
	OrgTokenEnv string = "MOLLIE_ORG_TOKEN"
)
```
Mollie constants holding values to initialize the client and create requests.

#### type APIClient

```go
type APIClient struct {
	Ctx               context.Context
	HTTPClient        httpClient
	BaseURL           *url.URL
	APIKey            string
	OrganizationToken string
}
```

APIClient for Mollie API plus information related to the different
authentication methods provided by the API.

#### func  NewClient

```go
func NewClient(ctx context.Context, baseClient httpClient, uri string) (mollie *APIClient, err error)
```
NewClient returns a fully qualified Mollie HTTP API client with context. It
receives a context, a httpClient and a uri to initialize the client but also
accepts nil in some cases.

If nil is passed to context, context.Background will be initialized. if nil is
passed as httpClient then the http.DefaultClient will be initialized. The uri
will be parsed with url.Parse function.

By default NewClient will lookup the environment for values to assign to the API
token (`MOLLIE_API_TOKEN`) and the Organization token (`MOLLIE_ORG_TOKEN`).

You can also set the token values programmatically by using the APIClient
WithAPIKey and WithOrganizationKey functions.

#### func (*APIClient) NewAPIRequest

```go
func (c *APIClient) NewAPIRequest(m string, uri string, body io.Reader, auth bool) (req *http.Request, err error)
```
NewAPIRequest is a wrapper around the http.NewRequest function. It takes the
same parameters plus a flag to indicate if the request needs to have the
authorization headers.

For setting up the headers it takes a hierarchical approach, this meaning that
if set the APIClient.OrganizationToken will be used, if this value is empty then
it will attempt to use the APIClient.APIKey, and if this value is also empty it
will return an error.

#### func (*APIClient) WithAPIKey

```go
func (c *APIClient) WithAPIKey(k string) error
```
WithAPIKey offers a convenient setter with some base validation to attach an API
key to an APIClient.

Ideally your API key will be provided from and environment variable or a secret
management engine.

#### func (*APIClient) WithOrganizationToken

```go
func (c *APIClient) WithOrganizationToken(t string) error
```
WithOrganizationToken offers a convenient token with some base validation to
attach a Mollie Organization Token to an APIClient.

Ideally your API key will be provided from and environment variable or a secret
management engine.

#### type Address

```go
type Address struct {
	StreetAndNumber  string `json:"streetAndNumber"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Region           string `json:"region,omitempty"`
	StreetAdditional string `json:"streetAdditional,omitempty"`
}
```

Address contains information to provide a means of physically locating a place
in the urban geographical space.

For Mollie the country must follow an ISO 3166-1 alpha-2 format, postal codes
must belong to the specified city and country. When providing an address object
as parameter to a request, the following conditions must be met:

    - If any of the fields is provided, all fields have to be provided
      with exception of the region field.
    - If only the region field is given, one should provide all the other
      fields as per the previous condition.
    - For certain PayPal payments the region field is required.

#### type Amount

```go
type Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}
```

Amount contains a currency code compliant to the ISO 4217 specification and a
value with the exact amount in the given currency.

#### type Date

```go
type Date struct {
	time.Time
}
```

Date is represented on `YYYY-MM-DD` format for Mollie requests, in order to
comply with this the time.Time struct is embedded on Date and custom marshal and
unmarshal will be provided.

#### func (*Date) UnmarshalJSON

```go
func (d *Date) UnmarshalJSON(b []byte) error
```
UnmarshalJSON overrides the default unmarshal action for the Date struct, as we
need links to be pointers to the time.Time struct.

#### type HAL

```go
type HAL map[string]Link
```

HAL is a map of Links that contains relevant resources according to the
response.

#### type HalURL

```go
type HalURL struct {
	*url.URL
}
```

HalURL has url.URL embedded to modify the default json marshal behaviour.

#### func (*HalURL) UnmarshalJSON

```go
func (hl *HalURL) UnmarshalJSON(b []byte) error
```
UnmarshalJSON overrides the default unmarshal action for the HalURL struct, as
we need links to be pointers to the url.URL struct.

#### type Link

```go
type Link struct {
	Href        HalURL `json:"href"`
	ContentType string `json:"type"`
}
```

Link for Mollie are commonly represented as objects with an href and type field.

#### type Locale

```go
type Locale string
```

Locale is a string representing the country and language in ISO 15897 format.

#### type PhoneNumber

```go
type PhoneNumber string
```

PhoneNumber is a string in the E.164 format. For example, `+31208202070`.

#### type QRCode

```go
type QRCode struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Source int `json:"src"`
}
```

QRCode represents an image of a QR code.

#### type ResponseErr

```go
type ResponseErr struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Field  string `json:"field,omitempty"`
	Links  HAL    `json:"_links"`
}
```

ResponseErr describes the content of the response body when the http call to
Mollie API is not successful.
