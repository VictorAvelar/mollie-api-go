package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CreatePaymentRefund describes the payload to create a refund associated to a payment.
type CreatePaymentRefund struct {
	Description string  `json:"description,omitempty"`
	Metadata    any     `json:"metadata,omitempty"`
	Amount      *Amount `json:"amount,omitempty"`
	PaymentRefundAccessTokenFields
	PaymentRefundMollieConnectFields
}

// CreateOrderRefund describes the payload to create a refund associated to an order.
type CreateOrderRefund struct {
	Description string             `json:"description,omitempty"`
	Metadata    any                `json:"metadata,omitempty"`
	Lines       []*OrderRefundLine `json:"lines,omitempty"`
	PaymentRefundAccessTokenFields
}

// OrderRefundLine describes the payload to create a refund associated to an order line.
type OrderRefundLine struct {
	Quantity int     `json:"quantity,omitempty"`
	ID       string  `json:"id,omitempty"`
	Amount   *Amount `json:"amount,omitempty"`
}

// PaymentRefundAccessTokenFields describes the access token available fields for a refund.
type PaymentRefundAccessTokenFields struct {
	Testmode bool `json:"testmode,omitempty"`
}

// PaymentRefundMollieConnectFields describes the Mollie Connect available fields for a refund.
type PaymentRefundMollieConnectFields struct {
	ReverseRouting   bool               `json:"reverseRouting,omitempty"`
	RoutingReversals []*RoutingReversal `json:"routingReversals,omitempty"`
}

// RoutingReversal describes the payload to be sent to the reverse routing endpoint.
type RoutingReversal struct {
	Amount *Amount `json:"amount,omitempty"`
	Source string  `json:"source,omitempty"`
}

// RoutingSource describes the source of the routing.
type RoutingSource struct {
	Type           string `json:"type,omitempty"`
	OrganizationID string `json:"organizationId,omitempty"`
}

// Refund describe a refund for a certain payment.
type Refund struct {
	Resource         string       `json:"resource,omitempty"`
	ID               string       `json:"id,omitempty"`
	SettlementID     string       `json:"settlementId,omitempty"`
	Description      string       `json:"description,omitempty"`
	PaymentID        string       `json:"paymentId,omitempty"`
	OrderID          string       `json:"orderId,omitempty"`
	Amount           *Amount      `json:"amount,omitempty"`
	SettlementAmount *Amount      `json:"settlementAmount,omitempty"`
	CreatedAt        *time.Time   `json:"createdAt,omitempty"`
	Lines            []*OrderLine `json:"lines,omitempty"`
	Metadata         any          `json:"metadata,omitempty"`
	Status           RefundStatus `json:"status,omitempty"`
	Links            RefundLinks  `json:"_links,omitempty"`
	PaymentRefundAccessTokenFields
	PaymentRefundMollieConnectFields
}

// RefundsList describes how a list of refunds will be retrieved by Mollie.
type RefundsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Refunds []*Refund
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// RefundStatus describes the status of the refund.
type RefundStatus string

// Valid refund status.
const (
	Queued     RefundStatus = "queued"
	Pending    RefundStatus = "pending"
	Processing RefundStatus = "processing"
	Refunded   RefundStatus = "refunded"
	Failed     RefundStatus = "failed"
)

// RefundLinks describes all the possible links to be returned with
// a Refund object.
type RefundLinks struct {
	Self          *URL `json:"self,omitempty"`
	Payment       *URL `json:"payment,omitempty"`
	Settlement    *URL `json:"settlement,omitempty"`
	Order         *URL `json:"order,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// PaymentRefundOptions describes payment refund endpoint valid query string parameters.
type PaymentRefundOptions struct {
	Embed []EmbedValue `url:"embed,omitempty"`
}

// RefundsListOptions describes payment and order refunds list endpoint valid query string parameters.
type RefundsListOptions struct {
	Limit     int          `url:"limit,omitempty"`
	From      string       `url:"from,omitempty"`
	ProfileID string       `url:"profileId,omitempty"`
	Embed     []EmbedValue `url:"embed,omitempty"`
}

// RefundsService instance operates over refund resources.
type RefundsService service

// List retrieves all refunds.
//
// See https://docs.mollie.com/reference/v2/refunds-api/list-refunds.
func (rs *RefundsService) List(ctx context.Context, opts *RefundsListOptions) (
	res *Response,
	rl *RefundsList,
	err error,
) {
	res, err = rs.client.get(ctx, "v2/refunds", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &rl); err != nil {
		return
	}

	return
}

// GetPaymentRefund retrieves a specific refund for a specific payment.
//
// See: https://docs.mollie.com/reference/v2/refunds-api/get-payment-refund
func (rs *RefundsService) GetPaymentRefund(
	ctx context.Context,
	paymentID, refundID string,
	opts *PaymentRefundOptions,
) (
	res *Response,
	refund *Refund,
	err error,
) {
	u := fmt.Sprintf("v2/payments/%s/refunds/%s", paymentID, refundID)

	res, err = rs.client.get(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &refund); err != nil {
		return
	}

	return
}

// ListPaymentRefunds retrieves all refunds for a specific payment.
//
// See: https://docs.mollie.com/reference/v2/refunds-api/list-payment-refunds
func (rs *RefundsService) ListPaymentRefunds(
	ctx context.Context,
	paymentID string,
	opts *RefundsListOptions,
) (
	res *Response,
	rl *RefundsList,
	err error,
) {
	u := fmt.Sprintf("v2/payments/%s/refunds", paymentID)

	res, err = rs.client.get(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &rl); err != nil {
		return
	}

	return
}

// Create a refund payment request.
//
// See https://docs.mollie.com/reference/v2/refunds-api/create-payment-refund
func (rs *RefundsService) CreatePaymentRefund(
	ctx context.Context,
	paymentID string,
	re CreatePaymentRefund,
	options *PaymentRefundOptions,
) (
	res *Response,
	rf *Refund,
	err error,
) {
	uri := fmt.Sprintf("v2/payments/%s/refunds", paymentID)

	if rs.client.HasAccessToken() && rs.client.config.testing {
		re.Testmode = true
	}

	res, err = rs.client.post(ctx, uri, re, options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &rf); err != nil {
		return
	}

	return
}

// CancelPaymentRefund cancels a refund for a specific payment.
//
// See https://docs.mollie.com/reference/v2/refunds-api/cancel-payment-refund
func (rs *RefundsService) CancelPaymentRefund(
	ctx context.Context, paymentID, refundID string,
) (res *Response, err error) {
	return rs.client.delete(ctx, fmt.Sprintf("v2/payments/%s/refunds/%s", paymentID, refundID), nil)
}

// CreateOrderRefund creates a refund for a specific order.
//
// See https://docs.mollie.com/reference/v2/refunds-api/create-order-refund
func (rs *RefundsService) CreateOrderRefund(
	ctx context.Context,
	orderID string,
	r CreateOrderRefund,
) (
	res *Response,
	rf *Refund,
	err error,
) {
	uri := fmt.Sprintf("v2/orders/%s/refunds", orderID)

	if rs.client.HasAccessToken() && rs.client.config.testing {
		r.Testmode = true
	}

	res, err = rs.client.post(ctx, uri, r, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &rf); err != nil {
		return
	}

	return
}

// ListOrderRefunds retrieves all refunds for a specific order.
//
// See https://docs.mollie.com/reference/v2/refunds-api/list-order-refunds
func (rs *RefundsService) ListOrderRefunds(
	ctx context.Context,
	orderID string,
	opts *RefundsListOptions,
) (
	res *Response,
	rl *RefundsList,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/refunds", orderID)

	res, err = rs.client.get(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &rl); err != nil {
		return
	}

	return
}
