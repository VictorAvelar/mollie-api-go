package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// SettlementsService operates over settlements resource.
type SettlementsService service

// SettlementStatus describes the status of the settlement.
type SettlementStatus string

// Available settlement statuses.
const (
	SettlementStatusOpen    SettlementStatus = "open"
	SettlementStatusPending SettlementStatus = "pending"
	SettlementStatusPaidOut SettlementStatus = "paidout"
	SettlementStatusFailed  SettlementStatus = "failed"
)

// SettlementRevenue objects contain the total revenue for each payment method during this period.
type SettlementRevenue struct {
	Description string        `json:"description,omitempty"`
	AmountNet   *Amount       `json:"amountNet,omitempty"`
	AmountVAT   *Amount       `json:"amountVat,omitempty"`
	AmountGross *Amount       `json:"amountGross,omitempty"`
	Count       int           `json:"count,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
}

// SettlementCosts contains information about costs related to a settlement.
type SettlementCosts struct {
	Description string        `json:"description,omitempty"`
	AmountNet   *Amount       `json:"amountNet,omitempty"`
	AmountVAT   *Amount       `json:"amountVat,omitempty"`
	AmountGross *Amount       `json:"amountGross,omitempty"`
	Count       int           `json:"count,omitempty"`
	Rate        *Rate         `json:"rate,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
}

// SettlementPeriod describe the settlement by month in full detail.
type SettlementPeriod struct {
	Revenue []*SettlementRevenue `json:"revenue,omitempty"`
	Costs   []*SettlementCosts   `json:"costs,omitempty"`
}

// SettlementObject nests as describes for settlement periods.
type SettlementObject map[string]map[string]SettlementPeriod

// SettlementLinks is an object with several URL objects relevant to the settlement.
type SettlementLinks struct {
	Self          *URL `json:"self,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Refunds       *URL `json:"refunds,omitempty"`
	Chargebacks   *URL `json:"chargebacks,omitempty"`
	Captures      *URL `json:"captures,omitempty"`
	Invoice       *URL `json:"invoice,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Settlement contains successful payments, together with refunds,
// captures and chargebacks into settlements.
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

// SettlementsListOptions contains query parameters for settlement lists.
type SettlementsListOptions struct {
	From  *ShortDate `url:"from,omitempty"`
	Limit int        `url:"limit,omitempty"`
	Embed EmbedValue `url:"embed,omitempty"`
}

// SettlementsList describes a list of settlements.
type SettlementsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Settlements []*Settlement
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// Get returns a settlement by its id or the bank reference id
//
// See: https://docs.mollie.com/reference/v2/settlements-api/get-settlement
func (ss *SettlementsService) Get(ctx context.Context, id string) (res *Response, s *Settlement, err error) {
	return ss.get(ctx, id)
}

// Next retrieves the details of the current settlement that has not yet been paid out.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/get-next-settlement
func (ss *SettlementsService) Next(ctx context.Context) (res *Response, s *Settlement, err error) {
	return ss.get(ctx, "next")
}

// Open retrieves the details of the open balance of the organization.
// This will return a settlement object representing your organization’s balance.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/get-open-settlement
func (ss *SettlementsService) Open(ctx context.Context) (res *Response, s *Settlement, err error) {
	return ss.get(ctx, "open")
}

// List retrieves all settlements, ordered from new to old
//
// See: https://docs.mollie.com/reference/v2/settlements-api/list-settlements
func (ss *SettlementsService) List(ctx context.Context, slo *SettlementsListOptions) (res *Response, sl *SettlementsList, err error) {
	res, err = ss.list(ctx, "", "", slo)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sl); err != nil {
		return
	}
	return
}

// GetPayments retrieves all payments included in a settlement.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/list-settlement-payments
func (ss *SettlementsService) GetPayments(ctx context.Context, id string, slo *SettlementsListOptions) (res *Response, pl *PaymentList, err error) {
	res, err = ss.list(ctx, id, "payments", slo)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}

// GetRefunds retrieves all refunds included in a settlement.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/list-settlement-refunds
func (ss *SettlementsService) GetRefunds(ctx context.Context, id string, slo *SettlementsListOptions) (res *Response, rl *RefundList, err error) {
	res, err = ss.list(ctx, id, "refunds", slo)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &rl); err != nil {
		return
	}
	return
}

// GetChargebacks retrieves all chargebacks included in a settlement.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/list-settlement-chargebacks
func (ss *SettlementsService) GetChargebacks(ctx context.Context, id string, slo *SettlementsListOptions) (res *Response, cl *ChargebacksList, err error) {
	res, err = ss.list(ctx, id, "chargebacks", slo)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}
	return
}

// GetCaptures retrieves all captures included in a settlement.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/list-settlement-captures
func (ss *SettlementsService) GetCaptures(ctx context.Context, id string, slo *SettlementsListOptions) (res *Response, cl *CapturesList, err error) {
	res, err = ss.list(ctx, id, "captures", slo)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}
	return
}

func (ss *SettlementsService) get(ctx context.Context, element string) (res *Response, s *Settlement, err error) {
	res, err = ss.client.get(ctx, fmt.Sprintf("v2/settlements/%s", element), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}

func (ss *SettlementsService) list(ctx context.Context, id string, category string, opts *SettlementsListOptions) (res *Response, err error) {
	u := "v2/settlements"

	if id != "" {
		u = fmt.Sprintf("%s/%s", u, id)

		if category != "" {
			u = fmt.Sprintf("%s/%s", u, category)
		}
	}

	res, err = ss.client.get(ctx, u, opts)
	if err != nil {
		return
	}
	return
}
