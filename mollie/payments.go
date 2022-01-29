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

// SequenceType indicates which type of payment this is in a recurring sequence.
type SequenceType string

// Valid sequence types.
const (
	OneOffSequence    SequenceType = "oneoff"
	FirstSequence     SequenceType = "first"
	RecurringSequence SequenceType = "recurring"
)

// Payment describes a transaction between a customer and a merchant.
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
}

// PaymentOptions describes payments endpoint valid query string parameters.
//
// See: https://docs.mollie.com/reference/v2/payments-api/get-payment
type PaymentOptions struct {
	Include string `url:"include,omitempty"`
	Embed   string `url:"embed,omitempty"`
}

// ListPaymentOptions describes list payments endpoint valid query string parameters.
type ListPaymentOptions struct {
	Limit     int    `url:"limit,omitempty"`
	Include   string `url:"include,omitempty"`
	Embed     string `url:"embed,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitempty"`
}

// PaymentsService instance operates over payment resources.
type PaymentsService service

// Get retrieves a single payment object by its payment token.
func (ps *PaymentsService) Get(ctx context.Context, id string, opts *PaymentOptions) (res *Response, p *Payment, err error) {
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
// See: https://docs.mollie.com/reference/v2/payments-api/create-payment#
func (ps *PaymentsService) Create(ctx context.Context, p Payment, opts *PaymentOptions) (res *Response, np *Payment, err error) {
	if ps.client.HasAccessToken() && ps.client.config.testing {
		p.TestMode = true
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
// See: https://docs.mollie.com/reference/v2/payments-api/cancel-payment
func (ps *PaymentsService) Cancel(ctx context.Context, id string) (res *Response, p *Payment, err error) {
	res, err = ps.client.delete(ctx, fmt.Sprintf("v2/payments/%s", id), nil)
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
// See: https://docs.mollie.com/reference/v2/payments-api/update-payment#
func (ps *PaymentsService) Update(ctx context.Context, id string, up Payment) (res *Response, p *Payment, err error) {
	res, err = ps.client.patch(ctx, fmt.Sprintf("v2/payments/%s", id), up, nil)
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
		Payments []Payment
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// List retrieves a list of payments associated with your account/organization.
//
// See: https://docs.mollie.com/reference/v2/payments-api/list-payments
func (ps *PaymentsService) List(ctx context.Context, opts *ListPaymentOptions) (res *Response, pl *PaymentList, err error) {
	res, err = ps.client.get(ctx, "v2/payments", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}
