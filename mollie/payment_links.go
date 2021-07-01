package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// PaymentLink is a resource that can be shared with your customers
// and will redirect them to them the payment page where they can
// complete the payment.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link
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

// PaymentLinkLinks describes all the possible links returned with
// a payment link struct.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link
type PaymentLinkLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	PaymentLink   *URL `json:"paymentLink,omitempty"`
	Next          *URL `json:"next,omitempty"`
	Previous      *URL `json:"previous,omitempty"`
}

// PaymentLinkOptions represents query string parameters to modify
// the payment links requests.
type PaymentLinkOptions struct {
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitemtpy"`
	Limit     int    `url:"limit,omitempty"`
}

type PaymentLinksList struct {
	Count    int              `json:"count,omitempty"`
	Links    PaymentLinkLinks `json:"_links,omitempty"`
	Embedded struct {
		PaymentLinks []*PaymentLink `json:"payment_links,omitempty"`
	} `json:"_embedded,omitempty"`
}

// PaymentLinksService operates over the payment link resource.
type PaymentLinksService service

// Get retrieves a single payment link object by its id/token.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link
func (pls *PaymentLinksService) Get(id string) (pl *PaymentLink, err error) {
	req, err := pls.client.NewAPIRequest(http.MethodGet, fmt.Sprintf("v2/payment-links/%s", id), nil)
	if err != nil {
		return
	}

	res, err := pls.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}

// Create generates payment links that by default, unlike regular payments, do not expire.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/create-payment-link
func (pls *PaymentLinksService) Create(p PaymentLink, opts *PaymentLinkOptions) (np *PaymentLink, err error) {
	u := "v2/payment-links"
	if opts != nil {
		v, _ := query.Values(opts)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}
	req, err := pls.client.NewAPIRequest(http.MethodPost, u, p)
	if err != nil {
		return
	}

	res, err := pls.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &np); err != nil {
		return
	}
	return
}

// List retrieves all payments links created with the current website profile,
// ordered from newest to oldest.
//
// See: https://docs.mollie.com/reference/v2/payment-links-api/list-payment-links
func (pls *PaymentLinksService) List(opts *PaymentLinkOptions) (pl *PaymentLinksList, err error) {
	u := "v2/payment-links"
	if opts != nil {
		v, _ := query.Values(opts)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}
	req, err := pls.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := pls.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}
