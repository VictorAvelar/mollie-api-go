package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// SettlementsService operates over settlements resource
type SettlementsService service

// SettlementStatus describes the status of the settlement
type SettlementStatus string

// Available settlement statuses
const (
	SettlementStatusOpen    = "string"
	SettlementStatusPending = "pending"
	SettlementStatusPaidOut = "paidout"
	SettlementStatusFailed  = "failed"
)

// SettlementRevenue objects contain the total revenue for each payment method during this period
type SettlementRevenue struct {
	Description string        `json:"description,omitempty"`
	AmountNet   Amount        `json:"amountNet,omitempty"`
	AmountVAT   Amount        `json:"amountVat,omitempty"`
	AmountGross Amount        `json:"amountGross,omitempty"`
	Count       int           `json:"count,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
}

type SettlementCosts struct {
	Description string        `json:"description,omitempty"`
	AmountNet   Amount        `json:"amountNet,omitempty"`
	AmountVAT   Amount        `json:"amountVat,omitempty"`
	AmountGross Amount        `json:"amountGross,omitempty"`
	Count       int           `json:"count,omitempty"`
	Rate        Rate          `json:"rate,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
}

// SettlementPeriod describe the settlement by month in full detail
type SettlementPeriod struct {
	Revenue []SettlementRevenue `json:"revenue,omitempty"`
	Costs   []SettlementCosts   `json:"costs,omitempty"`
}

// SettlementObject nests as describes for settlement periods
type SettlementObject map[string]map[string]SettlementPeriod

// SettlementLinks is an object with several URL objects relevant to the settlement
type SettlementLinks struct {
	Self          URL `json:"self,omitempty"`
	Payments      URL `json:"payments,omitempty"`
	Refunds       URL `json:"refunds,omitempty"`
	Chargebacks   URL `json:"chargebacks,omitempty"`
	Captures      URL `json:"captures,omitempty"`
	Invoice       URL `json:"invoice,omitempty"`
	Documentation URL `json:"documentation,omitempty"`
}

// Settlement contains successful payments, together with refunds,
// captures and chargebacks into settlements
type Settlement struct {
	ID        string           `json:"id,omitempty"`
	Resource  string           `json:"resource,omitempty"`
	Reference string           `json:"reference,omitempty"`
	CreatedAt *time.Time       `json:"createdAt,omitempty"`
	SettledAt *time.Time       `json:"settledAt,omitempty"`
	Status    SettlementStatus `json:"status,omitempty"`
	Amount    Amount           `json:"amount,omitempty"`
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
		Settlements []Settlement
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// Get returns a settlement by its id or the bank reference id
//
// See: https://docs.mollie.com/reference/v2/settlements-api/get-settlement
func (ss *SettlementsService) Get(id string) (s *Settlement, err error) {
	return ss.get(id)
}

// Next retrieves the details of the current settlement that has not yet been paid out.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/get-next-settlement
func (ss *SettlementsService) Next() (s *Settlement, err error) {
	return ss.get("next")
}

// Open retrieves the details of the open balance of the organization.
// This will return a settlement object representing your organization’s balance.
//
// See: https://docs.mollie.com/reference/v2/settlements-api/get-open-settlement
func (ss *SettlementsService) Open() (s *Settlement, err error) {
	return ss.get("open")
}

// List retrieves all settlements, ordered from new to old
//
// See: https://docs.mollie.com/reference/v2/settlements-api/list-settlements
func (ss *SettlementsService) List(slo *SettlementsListOptions) (sl *SettlementsList, err error) {
	res, err := ss.list("", "", slo)
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
func (ss *SettlementsService) GetPayments(id string, slo *SettlementsListOptions) (pl *PaymentList, err error) {
	res, err := ss.list(id, "payments", slo)
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
func (ss *SettlementsService) GetRefunds(id string, slo *SettlementsListOptions) (rl *RefundList, err error) {
	res, err := ss.list(id, "refunds", slo)
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
func (ss *SettlementsService) GetChargebacks(id string, slo *SettlementsListOptions) (cl *ChargebackList, err error) {
	res, err := ss.list(id, "chargebacks", slo)
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
func (ss *SettlementsService) GetCaptures(id string, slo *SettlementsListOptions) (cl *CapturesList, err error) {
	res, err := ss.list(id, "captures", slo)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}
	return
}

func (ss *SettlementsService) get(element string) (s *Settlement, err error) {
	u := fmt.Sprintf("v2/settlements/%s", element)
	req, err := ss.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}

func (ss *SettlementsService) list(id string, category string, slo *SettlementsListOptions) (res *Response, err error) {
	u := "v2/settlements"

	if id != "" {
		u = fmt.Sprintf("%s/%s", u, id)

		if category != "" {
			u = fmt.Sprintf("%s/%s", u, category)
		}
	}

	if slo != nil {
		v, _ := query.Values(slo)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	req, err := ss.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err = ss.client.Do(req)
	if err != nil {
		return
	}
	return
}
