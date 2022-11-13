package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// BalancesService allows you to retrieve real-time as well as historical
// information about your Mollie balance.
//
// Works with Organization access tokens and App access tokens.
//
// The API is in **BETA** so be careful and expect changes.
//
// See: https://docs.mollie.com/reference/v2/balances-api/overview
type BalancesService service

// BalanceStatus reflects whether a balance is operational or not.
type BalanceStatus string

// Possible values for type BalanceStatus.
const (
	BalanceActive   BalanceStatus = "active"
	BalanceInactive BalanceStatus = "inactive"
)

// TransferFrequency reflects the frequency at which the available amount
// on the balance will be settled to the configured transfer destination.
type TransferFrequency string

// Possible values for type TransferFrequency.
const (
	TransferDaily          TransferFrequency = "daily"
	TransferTwiceAWeek     TransferFrequency = "twice-a-week"
	TransferEveryMonday    TransferFrequency = "every-monday"
	TransferEveryTuesday   TransferFrequency = "every-tuesday"
	TransferEveryWednesday TransferFrequency = "every-wednesday"
	TransferEveryThursday  TransferFrequency = "every-thursday"
	TransferEveryFriday    TransferFrequency = "every-friday"
	TransferTwiceAMonth    TransferFrequency = "twice-a-month"
	TransferMonthly        TransferFrequency = "monthly"
	TransferNever          TransferFrequency = "never"
)

// TransferDestination where the available amount will be automatically
// transferred.
type TransferDestination struct {
	Type            string `json:"type,omitempty"`
	BankAccount     string `json:"bankAccount,omitempty"`
	BeneficiaryName string `json:"beneficiaryName,omitempty"`
}

// BalanceLinks holds URL objects relevant to the balance.
type BalanceLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Balance holds the payments processed with Mollie once fees have been deducted.
type Balance struct {
	ID                  string               `json:"id,omitempty"`
	Resource            string               `json:"resource,omitempty"`
	Currency            string               `json:"currency,omitempty"`
	TransferReference   string               `json:"transferReference,omitempty"`
	Status              BalanceStatus        `json:"status,omitempty"`
	TransferFrequency   TransferFrequency    `json:"transferFrequency,omitempty"`
	TransferThreshold   *Amount              `json:"transferThreshold,omitempty"`
	AvailableAmount     *Amount              `json:"availableAmount,omitempty"`
	PendingAmount       *Amount              `json:"pendingAmount,omitempty"`
	TransferDestination *TransferDestination `json:"transferDestination,omitempty"`
	CreatedAt           *time.Time           `json:"createdAt,omitempty"`
	Links               BalanceLinks         `json:"_links,omitempty"`
}

// BalancesList describes a list of captures.
type BalancesList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Balances []*Balance `json:"balances,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// BalanceListOptions contains valid query parameters
// for the list balances endpoint.
type BalanceListOptions struct {
	Currency string `url:"currency,omitempty"`
	From     string `url:"from,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

// BalanceGroupingFormat defines a grouping mechanism for transactions included
// in a balance report.
type BalanceGroupingFormat string

// Supported transaction grouping.
const (
	StatusBalancesGrouping        BalanceGroupingFormat = "status-balances"
	TransactionCategoriesGrouping BalanceGroupingFormat = "transaction-categories"
)

// BalanceReportOptions contains valid query parameters
// for the list balances endpoint.
type BalanceReportOptions struct {
	Grouping string     `url:"grouping,omitempty"`
	From     *ShortDate `url:"from,omitempty"`
	Until    *ShortDate `url:"until,omitempty"`
}

// Subtotal balance descriptor.
type Subtotal struct {
	TransactionType string      `json:"transactionType,omitempty"`
	Count           int         `json:"count,omitempty"`
	Amount          *Amount     `json:"amount,omitempty"`
	Subtotals       []*Subtotal `json:"subtotals,omitempty"`
}

// BalanceReportLinks holds URL objects relevant to the balance report.
type BalanceReportLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

type BalanceReport struct {
	Resource  string
	BalanceID string
	TimeZone  string
	From      *ShortDate
	Until     *ShortDate
	Grouping  BalanceGroupingFormat
}

// GetBalance retrieves a balance by its id.
func (bs *BalancesService) Get(ctx context.Context, balance string) (res *Response, b *Balance, err error) {
	return bs.get(ctx, balance)
}

// Primary retrieves the primary balance. This is the balance of your account’s
// primary currency, where all payments are settled to by default.
func (bs *BalancesService) Primary(ctx context.Context) (res *Response, b *Balance, err error) {
	const id = "primary"
	return bs.get(ctx, id)
}

// List retrieves all the organization’s balances, including the primary
// balance, ordered from newest to oldest.
func (bs *BalancesService) List(ctx context.Context, options *BalanceListOptions) (res *Response, bl *BalancesList, err error) {
	return bs.list(ctx, "v2/balances", options)
}

func (bs *BalancesService) get(ctx context.Context, balance string) (res *Response, b *Balance, err error) {
	u := fmt.Sprintf("v2/balances/%s", balance)

	res, err = bs.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &b); err != nil {
		return
	}

	return
}

func (bs *BalancesService) list(ctx context.Context, uri string, options interface{}) (res *Response, lb *BalancesList, err error) {
	res, err = bs.client.get(ctx, uri, options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &lb); err != nil {
		return
	}

	return
}
