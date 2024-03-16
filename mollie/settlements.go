package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

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
	Count       int           `json:"count,omitempty"`
	Description string        `json:"description,omitempty"`
	AmountNet   *Amount       `json:"amountNet,omitempty"`
	AmountVAT   *Amount       `json:"amountVat,omitempty"`
	AmountGross *Amount       `json:"amountGross,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
}

// SettlementCosts contains information about costs related to a settlement.
type SettlementCosts struct {
	Count       int           `json:"count,omitempty"`
	Description string        `json:"description,omitempty"`
	InvoiceID   string        `json:"invoiceId,omitempty"`
	AmountNet   *Amount       `json:"amountNet,omitempty"`
	AmountVAT   *Amount       `json:"amountVat,omitempty"`
	AmountGross *Amount       `json:"amountGross,omitempty"`
	Rate        *Rate         `json:"rate,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
}

// SettlementPeriod describe the settlement by month in full detail.
type SettlementPeriod struct {
	InvoiceID        string               `json:"invoiceId,omitempty"`
	InvoiceReference string               `json:"invoiceReference,omitempty"`
	Revenue          []*SettlementRevenue `json:"revenue,omitempty"`
	Costs            []*SettlementCosts   `json:"costs,omitempty"`
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
	ID        string            `json:"id,omitempty"`
	Resource  string            `json:"resource,omitempty"`
	Reference string            `json:"reference,omitempty"`
	InvoiceID string            `json:"invoiceId,omitempty"`
	CreatedAt *time.Time        `json:"createdAt,omitempty"`
	SettledAt *time.Time        `json:"settledAt,omitempty"`
	Amount    *Amount           `json:"amount,omitempty"`
	Periods   *SettlementObject `json:"periods,omitempty"`
	Status    SettlementStatus  `json:"status,omitempty"`
	Links     SettlementLinks   `json:"_links,omitempty"`
}

// ListSettlementsOptions contains query parameters for settlement lists.
type ListSettlementsOptions struct {
	From  string       `url:"from,omitempty"`
	Limit int          `url:"limit,omitempty"`
	Embed []EmbedValue `url:"embed,omitempty"`
}

// SettlementsList describes a list of settlements.
type SettlementsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Settlements []*Settlement
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// SettlementsService operates over settlements resource.
type SettlementsService service

// Get returns a settlement by its id or the bank reference id
//
// See: https://docs.mollie.com/reference/v2/settlements-api/get-settlement
func (ss *SettlementsService) Get(ctx context.Context, settlement string) (res *Response, s *Settlement, err error) {
	return ss.get(ctx, settlement)
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
func (ss *SettlementsService) List(ctx context.Context, slo *ListSettlementsOptions) (
	res *Response,
	sl *SettlementsList,
	err error,
) {
	res, err = ss.list(ctx, "", "", slo)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sl); err != nil {
		return
	}

	return
}

// ListPayments retrieves all payments included in a settlement.
// This API is an alias of the List payments.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/list-settlement-payments
func (ss *SettlementsService) ListPayments(ctx context.Context, settlement string, options *ListPaymentsOptions) (
	res *Response,
	pl *PaymentList,
	err error,
) {
	res, err = ss.list(ctx, settlement, "payments", options)
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
func (ss *SettlementsService) GetRefunds(ctx context.Context, settlement string, slo *ListSettlementsOptions) (
	res *Response,
	rl *RefundsList,
	err error,
) {
	res, err = ss.list(ctx, settlement, "refunds", slo)
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
func (ss *SettlementsService) GetChargebacks(ctx context.Context, settlement string, slo *ListChargebacksOptions) (
	res *Response,
	cl *ChargebacksList,
	err error,
) {
	res, err = ss.list(ctx, settlement, "chargebacks", slo)
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
func (ss *SettlementsService) GetCaptures(ctx context.Context, settlement string, slo *ListSettlementsOptions) (
	res *Response,
	cl *CapturesList,
	err error,
) {
	res, err = ss.list(ctx, settlement, "captures", slo)
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

func (ss *SettlementsService) list(ctx context.Context, settlement string, category string, opts any) (
	res *Response,
	err error,
) {
	uri := "v2/settlements"

	if settlement != "" {
		uri = fmt.Sprintf("%s/%s", uri, settlement)

		if category != "" {
			uri = fmt.Sprintf("%s/%s", uri, category)
		}
	}

	res, err = ss.client.get(ctx, uri, opts)
	if err != nil {
		return
	}

	return
}
