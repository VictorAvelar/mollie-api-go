package mollie

import (
	"time"
)

// Orders explain the items that customers need to pay for
type Orders struct {
	Resource            string         `json:"resource,omitempty"`
	ID                  string         `json:"id,omitempty"`
	ProfileID           string         `json:"profileId,omitempty"`
	Method              *PaymentMethod `json:"method,omitempty"`
	Mode                *Mode          `json:"mode,omitempty"`
	Amount              *Amount        `json:"amount,omitempty"`
	AmountCaptured      *Amount        `json:"amountCaptured,omitempty"`
	AmountRefunded      *Amount        `json:"amountRefunded,omitempty"`
	Status              *OrderStatus   `json:"status,omitempty"`
	IsCancelable        bool           `json:"isCancelable,omitempty"`
	BillingAddress      *OrderAddress  `json:"billingAddress,omitempty"`
	ConsumerDateOfBirth *time.Time     `json:"consumerDateOfBirth,omitempty"`
	OrderNumber         string         `json:"orderNumber,omitempty"`
	ShippingAddress     *OrderAddress  `json:"shippingAddress,omitempty"`
	Locale              *Locale        `json:"locale,omitempty"`
	Metadata            interface{}    `json:"metadata,omitempty"`
	RedirectURL         string         `json:"redirectUrl,omitempty"`
	Lines               *OrderLines    `json:"lines,omitempty"`
	WebhookURL          string         `json:"webhookUrl,omitempty"`
	CreatedAt           *time.Time     `json:"createdAt,omitempty"`
	ExpiresAt           *time.Time     `json:"expiresAt,omitempty"`
	ExpiredAt           *time.Time     `json:"expiredAt,omitempty"`
	PaidAt              *time.Time     `json:"paidAt,omitempty"`
	AuthorizedAt        *time.Time     `json:"authorizedAt,omitempty"`
	CanceledAt          *time.Time     `json:"canceledAt,omitempty"`
	CompletedAt         *time.Time     `json:"completedAt,omitempty"`
	Embedded            struct {
		Payments []Payment `json:"payments,omitempty"`
		Refunds  []Refund  `json:"refunds,omitempry"`
	} `json:"_embedded,omitempty"`
	Links *OrderLinks `json:"_links,omitempty"`
}

// OrderStatus describes the valid order status
type OrderStatus string

// Valid order status
const (
	Created    OrderStatus = "created"
	Paid       OrderStatus = "paid"
	Authorized OrderStatus = "authorized"
	Canceled   OrderStatus = "canceled"
	Shipping   OrderStatus = "shipping"
	Completed  OrderStatus = "completed"
	Expired    OrderStatus = "expired"
)

// OrderAddress identify both the address and the person the order is billed or shipped to.
type OrderAddress struct {
	OrganizationName string       `json:"organizationName,omitempty"`
	Title            string       `json:"title,omitempty"`
	GivenName        string       `json:"givenName,omitempty"`
	FamilyName       string       `json:"familyName,omitempty"`
	Email            string       `json:"email,omitempty"`
	Phone            *PhoneNumber `json:"phone,omitempty"`
	StreetAndNumber  string       `json:"stringAndNumber,omitempty"`
	StreetAdditional string       `json:"streetAdditional,omitempty"`
	PostalCode       string       `json:"postalCode,omitempty"`
	City             string       `json:"city,omitempty"`
	Region           string       `json:"region,omitempty"`
	Country          string       `json:"country,omitempty"`
}

// OrderLinks describes an object with several URL objects
// relevant to the order.
// Every URL object will contain an href and a type field.
type OrderLinks struct {
	Self          *URL `json:"self,omitempty"`
	Checkout      *URL `json:"checkout,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// OrderLines contain the actual things the customer bought.
type OrderLines struct {
	Resource           string           `json:"resource,omitempty"`
	ID                 string           `json:"id,omitempty"`
	OrderID            string           `json:"orderId,omitempty"`
	ProductType        *ProductType     `json:"type,omitempty"`
	Name               string           `json:"name,omitempty"`
	Status             *OrderLineStatus `json:"status,omitempty"`
	IsCancelable       bool             `json:"isCancelable,omitempty"`
	Quantity           int              `json:"quantity,omitempty"`
	QuantityShipped    int              `json:"quantityShipped,omitempty"`
	AmountShipped      *Amount          `json:"amountShipped,omitempty"`
	QuantityRefunded   int              `json:"quantityRefunded,omitempty"`
	AmountRefunded     *Amount          `json:"amountRefunded,omitempty"`
	QuantityCanceled   int              `json:"quantityCanceled,omitempty"`
	AmountCanceled     *Amount          `json:"amountCanceled,omitempty"`
	ShippableQuantity  int              `json:"shippableQuantity,omitempty"`
	RefundableQuantity int              `json:"refundableQuantity,omitempty"`
	CancelableQuantity int              `json:"cancelableQuantity,omitempty"`
	UnitPrice          *Amount          `json:"unitPrice,omitempty"`
	DiscountAmount     *Amount          `json:"discountAmount,omitempty"`
	TotalAmount        *Amount          `json:"totalAmount,omitempty"`
	VatRate            string           `json:"vatRate,omitempty"`
	VatAmount          *Amount          `json:"vatAmount,omitempty"`
	SKU                string           `json:"sku,omitempty"`
	CreatedAt          *time.Time       `json:"createdAt,omitempty"`
	Links              *OrderLineLinks  `json:"_links,omitempty"`
}

// ProductType describes the type of product bought, for example, a physical or a digital product.
type ProductType string

// Valid product type
const (
	Physical        ProductType = "physical"
	Discount        ProductType = "discount"
	Digital         ProductType = "digital"
	ShippingFee     ProductType = "shipping_fee"
	StoreCredit     ProductType = "store_credit"
	GiftCardProduct ProductType = "gift_card"
	Surcharge       ProductType = "surcharge"
)

// OrderLineStatus describes status of the order line
type OrderLineStatus string

// Valid order line status
const (
	OrderLineCreated    OrderLineStatus = "created"
	OrderLineAuthorized OrderLineStatus = "authorized"
	OrderLinePaid       OrderLineStatus = "paid"
	OrderLineShipping   OrderLineStatus = "shipping"
	OrderLineCanceled   OrderLineStatus = "canceled"
	OrderLineCompleted  OrderLineStatus = "completed"
)

// OrderLineLinks describes object with several URL objects relevant to the order line
type OrderLineLinks struct {
	ProductURL *URL `json:"productUrl,omitempty"`
	ImageURL   *URL `json:"imageUrl,omitempty"`
}
