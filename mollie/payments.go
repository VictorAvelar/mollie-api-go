package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// PaymentMethod is a payment method supported by Mollie.
type PaymentMethod string

// Supported payment methods
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

// Valid sequence types
const (
	OneOffSequence    SequenceType = "oneoff"
	FirstSequence     SequenceType = "first"
	RecurringSequence SequenceType = "recurring"
)

// Payment describes a transaction between a customer and a merchant
type Payment struct {
	Resource                        string          `json:"resource,omitempty"`
	ID                              string          `json:"id,omitempty"`
	Mode                            Mode            `json:"mode,omitempty"`
	CreatedAt                       *time.Time      `json:"createdAt,omitempty"`
	Status                          string          `json:"status,omitempty"`
	IsCancellable                   bool            `json:"isCancellable,omitempty"`
	AuthorizedAt                    *time.Time      `json:"authorizedAt,omitempty"`
	PaidAt                          *time.Time      `json:"paidAt,omitempty"`
	CanceledAt                      *time.Time      `json:"canceledAt,omitempty"`
	ExpiresAt                       *time.Time      `json:"expiresAt,omitempty"`
	ExpiredAt                       *time.Time      `json:"expiredAt,omitempty"`
	FailedAt                        *time.Time      `json:"failedAt,omitempty"`
	Amount                          *Amount         `json:"amount,omitempty"`
	AmountRefunded                  *Amount         `json:"amountRefunded,omitempty"`
	AmountRemaining                 *Amount         `json:"amountRemaining,omitempty"`
	AmountCaptured                  *Amount         `json:"amountCaptured,omitempty"`
	Description                     string          `json:"description,omitempty"`
	RedirectURL                     string          `json:"redirectUrl,omitempty"`
	WebhookURL                      string          `json:"webhookUrl,omitempty"`
	Method                          PaymentMethod   `json:"method,omitempty"`
	Metadata                        interface{}     `json:"metadata,omitempty"`
	Locale                          Locale          `json:"locale,omitempty"`
	CountryCode                     string          `json:"countryCode,omitempty"`
	ProfileID                       string          `json:"profileId,omitempty"`
	SettlementAmount                *Amount         `json:"settlementAmount,omitempty"`
	SettlementID                    string          `json:"settlementId,omitempty"`
	CustomerID                      string          `json:"customerId,omitempty"`
	SequenceType                    SequenceType    `json:"sequenceType,omitempty"`
	MandateID                       string          `json:"mandateId,omitempty"`
	OrderID                         string          `json:"orderId,omitempty"`
	ApplicationFee                  *ApplicationFee `json:"applicationFee,omitempty"`
	Links                           PaymentLinks    `json:"_links,omitempty"`
	Details                         *PaymentDetails `json:"details,omitempty"`
	RestrictPaymentMethodsToCountry Locale          `json:"restrictPaymentMethodsToCountry,omitempty"`
	SubscriptionID                  string          `json:"subscriptionId,omitempty"`
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
	Include   string `url:"include,omitempty"`
	Embed     string `url:"embed,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

// PaymentsService instance operates over payment resources
type PaymentsService service

// Get retrieves a single payment object by its payment token.
func (ps *PaymentsService) Get(id string, options *PaymentOptions) (p Payment, err error) {
	u := fmt.Sprintf("v2/payments/%s", id)
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}
	req, err := ps.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
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
func (ps *PaymentsService) Create(p Payment) (np Payment, err error) {
	u := "v2/payments"
	req, err := ps.client.NewAPIRequest(http.MethodPost, u, p)
	if err != nil {
		return
	}

	res, err := ps.client.Do(req)
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
func (ps *PaymentsService) Cancel(id string) (p Payment, err error) {
	u := fmt.Sprintf("v2/payments/%s", id)
	req, err := ps.client.NewAPIRequest(http.MethodDelete, u, nil)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
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
func (ps *PaymentsService) Update(id string, up Payment) (p Payment, err error) {
	u := fmt.Sprintf("v2/payments/%s", id)
	req, err := ps.client.NewAPIRequest(http.MethodPatch, u, up)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
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
func (ps *PaymentsService) List(options *ListPaymentOptions) (pl PaymentList, err error) {
	u := "v2/payments"
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}
	req, err := ps.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}
