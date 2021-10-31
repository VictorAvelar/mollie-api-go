# mollie

--
import "github.com/VictorAvelar/mollie-api-golang/v2/mollie"

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

Constants holding values for client initialization and request instantiation.

#### func CheckResponse

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

If any of the fields is provided, all fields have to be provided with exception
of the region field. If only the region field is given, one should provide all
the other fields as per the previous condition. For certain PayPal payments the
region field is required.

#### type Amount

```go
type Amount struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}
```

Amount represents a currency and value pair.

#### type ApplePaymentSession

```go
type ApplePaymentSession struct {
	EpochTimestamp    int    `json:"epochTimestamp,omitempty"`
	ExpiresAt         int    `json:"expiresAt,omitempty"`
	MerchantSessionID string `json:"merchantSessionIdentifier,omitempty"`
	Nonce             string `json:"nonce,omitempty"`
	MerchantID        string `json:"merchantIdentified,omitempty"`
	DomainName        string `json:"domainName,omitempty"`
	DisplayName       string `json:"displayName,omitempty"`
	Signature         string `json:"signature,omitempty"`
}
```

ApplePaymentSession contains information about an Apple pay session

#### type ApplePaymentSessionRequest

```go
type ApplePaymentSessionRequest struct {
	Domain        string `json:"domain,omitempty"`
	ValidationURL string `json:"validationUrl,omitempty"`
}
```

ApplePaymentSessionRequest contains the body parameters for requesting a valid
PaymentSession from Apple.

#### type ApplicationFee

```go
type ApplicationFee struct {
	Amount      *Amount `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
}
```

ApplicationFee allows you to split a payment between a platform and connected
merchant accounts.

#### type Capture

```go
type Capture struct {
	Resource         string       `json:"resource,omitempty"`
	ID               string       `json:"id,omitempty"`
	Mode             Mode         `json:"mode,omitempty"`
	Amount           *Amount      `json:"amount,omitempty"`
	SettlementAmount *Amount      `json:"settlementAmount,omitempty"`
	PaymentID        string       `json:"paymentId,omitempty"`
	ShipmentID       string       `json:"shipmentId,omitempty"`
	SettlementID     string       `json:"settlementId,omitempty"`
	CreatedAt        *time.Time   `json:"createdAt,omitempty"`
	Links            CaptureLinks `json:"links,omitempty"`
}
```

Capture describes a single capture Captures are used for payments that have the
authorize-then-capture flow

#### type CaptureLinks

```go
type CaptureLinks struct {
	Self          *URL `json:"self,omitempty"`
	Payment       *URL `json:"payment,omitempty"`
	Shipment      *URL `json:"shipment,omitempty"`
	Settlement    *URL `json:"settlement,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

CaptureLinks contains relevant links for a capture object

#### type CapturesList

```go
type CapturesList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Captures []*Capture
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

CapturesList describes a list of captures

#### type CapturesService

```go
type CapturesService service
```

CapturesService operates over captures resource

#### func (\*CapturesService) Get

```go
func (cs *CapturesService) Get(pID, cID string) (c *Capture, err error)
```

Get retrieves a single capture by its ID. Note the original payment’s ID is
needed as well.

See: https://docs.mollie.com/reference/v2/captures-api/get-capture

#### func (\*CapturesService) List

```go
func (cs *CapturesService) List(pID string) (cl *CapturesList, err error)
```

List retrieves all captures for a certain payment

See: https://docs.mollie.com/reference/v2/captures-api/list-captures

#### type CardLabel

```go
type CardLabel string
```

CardLabel Note that not all labels can be processed through Mollie.

```go
const (
	AmericaExpress CardLabel = "American Express"
	CartaSi        CardLabel = "Carta Si"
	CarteBleue     CardLabel = "Carte Bleue"
	Dankort        CardLabel = "Dankort"
	DinersClub     CardLabel = "Diners Club"
	Discover       CardLabel = "Discover"
	JCB            CardLabel = "JCB"
	Laser          CardLabel = "Laser"
	Maestro        CardLabel = "Maestro"
	Mastercard     CardLabel = "Mastercard"
	Unionpay       CardLabel = "Unionpay"
	Visa           CardLabel = "Visa"
	Empty          CardLabel = "null"
)
```

Available card labels

#### type CategoryCode

```go
type CategoryCode uint
```

CategoryCode specifies an industry or category

```go
const (
	BookMagazinesAndNewspapers          CategoryCode = 5192
	GeneralMerchandise                  CategoryCode = 5399
	FoodAndDrinks                       CategoryCode = 5499
	AutomotiveProducts                  CategoryCode = 5533
	ChildrenProducts                    CategoryCode = 5641
	ClothingAndShoes                    CategoryCode = 5651
	MarketplaceCrowdfundingAndDonations CategoryCode = 5262
	ElectronicsComputersAndSoftware     CategoryCode = 5732
	HostingOrVpnServices                CategoryCode = 5734
	Entertainment                       CategoryCode = 5735
	CreditsOrVouchersOrGiftCards        CategoryCode = 5815
	Alcohol                             CategoryCode = 5921
	JewelryAndAccessories               CategoryCode = 5944
	HealthAndBeautyProducts             CategoryCode = 5977
	FinancialServices                   CategoryCode = 6012
	Consultancy                         CategoryCode = 7299
	TravelRentalAndTransportation       CategoryCode = 7999
	AdvisingOrCoachingOrTraining        CategoryCode = 8299
	CharityAndDonations                 CategoryCode = 8398
	PoliticalParties                    CategoryCode = 8699
	Others                              CategoryCode = 0
)
```

Available category codes

#### type Chargeback

```go
type Chargeback struct {
	Resource         string          `json:"resource,omitempty"`
	ID               string          `json:"id,omitempty"`
	Amount           *Amount         `json:"amount,omitempty"`
	SettlementAmount *Amount         `json:"settlementAmount,omitempty"`
	CreatedAt        *time.Time      `json:"createdAt,omitempty"`
	ReversedAt       *time.Time      `json:"reversedAt,omitempty"`
	PaymentID        string          `json:"paymentId,omitempty"`
	Links            ChargebackLinks `json:"_links,omitempty"`
}
```

Chargeback describes a forced transaction reversal initiated by the cardholder's
bank

#### type ChargebackLinks

```go
type ChargebackLinks struct {
	Self          *URL `json:"self,omitempty"`
	Payment       *URL `json:"payment,omitempty"`
	Settlement    *URL `json:"settlement,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

ChargebackLinks describes all the possible links to be returned with a
chargeback object.

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

#### type ChargebacksService

```go
type ChargebacksService service
```

ChargebacksService instance operates over chargeback resources

#### func (\*ChargebacksService) Get

```go
func (cs *ChargebacksService) Get(paymentID, chargebackID string, options *ChargebackOptions) (p Chargeback, err error)
```

Get retrieves a single chargeback by its ID. Note the original payment’s ID is
needed as well.

See: https://docs.mollie.com/reference/v2/chargebacks-api/get-chargeback

#### func (\*ChargebacksService) List

```go
func (cs *ChargebacksService) List(options *ListChargebackOptions) (cl *ChargebackList, err error)
```

List retrieves a list of chargebacks associated with your account/organization.

See: https://docs.mollie.com/reference/v2/chargebacks-api/list-chargebacks

#### func (\*ChargebacksService) ListForPayment

```go
func (cs *ChargebacksService) ListForPayment(paymentID string, options *ListChargebackOptions) (cl *ChargebackList, err error)
```

ListForPayment retrieves a list of chargebacks associated with a single payment.

See: https://docs.mollie.com/reference/v2/chargebacks-api/list-chargebacks

#### type Client

```go
type Client struct {
	BaseURL *url.URL

	// Services
	Payments      *PaymentsService
	Chargebacks   *ChargebacksService
	Methods       *MethodsService
	Invoices      *InvoicesService
	Organizations *OrganizationsService
	Profiles      *ProfilesService
	Refunds       *RefundsService
	Shipments     *ShipmentsService
	Orders        *OrdersService
	Settlements   *SettlementsService
	Captures      *CapturesService
	Subscriptions *SubscriptionsService
	Customers     *CustomersService
	Miscellaneous *MiscellaneousService
	Mandates      *MandatesService
	Permissions   *PermissionsService
	Onboarding    *OnboardingService
	PaymentLinks  *PaymentLinksService
	Partners      *PartnerService
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

#### func (\*Client) HasAccessToken

```go
func (c *Client) HasAccessToken() bool
```

HasAccessToken will return true when the provided authentication token complies
with the access token REGEXP match check. This will enable TestMode inside the
request body.

See: https://github.com/VictorAvelar/mollie-api-go/issues/123

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

#### type Commission

```go
type Commission struct {
	Count       int     `json:"count,omitempty"`
	TotalAmount *Amount `json:"totalAmount,omitempty"`
}
```

Commission describes a partner take from any operation on Mollie's API.

#### type Config

```go
type Config struct {
}
```

Config contains information that helps during the setup of a new Mollie client.

#### func NewConfig

```go
func NewConfig(t bool, auth string) *Config
```

NewConfig builds a Mollie configuration object, it takes t to indicate if our
client is meant to create requests for testing and auth to indicate the
authentication method we want to use.

#### type CreateShipmentRequest

```go
type CreateShipmentRequest struct {
	Lines    []OrderLine      `json:"lines,omitempty"`
	Tracking ShipmentTracking `json:"tracking,omitempty"`
	TestMode bool             `json:"testmode,omitempty"`
}
```

CreateShipmentRequest defines information required to create a new shipment

#### type Customer

```go
type Customer struct {
	Resource  string                 `json:"resource,omitempty"`
	ID        string                 `json:"id,omitempty"`
	Mode      Mode                   `json:"mode,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Email     string                 `json:"email,omitempty"`
	Locale    Locale                 `json:"locale,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt *time.Time             `json:"createdAt,omitempty"`
	Links     CustomerLinks          `json:"_links,omitempty"`
}
```

Customer represents buyers

#### type CustomerLinks

```go
type CustomerLinks struct {
	Self          *URL `json:"self,omitempty"`
	Mandates      *URL `json:"mandates,omitempty"`
	Subscriptions *URL `json:"subscriptions,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
}
```

CustomerLinks contains the HAL resources for a customer response

#### type CustomersList

```go
type CustomersList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Customers []Customer `json:"customers,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}
```

CustomersList contains a embedded list of customers wrapped in a standard Mollie
paginated response.

#### type CustomersService

```go
type CustomersService service
```

CustomersService operates over the customer resource

#### func (\*CustomersService) Create

```go
func (cs *CustomersService) Create(c Customer) (cc *Customer, err error)
```

Create creates a simple minimal representation of a customer in the Mollie API
to use for the Mollie Checkout and Recurring features

See: https://docs.mollie.com/reference/v2/customers-api/create-customer

#### func (\*CustomersService) CreatePayment

```go
func (cs *CustomersService) CreatePayment(id string, p Payment) (pp *Payment, err error)
```

CreatePayment creates a payment for the customer.

See: https://docs.mollie.com/reference/v2/customers-api/create-customer-payment

#### func (\*CustomersService) Delete

```go
func (cs *CustomersService) Delete(id string) (err error)
```

Delete a customer. All mandates and subscriptions created for this customer will
be canceled as well.

See: https://docs.mollie.com/reference/v2/customers-api/delete-customer

#### func (\*CustomersService) Get

```go
func (cs *CustomersService) Get(id string) (c *Customer, err error)
```

Get finds a customer by its ID

See: https://docs.mollie.com/reference/v2/customers-api/get-customer

#### func (\*CustomersService) GetPayments

```go
func (cs *CustomersService) GetPayments(id string, options *ListCustomersOptions) (pl *PaymentList, err error)
```

GetPayments retrieves all payments linked to the customer.

See: https://docs.mollie.com/reference/v2/customers-api/list-customer-payments

#### func (\*CustomersService) List

```go
func (cs *CustomersService) List(options *ListCustomersOptions) (cl *CustomersList, err error)
```

List retrieves all customers created.

See: https://docs.mollie.com/reference/v2/customers-api/list-customers

#### func (\*CustomersService) Update

```go
func (cs *CustomersService) Update(id string, c Customer) (cc *Customer, err error)
```

Update mutates an existing customer

See: https://docs.mollie.com/reference/v2/customers-api/update-customer

#### type EligibilityReasons

```go
type EligibilityReasons string
```

EligibilityReasons for paypal seller protection

```go
const (
	Eligible                        EligibilityReasons = "Eligible"
	Ineligible                      EligibilityReasons = "Ineligible"
	INROnly                         EligibilityReasons = "Partially Eligible - INR Only"
	AuthOnly                        EligibilityReasons = "Partially Eligible - Unauth Only"
	PartiallyEligible               EligibilityReasons = "PartiallyEligible"
	EligibilityNone                 EligibilityReasons = "None"
	ActiveFraudControlUnAuthPremium EligibilityReasons = "Active Fraud Control - Unauth Premium Eligible"
)
```

Available eligibility reasons for Paypal sellers protection

#### type EmbedValue

```go
type EmbedValue string
```

EmbedValue describes the valid value of embed query string.

```go
const (
	EmbedPayment     EmbedValue = "payment"
	EmbedRefund      EmbedValue = "refund"
	EmbedShipments   EmbedValue = "shipments"
	EmbedChangebacks EmbedValue = "chanrgebacks"
)
```

Valid Embed query string value.

#### type Error

```go
type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Content  string         `json:"content,omitempty"`
	Response *http.Response `json:"response"` // the full response that produced the error
}
```

Error reports details on a failed API request.

#### func (\*Error) Error

```go
func (e *Error) Error() string
```

Error function complies with the error interface

#### type FailureReason

```go
type FailureReason string
```

FailureReason provides a description on a failed payment

```go
const (
	ReasonAuthenticationAbandoned      FailureReason = "authentication_abandoned"
	ReasonAuthenticationUnavailableACS FailureReason = "authentication_unavailable_acs"
	ReasonInvalidCardNumber            FailureReason = "invalid_card_number"
	ReasonInvalidCCV                   FailureReason = "invalid_ccv"
	ReasonInvalidCardHolderName        FailureReason = "invalid_card_holder_name"
	ReasonCardExpired                  FailureReason = "card_expired"
	ReasonInvalidCardType              FailureReason = "invalid_card_type"
	ReasonRefusedByIssuer              FailureReason = "refused_by_issuer"
	ReasonInsufficientFunds            FailureReason = "insufficient_funds"
	ReasonInactiveCard                 FailureReason = "inactive_card"
	ReasonUnknown                      FailureReason = "unknown_reason"
	ReasonPossibleFraud                FailureReason = "possible_fraud"
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

#### type GetPartnerClientOptions

```go
type GetPartnerClientOptions struct {
	Embed string `url:"embed,omitempty"`
}
```

GetPartnerClientOptions contains valid query parameters for the get clients
endpoint.

#### type GiftCardEnabled

```go
type GiftCardEnabled struct {
	Resource    string               `json:"resource,omitempty"`
	ID          GiftCardIssuer       `json:"id,omitempty"`
	Description string               `json:"description,omitempty"`
	Status      GiftCardIssuerStatus `json:"status,omitempty"`
	Links       GiftCardLinks        `json:"_links,omitempty"`
}
```

GiftCardEnabled describes the response of a gift card issuer enable operation.

#### type GiftCardIssuer

```go
type GiftCardIssuer string
```

GiftCardIssuer type describes issuers supported by mollie.

```go
const (
	BloemenCadeuKaart          GiftCardIssuer = "bloemencadeaukaart"
	BloemPlantGiftCard         GiftCardIssuer = "bloemplantgiftcard"
	Boekenbon                  GiftCardIssuer = "boekenbon"
	DecaudeuKaart              GiftCardIssuer = "decadeaukaart"
	DelokaleDecauKaart         GiftCardIssuer = "delokalecadeaukaart"
	Dinercadeau                GiftCardIssuer = "dinercadeau"
	Fashioncheque              GiftCardIssuer = "fashioncheque"
	Festivalcadeau             GiftCardIssuer = "festivalcadeau"
	Good4fun                   GiftCardIssuer = "good4fun"
	HuistuinCadeauKaart        GiftCardIssuer = "huistuincadeaukaart"
	KlusCadeu                  GiftCardIssuer = "kluscadeau"
	Kunstencultuurcadeaukaart  GiftCardIssuer = "kunstencultuurcadeaukaart"
	Nationalebioscoopbon       GiftCardIssuer = "nationalebioscoopbon"
	Nationaleentertainmentcard GiftCardIssuer = "nationaleentertainmentcard"
	Nationalegolfbon           GiftCardIssuer = "nationalegolfbon"
	Ohmygood                   GiftCardIssuer = "ohmygood"
	Podiumcadeaukaart          GiftCardIssuer = "podiumcadeaukaart"
	Reiscadeau                 GiftCardIssuer = "reiscadeau"
	Restaurantcadeau           GiftCardIssuer = "restaurantcadeau"
	SodexoSportCulturePass     GiftCardIssuer = "sodexosportculturepass"
	Sportenfitcadeau           GiftCardIssuer = "sportenfitcadeau"
	Sustainablefashion         GiftCardIssuer = "sustainablefashion"
	Travelcheq                 GiftCardIssuer = "travelcheq"
	Vvvgiftcard                GiftCardIssuer = "vvvgiftcard"
	Vvvdinercheque             GiftCardIssuer = "vvvdinercheque"
	Vvvlekkerweg               GiftCardIssuer = "vvvlekkerweg"
	Webshopgiftcard            GiftCardIssuer = "webshopgiftcard"
	Yourgift                   GiftCardIssuer = "yourgift"
)
```

Supported gift card issuers

#### type GiftCardIssuerStatus

```go
type GiftCardIssuerStatus string
```

GiftCardIssuerStatus describes the status of a gift card issuer in your account.

```go
const (
	PendingIssuer GiftCardIssuerStatus = "pending-issuer"
	EnabledIssuer GiftCardIssuerStatus = "enabled"
)
```

Valid issuer statuses

#### type GiftCardLinks

```go
type GiftCardLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

GiftCardLinks are links embedded when a gift card is enabled.

#### type Image

```go
type Image struct {
	Size1x string `json:"size1X,omitempty"`
	Size2X string `json:"size2X,omitempty"`
	Svg    string `json:"svg,omitempty"`
}
```

Image describes a generic image resource retrieved by Mollie.

#### type Invoice

```go
type Invoice struct {
	Resource    string        `json:"resource,omitempty"`
	ID          string        `json:"id,omitempty"`
	Reference   string        `json:"reference,omitempty"`
	VatNumber   string        `json:"vatNumber,omitempty"`
	Status      InvoiceStatus `json:"status,omitempty"`
	IssuedAt    string        `json:"issuedAt,omitempty"`
	PaidAt      string        `json:"paidAt,omitempty"`
	DueAt       string        `json:"dueAt,omitempty"`
	NetAmount   *Amount       `json:"netAmount,omitempty"`
	VatAmount   *Amount       `json:"vatAmount,omitempty"`
	GrossAmount *Amount       `json:"grossAmount,omitempty"`
	Lines       []*LineItem   `json:"lines,omitempty"`
	Links       InvoiceLinks  `json:"_links,omitempty"`
}
```

Invoice describes an invoice details

#### type InvoiceLinks

```go
type InvoiceLinks struct {
	Self          *URL `json:"self,omitempty"`
	PDF           *URL `json:"pdf,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

InvoiceLinks describes all the possible links to be returned with a invoice
object.

#### type InvoiceList

```go
type InvoiceList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Invoices []Invoice `json:"invoices"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

InvoiceList describes how a list of invoices will be retrieved by Mollie.

#### type InvoiceStatus

```go
type InvoiceStatus string
```

InvoiceStatus status of the invoice

```go
const (
	InvoiceStatusOpen    InvoiceStatus = "open"
	InvoiceStatusPaid    InvoiceStatus = "paid"
	InvoiceStatusOverdue InvoiceStatus = "overdue"
)
```

Valid status of the invoice

#### type InvoicesService

```go
type InvoicesService service
```

InvoicesService instance operates over invoice resources

#### func (\*InvoicesService) Get

```go
func (is *InvoicesService) Get(id string) (i Invoice, err error)
```

Get retrieve details of an invoice, using the invoice’s identifier.

#### func (\*InvoicesService) List

```go
func (is *InvoicesService) List(options *ListInvoiceOptions) (il InvoiceList, err error)
```

List retrieves a list of invoices associated with your account/organization.

#### type LineItem

```go
type LineItem struct {
	Period        string  `json:"period,omitempty"`
	Description   string  `json:"description,omitempty"`
	Count         int64   `json:"count,omitempty"`
	VatPercentage float64 `json:"vatPercentage,omitempty"`
	Amount        *Amount `json:"amount,omitempty"`
}
```

LineItem product details

#### type ListChargebackOptions

```go
type ListChargebackOptions struct {
	Include   string `url:"include,omitempty"`
	Embed     string `url:"embed,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}
```

ListChargebackOptions describes list chargebacks endpoint valid query string
parameters.

#### type ListCustomersOptions

```go
type ListCustomersOptions struct {
	From         string       `url:"from,omitempty"`
	Limit        int          `url:"limit,omitempty"`
	ProfileID    string       `url:"profileId,omitempty"`
	SequenceType SequenceType `url:"sequenceType,omitempty"`
	RedirectURL  string       `url:"redirectUrl,omitempty"`
}
```

ListCustomersOptions contains valid query parameters for the list customers
endpoint.

#### type ListInvoiceOptions

```go
type ListInvoiceOptions struct {
	Reference string `url:"reference,omitempty"`
	Year      string `url:"year,omitempty"`
	From      string `url:"from,omitempty"`
	Limit     int64  `url:"limit,omitempty"`
}
```

ListInvoiceOptions describes list invoices endpoint valid query string
parameters.

#### type ListMandatesOptions

```go
type ListMandatesOptions struct {
	From  string `url:"from,omitempty"`
	Limit int    `url:"limit,omitempty"`
}
```

ListMandatesOptions contains valid query parameters to filter the List mandates
actions. From is a mandate id to offset from (inclusive) Limit is the max number
of mandates to retrieve

#### type ListMethods

```go
type ListMethods struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Methods []*PaymentMethodInfo
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

ListMethods describes a list of paginated payment methods.

#### type ListPartnerClientsOptions

```go
type ListPartnerClientsOptions struct {
	From  int `url:"from,omitempty"`
	Limit int `url:"limit,omitempty"`
	Year  int `url:"year,omitempty"`
	Month int `url:"month,omitempty"`
}
```

ListPartnerClientsOptions contains valid query parameters for the list clients
endpoint.

#### type ListPaymentOptions

```go
type ListPaymentOptions struct {
	Limit     int    `url:"limit,omitempty"`
	Include   string `url:"include,omitempty"`
	Embed     string `url:"embed,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitempty"`
}
```

ListPaymentOptions describes list payments endpoint valid query string
parameters.

#### type ListRefundOptions

```go
type ListRefundOptions struct {
	From      string     `url:"from,omitempty"`
	Limit     int        `url:"limit,omitempty"`
	ProfileID string     `url:"profileId,omitempty"`
	Embed     EmbedValue `url:"embed,omitempty"`
}
```

ListRefundOptions describes list refund endpoint valid query string parameters.

See: https://docs.mollie.com/reference/v2/refunds-api/list-refunds.

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

#### type Mandate

```go
type Mandate struct {
	ID               string         `json:"id,omitempty"`
	Resource         string         `json:"resource,omitempty"`
	Method           PaymentMethod  `json:"method,omitempty"`
	ConsumerName     string         `json:"consumerName,omitempty"`
	ConsumerAccount  string         `json:"consumerAccount,omitempty"`
	ConsumerBic      string         `json:"consumerBic,omitempty"`
	SignatureDate    *ShortDate     `json:"signatureDate,omitempty"`
	MandateReference string         `json:"mandateReference,omitempty"`
	Mode             Mode           `json:"mode,omitempty"`
	Status           MandateStatus  `json:"status,omitempty"`
	CreatedAt        *time.Time     `json:"createdAt,omitempty"`
	Details          MandateDetails `json:"details,omitempty"`
	Links            MandateLinks   `json:"_links,omitempty"`
}
```

Mandate allow you to charge a customer’s credit card or bank account
recurrently.

#### type MandateDetails

```go
type MandateDetails struct {
	ConsumerName    string     `json:"consumerName,omitempty"`
	ConsumerAccount string     `json:"consumerAccount,omitempty"`
	ConsumerBic     string     `json:"consumerBic,omitempty"`
	CardHolder      string     `json:"cardHolder,omitempty"`
	CardNumber      string     `json:"cardNumber,omitempty"`
	CardLabel       CardLabel  `json:"cardLabel,omitempty"`
	CardFingerprint string     `json:"cardFingerprint,omitempty"`
	CardExpiryDate  *ShortDate `json:"cardExpiryDate,omitempty"`
}
```

MandateDetails are possible values inside the mandate.details field

#### type MandateLinks

```go
type MandateLinks struct {
	Self          *URL `json:"self,omitempty"`
	Customer      *URL `json:"customer,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

MandateLinks response objects

#### type MandateList

```go
type MandateList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Mandates []Mandate
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

MandateList describes how a list of mandates will be retrieved by Mollie.

#### type MandateStatus

```go
type MandateStatus string
```

MandateStatus for the Mandate object

```go
const (
	PendingMandate MandateStatus = "pending"
	ValidMandate   MandateStatus = "valid"
	InvalidMandate MandateStatus = "invalid"
)
```

Valid mandate statuses

#### type MandatesService

```go
type MandatesService service
```

MandatesService operates over customer mandates endpoints.

#### func (\*MandatesService) Create

```go
func (ms *MandatesService) Create(cID string, mandate Mandate) (mr *Mandate, err error)
```

Create a mandate for a specific customer. Mandates allow you to charge a
customer’s credit card or bank account recurrently.

See: https://docs.mollie.com/reference/v2/mandates-api/create-mandate

#### func (\*MandatesService) Get

```go
func (ms *MandatesService) Get(cID, mID string) (mr *Mandate, err error)
```

Get retrieves a mandate by its ID and its customer’s ID. The mandate will either
contain IBAN or credit card details, depending on the type of mandate.

See: https://docs.mollie.com/reference/v2/mandates-api/get-mandate

#### func (\*MandatesService) List

```go
func (ms *MandatesService) List(cID string, opt *ListMandatesOptions) (ml MandateList, err error)
```

List retrieves all mandates for the given customerId, ordered from newest to
oldest.

See: https://docs.mollie.com/reference/v2/mandates-api/list-mandates

#### func (\*MandatesService) Revoke

```go
func (ms *MandatesService) Revoke(cID, mID string) (err error)
```

Revoke a customer’s mandate. You will no longer be able to charge the consumer’s
bank account or credit card with this mandate and all connected subscriptions
will be canceled.

See: https://docs.mollie.com/reference/v2/mandates-api/revoke-mandate

#### type MethodsLinks

```go
type MethodsLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
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
	AmountCurrency string       `url:"amount[currency],omitempty"`
	AmountValue    string       `url:"amount[value],omitempty"`
	Resource       string       `url:"resource,omitempty"`
	BillingCountry string       `url:"billingCountry,omitempty"`
	IncludeWallets string       `url:"includeWallets,omitempty"`
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

#### type MiscellaneousService

```go
type MiscellaneousService service
```

MiscellaneousService operates over the resources described in Mollie's
miscellaneous API endpoints section

#### func (\*MiscellaneousService) ApplePaymentSession

```go
func (ms *MiscellaneousService) ApplePaymentSession(asr *ApplePaymentSessionRequest) (aps *ApplePaymentSession, err error)
```

ApplePaymentSession returns an Apple Payment Session object valid for one
transaction.

See:
https://docs.mollie.com/reference/v2/wallets-api/request-apple-pay-payment-session

#### type Mode

```go
type Mode string
```

Mode contains information about the creation environment.

```go
const (
	LiveMode Mode = "live"
	TestMode Mode = "test"
)
```

Valid modes

#### type Onboarding

```go
type Onboarding struct {
	Resource              string           `json:"reference,omitempty"`
	Name                  string           `json:"name,omitempty"`
	SignedUpAt            *time.Time       `json:"signedUpAt,omitempty"`
	Status                OnboardingStatus `json:"status,omitempty"`
	CanReceivePayments    bool             `json:"canReceivePayments,omitempty"`
	CanReveiceSettlements bool             `json:"canReceiveSettlements,omitempty"`
	Links                 OnboardingLinks  `json:"_links,omitempty"`
}
```

Onboarding data for an organization.

#### type OnboardingData

```go
type OnboardingData struct {
	Organization struct {
		Name               string   `json:"name,omitempty"`
		Address            *Address `json:"address,omitempty"`
		RegistrationNumber string   `json:"registrationNumber,omitempty"`
		VatNumber          string   `json:"vatNumber,omitempty"`
		VatRegulation      string   `json:"vatRegulation,omitempty"`
	} `json:"organization,omitempty"`
	Profile struct {
		Name         string       `json:"name,omitempty"`
		URL          string       `json:"url,omitempty"`
		Email        string       `json:"email,omitempty"`
		Description  string       `json:"description,omitempty"`
		Phone        string       `json:"phone,omitempty"`
		CategoryCode CategoryCode `json:"categoryCode,omitempty"`
	} `json:"profile,omitempty"`
}
```

OnboardingData request possible values.

Please note that even though all parameters are optional, at least one of them
needs to be provided in the request.

Information that the merchant has entered in their dashboard will not be
overwritten.

#### type OnboardingLinks

```go
type OnboardingLinks struct {
	Self          *URL `json:"self,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
	Organization  *URL `json:"organization,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

OnboardingLinks contains URL objects relevant to the onboarding status.

#### type OnboardingService

```go
type OnboardingService service
```

OnboardingService operates over the onboarding API.

#### func (\*OnboardingService) GetOnboardingStatus

```go
func (os *OnboardingService) GetOnboardingStatus() (o *Onboarding, err error)
```

GetOnboardingStatus gets the status of onboarding of the authenticated
organization.

See: https://docs.mollie.com/reference/v2/onboarding-api/get-onboarding-status

#### func (\*OnboardingService) SubmitOnboardingData

```go
func (os *OnboardingService) SubmitOnboardingData(d *OnboardingData) (err error)
```

SubmitOnboardingData sends data that will be prefilled in the merchant’s
onboarding. Please note that the data you submit will only be processed when the
onboarding status is needs-data.

See: https://docs.mollie.com/reference/v2/onboarding-api/submit-onboarding-data

#### type OnboardingStatus

```go
type OnboardingStatus string
```

OnboardingStatus describes status of the organization’s onboarding process.

```go
const (
	NeedsDataOnboardingStatus OnboardingStatus = "needs-data"
	InReviewOnboardingStatus  OnboardingStatus = "in-review"
	CompletedOnboardingStatus OnboardingStatus = "completed"
)
```

Possible status values.

#### type Order

```go
type Order struct {
	Resource                                 string        `json:"resource,omitempty"`
	ID                                       string        `json:"id,omitempty"`
	ProfileID                                string        `json:"profileId,omitempty"`
	TestMode                                 bool          `json:"testmode,omitempty"`
	Method                                   PaymentMethod `json:"method,omitempty"`
	Mode                                     Mode          `json:"mode,omitempty"`
	Amount                                   *Amount       `json:"amount,omitempty"`
	AmountCaptured                           *Amount       `json:"amountCaptured,omitempty"`
	AmountRefunded                           *Amount       `json:"amountRefunded,omitempty"`
	Status                                   OrderStatus   `json:"status,omitempty"`
	IsCancelable                             bool          `json:"isCancelable,omitempty"`
	BillingAddress                           *OrderAddress `json:"billingAddress,omitempty"`
	ConsumerDateOfBirth                      *ShortDate    `json:"consumerDateOfBirth,omitempty"`
	OrderNumber                              string        `json:"orderNumber,omitempty"`
	ShippingAddress                          OrderAddress  `json:"shippingAddress,omitempty"`
	Locale                                   Locale        `json:"locale,omitempty"`
	Metadata                                 interface{}   `json:"metadata,omitempty"`
	RedirectURL                              string        `json:"redirectUrl,omitempty"`
	Lines                                    []*OrderLine  `json:"lines,omitempty"`
	WebhookURL                               string        `json:"webhookUrl,omitempty"`
	CreatedAt                                *time.Time    `json:"createdAt,omitempty"`
	ExpiresAt                                *time.Time    `json:"expiresAt,omitempty"`
	ExpiredAt                                *time.Time    `json:"expiredAt,omitempty"`
	PaidAt                                   *time.Time    `json:"paidAt,omitempty"`
	AuthorizedAt                             *time.Time    `json:"authorizedAt,omitempty"`
	CanceledAt                               *time.Time    `json:"canceledAt,omitempty"`
	CompletedAt                              *time.Time    `json:"completedAt,omitempty"`
	Links                                    OrderLinks    `json:"_links,omitempty"`
	OrderPayment                             *OrderPayment `json:"payment,omitempty"`
	Description                              string        `json:"description,omitempty"`
	ShopperCountryMustMatchTheBillingCountry bool          `json:"shopperCountryMustMatchTheBillingCountry,omitempty"`
	Sku                                      string        `json:"sku,omitempty"`
	Embedded                                 struct {
		Payments []*Payment `json:"payments,omitempty"`
		Refunds  []*Refund  `json:"refunds,omitempty"`
	} `json:"_embedded,omitempty"`
}
```

Order explain the items that customers need to pay for.

#### type OrderAddress

```go
type OrderAddress struct {
	OrganizationName string      `json:"organizationName,omitempty"`
	Title            string      `json:"title,omitempty"`
	GivenName        string      `json:"givenName,omitempty"`
	FamilyName       string      `json:"familyName,omitempty"`
	Email            string      `json:"email,omitempty"`
	Phone            PhoneNumber `json:"phone,omitempty"`
	StreetAndNumber  string      `json:"streetAndNumber,omitempty"`
	StreetAdditional string      `json:"streetAdditional,omitempty"`
	PostalCode       string      `json:"postalCode,omitempty"`
	City             string      `json:"city,omitempty"`
	Region           string      `json:"region,omitempty"`
	Country          string      `json:"country,omitempty"`
}
```

OrderAddress identify both the address and the person the order is billed or
shipped to.

#### type OrderLine

```go
type OrderLine struct {
	Resource           string          `json:"resource,omitempty"`
	ID                 string          `json:"id,omitempty"`
	OrderID            string          `json:"orderId,omitempty"`
	ProductType        ProductType     `json:"type,omitempty"`
	Name               string          `json:"name,omitempty"`
	Amount             *Amount         `json:"amount,omitempty"`
	Status             OrderLineStatus `json:"status,omitempty"`
	IsCancelable       bool            `json:"isCancelable,omitempty"`
	Quantity           int             `json:"quantity,omitempty"`
	QuantityShipped    int             `json:"quantityShipped,omitempty"`
	AmountShipped      *Amount         `json:"amountShipped,omitempty"`
	QuantityRefunded   int             `json:"quantityRefunded,omitempty"`
	AmountRefunded     *Amount         `json:"amountRefunded,omitempty"`
	QuantityCanceled   int             `json:"quantityCanceled,omitempty"`
	AmountCanceled     *Amount         `json:"amountCanceled,omitempty"`
	ShippableQuantity  int             `json:"shippableQuantity,omitempty"`
	RefundableQuantity int             `json:"refundableQuantity,omitempty"`
	CancelableQuantity int             `json:"cancelableQuantity,omitempty"`
	UnitPrice          *Amount         `json:"unitPrice,omitempty"`
	DiscountAmount     *Amount         `json:"discountAmount,omitempty"`
	TotalAmount        *Amount         `json:"totalAmount,omitempty"`
	VatRate            string          `json:"vatRate,omitempty"`
	VatAmount          *Amount         `json:"vatAmount,omitempty"`
	SKU                string          `json:"sku,omitempty"`
	CreatedAt          *time.Time      `json:"createdAt,omitempty"`
	Links              OrderLineLinks  `json:"_links,omitempty"`
	ImageURL           string          `json:"imageUrl,omitempty"`
	ProductURL         string          `json:"productUrl,omitempty"`
}
```

OrderLine contain the actual things the customer bought.

#### type OrderLineLinks

```go
type OrderLineLinks struct {
	ProductURL *URL `json:"productUrl,omitempty"`
	ImageURL   *URL `json:"imageUrl,omitempty"`
}
```

OrderLineLinks describes object with several URL objects relevant to the order
line.

#### type OrderLineStatus

```go
type OrderLineStatus string
```

OrderLineStatus describes status of the order line.

```go
const (
	OrderLineCreated    OrderLineStatus = "created"
	OrderLineAuthorized OrderLineStatus = "authorized"
	OrderLinePaid       OrderLineStatus = "paid"
	OrderLineShipping   OrderLineStatus = "shipping"
	OrderLineCanceled   OrderLineStatus = "canceled"
	OrderLineCompleted  OrderLineStatus = "completed"
)
```

Valid order line status.

#### type OrderLinks

```go
type OrderLinks struct {
	Self          *URL `json:"self,omitempty"`
	Checkout      *URL `json:"checkout,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
}
```

OrderLinks describes an object with several URL objects relevant to the order.
Every URL object will contain an href and a type field.

#### type OrderList

```go
type OrderList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Orders []*Order `json:"orders,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}
```

OrderList for containing the response of list orders

#### type OrderListOptions

```go
type OrderListOptions struct {
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}
```

OrderListOptions describes order endpoint valid query string parameters.

#### type OrderListRefund

```go
type OrderListRefund struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Refunds []*Refund `json:"refund,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}
```

OrderListRefund for containing the response of list orders

#### type OrderListRefundOptions

```go
type OrderListRefundOptions struct {
	From  string     `url:"from,omitempty"`
	Limit int        `url:"limit,omitempty"`
	Embed EmbedValue `url:"embed,omitempty"`
}
```

OrderListRefundOptions describes order endpoint valid query string parameters.

#### type OrderOptions

```go
type OrderOptions struct {
	Embed     []EmbedValue `url:"embed,omitempty"`
	ProfileID string       `url:"profileId,omitempty"`
}
```

OrderOptions describes order endpoint valid query string parameters.

#### type OrderPayment

```go
type OrderPayment struct {
	ConsumerAccount   string          `json:"consumerAccount,omitempty"`
	CustomerID        string          `json:"customerId,omitempty"`
	CustomerReference string          `json:"customerReference,omitempty"`
	Issuer            string          `json:"issuer,omitempty"`
	MandateID         string          `json:"mandateId,omitempty"`
	SequenceType      SequenceType    `json:"sequenceType,omitempty"`
	VoucherNumber     string          `json:"voucherNumber,omitempty"`
	VoucherPin        string          `json:"voucherPin,omitempty"`
	WebhookURL        string          `json:"webhookUrl,omitempty"`
	ApplicationFee    *ApplicationFee `json:"applicationFee,omitempty"`
	Method            PaymentMethod   `json:"method,omitempty"`
}
```

OrderPayment describes payment specific parameters that can be passed during
order creation

#### type OrderStatus

```go
type OrderStatus string
```

OrderStatus describes the valid order status.

```go
const (
	Created    OrderStatus = "created"
	Paid       OrderStatus = "paid"
	Authorized OrderStatus = "authorized"
	Canceled   OrderStatus = "canceled"
	Shipping   OrderStatus = "shipping"
	Completed  OrderStatus = "completed"
	Expired    OrderStatus = "expired"
)
```

Valid order status.

#### type OrdersService

```go
type OrdersService service
```

OrdersService instance operates over refund resources.

#### func (\*OrdersService) Cancel

```go
func (ors *OrdersService) Cancel(orderID string) (order *Order, err error)
```

Cancel try to cancel the order that fulfill certain requirements.

See https://docs.mollie.com/reference/v2/orders-api/cancel-order

#### func (\*OrdersService) CancelOrderLines

```go
func (ors *OrdersService) CancelOrderLines(orderID string, orderLines []OrderLine) (err error)
```

CancelOrderLines can be used to cancel one or more order lines that were
previously authorized using a pay after delivery payment method. Use the Cancel
Order API if you want to cancel the entire order or the remainder of the order.

See https://docs.mollie.com/reference/v2/orders-api/cancel-order-lines

#### func (\*OrdersService) Create

```go
func (ors *OrdersService) Create(ord Order, opt *OrderOptions) (order *Order, err error)
```

Create an order will automatically create the required payment to allow your
customer to pay for the order.

See https://docs.mollie.com/reference/v2/orders-api/create-order

#### func (\*OrdersService) CreateOrderPayment

```go
func (ors *OrdersService) CreateOrderPayment(orderID string, ordPay *OrderPayment) (payment *Payment, err error)
```

CreateOrderPayment can only be created while the status of the order is created,
and when the status of the existing payment is either expired, canceled or
failed. See https://docs.mollie.com/reference/v2/orders-api/create-order-payment

#### func (\*OrdersService) CreateOrderRefund

```go
func (ors *OrdersService) CreateOrderRefund(orderID string, order *Order) (refund Refund, err error)
```

CreateOrderRefund using the Orders API, refunds should be made against the
order. See https://docs.mollie.com/reference/v2/orders-api/create-order-refund

#### func (\*OrdersService) Get

```go
func (ors *OrdersService) Get(orID string, opt *OrderOptions) (order *Order, err error)
```

Get retrieve a single order by its ID.

See https://docs.mollie.com/reference/v2/orders-api/get-order

#### func (\*OrdersService) List

```go
func (ors *OrdersService) List(opt *OrderListOptions) (ordList *OrderList, err error)
```

List is to retrieve all orders.

See https://docs.mollie.com/reference/v2/orders-api/list-orders

#### func (\*OrdersService) ListOrderRefunds

```go
func (ors *OrdersService) ListOrderRefunds(orderID string, opt *OrderListRefundOptions) (orderListRefund OrderListRefund, err error)
```

ListOrderRefunds retrieve all order refunds. See
https://docs.mollie.com/reference/v2/orders-api/list-order-refunds

#### func (\*OrdersService) Update

```go
func (ors *OrdersService) Update(orderID string, ord Order) (order *Order, err error)
```

Update is used to update the billing and/or shipping address of an order.

See https://docs.mollie.com/reference/v2/orders-api/update-order

#### func (\*OrdersService) UpdateOrderLine

```go
func (ors *OrdersService) UpdateOrderLine(orderID string, orderLineID string, orderLine OrderLine) (order *Order, err error)
```

UpdateOrderLine can be used to update an order line.

See https://docs.mollie.com/reference/v2/orders-api/update-orderline

#### type Organization

```go
type Organization struct {
	Resource           string            `json:"resource,omitempty"`
	ID                 string            `json:"id,omitempty"`
	Name               string            `json:"name,omitempty"`
	Email              string            `json:"email,omitempty"`
	Locale             string            `json:"locale,omitempty"`
	Address            *Address          `json:"address,omitempty"`
	RegistrationNumber string            `json:"registrationNumber,omitempty"`
	VatNumber          string            `json:"vatNumber,omitempty"`
	VatRegulation      string            `json:"vatRegulation,omitempty"`
	Links              OrganizationLinks `json:"_links,omitempty"`
}
```

Organization describes an organization detail

#### type OrganizationLinks

```go
type OrganizationLinks struct {
	Self          *URL `json:"self,omitempty"`
	Chargebacks   *URL `json:"chargebacks,omitempty"`
	Customers     *URL `json:"customers,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
	Invoices      *URL `json:"invoices,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Profiles      *URL `json:"profiles,omitempty"`
	Refunds       *URL `json:"refunds,omitempty"`
	Settlements   *URL `json:"settlements,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

OrganizationLinks describes all the possible links to be returned with a
organization object.

#### type OrganizationsService

```go
type OrganizationsService service
```

OrganizationsService instance operates over organization resources

#### func (\*OrganizationsService) Get

```go
func (os *OrganizationsService) Get(id string) (o *Organization, err error)
```

Get retrieve an organization by its id.

#### func (\*OrganizationsService) GetCurrent

```go
func (os *OrganizationsService) GetCurrent() (o *Organization, err error)
```

GetCurrent retrieve the currently authenticated organization

#### type PaginationLinks

```go
type PaginationLinks struct {
	Self          URL `json:"self,omitempty"`
	Previous      URL `json:"previous,omitempty"`
	Next          URL `json:"next,omitempty"`
	Documentation URL `json:"documentation,omitempty"`
}
```

PaginationLinks describes the hal component of paginated responses.

#### type PartnerClient

```go
type PartnerClient struct {
	Resource              string             `json:"resource,omitempty"`
	ID                    string             `json:"id,omitempty"`
	OrganizationCreatedAt *time.Time         `json:"organizationCreatedAt,omitempty"`
	Commission            Commission         `json:"commission,omitempty"`
	Links                 PartnerClientLinks `json:"_links,omitempty"`
}
```

PartnerClient describes a partner client.

#### type PartnerClientLinks

```go
type PartnerClientLinks struct {
	Self          *URL `json:"self,omitempty"`
	Organization  *URL `json:"organization,omitempty"`
	Onboarding    *URL `json:"onboarding,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

PartnerClientLinks contains URL objects relevant to the client.

#### type PartnerClientList

```go
type PartnerClientList struct {
	Count          int `json:"count,omitempty"`
	PartnerClients struct {
		Clients []*PartnerClient `json:"clients,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

PartnerClientList describes a list of partner clients.

#### type PartnerService

```go
type PartnerService service
```

PartnerService operates over the partners API.

#### func (\*PartnerService) Get

```go
func (ps *PartnerService) Get(id string, opts *GetPartnerClientOptions) (pc *PartnerClient, err error)
```

Get retrieves a single client by its ID.

See: https://docs.mollie.com/reference/v2/partners-api/get-client

#### func (\*PartnerService) List

```go
func (ps *PartnerService) List(opts *ListPartnerClientsOptions) (pc *PartnerClientList, err error)
```

Get retrieves all clients.

See: https://docs.mollie.com/reference/v2/partners-api/list-clients

#### type Payment

```go
type Payment struct {
	IsCancellable                   bool                   `json:"isCancellable,omitempty"`
	TestMode                        bool                   `json:"testmode,omitempty"`
	DigitalGoods                    bool                   `json:"digitalGoods,omitempty"`
	ApplePayPaymentToken            string                 `json:"applePayPaymentToken,omitempty"`
	BillingEmail                    string                 `json:"billingEmail,omitempty"`
	CardToken                       string                 `json:"cardToken,omitempty"`
	Issuer                          string                 `json:"issuer,omitempty"`
	VoucherNumber                   string                 `json:"voucherNumber,omitempty"`
	VoucherPin                      string                 `json:"voucherPin,omitempty"`
	ExtraMerchantData               string                 `json:"extraMerchantData,omitempty"`
	SessionID                       string                 `json:"sessionId,omitempty"`
	CustomerReference               string                 `json:"customerReference,omitempty"`
	ConsumerName                    string                 `json:"consumerName,omitempty"`
	ConsumerAccount                 string                 `json:"consumerAccount,omitempty"`
	WebhookURL                      string                 `json:"webhookUrl,omitempty"`
	Resource                        string                 `json:"resource,omitempty"`
	ID                              string                 `json:"id,omitempty"`
	MandateID                       string                 `json:"mandateId,omitempty"`
	OrderID                         string                 `json:"orderId,omitempty"`
	ProfileID                       string                 `json:"profileId,omitempty"`
	SettlementID                    string                 `json:"settlementId,omitempty"`
	CustomerID                      string                 `json:"customerId,omitempty"`
	Status                          string                 `json:"status,omitempty"`
	Description                     string                 `json:"description,omitempty"`
	RedirectURL                     string                 `json:"redirectUrl,omitempty"`
	CountryCode                     string                 `json:"countryCode,omitempty"`
	SubscriptionID                  string                 `json:"subscriptionId,omitempty"`
	Metadata                        interface{}            `json:"metadata,omitempty"`
	Amount                          *Amount                `json:"amount,omitempty"`
	AmountRefunded                  *Amount                `json:"amountRefunded,omitempty"`
	AmountRemaining                 *Amount                `json:"amountRemaining,omitempty"`
	AmountCaptured                  *Amount                `json:"amountCaptured,omitempty"`
	SettlementAmount                *Amount                `json:"settlementAmount,omitempty"`
	ApplicationFee                  *ApplicationFee        `json:"applicationFee,omitempty"`
	Details                         *PaymentDetails        `json:"details,omitempty"`
	CreatedAt                       *time.Time             `json:"createdAt,omitempty"`
	AuthorizedAt                    *time.Time             `json:"authorizedAt,omitempty"`
	PaidAt                          *time.Time             `json:"paidAt,omitempty"`
	CanceledAt                      *time.Time             `json:"canceledAt,omitempty"`
	ExpiresAt                       *time.Time             `json:"expiresAt,omitempty"`
	ExpiredAt                       *time.Time             `json:"expiredAt,omitempty"`
	FailedAt                        *time.Time             `json:"failedAt,omitempty"`
	DueDate                         *ShortDate             `json:"dueDate,omitempty"`
	BillingAddress                  *Address               `json:"billingAddress,omitempty"`
	ShippingAddress                 *PaymentDetailsAddress `json:"shippingAddress,omitempty"`
	Mode                            Mode                   `json:"mode,omitempty"`
	Locale                          Locale                 `json:"locale,omitempty"`
	RestrictPaymentMethodsToCountry Locale                 `json:"restrictPaymentMethodsToCountry,omitempty"`
	Method                          PaymentMethod          `json:"method,omitempty"`
	Links                           PaymentLinks           `json:"_links,omitempty"`
	SequenceType                    SequenceType           `json:"sequenceType,omitempty"`
}
```

Payment describes a transaction between a customer and a merchant.

#### type PaymentDetails

```go
type PaymentDetails struct {
	PaypalDigitalGoods bool                   `json:"digitalGoods,omitempty"`
	BankAccount        string                 `json:"bankAccount,omitempty"`
	BankBIC            string                 `json:"bankBic,omitempty"`
	BankName           string                 `json:"bankName,omitempty"`
	BankReason         string                 `json:"bankReason,omitempty"`
	BatchReference     string                 `json:"batchReference,omitempty"`
	BillingEmail       string                 `json:"billingEmail,omitempty"`
	CardAudience       string                 `json:"cardAudience,omitempty"`
	CardCountryCode    string                 `json:"cardCountryCode,omitempty"`
	CardFingerPrint    string                 `json:"cardFingerPrint,omitempty"`
	CardHolder         string                 `json:"cardHolder,omitempty"`
	CardLabel          string                 `json:"cardLabel,omitempty"`
	CardNumber         string                 `json:"cardNumber,omitempty"`
	CardSecurity       string                 `json:"cardSecurity,omitempty"`
	ConsumerAccount    string                 `json:"consumerAccount,omitempty"`
	ConsumerBIC        string                 `json:"consumerBic,omitempty"`
	ConsumerName       string                 `json:"consumerName,omitempty"`
	ConsumerReference  string                 `json:"consumerReference,omitempty"`
	CreditorIdentifier string                 `json:"creditorIdentifier,omitempty"`
	EndToEndIdentifier string                 `json:"endToEndIdentifier,omitempty"`
	FileReference      string                 `json:"fileReference,omitempty"`
	MandateReference   string                 `json:"mandateReference,omitempty"`
	PaypalReference    string                 `json:"paypalReference,omitempty"`
	PaypalPayerID      string                 `json:"paypalPayerId,omitempty"`
	TransferReference  string                 `json:"transferReference,omitempty"`
	VoucherNumber      string                 `json:"voucherNumber,omitempty"`
	Wallet             string                 `json:"wallet,omitempty"`
	PaypalFee          Amount                 `json:"paypalFee,omitempty"`
	FailureReason      FailureReason          `json:"failureReason,omitempty"`
	FeeRegion          FeeRegion              `json:"feeRegion,omitempty"`
	RemainderMethod    PaymentMethod          `json:"remainderMethod,omitempty"`
	SellerProtection   EligibilityReasons     `json:"sellerProtection,omitempty"`
	DueDate            *ShortDate             `json:"dueDate,omitempty"`
	GiftCards          []*UsedGiftCard        `json:"giftCards,omitempty"`
	QRCode             *QRCode                `json:"qrCode,omitempty"`
	RemainderAmount    *Amount                `json:"remainderAmount,omitempty"`
	ShippingAddress    *PaymentDetailsAddress `json:"shippingAddress,omitempty"`
	SignatureDate      *ShortDate             `json:"signatureDate,omitempty"`
	Links              struct {
		Status    *URL `json:"status,omitempty"`
		PayOnline *URL `json:"payOnline,omitempty"`
	} `json:"_links,omitempty"`
}
```

PaymentDetails contains details for the specified payment method

#### type PaymentDetailsAddress

```go
type PaymentDetailsAddress struct {
	OrganizationName string      `json:"organizationName,omitempty"`
	Title            string      `json:"title,omitempty"`
	GivenName        string      `json:"givenName,omitempty"`
	FamilyName       string      `json:"familyName,omitempty"`
	Email            string      `json:"email,omitempty"`
	Phone            PhoneNumber `json:"phone,omitempty"`
	StreetAndNumber  string      `json:"streetAndNumber,omitempty"`
	StreetAdditional string      `json:"streetAdditional,omitempty"`
	PostalCode       string      `json:"postalCode,omitempty"`
	City             string      `json:"city,omitempty"`
	Region           string      `json:"region,omitempty"`
	Country          string      `json:"country,omitempty"`
}
```

PaymentDetailsAddress identify both the address and the person the payment is
shipped to.

#### type PaymentLink

```go
type PaymentLink struct {
	ID          string           `json:"id,omitempty"`
	Resource    string           `json:"resource,omitempty"`
	Description string           `json:"description,omitempty"`
	ProfileID   string           `json:"profileId,omitempty"`
	RedirectURL string           `json:"redirectUrl,omitempty"`
	WebhookURL  string           `json:"webhookUrl,omitempty"`
	Mode        Mode             `json:"mode,omitempty"`
	Amount      Amount           `json:"amount,omitempty"`
	CreatedAt   *time.Time       `json:"createdAt,omitempty"`
	PaidAt      *time.Time       `json:"paidAt,omitempty"`
	UpdatedAt   *time.Time       `json:"updatedAt,omitempty"`
	ExpiresAt   *time.Time       `json:"expiresAt,omitempty"`
	Links       PaymentLinkLinks `json:"_links,omitempty"`
}
```

PaymentLink is a resource that can be shared with your customers and will
redirect them to them the payment page where they can complete the payment.

See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link

#### type PaymentLinkLinks

```go
type PaymentLinkLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	PaymentLink   *URL `json:"paymentLink,omitempty"`
	Next          *URL `json:"next,omitempty"`
	Previous      *URL `json:"previous,omitempty"`
}
```

PaymentLinkLinks describes all the possible links returned with a payment link
struct.

See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link

#### type PaymentLinkOptions

```go
type PaymentLinkOptions struct {
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitemtpy"`
	Limit     int    `url:"limit,omitempty"`
}
```

PaymentLinkOptions represents query string parameters to modify the payment
links requests.

#### type PaymentLinks

```go
type PaymentLinks struct {
	Self               *URL `json:"self,omitempty"`
	Checkout           *URL `json:"checkout,omitempty"`
	ChangePaymentState *URL `json:"changePaymentState,omitempty"`
	Refunds            *URL `json:"refunds,omitempty"`
	ChargeBacks        *URL `json:"chargebacks,omitempty"`
	Captures           *URL `json:"captures,omitempty"`
	Settlement         *URL `json:"settlement,omitempty"`
	Documentation      *URL `json:"documentation,omitempty"`
	Mandate            *URL `json:"mandate,omitempty"`
	Subscription       *URL `json:"subscription,omitempty"`
	Customer           *URL `json:"customer,omitempty"`
	Order              *URL `json:"order,omitempty"`
	Dashboard          *URL `json:"dashboard,omitempty"`
}
```

PaymentLinks describes all the possible links to be returned with a payment
object.

#### type PaymentLinksList

```go
type PaymentLinksList struct {
	Count    int              `json:"count,omitempty"`
	Links    PaymentLinkLinks `json:"_links,omitempty"`
	Embedded struct {
		PaymentLinks []*PaymentLink `json:"payment_links,omitempty"`
	} `json:"_embedded,omitempty"`
}
```

#### type PaymentLinksService

```go
type PaymentLinksService service
```

PaymentLinksService operates over the payment link resource.

#### func (\*PaymentLinksService) Create

```go
func (pls *PaymentLinksService) Create(p PaymentLink, opts *PaymentLinkOptions) (np *PaymentLink, err error)
```

Create generates payment links that by default, unlike regular payments, do not
expire.

See: https://docs.mollie.com/reference/v2/payment-links-api/create-payment-link

#### func (\*PaymentLinksService) Get

```go
func (pls *PaymentLinksService) Get(id string) (pl *PaymentLink, err error)
```

Get retrieves a single payment link object by its id/token.

See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link

#### func (\*PaymentLinksService) List

```go
func (pls *PaymentLinksService) List(opts *PaymentLinkOptions) (pl *PaymentLinksList, err error)
```

List retrieves all payments links created with the current website profile,
ordered from newest to oldest.

See: https://docs.mollie.com/reference/v2/payment-links-api/list-payment-links

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
	Resource      string                  `json:"resource,omitempty"`
	ID            string                  `json:"id,omitempty"`
	Description   string                  `json:"description,omitempty"`
	MinimumAmount *Amount                 `json:"minimumAmount,omitempty"`
	MaximumAmount *Amount                 `json:"maximumAmount,omitempty"`
	Image         *Image                  `json:"image,omitempty"`
	Pricing       []*PaymentMethodPricing `json:"pricing,omitempty"`
	Issuers       []*PaymentMethodIssuer  `json:"issuers,omitempty"`
	Status        *PaymentMethodStatus    `json:"status,omitempty"`
	Links         MethodsLinks            `json:"_links,omitempty"`
}
```

PaymentMethodInfo describes a single method with details.

#### type PaymentMethodIssuer

```go
type PaymentMethodIssuer struct {
	Resource string `json:"resource,omitempty"`
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Image    Image  `json:"image,omitempty"`
}
```

PaymentMethodIssuer available for the payment method (for iDEAL, KBC/CBC payment
button, gift cards, or meal vouchers).

#### type PaymentMethodPricing

```go
type PaymentMethodPricing struct {
	Description string    `json:"description,omitempty"`
	Fixed       *Amount   `json:"fixed,omitempty"`
	Variable    string    `json:"variable,omitempty"`
	FeeRegion   FeeRegion `json:"feeRegion,omitempty"`
}
```

PaymentMethodPricing contains information about commissions and fees applicable
to a payment method.

#### type PaymentMethodStatus

```go
type PaymentMethodStatus string
```

PaymentMethodStatus tels the status that the method is in. Possible values:
activated pending-boarding pending-review pending-external rejected.

```go
const (
	PaymentMethodActivated       PaymentMethodStatus = "activated"
	PaymentMethodPendingBoarding PaymentMethodStatus = "pending-boarding"
	PaymentMethodPendingReview   PaymentMethodStatus = "pending-review"
	PaymentMethodPendingExternal PaymentMethodStatus = "pending-external"
	PaymentMethodRejected        PaymentMethodStatus = "pending-rejected"
)
```

Available payment method statuses

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
func (ps *PaymentsService) Create(p Payment, options *PaymentOptions) (np Payment, err error)
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

#### type Permission

```go
type Permission struct {
	Resource    string          `json:"resource,omitempty"`
	ID          PermissionGrant `json:"id,omitempty"`
	Description string          `json:"description,omitempty"`
	Granted     bool            `json:"granted,omitempty"`
	Links       PermissionLinks `json:"_links,omitempty"`
}
```

Permission represents an action that can be performed by any API actor.

#### type PermissionGrant

```go
type PermissionGrant string
```

PermissionGrant defines supported permissions.

```go
const (
	PaymentsRead       PermissionGrant = "payments.read"
	PaymentsWrite      PermissionGrant = "payments.write"
	RefundsRead        PermissionGrant = "refunds.read"
	RefundsWrite       PermissionGrant = "refunds.write"
	CustomersRead      PermissionGrant = "customers.read"
	CustomersWrite     PermissionGrant = "customers.write"
	MandatesRead       PermissionGrant = "mandates.read"
	MandatesWrite      PermissionGrant = "mandates.write"
	SubscriptionsRead  PermissionGrant = "subscriptions.read"
	SubscriptionsWrite PermissionGrant = "subscriptions.write"
	ProfilesRead       PermissionGrant = "profiles.read"
	ProfilesWrite      PermissionGrant = "profiles.write"
	InvoicesRead       PermissionGrant = "invoices.read"
	OrdersRead         PermissionGrant = "orders.read"
	OrdersWrite        PermissionGrant = "orders.write"
	ShipmentsRead      PermissionGrant = "shipments.read"
	ShipmentsWrite     PermissionGrant = "shipments.write"
	OrganizationsRead  PermissionGrant = "organizations.read"
	OrganizationsWrite PermissionGrant = "organizations.write"
	OnboardingRead     PermissionGrant = "onbording.read"
	OnboardingWrite    PermissionGrant = "onbording.write"
)
```

Available permission grants.

#### type PermissionLinks

```go
type PermissionLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

PermissionLinks contains URL objects that make reference to an http address
related to permissions.

#### type PermissionsList

```go
type PermissionsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Permissions []*Permission `json:"permissions,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PermissionLinks `json:"_links,omitempty"`
}
```

PermissionsList lists all the permissions given to an API actor.

#### type PermissionsService

```go
type PermissionsService service
```

PermissionsService operates over permission resources.

#### func (\*PermissionsService) Get

```go
func (ps *PermissionsService) Get(id string) (p *Permission, err error)
```

Get returns a permission by its id.

See: https://docs.mollie.com/reference/v2/permissions-api/get-permission

#### func (\*PermissionsService) List

```go
func (ps *PermissionsService) List() (pl *PermissionsList, err error)
```

List retrieves all permissions available with the current app access token. The
list is not paginated.

See: https://docs.mollie.com/reference/v2/permissions-api/list-permissions

#### type PhoneNumber

```go
type PhoneNumber string
```

PhoneNumber represents a phone number in the E.164 format.

#### type ProductType

```go
type ProductType string
```

ProductType describes the type of product bought, for example, a physical or a
digital product.

```go
const (
	Physical        ProductType = "physical"
	Discount        ProductType = "discount"
	Digital         ProductType = "digital"
	ShippingFee     ProductType = "shipping_fee"
	StoreCredit     ProductType = "store_credit"
	GiftCardProduct ProductType = "gift_card"
	Surcharge       ProductType = "surcharge"
)
```

Valid product type.

#### type Profile

```go
type Profile struct {
	ID           string       `json:"id,omitempty"`
	CategoryCode CategoryCode `json:"categoryCode,omitempty"`
	CreatedAt    *time.Time   `json:"createdAt,omitempty"`
	Email        string       `json:"email,omitempty"`
	Mode         Mode         `json:"mode,omitempty"`
	Name         string       `json:"name,omitempty"`
	Phone        PhoneNumber  `json:"phone,omitempty"`
	Resource     string       `json:"resource,omitempty"`
	Review       struct {
		Status string `json:"status,omitempty"`
	} `json:"review,omitempty"`
	Status  ProfileStatus `json:"status,omitempty"`
	Website string        `json:"website,omitempty"`
	Links   ProfileLinks  `json:"_links,omitempty"`
}
```

Profile will usually reflect the trademark or brand name of the profile’s
website or application.

#### type ProfileLinks

```go
type ProfileLinks struct {
	Self               *URL `json:"self,omitempty"`
	Chargebacks        *URL `json:"chargebacks,omitempty"`
	Methods            *URL `json:"methods,omitempty"`
	Refunds            *URL `json:"refunds,omitempty"`
	CheckoutPreviewURL *URL `json:"checkoutPreviewUrl,omitempty"`
	Documentation      *URL `json:"documentation,omitempty"`
	Dashboard          *URL `json:"dashboard,omitempty"`
}
```

ProfileLinks contains URL's to relevant information related to a profile.

#### type ProfileList

```go
type ProfileList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Profiles []*Profile `json:"profiles,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

ProfileList contains a list of profiles for your account.

#### type ProfileListOptions

```go
type ProfileListOptions struct {
	From  string `url:"from,omitempty"`
	Limit uint   `url:"limit,omitempty"`
}
```

ProfileListOptions are optional query string parameters for the list profiles
request

#### type ProfileStatus

```go
type ProfileStatus string
```

ProfileStatus determines whether the profile is able to receive live payments

```go
const (
	StatusUnverified ProfileStatus = "unverified"
	StatusVerified   ProfileStatus = "verified"
	StatusBlocked    ProfileStatus = "blocked"
)
```

Possible profile statuses

#### type ProfilesService

```go
type ProfilesService service
```

ProfilesService operates over profile resource

#### func (\*ProfilesService) Create

```go
func (ps *ProfilesService) Create(np *Profile) (p *Profile, err error)
```

Create stores a new profile in your Mollie account.

#### func (\*ProfilesService) Current

```go
func (ps *ProfilesService) Current() (p *Profile, err error)
```

Current returns the profile belonging to the API key. This method only works
when using API keys.

#### func (\*ProfilesService) Delete

```go
func (ps *ProfilesService) Delete(id string) (err error)
```

Delete enables profile deletions, rendering the profile unavailable for further
API calls and transactions.

#### func (\*ProfilesService) DisableGiftCardIssuer

```go
func (ps *ProfilesService) DisableGiftCardIssuer(profileID string, issuer GiftCardIssuer) (err error)
```

DisableGiftCardIssuer deactivates the requested giftcard issuer for the provided
profile id.

See: https://docs.mollie.com/reference/v2/profiles-api/disable-gift-card-issuer

#### func (\*ProfilesService) DisableGiftCardIssuerForCurrent

```go
func (ps *ProfilesService) DisableGiftCardIssuerForCurrent(issuer GiftCardIssuer) (err error)
```

DisableGiftCardIssuerForCurrent deactivates the specified issuer for the curent
profile (token owner).

See: https://docs.mollie.com/reference/v2/profiles-api/disable-gift-card-issuer

#### func (\*ProfilesService) DisablePaymentMethod

```go
func (ps *ProfilesService) DisablePaymentMethod(id string, pm PaymentMethod) (err error)
```

DisablePaymentMethod disables a payment method on a specific or authenticated
profile. If you're using API tokens for authentication, pass "me" as id.

#### func (\*ProfilesService) EnableGiftCardIssuer

```go
func (ps *ProfilesService) EnableGiftCardIssuer(profileID string, issuer GiftCardIssuer) (gc *GiftCardEnabled, err error)
```

EnableGiftCardIssuer activates the requested giftcard issuer for the provided
profile id.

See: https://docs.mollie.com/reference/v2/profiles-api/enable-gift-card-issuer

#### func (\*ProfilesService) EnableGiftCardIssuerForCurrent

```go
func (ps *ProfilesService) EnableGiftCardIssuerForCurrent(issuer GiftCardIssuer) (gc *GiftCardEnabled, err error)
```

EnableGiftCardIssuerForCurrent activates the specified issuer for the curent
profile (token owner).

See: https://docs.mollie.com/reference/v2/profiles-api/enable-gift-card-issuer

#### func (\*ProfilesService) EnablePaymentMethod

```go
func (ps *ProfilesService) EnablePaymentMethod(id string, pm PaymentMethod) (pmi *PaymentMethodInfo, err error)
```

EnablePaymentMethod enables a payment method on a specific or authenticated
profile. If you're using API tokens for authentication, pass "me" as id.

#### func (\*ProfilesService) Get

```go
func (ps *ProfilesService) Get(id string) (p *Profile, err error)
```

Get retrieves the a profile by ID.

#### func (\*ProfilesService) List

```go
func (ps *ProfilesService) List(options *ProfileListOptions) (pl *ProfileList, err error)
```

List returns all the profiles for the authenticated account

#### func (\*ProfilesService) Update

```go
func (ps *ProfilesService) Update(id string, up *Profile) (p *Profile, err error)
```

Update allows you to perform mutations on a profile

#### type QRCode

```go
type QRCode struct {
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	Src    string `json:"src,omitempty"`
}
```

QRCode object represents an image of a QR code.

#### type Rate

```go
type Rate struct {
	Fixed    *Amount `json:"fixed,omitempty"`
	Variable string  `json:"variable,omitempty"`
}
```

Rate describes service rates, further divided into fixed and percentage costs.

#### type Refund

```go
type Refund struct {
	Resource         string       `json:"resource,omitempty"`
	ID               string       `json:"id,omitempty"`
	Amount           *Amount      `json:"amount,omitempty"`
	SettlementID     string       `json:"settlementId,omitempty"`
	SettlementAmount *Amount      `json:"settlementAmount,omitempty"`
	Description      string       `json:"description,omitempty"`
	Metadata         interface{}  `json:"metadata,omitempty"`
	Status           RefundStatus `json:"status,omitempty"`
	Lines            []*OrderLine `json:"lines,omitempty"`
	PaymentID        string       `json:"paymentId,omitempty"`
	OrderID          string       `json:"orderId,omitempty"`
	CreatedAt        *time.Time   `json:"createdAt,omitempty"`
	TestMode         bool         `json:"testmode,omitempty"`
	Links            RefundLinks  `json:"_links,omitempty"`
}
```

Refund describe a refund for a certain payment.

#### type RefundLinks

```go
type RefundLinks struct {
	Self          *URL `json:"self,omitempty"`
	Payment       *URL `json:"payment,omitempty"`
	Settlement    *URL `json:"settlement,omitempty"`
	Order         *URL `json:"order,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

RefundLinks describes all the possible links to be returned with a Refund
object.

#### type RefundList

```go
type RefundList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Refunds []*Refund
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

RefundList describes how a list of refunds will be retrieved by Mollie.

#### type RefundOptions

```go
type RefundOptions struct {
	Embed EmbedValue `url:"embed,omitempty"`
}
```

RefundOptions describes refund endpoint valid query string parameters.

See: https://docs.mollie.com/reference/v2/refunds-api/get-refund.

#### type RefundStatus

```go
type RefundStatus string
```

RefundStatus describes the status of the refund.

```go
const (
	Queued     RefundStatus = "queued"
	Pending    RefundStatus = "pending"
	Processing RefundStatus = "processing"
	Refunded   RefundStatus = "refunded"
	Failed     RefundStatus = "failed"
)
```

Valid refund status.

#### type RefundsService

```go
type RefundsService service
```

RefundsService instance operates over refund resources.

#### func (\*RefundsService) Cancel

```go
func (rs *RefundsService) Cancel(paymentID, refundID string, options *RefundOptions) (err error)
```

Cancel try to cancel the refund request. The refund can only be canceled while
the refund’s status is either queued or pending. See
https://docs.mollie.com/reference/v2/refunds-api/cancel-refund

#### func (\*RefundsService) Create

```go
func (rs *RefundsService) Create(paymentID string, re Refund, options *RefundOptions) (rf Refund, err error)
```

Create a refund payment request.

See https://docs.mollie.com/reference/v2/refunds-api/create-refund.

#### func (\*RefundsService) Get

```go
func (rs *RefundsService) Get(paymentID, refundID string, options *RefundOptions) (refund Refund, err error)
```

Get retrieve a single refund by its ID.

If you do not know the original payment’s ID, you can use the List payment
refunds endpoint.

#### func (\*RefundsService) ListRefund

```go
func (rs *RefundsService) ListRefund(options *ListRefundOptions) (rl *RefundList, err error)
```

ListRefund calls the top level https://api.mollie.com/v2/refunds.

See https://docs.mollie.com/reference/v2/refunds-api/list-refunds.

#### func (\*RefundsService) ListRefundPayment

```go
func (rs *RefundsService) ListRefundPayment(paymentID string, options *ListRefundOptions) (rl *RefundList, err error)
```

ListRefundPayment calls the payment-specific
https://api.mollie.com/v2/payments/*paymentId*/refunds. Only refunds for that
specific payment are returned. See
https://docs.mollie.com/reference/v2/refunds-api/list-refunds

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

#### type Settlement

```go
type Settlement struct {
	ID        string           `json:"id,omitempty"`
	Resource  string           `json:"resource,omitempty"`
	Reference string           `json:"reference,omitempty"`
	CreatedAt *time.Time       `json:"createdAt,omitempty"`
	SettledAt *time.Time       `json:"settledAt,omitempty"`
	Status    SettlementStatus `json:"status,omitempty"`
	Amount    *Amount          `json:"amount,omitempty"`
	Periods   SettlementObject `json:"periods,omitempty"`
	InvoiceID string           `json:"invoiceId,omitempty"`
	Links     SettlementLinks  `json:"_links,omitempty"`
}
```

Settlement contains successful payments, together with refunds, captures and
chargebacks into settlements

#### type SettlementCosts

```go
type SettlementCosts struct {
	Description string        `json:"description,omitempty"`
	AmountNet   *Amount       `json:"amountNet,omitempty"`
	AmountVAT   *Amount       `json:"amountVat,omitempty"`
	AmountGross *Amount       `json:"amountGross,omitempty"`
	Count       int           `json:"count,omitempty"`
	Rate        *Rate         `json:"rate,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
}
```

SettlementCosts contains information about costs related to a settlement

#### type SettlementLinks

```go
type SettlementLinks struct {
	Self          *URL `json:"self,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Refunds       *URL `json:"refunds,omitempty"`
	Chargebacks   *URL `json:"chargebacks,omitempty"`
	Captures      *URL `json:"captures,omitempty"`
	Invoice       *URL `json:"invoice,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

SettlementLinks is an object with several URL objects relevant to the settlement

#### type SettlementObject

```go
type SettlementObject map[string]map[string]SettlementPeriod
```

SettlementObject nests as describes for settlement periods

#### type SettlementPeriod

```go
type SettlementPeriod struct {
	Revenue []*SettlementRevenue `json:"revenue,omitempty"`
	Costs   []*SettlementCosts   `json:"costs,omitempty"`
}
```

SettlementPeriod describe the settlement by month in full detail

#### type SettlementRevenue

```go
type SettlementRevenue struct {
	Description string        `json:"description,omitempty"`
	AmountNet   *Amount       `json:"amountNet,omitempty"`
	AmountVAT   *Amount       `json:"amountVat,omitempty"`
	AmountGross *Amount       `json:"amountGross,omitempty"`
	Count       int           `json:"count,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
}
```

SettlementRevenue objects contain the total revenue for each payment method
during this period

#### type SettlementStatus

```go
type SettlementStatus string
```

SettlementStatus describes the status of the settlement

```go
const (
	SettlementStatusOpen    SettlementStatus = "open"
	SettlementStatusPending SettlementStatus = "pending"
	SettlementStatusPaidOut SettlementStatus = "paidout"
	SettlementStatusFailed  SettlementStatus = "failed"
)
```

Available settlement statuses

#### type SettlementsList

```go
type SettlementsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Settlements []*Settlement
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

SettlementsList describes a list of settlements.

#### type SettlementsListOptions

```go
type SettlementsListOptions struct {
	From  *ShortDate `url:"from,omitempty"`
	Limit int        `url:"limit,omitempty"`
	Embed EmbedValue `url:"embed,omitempty"`
}
```

SettlementsListOptions contains query parameters for settlement lists.

#### type SettlementsService

```go
type SettlementsService service
```

SettlementsService operates over settlements resource

#### func (\*SettlementsService) Get

```go
func (ss *SettlementsService) Get(id string) (s *Settlement, err error)
```

Get returns a settlement by its id or the bank reference id

See: https://docs.mollie.com/reference/v2/settlements-api/get-settlement

#### func (\*SettlementsService) GetCaptures

```go
func (ss *SettlementsService) GetCaptures(id string, slo *SettlementsListOptions) (cl *CapturesList, err error)
```

GetCaptures retrieves all captures included in a settlement.

See:
https://docs.mollie.com/reference/v2/settlements-api/list-settlement-captures

#### func (\*SettlementsService) GetChargebacks

```go
func (ss *SettlementsService) GetChargebacks(id string, slo *SettlementsListOptions) (cl *ChargebackList, err error)
```

GetChargebacks retrieves all chargebacks included in a settlement.

See:
https://docs.mollie.com/reference/v2/settlements-api/list-settlement-chargebacks

#### func (\*SettlementsService) GetPayments

```go
func (ss *SettlementsService) GetPayments(id string, slo *SettlementsListOptions) (pl *PaymentList, err error)
```

GetPayments retrieves all payments included in a settlement.

See:
https://docs.mollie.com/reference/v2/settlements-api/list-settlement-payments

#### func (\*SettlementsService) GetRefunds

```go
func (ss *SettlementsService) GetRefunds(id string, slo *SettlementsListOptions) (rl *RefundList, err error)
```

GetRefunds retrieves all refunds included in a settlement.

See:
https://docs.mollie.com/reference/v2/settlements-api/list-settlement-refunds

#### func (\*SettlementsService) List

```go
func (ss *SettlementsService) List(slo *SettlementsListOptions) (sl *SettlementsList, err error)
```

List retrieves all settlements, ordered from new to old

See: https://docs.mollie.com/reference/v2/settlements-api/list-settlements

#### func (\*SettlementsService) Next

```go
func (ss *SettlementsService) Next() (s *Settlement, err error)
```

Next retrieves the details of the current settlement that has not yet been paid
out.

See: https://docs.mollie.com/reference/v2/settlements-api/get-next-settlement

#### func (\*SettlementsService) Open

```go
func (ss *SettlementsService) Open() (s *Settlement, err error)
```

Open retrieves the details of the open balance of the organization. This will
return a settlement object representing your organization’s balance.

See: https://docs.mollie.com/reference/v2/settlements-api/get-open-settlement

#### type Shipment

```go
type Shipment struct {
	Resource  string            `json:"resource,omitempty"`
	ID        string            `json:"id,omitempty"`
	OrderID   string            `json:"orderId,omitempty"`
	TestMode  bool              `json:"testmode,omitempty"`
	CreatedAt *time.Time        `json:"createdAt,omitempty"`
	Tracking  *ShipmentTracking `json:"tracking,omitempty"`
	Lines     []*OrderLine      `json:"lines,omitempty"`
	Links     ShipmentLinks     `json:"_links,omitempty"`
}
```

Shipment contains information about a user service/product delivery and is used
in the figurative sense here. It can also mean that a service was provided or
digital content was delivered.

#### type ShipmentLinks

```go
type ShipmentLinks struct {
	Self          *URL `json:"self,omitempty"`
	Order         *URL `json:"order,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

ShipmentLinks contains URL objects with shipment relevant information for the
user.

#### type ShipmentTracking

```go
type ShipmentTracking struct {
	Carrier string `json:"carrier,omitempty"`
	Code    string `json:"code,omitempty"`
	URL     string `json:"url,omitempty"`
}
```

ShipmentTracking contains shipment tracking details.

#### type ShipmentsList

```go
type ShipmentsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Shipments []Shipment
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

ShipmentsList describes how a list of payments will be retrieved by Mollie.

#### type ShipmentsService

```go
type ShipmentsService service
```

ShipmentsService operates on shipments endpoints.

#### func (\*ShipmentsService) Create

```go
func (ss *ShipmentsService) Create(oID string, cs CreateShipmentRequest) (s *Shipment, err error)
```

Create can be used to ship order lines.

See: https://docs.mollie.com/reference/v2/shipments-api/create-shipment

#### func (\*ShipmentsService) Get

```go
func (ss *ShipmentsService) Get(oID string, sID string) (s *Shipment, err error)
```

Get retrieves a single shipment and the order lines shipped by a shipment’s ID.

See: https://docs.mollie.com/reference/v2/shipments-api/get-shipment#

#### func (\*ShipmentsService) List

```go
func (ss *ShipmentsService) List(oID string) (sl *ShipmentsList, err error)
```

List retrieves all shipments for an order.

See: https://docs.mollie.com/reference/v2/shipments-api/list-shipments

#### func (\*ShipmentsService) Update

```go
func (ss *ShipmentsService) Update(oID string, sID string, st ShipmentTracking) (s *Shipment, err error)
```

Update can be used to update the tracking information of a shipment

See: https://docs.mollie.com/reference/v2/shipments-api/update-shipment

#### type ShortDate

```go
type ShortDate struct {
	time.Time
}
```

ShortDate is a string representing a date in YYYY-MM-DD format.

#### func (\*ShortDate) MarshalJSON

```go
func (d *ShortDate) MarshalJSON() ([]byte, error)
```

MarshalJSON overrides the default marshal action for the Date struct. Returns
date as YYYY-MM-DD formatted string.

#### func (\*ShortDate) UnmarshalJSON

```go
func (d *ShortDate) UnmarshalJSON(b []byte) error
```

UnmarshalJSON overrides the default unmarshal action for the Date struct, as we
need links to be pointers to the time.Time struct.

#### type Subscription

```go
type Subscription struct {
	Resource        string                 `json:"resource,omitempty"`
	ID              string                 `json:"id,omitempty"`
	MandateID       string                 `json:"mandateId,omitempty"`
	Mode            Mode                   `json:"mode,omitempty"`
	CreatedAT       *time.Time             `json:"createdAt,omitempty"`
	Status          SubscriptionStatus     `json:"status,omitempty"`
	Amount          *Amount                `json:"amount,omitempty"`
	Times           int                    `json:"times,omitempty"`
	TimesRemaining  int                    `json:"timesRemaining,omitempty"`
	Interval        string                 `json:"interval,omitempty"`
	StartDate       *ShortDate             `json:"startDate,omitempty"`
	NextPaymentDate *ShortDate             `json:"nextPaymentDate,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Method          PaymentMethod          `json:"method,omitempty"`
	CanceledAt      *time.Time             `json:"canceledAt,omitempty"`
	WebhookURL      string                 `json:"webhookUrl,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	ApplicationFee  *ApplicationFee        `json:"applicationFee,omitempty"`
	TestMode        bool                   `json:"testmode,omitempty"`
	Links           SubscriptionLinks      `json:"_links,omitempty"`
}
```

Subscription contains information about a customer subscription

#### type SubscriptionLinks

```go
type SubscriptionLinks struct {
	Self          *URL `json:"self,omitempty"`
	Customer      *URL `json:"customer,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
```

SubscriptionLinks contains several URL objects relevant to the subscription

#### type SubscriptionList

```go
type SubscriptionList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Subscriptions []*Subscription
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}
```

SubscriptionList describes the response for subscription list endpoints

#### type SubscriptionListOptions

```go
type SubscriptionListOptions struct {
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}
```

SubscriptionListOptions holds query string parameters valid for subscription
lists

#### type SubscriptionStatus

```go
type SubscriptionStatus string
```

SubscriptionStatus contains references to valid subscription statuses

```go
const (
	SubscriptionStatusPending   SubscriptionStatus = "pending"
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusCanceled  SubscriptionStatus = "canceled"
	SubscriptionStatusSuspended SubscriptionStatus = "suspended"
	SubscriptionStatusCompleted SubscriptionStatus = "completed"
)
```

Available subscription statuses

#### type SubscriptionsService

```go
type SubscriptionsService service
```

SubscriptionsService operates over subscriptions resource

#### func (\*SubscriptionsService) All

```go
func (ss *SubscriptionsService) All(options *SubscriptionListOptions) (sl *SubscriptionList, err error)
```

All retrieves all subscriptions, ordered from newest to oldest. By using an API
key all the subscriptions created with the current website profile will be
returned. In the case of an OAuth Access Token relies the website profile on the
profileId field

See:
https://docs.mollie.com/reference/v2/subscriptions-api/list-all-subscriptions

#### func (\*SubscriptionsService) Create

```go
func (ss *SubscriptionsService) Create(cID string, sc *Subscription) (s *Subscription, err error)
```

Create stores a new subscription for a given customer

See: https://docs.mollie.com/reference/v2/subscriptions-api/create-subscription

#### func (\*SubscriptionsService) Delete

```go
func (ss *SubscriptionsService) Delete(cID, sID string) (s *Subscription, err error)
```

Delete cancels a subscription

See: https://docs.mollie.com/reference/v2/subscriptions-api/cancel-subscription

#### func (\*SubscriptionsService) Get

```go
func (ss *SubscriptionsService) Get(cID, sID string) (s *Subscription, err error)
```

Get retrieves a customer's subscription

See: https://docs.mollie.com/reference/v2/subscriptions-api/get-subscription

#### func (\*SubscriptionsService) GetPayments

```go
func (ss *SubscriptionsService) GetPayments(cID, sID string, options *SubscriptionListOptions) (sl *PaymentList, err error)
```

GetPayments retrieves all payments of a specific subscriptions of a customer

See:
https://docs.mollie.com/reference/v2/subscriptions-api/list-subscriptions-payments

#### func (\*SubscriptionsService) List

```go
func (ss *SubscriptionsService) List(cID string, options *SubscriptionListOptions) (sl *SubscriptionList, err error)
```

List retrieves all subscriptions of a customer

See: https://docs.mollie.com/reference/v2/subscriptions-api/list-subscriptions

#### func (\*SubscriptionsService) Update

```go
func (ss *SubscriptionsService) Update(cID, sID string, sc *Subscription) (s *Subscription, err error)
```

Update changes fields on a subscription object

See: https://docs.mollie.com/reference/v2/subscriptions-api/update-subscription

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
