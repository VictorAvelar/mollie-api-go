package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Chargeback describes a forced transaction reversal initiated by the cardholder's bank.
type Chargeback struct {
	Resource         string            `json:"resource,omitempty"`
	ID               string            `json:"id,omitempty"`
	PaymentID        string            `json:"paymentId,omitempty"`
	Amount           *Amount           `json:"amount,omitempty"`
	SettlementAmount *Amount           `json:"settlementAmount,omitempty"`
	Reason           *ChargebackReason `json:"reason,omitempty"`
	CreatedAt        *time.Time        `json:"createdAt,omitempty"`
	ReversedAt       *time.Time        `json:"reversedAt,omitempty"`
	Links            ChargebackLinks   `json:"_links,omitempty"`
	ChargebackAccessTokenFields
}

// ChargebackReason describes the reason for the chargeback as given by the bank.
type ChargebackReason struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

// ChargebackAccessTokenFields describes the fields to be used to create a chargeback access token.
type ChargebackAccessTokenFields struct {
	ProfileID string `json:"profileId,omitempty"`
	Testmode  bool   `json:"testmode,omitempty"`
}

// ChargebackLinks describes all the possible links to be returned with
// a chargeback object.
type ChargebackLinks struct {
	Self          *URL `json:"self,omitempty"`
	Payment       *URL `json:"payment,omitempty"`
	Settlement    *URL `json:"settlement,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// ChargebackOptions describes chargeback endpoint valid query string parameters.
type ChargebackOptions struct {
	Include []IncludeValue `url:"include,omitempty"`
	Embed   []EmbedValue   `url:"embed,omitempty"`
}

// ListChargebacksOptions describes list chargebacks endpoint valid query string parameters.
type ListChargebacksOptions struct {
	From      string         `url:"from,omitempty"`
	Limit     int            `url:"limit,omitempty"`
	Include   []IncludeValue `url:"include,omitempty"`
	Embed     []EmbedValue   `url:"embed,omitempty"`
	ProfileID string         `url:"profileId,omitempty"`
}

// ChargebacksList describes how a list of chargebacks will be retrieved by Mollie.
type ChargebacksList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Chargebacks []*Chargeback
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// ChargebacksService instance operates over chargeback resources.
type ChargebacksService service

// Get retrieves a single chargeback by its ID.
// Note the original payment’s ID is needed as well.
//
// See: https://docs.mollie.com/reference/get-chargeback
func (cs *ChargebacksService) Get(ctx context.Context, payment, chargeback string, opts *ChargebackOptions) (
	res *Response,
	p *Chargeback,
	err error,
) {
	u := fmt.Sprintf("v2/payments/%s/chargebacks/%s", payment, chargeback)

	res, err = cs.client.get(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}

	return
}

// List retrieves a list of chargebacks associated with your account/organization.
//
// See: https://docs.mollie.com/reference/list-chargebacks
func (cs *ChargebacksService) List(ctx context.Context, options *ListChargebacksOptions) (
	res *Response,
	cl *ChargebacksList,
	err error,
) {
	return cs.list(ctx, "v2/chargebacks", options)
}

// ListForPayment retrieves a list of chargebacks associated with a single payment.
//
// See: https://docs.mollie.com/reference/list-chargebacks
func (cs *ChargebacksService) ListForPayment(ctx context.Context, payment string, options *ListChargebacksOptions) (
	res *Response,
	cl *ChargebacksList,
	err error,
) {
	return cs.list(ctx, fmt.Sprintf("v2/payments/%s/chargebacks", payment), options)
}

// encapsulates the shared list methods logic.
func (cs *ChargebacksService) list(ctx context.Context, uri string, options interface{}) (
	res *Response,
	cl *ChargebacksList,
	err error,
) {
	res, err = cs.client.get(ctx, uri, options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}

	return
}
