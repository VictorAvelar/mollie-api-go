package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Refund describe a refund for a certain payment
type Refund struct {
	Resource         string        `json:"resource,omitempty"`
	ID               string        `json:"id,omitempty"`
	Amount           *Amount       `json:"amount,omitempty"`
	SettlementID     string        `json:"settlementId,omitempty"`
	SettlementAmount *Amount       `json:"settlementAmount,omitempty"`
	Description      string        `json:"description,omitempty"`
	Metadata         interface{}   `json:"metadata,omitempty"`
	Status           *RefundStatus `json:"status,omitempty"`
	Lines            *OrderLines   `json:"lines,omitempty"` //TODO: this property should be from order.go
	PaymentID        string        `json:"paymentId,omitempty"`
	OrderID          string        `json:"orderId,omitempty"`
	CreatedAt        *time.Time    `json:"createdAt,omitempty"`
	Links            *RefundLinks  `json:"_links,omitempty"`
}

// RefundStatus describes the status of the refund
type RefundStatus string

// Valid refund status
const (
	Queued     RefundStatus = "queued"
	Pending    RefundStatus = "pending"
	Processing RefundStatus = "processing"
	Refunded   RefundStatus = "refunded"
	Failed     RefundStatus = "failed"
)

// OrderLines describes an array of order line objects
type OrderLines struct {
	Quantity       int     `json:"quantity,omitempty"`
	DiscountAmount *Amount `json:"discountAmount,omitempty"`
	VatAmount      *Amount `json:"vatAmount,omitempty"`
	TotalAmount    *Amount `json:"totalAmount,omitempty"`
}

// RefundLinks describes all the possible links to be returned with
// a Refund object
type RefundLinks struct {
	Self       *URL `json:"self,omitempty"`
	Payment    *URL `json:"payment,omitempty"`
	Settlement *URL `json:"settlement,omitempty"`
	Order      *URL `json:"order,omitempty"`
}

// RefundOptions describes refund endpoint valid query string parameters.
//
// See: https://docs.mollie.com/reference/v2/refunds-api/get-refund
type RefundOptions struct {
	Embed    string `url:"embed,omitempty"`
	TestMode bool   `url:"testmode,omitempty"`
}

// RefundsService instance operates over refund resources
type RefundsService service

// Get Retrieve a single refund by its ID
//
// If you do not know the original payment’s ID, you can use the List payment refunds endpoint.
func (rs *RefundsService) Get(paymentID, refundID string, options *RefundOptions) (refund Refund, err error) {
	u := fmt.Sprintf("v2/payments/%s/refunds/%s", paymentID, refundID)
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v)
	}

	req, err := rs.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := rs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &refund); err != nil {
		return
	}

	return
}

var (
	requiredCreateParamRefund = "Parameter required for creating a refund: %+v"
)

// Create request a payment refund
//
// See https://docs.mollie.com/reference/v2/refunds-api/create-refund
func (rs *RefundsService) Create(paymentID string, re Refund, options *RefundOptions) (rf Refund, err error) {
	if re.Amount == nil {
		return re, fmt.Errorf(requiredCreateParamRefund, re.Amount)
	}

	if re.Amount.Currency == "" {
		return re, fmt.Errorf(requiredCreateParamRefund, re.Amount.Currency)
	}

	if re.Amount.Value == "" {
		return re, fmt.Errorf(requiredCreateParamRefund, re.Amount.Value)
	}

	u := fmt.Sprintf("v2/payments/%s/refunds", paymentID)
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v)
	}

	req, err := rs.client.NewAPIRequest(http.MethodPost, u, re)
	if err != nil {
		return
	}

	res, err := rs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &rf); err != nil {
		return
	}

	return
}

// Cancel try to cancel the refund request
// The refund can only be canceled while the refund’s status is either queued or pending
// See https://docs.mollie.com/reference/v2/refunds-api/cancel-refund
func (rs *RefundsService) Cancel(paymentID, refundID string, options *RefundOptions) (err error) {
	u := fmt.Sprintf("v2/payments/%s/refunds/%s", paymentID, refundID)
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v)
	}

	req, err := rs.client.NewAPIRequest(http.MethodDelete, u, nil)
	if err != nil {
		return
	}

	_, err = rs.client.Do(req)
	if err != nil {
		return
	}

	return
}
