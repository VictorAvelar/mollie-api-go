package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// PaymentMethod is a payment method supported by Mollie.
type PaymentMethod string

// Supported payment methods.
const (
	Bancontact     PaymentMethod = "bancontact"
	BankTransfer   PaymentMethod = "banktransfer"
	Belfius        PaymentMethod = "belfius"
	CBC            PaymentMethod = "cbc"
	CreditCard     PaymentMethod = "creditcard"
	DirectDebit    PaymentMethod = "directdebit"
	EPS            PaymentMethod = "eps"
	GiftCard       PaymentMethod = "giftcard"
	GiroPay        PaymentMethod = "giropay"
	IDeal          PaymentMethod = "ideal"
	KBC            PaymentMethod = "kbc"
	KlarnaPayLater PaymentMethod = "klarnapaylater"
	KlarnaSliceIt  PaymentMethod = "klarnasliceit"
	MyBank         PaymentMethod = "mybank"
	PayPal         PaymentMethod = "paypal"
	PaySafeCard    PaymentMethod = "paysafecard"
	PRZelewy24     PaymentMethod = "przelewy24"
	Sofort         PaymentMethod = "sofort"
	Voucher        PaymentMethod = "voucher"
)

// SequenceType indicates which type of payment this is in a recurring sequence.
type SequenceType string

// Valid sequence types.
const (
	OneOffSequence    SequenceType = "oneoff"
	FirstSequence     SequenceType = "first"
	RecurringSequence SequenceType = "recurring"
)

// PaymentRouting describes the routing of a payment.
type PaymentRouting struct {
	Destination PaymentDestination `json:"destination,omitempty"`
	Amount      *Amount            `json:"amount,omitempty"`
	ReleaseDate *ShortDate         `json:"releaseDate,omitempty"`
}

// PaymentDestination describes the destination of a payment.
type PaymentDestination struct {
	Kind           string `json:"type,omitempty"`
	OrganizationID string `json:"organizationId,omitempty"`
}

// CreatePayment describes the payload to be sent to the Mollie API when
// creating or updating a new payment.
//
// Some fields are only valid for specific payment methods, and are
// documented in the Mollie API reference.
//
// See: https://docs.mollie.com/reference/create-payment#payment-method-specific-parameters
type CreatePayment struct {
	Description                     string          `json:"description,omitempty"`
	RedirectURL                     string          `json:"redirectUrl,omitempty"`
	CancelURL                       string          `json:"cancelUrl,omitempty"`
	WebhookURL                      string          `json:"webhookUrl,omitempty"`
	RestrictPaymentMethodsToCountry string          `json:"restrictPaymentMethodsToCountry,omitempty"`
	Amount                          *Amount         `json:"amount,omitempty"`
	Locale                          Locale          `json:"locale,omitempty"`
	Method                          []PaymentMethod `json:"method,omitempty"`
	Metadata                        any             `json:"metadata,omitempty"`

	// Beta fields
	Lines []PaymentLines `json:"lines,omitempty"`

	// PaymentMethods specific fields
	DigitalGoods         bool       `json:"digitalGoods,omitempty"`
	ApplePayPaymentToken string     `json:"applePayPaymentToken,omitempty"`
	BillingEmail         string     `json:"billingEmail,omitempty"`
	CardToken            string     `json:"cardToken,omitempty"`
	VoucherNumber        string     `json:"voucherNumber,omitempty"`
	VoucherPin           string     `json:"voucherPin,omitempty"`
	Issuer               string     `json:"issuer,omitempty"`
	ExtraMerchantData    string     `json:"extraMerchantData,omitempty"`
	SessionID            string     `json:"sessionId,omitempty"`
	CustomerReference    string     `json:"customerReference,omitempty"`
	TerminalID           string     `json:"terminalId,omitempty"`
	ConsumerName         string     `json:"consumerName,omitempty"`
	ConsumerAccount      string     `json:"consumerAccount,omitempty"`
	DueDate              *ShortDate `json:"dueDate,omitempty"`
	ShippingAddress      *Address   `json:"shippingAddress,omitempty"`
	BillingAddress       *Address   `json:"billingAddress,omitempty"`
	Company              *Company   `json:"company,omitempty"`

	// Other case specific fields
	CreateRecurrentPaymentFields
	CreatePreAuthorizedPaymentFields
	CreatePaymentAccessTokenFields
	CreateMollieConnectPaymentFields
}

// CreateRecurrentPaymentFields describes the fields to be sent to the Mollie API when
// creating a new recurrent payment.
//
// See: https://docs.mollie.com/reference/create-payment#parameters-for-recurring-payments
type CreateRecurrentPaymentFields struct {
	CustomerID   string       `json:"customerId,omitempty"`
	MandateID    string       `json:"mandateId,omitempty"`
	SequenceType SequenceType `json:"sequenceType,omitempty"`
}

// CreatePreAuthorizedPaymentFields describes the fields to be sent to the Mollie API when
// creating a new pre-authorized payment.
//
// See: https://docs.mollie.com/reference/create-payment#parameters-for-pre-authorized-payments
type CreatePreAuthorizedPaymentFields struct {
	CaptureDelay string      `json:"captureDelay,omitempty"`
	CaptureMode  CaptureMode `json:"captureMode,omitempty"`
}

// CreatePaymentAccessTokenFields describes the fields to be sent to the Mollie API when
// creating a new payment using an access token.
//
// See: https://docs.mollie.com/reference/create-payment#access-token-parameters
type CreatePaymentAccessTokenFields struct {
	ProfileID string `json:"profileId,omitempty"`
	Testmode  bool   `json:"testmode,omitempty"`
}

// CreateMollieConnectPaymentFields describes the fields to be sent to the Mollie API when
// creating a new payment using Mollie Connect.
//
// See: https://docs.mollie.com/reference/create-payment#mollie-connect-parameters
type CreateMollieConnectPaymentFields struct {
	ApplicationFee *ApplicationFee   `json:"applicationFee,omitempty"`
	Routing        []*PaymentRouting `json:"routing,omitempty"`
}

// PaymentLineType reflects the type of product bought.
type PaymentLineType string

// Supported payment line types.
const (
	PhysicalProductLine PaymentLineType = "physical"
	DiscountProductLine PaymentLineType = "discount"
	DigitalProductLine  PaymentLineType = "digital"
	ShippingFeeLine     PaymentLineType = "shipping_fee"
	StoreCreditLine     PaymentLineType = "store_credit"
	GiftCardLine        PaymentLineType = "gift_card"
	SurchargeLine       PaymentLineType = "surcharge"
	Tip                 PaymentLineType = "tip"
)

// PaymentLines describes the payment lines to be sent to the Mollie API when
// creating a new payment.
type PaymentLines struct {
	Quantity       int             `json:"quantity,omitempty"`
	Description    string          `json:"description,omitempty"`
	QuantityUnit   string          `json:"quantityUnit,omitempty"`
	SKU            string          `json:"sku,omitempty"`
	ImageURL       string          `json:"imageUrl,omitempty"`
	ProductURL     string          `json:"productUrl,omitempty"`
	VATRate        string          `json:"vatRate,omitempty"`
	UnitPrice      *Amount         `json:"unitPrice,omitempty"`
	DiscountAmount *Amount         `json:"discountAmount,omitempty"`
	TotalAmount    *Amount         `json:"totalAmount,omitempty"`
	VATAmount      *Amount         `json:"vatAmount,omitempty"`
	Type           PaymentLineType `json:"type,omitempty"`
}

// UpdatePayment describes the payload to be sent to the Mollie API when
// updating a payment.
//
// See: https://docs.mollie.com/reference/update-payment
// See: https://docs.mollie.com/reference/update-payment#payment-method-specific-parameters
type UpdatePayment struct {
	Description                     string        `json:"description,omitempty"`
	RedirectURL                     string        `json:"redirectUrl,omitempty"`
	CancelURL                       string        `json:"cancelUrl,omitempty"`
	WebhookURL                      string        `json:"webhookUrl,omitempty"`
	Metadata                        any           `json:"metadata,omitempty"`
	Method                          PaymentMethod `json:"method,omitempty"`
	Locale                          Locale        `json:"locale,omitempty"`
	RestrictPaymentMethodsToCountry string        `json:"restrictPaymentMethodsToCountry,omitempty"`

	// PaymentMethods specific fields
	BillingEmail string     `json:"billingEmail,omitempty"`
	DueDate      *ShortDate `json:"dueDate,omitempty"`
	Issuer       string     `json:"issuer,omitempty"`
}

// Payment describes a transaction between a customer and a merchant.
type Payment struct {
	Resource                        string        `json:"resource,omitempty"`
	ID                              string        `json:"id,omitempty"`
	Status                          string        `json:"status,omitempty"`
	Description                     string        `json:"description,omitempty"`
	CancelURL                       string        `json:"cancelUrl,omitempty"`
	WebhookURL                      string        `json:"webhookUrl,omitempty"`
	CountryCode                     string        `json:"countryCode,omitempty"`
	RestrictPaymentMethodsToCountry string        `json:"restrictPaymentMethodsToCountry,omitempty"`
	ProfileID                       string        `json:"profileId,omitempty"`
	SettlementID                    string        `json:"settlementId,omitempty"`
	OrderID                         string        `json:"orderId,omitempty"`
	IsCancelable                    bool          `json:"isCancelable,omitempty"`
	Mode                            Mode          `json:"mode,omitempty"`
	Locale                          Locale        `json:"locale,omitempty"`
	Method                          PaymentMethod `json:"method,omitempty"`
	Metadata                        any           `json:"metadata,omitempty"`
	Links                           PaymentLinks  `json:"_links,omitempty"`
	CreatedAt                       *time.Time    `json:"createdAt,omitempty"`
	AuthorizedAt                    *time.Time    `json:"authorizedAt,omitempty"`
	PaidAt                          *time.Time    `json:"paidAt,omitempty"`
	CanceledAt                      *time.Time    `json:"canceledAt,omitempty"`
	ExpiresAt                       *time.Time    `json:"expiresAt,omitempty"`
	ExpiredAt                       *time.Time    `json:"expiredAt,omitempty"`
	FailedAt                        *time.Time    `json:"failedAt,omitempty"`
	Amount                          *Amount       `json:"amount,omitempty"`
	AmountRefunded                  *Amount       `json:"amountRefunded,omitempty"`
	AmountRemaining                 *Amount       `json:"amountRemaining,omitempty"`
	AmountCaptured                  *Amount       `json:"amountCaptured,omitempty"`
	AmountChargedBack               *Amount       `json:"amountChargedback,omitempty"`
	SettlementAmount                *Amount       `json:"settlementAmount,omitempty"`

	// Beta fields
	Lines []PaymentLines `json:"lines,omitempty"`

	// PaymentMethods specific fields
	Details PaymentDetails `json:"details,omitempty"`

	// Other case specific fields
	RecurrentPaymentFields
	PreAuthorizedPaymentFields
	MollieConnectPaymentFields
	AccessTokenPaymentFields
}

// RecurrentPaymentFields describes the fields specific to recurrent payments.
type RecurrentPaymentFields struct {
	SequenceType   SequenceType `json:"sequenceType,omitempty"`
	CustomerID     string       `json:"customerId,omitempty"`
	MandateID      string       `json:"mandateId,omitempty"`
	SubscriptionID string       `json:"subscriptionId,omitempty"`
}

// PreAuthorizedPaymentFields describes the fields specific to pre-authorized payments.
type PreAuthorizedPaymentFields struct {
	CaptureDelay  string      `json:"captureDelay,omitempty"`
	CaptureMode   CaptureMode `json:"captureMode,omitempty"`
	CaptureBefore *time.Time  `json:"captureBefore,omitempty"`
}

// MollieConnectPaymentFields describes the fields specific to Mollie Connect payments.
type MollieConnectPaymentFields struct {
	ApplicationFee *ApplicationFee `json:"applicationFee,omitempty"`
}

// AccessTokenPaymentFields describes the fields specific to payments created using an access token.
type AccessTokenPaymentFields struct {
	Testmode bool `json:"testmode,omitempty"`
}

// PaymentLinks describes all the possible links to be returned with
// a payment object.
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
	MobileAppCheckout  *URL `json:"mobileAppCheckout,omitempty"`
	Terminal           *URL `json:"terminal,omitempty"`
}

// PaymentOptions describes payments endpoint valid query string parameters.
//
// See: https://docs.mollie.com/reference/get-payment
type PaymentOptions struct {
	Include []IncludeValue `url:"include,omitempty"`
	Embed   []EmbedValue   `url:"embed,omitempty"`
}

// ListPaymentsOptions describes list payments endpoint valid query string parameters.
type ListPaymentsOptions struct {
	Limit     int            `url:"limit,omitempty"`
	Include   []IncludeValue `url:"include,omitempty"`
	Embed     []EmbedValue   `url:"embed,omitempty"`
	ProfileID string         `url:"profileId,omitempty"`
	From      string         `url:"from,omitempty"`
}

// PaymentsService instance operates over payment resources.
type PaymentsService service

// Get retrieves a single payment object by its payment token.
func (ps *PaymentsService) Get(ctx context.Context, id string, opts *PaymentOptions) (
	res *Response,
	p *Payment,
	err error,
) {
	res, err = ps.client.get(ctx, fmt.Sprintf("v2/payments/%s", id), opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}

	return
}

// Create stores a new payment object attached to your Mollie account.
//
// See: https://docs.mollie.com/reference/create-payment#
func (ps *PaymentsService) Create(ctx context.Context, p CreatePayment, opts *PaymentOptions) (
	res *Response,
	np *Payment,
	err error,
) {
	if ps.client.HasAccessToken() && ps.client.config.testing {
		p.Testmode = true
	}

	res, err = ps.client.post(ctx, "v2/payments", p, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &np); err != nil {
		return
	}

	return
}

// Cancel removes a payment (if possible) from your Mollie account.
//
// See: https://docs.mollie.com/reference/cancel-payment
func (ps *PaymentsService) Cancel(ctx context.Context, id string) (res *Response, p *Payment, err error) {
	res, err = ps.client.delete(ctx, fmt.Sprintf("v2/payments/%s", id))
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}

	return
}

// Update can be used to update some details of a created payment.
//
// See: https://docs.mollie.com/reference/update-payment
func (ps *PaymentsService) Update(ctx context.Context, id string, up UpdatePayment) (
	res *Response,
	p *Payment,
	err error,
) {
	res, err = ps.client.patch(ctx, fmt.Sprintf("v2/payments/%s", id), up)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}

	return
}

// PaymentList describes how a list of payments will be retrieved by Mollie.
type PaymentList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Payments []*Payment
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// List retrieves a list of payments associated with your account/organization.
//
// See: https://docs.mollie.com/reference/list-payments
func (ps *PaymentsService) List(ctx context.Context, opts *ListPaymentsOptions) (
	res *Response,
	pl *PaymentList,
	err error,
) {
	res, err = ps.client.get(ctx, "v2/payments", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}

	return
}
