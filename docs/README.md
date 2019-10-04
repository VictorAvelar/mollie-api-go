# mollie
--
    import "github.com/VictorAvelar/molli-api-go/mollie"

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
	BaseURL            string = "https://api.mollie.com/"
	AuthHeader         string = "Authorization"
	TokenType          string = "Bearer"
	APITokenEnv        string = "MOLLIE_API_TOKEN"
	OrgTokenEnv        string = "MOLLIE_ORG_TOKEN"
	RequestContentType string = "application/json"
)
```
Mollie constants holding values to initialize the client and create requests.

#### func  CheckResponse

```go
func CheckResponse(r *http.Response) error
```
CheckResponse checks the API response for errors, and returns them if present. A
response is considered an error if it has a status code outside the 200 range.
API error responses are expected to have either no response body, or a JSON
response body.

#### type Address

```go
type Address struct {
	StreetAndNumber  string `json:"streetAndNumber,omitempty"`
	StreetAdditional string `json:"streetAdditional,omitempty"`
	PostalCode       string `json:"postalCode,omitempty"`
	City             string `json:"city,omitempty"`
	Region           string `json:"region,omitempty"`
	Country          string `json:"country,omitempty"`
}
```

Address provides a human friendly representation of a geographical space.

When providing an address object as parameter to a request, the following
conditions must be met:

-If any of the fields is provided, all fields have to be provided with exception
of the region field. -If only the region field is given, one should provide all
the other fields as per the previous condition. -For certain PayPal payments the
region field is required.

#### type Amount

```go
type Amount struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}
```

Amount represents a currency and value pair.

#### type ApplicationFee

```go
type ApplicationFee struct {
	Amount      Amount `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}
```

ApplicationFee allows you to split a payment between a platform and connected
merchant accounts.

#### type Chargeback
```go
type Chargeback struct {
	Resource         string           `json:"resource,omitempty"`
	ID               string           `json:"id,omitempty"`
	Amount           *Amount          `json:"amount,omitempty"`
	SettlementAmount *Amount          `json:"settlementAmount,omitempty"`
	CreatedAt        *time.Time       `json:"createdAt,omitempty"`
	ReversedAt       *time.Time       `json:"reversedAt,omitempty"`
	ChargebackID     string           `json:"chargebackId,omitempty"`
	Links            *ChargebackLinks `json:"_links,omitempty"`
}
```

Chargeback describes a forced transaction reversal initiated by the cardholder's bank

#### type ChargebackLinks

```go
type ChargebackLinks struct {
	Self          URL `json:"self,omitempty"`
	Chargeback    URL `json:"chargeback,omitempty"`
	Settlement    URL `json:"settlement,omitempty"`
	Documentation URL `json:"documentation,omitempty"`
}
```

ChargebackLinks describes all the possible links to be returned with
a chargeback object.

#### type ChargebackList

```go
type ChargebackList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Chargebacks []Chargeback
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

ChargebackList describes how a list of chargebacks will be retrieved by Mollie.

#### type ChargebackOptions

```go
type ChargebackOptions struct {
	Include string `url:"include,omitempty"`
	Embed   string `url:"embed,omitempty"`
}
```

ChargebackOptions describes chargeback endpoint valid query string parameters.

See: https://docs.mollie.com/reference/v2/chargebacks-api/get-chargeback

#### type ChargebacksService

```go
type ChargebacksService service
```

ChargebacksService instance operates over chargeback resources

#### func (\*ChargebacksService) Get

```go
func (cs *ChargebacksService) Get(paymentID, chargebackID string, options *ChargebackOptions) (p Chargeback, err error)
```

Get retrieves a single chargeback by its ID. Note the original payment’s ID is needed as well.

If you do not know the original payment’s ID, you can use the List function

#### func (\*ChargebacksService) List

```go
func (cs *ChargebacksService) List(options *ListChargebackOptions) (pl ChargebackList, err error)
```

List retrieves a list of chargebacks associated with your account/organization.

See: https://docs.mollie.com/reference/v2/chargebacks-api/list-chargebacks

#### func (\*ChargebacksService) ListForPayment

```go
func (cs *ChargebacksService) ListForPayment(paymentID string, options *ListChargebackOptions) (pl ChargebackList, err error)
```

ListForPayment retrieves a list of chargebacks associated with a single payment.

See: https://docs.mollie.com/reference/v2/chargebacks-api/list-chargebacks

#### type Client

```go
type Client struct {
	BaseURL *url.URL

  // Services
	Payments    *PaymentsService
	Chargebacks *ChargebacksService
	Methods     *MethodsService
}
```

Client manages communication with Mollie's API.

#### func NewClient

```go
func NewClient(baseClient *http.Client, c *Config) (mollie *Client, err error)
```
NewClient returns a new Mollie HTTP API client. You can pass a previously build
http client, if none is provided then http.DefaultClient will be used.

NewClient will lookup the environment for values to assign to the API token
(`MOLLIE_API_TOKEN`) and the Organization token (`MOLLIE_ORG_TOKEN`) according
to the provided Config object.

You can also set the token values programmatically by using the Client
WithAPIKey and WithOrganizationKey functions.

#### func (\*Client) Do

```go
func (c *Client) Do(req *http.Request) (*Response, error)
```
Do sends an API request and returns the API response or returned as an error if
an API error has occurred.

#### func (\*Client) NewAPIRequest

```go
func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error)
```
NewAPIRequest is a wrapper around the http.NewRequest function.

It will setup the authentication headers/parameters according to the client
config.

#### func (\*Client) WithAuthenticationValue

```go
func (c *Client) WithAuthenticationValue(k string) error
```
WithAuthenticationValue offers a convenient setter for any of the valid
authentication tokens provided by Mollie.

Ideally your API key will be provided from and environment variable or a secret
management engine. This should only be used when environment variables are
"impossible" to be used.

#### type Config

```go
type Config struct {
}
```

Config contains information that helps during the setup of a new Mollie client.

#### func  NewConfig

```go
func NewConfig(t bool, auth string) *Config
```
NewConfig build a Mollie configuration object, it takes t to indicate if our
client is meant to create requests for testing and auth to indicate the
authentication method we want to use.

#### type Error

```go
type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Response *http.Response `json:"response"` // the full response that produced the error
}
```

Error reports details on a failed API request. The success or failure of each
HTTP request is shown in the status field of the HTTP response header, which
contains standard HTTP status codes: - a 2xx code for success - a 4xx or 5xx
code for failure

#### func (\*Error) Error

```go
func (e *Error) Error() string
```
Error functions implement the Error interface on the zuora.Error struct.

#### type FailureReason

```go
type FailureReason string
```

FailureReason provides a description on a failed payment

```go
const (
	ReasonInvalidCardNumber     FailureReason = "invalid_card_number"
	ReasonInvalidCCV            FailureReason = "invalid_ccv"
	ReasonInvalidCardHolderName FailureReason = "invalid_card_holder_name"
	ReasonCardExpired           FailureReason = "card_expired"
	ReasonInvalidCardType       FailureReason = "invalid_card_type"
	ReasonRefusedByIssuer       FailureReason = "refused_by_issuer"
	ReasonInsufficientFunds     FailureReason = "insufficient_funds"
	ReasonInactiveCard          FailureReason = "inactive_card"
	ReasonUnknown               FailureReason = "unknown_reason"
	ReasonPossibleFraud         FailureReason = "possible_fraud"
)
```
Possible failure reasons

#### type FeeRegion

```go
type FeeRegion string
```

FeeRegion contains the fee region for the payment.

```go
const (
	AmericanExpress FeeRegion = "american-express"
	CarteBancaire   FeeRegion = "carte-bancaire"
	IntraEU         FeeRegion = "intra-eu"
	MaestroRegion   FeeRegion = "maestro"
	Other           FeeRegion = "other"
)
```
Valid Fee regions

#### type Image

```go
type Image struct {
	Size1x string `json:"size1X,omitempty"`
	Size2X string `json:"size2X,omitempty"`
	Svg    string `json:"svg,omitempty"`
}
```

Image describes a generic image resource retrieved by Mollie.

#### type ListChargebackOptions

```go
type ListChargebackOptions struct {
	Include   string `url:"include,omitempty"`
	Embed     string `url:"embed,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}
```

ListChargebackOptions describes list chargebacks endpoint valid query string parameters.

#### type ListMethods

```go
type ListMethods struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Methods []PaymentMethodInfo
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

ListMethods describes a list of paginated payment methods.

#### type ListPaymentOptions

```go
type ListPaymentOptions struct {
	Include   string `url:"include,omitempty"`
	Embed     string `url:"embed,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}
```

ListPaymentOptions describes list payments endpoint valid query string
parameters.

#### type Locale

```go
type Locale string
```

Locale represents a country and language in ISO-15897 format.

```go
const (
	English       Locale = "en_US"
	Dutch         Locale = "nl_NL"
	DutchBelgium  Locale = "nl_BE"
	French        Locale = "fr_FR"
	FrenchBelgium Locale = "fr_BE"
	German        Locale = "de_DE"
	GermanAustria Locale = "de_AT"
	GermanSwiss   Locale = "de_CH"
	Spanish       Locale = "es_ES"
	Catalan       Locale = "ca_ES"
	Portuguese    Locale = "pt_PT"
	Italian       Locale = "it_IT"
	Norwegian     Locale = "nb_NO"
	Swedish       Locale = "sv_SE"
	Finish        Locale = "fi_FI"
	Danish        Locale = "da_DK"
	Icelandic     Locale = "is_IS"
	Hungarian     Locale = "hu_HU"
	Polish        Locale = "pl_PL"
	Latvian       Locale = "lv_LV"
	Lithuanian    Locale = "lt_LT"
)
```
Mollie supported locales

#### type MethodsLinks

```go
type MethodsLinks struct {
	Self URL `json:"self,omitempty"`
	Docs URL `json:"documentation,omitempty"`
}
```

MethodsLinks describes links attached to methods service responses.

#### type MethodsOptions

```go
type MethodsOptions struct {
	Locale    Locale `url:"locale,omitempty"`
	Currency  string `url:"currency,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
	Include   string `url:"include,omitempty"`
	// Use for List method only
	SequenceType   SequenceType `url:"sequenceType,omitempty"`
	Amount         Amount       `url:"amount,omitempty"`
	Resource       string       `url:"resource,omitempty"`
	BillingCountry string       `url:"billingCountry,omitempty"`
	IncludeWallets string       `json:"includeWallets,omitempty"`
}
```

MethodsOptions are applicable query string parameters to methods service
endpoints.

#### type MethodsService

```go
type MethodsService service
```

MethodsService operates on methods endpoints

#### func (\*MethodsService) All

```go
func (ms *MethodsService) All(options *MethodsOptions) (pm *ListMethods, err error)
```
All retrieves all the payment methods enabled for your account/organization

See: https://docs.mollie.com/reference/v2/methods-api/list-all-methods

#### func (\*MethodsService) Get

```go
func (ms *MethodsService) Get(id string, options *MethodsOptions) (pmi *PaymentMethodInfo, err error)
```
Get returns information about the payment method specified by id, it also
receives a pointer to the method options containing applicable query string
parameters

See: https://docs.mollie.com/reference/v2/methods-api/get-method

#### func (\*MethodsService) List

```go
func (ms *MethodsService) List(options *MethodsOptions) (pm *ListMethods, err error)
```
List retrieves all enabled payment methods. The results are not paginated.

See: https://docs.mollie.com/reference/v2/methods-api/list-methods

#### type Mode

```go
type Mode string
```

Mode contains information about the payment creation environment.

```go
const (
	LiveMode Mode = "live"
	TestMode Mode = "test"
)
```
Valid modes

#### type PaginationLinks

```go
type PaginationLinks struct {
	Self     *URL `json:"self,omitempty"`
	Previous *URL `json:"previous,omitempty"`
	Next     *URL `json:"next,omitempty"`
	Docs     *URL `json:"documentation,omitempty"`
}
```

PaginationLinks describes the hal component of paginated responses.

#### type Payment

```go
type Payment struct {
	Resource         string          `json:"resource,omitempty"`
	ID               string          `json:"id,omitempty"`
	Mode             Mode            `json:"mode,omitempty"`
	CreatedAt        *time.Time      `json:"createdAt,omitempty"`
	Status           string          `json:"status,omitempty"`
	IsCancellable    bool            `json:"isCancellable,omitempty"`
	AuthorizedAt     *time.Time      `json:"authorizedAt,omitempty"`
	PaidAt           *time.Time      `json:"paidAt,omitempty"`
	CanceledAt       *time.Time      `json:"canceledAt,omitempty"`
	ExpiresAt        *time.Time      `json:"expiresAt,omitempty"`
	ExpiredAt        *time.Time      `json:"expiredAt,omitempty"`
	FailedAt         *time.Time      `json:"failedAt,omitempty"`
	Amount           *Amount         `json:"amount,omitempty"`
	AmountRefunded   *Amount         `json:"amountRefunded,omitempty"`
	AmountRemaining  *Amount         `json:"amountRemaining,omitempty"`
	AmountCaptured   *Amount         `json:"amountCaptured,omitempty"`
	Description      string          `json:"description,omitempty"`
	RedirectURL      string          `json:"redirectUrl,omitempty"`
	WebhookURL       string          `json:"webhookUrl,omitempty"`
	Method           *PaymentMethod  `json:"method,omitempty"`
	Metadata         interface{}     `json:"metadata,omitempty"`
	Locale           *Locale         `json:"locale,omitempty"`
	CountryCode      string          `json:"countryCode,omitempty"`
	ProfileID        string          `json:"profileId,omitempty"`
	SettlementAmount *Amount         `json:"settlementAmount,omitempty"`
	SettlementID     string          `json:"settlementId,omitempty"`
	CustomerID       string          `json:"customerId,omitempty"`
	SequenceType     *SequenceType   `json:"sequenceType,omitempty"`
	MandateID        string          `json:"mandateId,omitempty"`
	OrderID          string          `json:"orderId,omitempty"`
	ApplicationFee   *ApplicationFee `json:"applicationFee,omitempty"`
	Links            *PaymentLinks   `json:"_links,omitempty"`
	Details          *PaymentDetails `json:"details,omitempty"`
}
```

Payment describes a transaction between a customer and a merchant

#### type PaymentDetails

```go
type PaymentDetails struct {
	BankAccount        string         `json:"bankAccount,omitempty"`
	BankBIC            string         `json:"bankBic,omitempty"`
	BankName           string         `json:"bankName,omitempty"`
	BankReason         string         `json:"bankReason,omitempty"`
	BatchReference     string         `json:"batchReference,omitempty"`
	BillingEmail       string         `json:"billingEmail,omitempty"`
	CardAudience       string         `json:"cardAudience,omitempty"`
	CardCountryCode    string         `json:"cardCountryCode,omitempty"`
	CardFingerPrint    string         `json:"cardFingerPrint,omitempty"`
	CardHolder         string         `json:"cardHolder,omitempty"`
	CardLabel          string         `json:"cardLabel,omitempty"`
	CardNumber         string         `json:"cardNumber,omitempty"`
	CardSecurity       string         `json:"cardSecurity,omitempty"`
	ConsumerAccount    string         `json:"consumerAccount,omitempty"`
	ConsumerBIC        string         `json:"consumerBic,omitempty"`
	ConsumerName       string         `json:"consumerName,omitempty"`
	ConsumerReference  string         `json:"consumerReference,omitempty"`
	CreditorIdentifier string         `json:"creditorIdentifier,omitempty"`
	DueDate            *ShortDate     `json:"dueDate,omitempty"`
	EndToEndIdentifier string         `json:"endToEndIdentifier,omitempty"`
	FailureReason      *FailureReason `json:"failureReason,omitempty"`
	FeeRegion          *FeeRegion     `json:"feeRegion,omitempty"`
	FileReference      string         `json:"fileReference,omitempty"`
	GiftCards          []UsedGiftCard `json:"giftCards,omitempty"`
	MandateReference   string         `json:"mandateReference,omitempty"`
	PayPalReference    string         `json:"payPalReference,omitempty"`
	QRCode             *QRCode        `json:"qrCode,omitempty"`
	RemainderAmount    *Amount        `json:"remainderAmount,omitempty"`
	RemainderMethod    *PaymentMethod `json:"remainderMethod,omitempty"`
	SignatureDate      *ShortDate     `json:"signatureDate,omitempty"`
	TransferReference  string         `json:"transferReference,omitempty"`
	VoucherNumber      string         `json:"voucherNumber,omitempty"`
	Wallet             string         `json:"wallet,omitempty"`
	Links              struct {
		Status    URL `json:"status,omitempty"`
		PayOnline URL `json:"payOnline,omitempty"`
	} `json:"_links,omitempty"`
}
```

PaymentDetails contains details for the specified payment method

#### type PaymentLinks

```go
type PaymentLinks struct {
	Self               URL `json:"self,omitempty"`
	Checkout           URL `json:"checkout,omitempty"`
	ChangePaymentState URL `json:"changePaymentState,omitempty"`
	Refunds            URL `json:"refunds,omitempty"`
	ChargeBacks        URL `json:"chargebacks,omitempty"`
	Captures           URL `json:"captures,omitempty"`
	Settlement         URL `json:"settlement,omitempty"`
	Documentation      URL `json:"documentation,omitempty"`
	Mandate            URL `json:"mandate,omitempty"`
	Subscription       URL `json:"subscription,omitempty"`
	Customer           URL `json:"customer,omitempty"`
	Order              URL `json:"order,omitempty"`
}
```

PaymentLinks describes all the possible links to be returned with a payment
object.

#### type PaymentList

```go
type PaymentList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Payments []Payment
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

PaymentList describes how a list of payments will be retrieved by Mollie.

#### type PaymentMethod

```go
type PaymentMethod string
```

PaymentMethod is a payment method supported by Mollie.

```go
const (
	Bancontact     PaymentMethod = "bancontact"
	BankTransfer   PaymentMethod = "banktransfer"
	Belfius        PaymentMethod = "belfius"
	CreditCard     PaymentMethod = "creditcard"
	DirectDebit    PaymentMethod = "directdebit"
	EPS            PaymentMethod = "eps"
	GiftCard       PaymentMethod = "giftcard"
	GiroPay        PaymentMethod = "giropay"
	IDeal          PaymentMethod = "ideal"
	INGHomePay     PaymentMethod = "inghomepay"
	KBC            PaymentMethod = "kbc"
	KlarnaPayLater PaymentMethod = "klarnapaylater"
	KlarnaLiceit   PaymentMethod = "klarnaliceit"
	MyBank         PaymentMethod = "mybank"
	PayPal         PaymentMethod = "paypal"
	PaySafeCard    PaymentMethod = "paysafecard"
	PRZelewy24     PaymentMethod = "przelewy24"
	Sofort         PaymentMethod = "sofort"
)
```
Supported payment methods

#### type PaymentMethodInfo

```go
type PaymentMethodInfo struct {
	Resource      string                 `json:"resource,omitempty"`
	ID            string                 `json:"id,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MinimumAmount *Amount                `json:"minimumAmount,omitempty"`
	MaximumAmount *Amount                `json:"maximumAmount,omitempty"`
	Image         Image                  `json:"image,omitempty"`
	Pricing       []PaymentMethodPricing `json:"pricing,omitempty"`
	Links         MethodsLinks           `json:"_links,omitempty"`
}
```

PaymentMethodInfo describes a single method with details.

#### type PaymentMethodPricing

```go
type PaymentMethodPricing struct {
	Description string  `json:"description,omitempty"`
	Fixed       *Amount `json:"fixed,omitempty"`
	Variable    string  `json:"variable,omitempty"`
}
```

PaymentMethodPricing contains information about commissions and fees applicable
to a payment method.

#### type PaymentOptions

```go
type PaymentOptions struct {
	Include string `url:"include,omitempty"`
	Embed   string `url:"embed,omitempty"`
}
```

PaymentOptions describes payments endpoint valid query string parameters.

See: https://docs.mollie.com/reference/v2/payments-api/get-payment

#### type PaymentsService

```go
type PaymentsService service
```

PaymentsService instance operates over payment resources

#### func (\*PaymentsService) Cancel

```go
func (ps *PaymentsService) Cancel(id string) (p Payment, err error)
```
Cancel removes a payment (if possible) from your Mollie account.

See: https://docs.mollie.com/reference/v2/payments-api/cancel-payment

#### func (\*PaymentsService) Create

```go
func (ps *PaymentsService) Create(p Payment) (np Payment, err error)
```
Create stores a new payment object attached to your Mollie account.

See: https://docs.mollie.com/reference/v2/payments-api/create-payment#

#### func (\*PaymentsService) Get

```go
func (ps *PaymentsService) Get(id string, options *PaymentOptions) (p Payment, err error)
```
Get retrieves a single payment object by its payment token.

#### func (\*PaymentsService) List

```go
func (ps *PaymentsService) List(options *ListPaymentOptions) (pl PaymentList, err error)
```
List retrieves a list of payments associated with your account/organization.

See: https://docs.mollie.com/reference/v2/payments-api/list-payments

#### func (\*PaymentsService) Update

```go
func (ps *PaymentsService) Update(id string, up Payment) (p Payment, err error)
```
Update can be used to update some details of a created payment.

See: https://docs.mollie.com/reference/v2/payments-api/update-payment#

#### type PhoneNumber

```go
type PhoneNumber string
```

PhoneNumber represents a phone number in the E.164 format.

#### type QRCode

```go
type QRCode struct {
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	Src    string `json:"src,omitempty"`
}
```

QRCode object represents an image of a QR code.

#### type Response

```go
type Response struct {
	*http.Response
}
```

Response is a Mollie API response. This wraps the standard http.Response
returned from Mollie and provides convenient access to things like pagination
links.

#### type SequenceType

```go
type SequenceType string
```

SequenceType indicates which type of payment this is in a recurring sequence.

```go
const (
	OneOffSequence    SequenceType = "oneoff"
	FirstSequence     SequenceType = "first"
	RecurringSequence SequenceType = "recurring"
)
```
Valid sequence types

#### type ShortDate

```go
type ShortDate struct {
	time.Time
}
```

ShortDate is a string representing a date in YYYY-MM-DD format.

#### func (\*ShortDate) UnmarshalJSON

```go
func (d *ShortDate) UnmarshalJSON(b []byte) error
```
UnmarshalJSON overrides the default unmarshal action for the Date struct, as we
need links to be pointers to the time.Time struct.

#### type URL

```go
type URL struct {
	Href string `json:"href,omitempty"`
	Type string `json:"type,omitempty"`
}
```

URL in Mollie are commonly represented as objects with an href and type field.

#### type UsedGiftCard

```go
type UsedGiftCard struct {
	Issuer        string  `json:"issuer,omitempty"`
	Amount        *Amount `json:"amount,omitempty"`
	VoucherNumber string  `json:"voucherNumber,omitempty"`
}
```

UsedGiftCard describes a used gift card.
