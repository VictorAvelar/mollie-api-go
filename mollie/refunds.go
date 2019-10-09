package mollie

import "time"

// Refunds describe a refund for a certain payment
type Refunds struct {
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
