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
	Grouping BalanceGroupingFormat `url:"grouping,omitempty"`
	From     *ShortDate            `url:"from,omitempty"`
	Until    *ShortDate            `url:"until,omitempty"`
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

// BalanceReport contains the common fields between
// different balance grouping options.
type BalanceReport struct {
	Resource  string                       `json:"resource,omitempty"`
	BalanceID string                       `json:"balanceId,omitempty"`
	TimeZone  string                       `json:"timeZone,omitempty"`
	From      *ShortDate                   `json:"from,omitempty"`
	Until     *ShortDate                   `json:"until,omitempty"`
	Totals    *BalanceReportTotalsGrouping `json:"totals,omitempty"`
	Grouping  BalanceGroupingFormat        `json:"grouping,omitempty"`
	Links     BalanceReportLinks           `json:"_links,omitempty"`
}

// BalanceAmount wraps the std Amount type.
type BalanceAmount struct {
	Amount *Amount `json:"amount,omitempty"`
}

// BalanceReportDetail contains the breakdown
// categories when grouping balance transactions.
type BalanceReportDetail struct {
	Open                 *BalanceAmount `json:"open,omitempty"`
	Pending              *BalanceAmount `json:"pending,omitempty"`
	MovedToAvailable     *BalanceAmount `json:"movedToAvailable,omitempty"`
	ImmediatelyAvailable *BalanceAmount `json:"immediatelyAvailable,omitempty"`
	Close                *BalanceAmount `json:"close,omitempty"`
}

// BalanceReportTotalsGrouping contains the per totals
// grouped balances for the requested period.
type BalanceReportTotalsGrouping struct {
	PendingBalance   *BalanceReportDetail `json:"pendingBalance,omitempty"`
	AvailableBalance *BalanceReportDetail `json:"availableBalance,omitempty"`
	Open             *BalanceReportDetail `json:"open,omitempty"`
	Payments         *BalanceReportDetail `json:"payments,omitempty"`
	Refunds          *BalanceReportDetail `json:"refunds,omitempty"`
	Chargebacks      *BalanceReportDetail `json:"chargebacks,omitempty"`
	Capital          *BalanceReportDetail `json:"capital,omitempty"`
	Transfers        *BalanceReportDetail `json:"transfers,omitempty"`
	FeePrePayments   *BalanceReportDetail `json:"fee-prepayments,omitempty"`
	Corrections      *BalanceReportDetail `json:"corrections,omitempty"`
	Close            *BalanceReportDetail `json:"close,omitempty"`
}

// BalanceTransaction represents a the movement on your balance.
type BalanceTransaction struct {
	Resource        string        `json:"resource,omitempty"`
	ID              string        `json:"id,omitempty"`
	TransactionType string        `json:"transactionType,omitempty"`
	ResultAmount    *Amount       `json:"resultAmount,omitempty"`
	InitialAmount   *Amount       `json:"initialAmount,omitempty"`
	Deductions      *Amount       `json:"deductions,omitempty"`
	CreatedAt       *time.Time    `json:"createdAt,omitempty"`
	Context         ContextValues `json:"context,omitempty"`
}

type (
	// TransactionType specifies the reason for the movement.
	TransactionType string
	// ContextValue represents a relevant value in the system
	// associated with a BalanceTransaction.
	ContextValue string
)

// Known and supported transaction types.
const (
	PaymentTransaction                     TransactionType = "payment"
	CaptureTransaction                     TransactionType = "capture"
	UnauthorizedDirectDebitTransaction     TransactionType = "unauthorized-direct-debit"
	FailedPaymentTransaction               TransactionType = "failed-payment"
	RefundTransaction                      TransactionType = "refund-transaction"
	ReturnedRefundTransaction              TransactionType = "returned-refund"
	ChargebackTransaction                  TransactionType = "chargeback"
	ChargebackReversalTransaction          TransactionType = "chargeback-reversal"
	OutgoingTransferTransaction            TransactionType = "outgoing-transfer"
	CanceledOutgoingTransfer               TransactionType = "canceled-outgoing-transfer"
	ReturnedTransferTransaction            TransactionType = "returned-transfer"
	InvoiceCompensationTransferTransaction TransactionType = "invoice-compensation"
	BalanceCorrectionTransaction           TransactionType = "balance-correction"
	ApplicationFeeTransaction              TransactionType = "application-fee"
	SplitPaymentTransaction                TransactionType = "split-payment"
	PlatformPaymentRefundTransaction       TransactionType = "platform-payment-refund"
	PlatformPaymentChargeback              TransactionType = "platform-payment-chargeback"
)

// BalanceTransactionsList contains an array of embedded transactions.
type BalanceTransactionsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		BalanceTransactions []*BalanceTransaction `json:"balance_transactions,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// BalanceTransactionsListOptions are valid query parameters for list
// balance transactions requests.
type BalanceTransactionsListOptions struct {
	From  string `url:"from,omitempty"`
	Limit int    `url:"limit,omitempty"`
}

// GetBalance retrieves a balance by its id.
//
// See: https://docs.mollie.com/reference/v2/balances-api/get-balance
func (bs *BalancesService) Get(ctx context.Context, balance string) (res *Response, b *Balance, err error) {
	return bs.get(ctx, balance)
}

// Primary retrieves the primary balance. This is the balance of your account’s
// primary currency, where all payments are settled to by default.
//
// See: https://docs.mollie.com/reference/v2/balances-api/get-primary-balance
func (bs *BalancesService) Primary(ctx context.Context) (res *Response, b *Balance, err error) {
	const id = "primary"

	return bs.get(ctx, id)
}

// List retrieves all the organization’s balances, including the primary
// balance, ordered from newest to oldest.
//
// See: https://docs.mollie.com/reference/v2/balances-api/list-balances
func (bs *BalancesService) List(ctx context.Context, options *BalanceListOptions) (
	res *Response,
	bl *BalancesList,
	err error,
) {
	return bs.list(ctx, "v2/balances", options)
}

// GetReport returns the balance report for the specified balance id.
//
// See: https://docs.mollie.com/reference/v2/balances-api/get-balance-report
func (bs *BalancesService) GetReport(ctx context.Context, balance string, options *BalanceReportOptions) (
	res *Response,
	br *BalanceReport,
	err error,
) {
	return bs.getReport(ctx, balance, options)
}

// GetPrimaryReport returns the report for the primary balance.
//
// See: https://docs.mollie.com/reference/v2/balances-api/get-primary-balance-report
func (bs *BalancesService) GetPrimaryReport(ctx context.Context, options *BalanceReportOptions) (
	res *Response,
	br *BalanceReport,
	err error,
) {
	return bs.getReport(ctx, "primary", options)
}

// GetTransactionsList retrieves a list of movements (transactions) for the
// specified balance.
//
// See: https://docs.mollie.com/reference/v2/balances-api/list-balance-transactions
func (bs *BalancesService) GetTransactionsList(
	ctx context.Context,
	balance string,
	options *BalanceTransactionsListOptions,
) (
	res *Response,
	btl *BalanceTransactionsList,
	err error,
) {
	return bs.listTransactions(ctx, balance, options)
}

// GetPrimaryTransactionsList retrieves the list of movements (transactions) for the
// primary balance of the account.
//
// See: https://docs.mollie.com/reference/v2/balances-api/list-primary-balance-transactions
func (bs *BalancesService) GetPrimaryTransactionsList(ctx context.Context, options *BalanceTransactionsListOptions) (
	res *Response,
	btl *BalanceTransactionsList,
	err error,
) {
	return bs.listTransactions(ctx, "primary", options)
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

func (bs *BalancesService) list(
	ctx context.Context,
	uri string,
	options interface{},
) (res *Response, lb *BalancesList, err error) {
	res, err = bs.client.get(ctx, uri, options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &lb); err != nil {
		return
	}

	return
}

func (bs *BalancesService) getReport(
	ctx context.Context,
	balance string,
	options *BalanceReportOptions,
) (res *Response, br *BalanceReport, err error) {
	u := fmt.Sprintf("v2/balances/%s/report", balance)

	res, err = bs.client.get(ctx, u, options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &br); err != nil {
		return
	}

	return
}

func (bs *BalancesService) listTransactions(
	ctx context.Context,
	balance string,
	options *BalanceTransactionsListOptions,
) (res *Response, btl *BalanceTransactionsList, err error) {
	u := fmt.Sprintf("v2/balances/%s/transactions", balance)

	res, err = bs.client.get(ctx, u, options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &btl); err != nil {
		return
	}

	return
}
