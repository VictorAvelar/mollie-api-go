package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Order explain the items that customers need to pay for.
type Order struct {
	TestMode                                 bool          `json:"testmode,omitempty"`
	IsCancelable                             bool          `json:"isCancelable,omitempty"`
	ShopperCountryMustMatchTheBillingCountry bool          `json:"shopperCountryMustMatchTheBillingCountry,omitempty"`
	Resource                                 string        `json:"resource,omitempty"`
	ID                                       string        `json:"id,omitempty"`
	ProfileID                                string        `json:"profileId,omitempty"`
	OrderNumber                              string        `json:"orderNumber,omitempty"`
	RedirectURL                              string        `json:"redirectUrl,omitempty"`
	WebhookURL                               string        `json:"webhookUrl,omitempty"`
	Description                              string        `json:"description,omitempty"`
	Sku                                      string        `json:"sku,omitempty"`
	CancelURL                                string        `json:"cancelUrl,omitempty"`
	Metadata                                 interface{}   `json:"metadata,omitempty"`
	Mode                                     Mode          `json:"mode,omitempty"`
	Method                                   PaymentMethod `json:"method,omitempty"`
	Status                                   OrderStatus   `json:"status,omitempty"`
	Locale                                   Locale        `json:"locale,omitempty"`
	ShippingAddress                          *OrderAddress `json:"shippingAddress,omitempty"`
	Links                                    OrderLinks    `json:"_links,omitempty"`
	Amount                                   *Amount       `json:"amount,omitempty"`
	AmountCaptured                           *Amount       `json:"amountCaptured,omitempty"`
	AmountRefunded                           *Amount       `json:"amountRefunded,omitempty"`
	BillingAddress                           *OrderAddress `json:"billingAddress,omitempty"`
	ConsumerDateOfBirth                      *ShortDate    `json:"consumerDateOfBirth,omitempty"`
	CreatedAt                                *time.Time    `json:"createdAt,omitempty"`
	ExpiresAt                                *time.Time    `json:"expiresAt,omitempty"`
	ExpiredAt                                *time.Time    `json:"expiredAt,omitempty"`
	PaidAt                                   *time.Time    `json:"paidAt,omitempty"`
	AuthorizedAt                             *time.Time    `json:"authorizedAt,omitempty"`
	CanceledAt                               *time.Time    `json:"canceledAt,omitempty"`
	CompletedAt                              *time.Time    `json:"completedAt,omitempty"`
	OrderPayment                             *OrderPayment `json:"payment,omitempty"`
	Lines                                    []*OrderLine  `json:"lines,omitempty"`
	Embedded                                 struct {
		Payments []*Payment `json:"payments,omitempty"`
		Refunds  []*Refund  `json:"refunds,omitempty"`
	} `json:"_embedded,omitempty"`
}

// OrderPayment describes payment specific parameters that can be passed during order creation.
type OrderPayment struct {
	ConsumerAccount   string          `json:"consumerAccount,omitempty"`
	CustomerID        string          `json:"customerId,omitempty"`
	CustomerReference string          `json:"customerReference,omitempty"`
	Issuer            string          `json:"issuer,omitempty"`
	MandateID         string          `json:"mandateId,omitempty"`
	SequenceType      SequenceType    `json:"sequenceType,omitempty"`
	VoucherNumber     string          `json:"voucherNumber,omitempty"`
	VoucherPin        string          `json:"voucherPin,omitempty"`
	WebhookURL        string          `json:"webhookUrl,omitempty"`
	ApplicationFee    *ApplicationFee `json:"applicationFee,omitempty"`
	Method            PaymentMethod   `json:"method,omitempty"`
}

// OrderStatus describes the valid order status.
type OrderStatus string

// Valid order status.
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
	OrganizationName string      `json:"organizationName,omitempty"`
	Title            string      `json:"title,omitempty"`
	GivenName        string      `json:"givenName,omitempty"`
	FamilyName       string      `json:"familyName,omitempty"`
	Email            string      `json:"email,omitempty"`
	Phone            PhoneNumber `json:"phone,omitempty"`
	StreetAndNumber  string      `json:"streetAndNumber,omitempty"`
	StreetAdditional string      `json:"streetAdditional,omitempty"`
	PostalCode       string      `json:"postalCode,omitempty"`
	City             string      `json:"city,omitempty"`
	Region           string      `json:"region,omitempty"`
	Country          string      `json:"country,omitempty"`
}

// OrderLinks describes an object with several URL objects
// relevant to the order.
// Every URL object will contain an href and a type field.
type OrderLinks struct {
	Self          *URL `json:"self,omitempty"`
	Checkout      *URL `json:"checkout,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
}

// OrderLine contain the actual things the customer bought.
type OrderLine struct {
	Resource           string          `json:"resource,omitempty"`
	ID                 string          `json:"id,omitempty"`
	OrderID            string          `json:"orderId,omitempty"`
	ProductType        ProductType     `json:"type,omitempty"`
	Name               string          `json:"name,omitempty"`
	Amount             *Amount         `json:"amount,omitempty"`
	Status             OrderLineStatus `json:"status,omitempty"`
	IsCancelable       bool            `json:"isCancelable,omitempty"`
	Quantity           int             `json:"quantity,omitempty"`
	QuantityShipped    int             `json:"quantityShipped,omitempty"`
	AmountShipped      *Amount         `json:"amountShipped,omitempty"`
	QuantityRefunded   int             `json:"quantityRefunded,omitempty"`
	AmountRefunded     *Amount         `json:"amountRefunded,omitempty"`
	QuantityCanceled   int             `json:"quantityCanceled,omitempty"`
	AmountCanceled     *Amount         `json:"amountCanceled,omitempty"`
	ShippableQuantity  int             `json:"shippableQuantity,omitempty"`
	RefundableQuantity int             `json:"refundableQuantity,omitempty"`
	CancelableQuantity int             `json:"cancelableQuantity,omitempty"`
	UnitPrice          *Amount         `json:"unitPrice,omitempty"`
	DiscountAmount     *Amount         `json:"discountAmount,omitempty"`
	TotalAmount        *Amount         `json:"totalAmount,omitempty"`
	VatRate            string          `json:"vatRate,omitempty"`
	VatAmount          *Amount         `json:"vatAmount,omitempty"`
	SKU                string          `json:"sku,omitempty"`
	CreatedAt          *time.Time      `json:"createdAt,omitempty"`
	Links              OrderLineLinks  `json:"_links,omitempty"`
	ImageURL           string          `json:"imageUrl,omitempty"`
	ProductURL         string          `json:"productUrl,omitempty"`
	Metadata           interface{}     `json:"metadata,omitempty"`
}

// OrderLineOperation describes supported operations when managing order lines.
type OrderLineOperation string

// Supported order lines operation types.
const (
	AddOrderLine    OrderLineOperation = "add"
	UpdateOrderLine OrderLineOperation = "update"
	CancelOrderLine OrderLineOperation = "cancel"
)

// OrderLineOperationProductCategory contains the product category.
type OrderLineOperationProductCategory string

// Product category possible values.
const (
	MealProductCategory OrderLineOperationProductCategory = "meal"
	EcoProductCategory  OrderLineOperationProductCategory = "eco"
	GiftProductCategory OrderLineOperationProductCategory = "gift"
)

// OrderLineOperationData contains the order line’s details for an update operation.
type OrderLineOperationData struct {
	Quantity       int                               `json:"quantity,omitempty"`
	ID             string                            `json:"id,omitempty"`
	Name           string                            `json:"name,omitempty"`
	SKU            string                            `json:"sku,omitempty"`
	ImageURL       string                            `json:"imageUrl,omitempty"`
	ProductURL     string                            `json:"productUrl,omitempty"`
	VATRate        string                            `json:"vatRate,omitempty"`
	Type           string                            `json:"type,omitempty"`
	Category       OrderLineOperationProductCategory `json:"category,omitempty"`
	Amount         *Amount                           `json:"amount,omitempty"`
	UnitPrice      *Amount                           `json:"unitPrice,omitempty"`
	DiscountAmount *Amount                           `json:"discountAmount,omitempty"`
	VATAmount      *Amount                           `json:"vatAmount,omitempty"`
	TotalAmount    *Amount                           `json:"totalAmount,omitempty"`
	Metadata       interface{}                       `json:"metadata,omitempty"`
}

// OrderLineChangeInstruction contains details on what needs to be changed when managing order lines.
type OrderLineChangeInstruction struct {
	Operation OrderLineOperation      `json:"operation,omitempty"`
	Data      *OrderLineOperationData `json:"data,omitempty"`
}

// OrderLineOperations contains the operations to be performed when managing order lines.
type OrderLineOperations struct {
	Operations []*OrderLineChangeInstruction `json:"operations,omitempty"`
}

// OrderList for containing the response of list orders.
type OrderList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Orders []*Order `json:"orders,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}

// OrderListRefund for containing the response of list orders.
type OrderListRefund struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Refunds []*Refund `json:"refund,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}

// ProductType describes the type of product bought, for example, a physical or a digital product.
type ProductType string

// Valid product type.
const (
	Physical        ProductType = "physical"
	Discount        ProductType = "discount"
	Digital         ProductType = "digital"
	ShippingFee     ProductType = "shipping_fee"
	StoreCredit     ProductType = "store_credit"
	GiftCardProduct ProductType = "gift_card"
	Surcharge       ProductType = "surcharge"
)

// OrderLineStatus describes status of the order line.
type OrderLineStatus string

// Valid order line status.
const (
	OrderLineCreated    OrderLineStatus = "created"
	OrderLineAuthorized OrderLineStatus = "authorized"
	OrderLinePaid       OrderLineStatus = "paid"
	OrderLineShipping   OrderLineStatus = "shipping"
	OrderLineCanceled   OrderLineStatus = "canceled"
	OrderLineCompleted  OrderLineStatus = "completed"
)

// OrderLineLinks describes object with several URL objects relevant to the order line.
type OrderLineLinks struct {
	ProductURL *URL `json:"productUrl,omitempty"`
	ImageURL   *URL `json:"imageUrl,omitempty"`
}

// OrderOptions describes order endpoint valid query string parameters.
type OrderOptions struct {
	Embed     []EmbedValue `url:"embed,omitempty"`
	ProfileID string       `url:"profileId,omitempty"`
}

// OrderListOptions describes order endpoint valid query string parameters.
type OrderListOptions struct {
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

// OrderListRefundOptions describes order endpoint valid query string parameters.
type OrderListRefundOptions struct {
	From  string     `url:"from,omitempty"`
	Limit int        `url:"limit,omitempty"`
	Embed EmbedValue `url:"embed,omitempty"`
}

// OrdersService instance operates over order resources.
type OrdersService service

// Get retrieve a single order by its ID.
//
// See https://docs.mollie.com/reference/v2/orders-api/get-order
func (ors *OrdersService) Get(ctx context.Context, orID string, opts *OrderOptions) (
	res *Response,
	order *Order,
	err error,
) {
	res, err = ors.client.get(ctx, fmt.Sprintf("v2/orders/%s", orID), opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// Create an order will automatically create the required payment to allow your customer to pay for the order.
//
// See https://docs.mollie.com/reference/v2/orders-api/create-order
func (ors *OrdersService) Create(ctx context.Context, ord Order, opts *OrderOptions) (
	res *Response,
	order *Order,
	err error,
) {
	if ors.client.HasAccessToken() && ors.client.config.testing {
		ord.TestMode = true
	}

	res, err = ors.client.post(ctx, "v2/orders", ord, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// Update is used to update the billing and/or shipping address of an order.
//
// See https://docs.mollie.com/reference/v2/orders-api/update-order
func (ors *OrdersService) Update(ctx context.Context, orderID string, ord Order) (
	res *Response,
	order *Order,
	err error,
) {
	res, err = ors.client.patch(ctx, fmt.Sprintf("v2/orders/%s", orderID), ord, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// Cancel try to cancel the order that fulfill certain requirements.
//
// See https://docs.mollie.com/reference/v2/orders-api/cancel-order
func (ors *OrdersService) Cancel(ctx context.Context, orderID string) (res *Response, order *Order, err error) {
	res, err = ors.client.delete(ctx, fmt.Sprintf("v2/orders/%s", orderID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// List is to retrieve all orders.
//
// See https://docs.mollie.com/reference/v2/orders-api/list-orders
func (ors *OrdersService) List(ctx context.Context, opts *OrderListOptions) (
	res *Response,
	ordList *OrderList,
	err error,
) {
	res, err = ors.client.get(ctx, "v2/orders", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &ordList); err != nil {
		return
	}

	return
}

// UpdateOrderLine can be used to update an order line.
//
// See https://docs.mollie.com/reference/v2/orders-api/update-orderline
func (ors *OrdersService) UpdateOrderLine(
	ctx context.Context,
	orderID string,
	orderLineID string,
	orderLine OrderLine) (
	res *Response,
	order *Order,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/lines/%s", orderID, orderLineID)

	res, err = ors.client.patch(ctx, u, orderLine, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// CancelOrderLines can be used to cancel one or more order lines
// that were previously authorized using a pay after delivery payment method.
// Use the Cancel Order API if you want to cancel the entire order or the remainder of the order.
//
// See https://docs.mollie.com/reference/v2/orders-api/cancel-order-lines
func (ors *OrdersService) CancelOrderLines(ctx context.Context, orderID string, orderLines []OrderLine) (
	res *Response,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/lines", orderID)

	res, err = ors.client.delete(ctx, u, nil)
	if err != nil {
		return
	}

	return
}

// CreateOrderPayment can only be created while the status of the order is created,
// and when the status of the existing payment is either expired, canceled or failed.
//
// See https://docs.mollie.com/reference/v2/orders-api/create-order-payment
func (ors *OrdersService) CreateOrderPayment(ctx context.Context, orderID string, ordPay *OrderPayment) (
	res *Response,
	payment *Payment,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/payments", orderID)

	res, err = ors.client.post(ctx, u, ordPay, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &payment); err != nil {
		return
	}

	return
}

// CreateOrderRefund using the Orders API, refunds should be made against the order.
//
// See https://docs.mollie.com/reference/v2/orders-api/create-order-refund
func (ors *OrdersService) CreateOrderRefund(ctx context.Context, orderID string, order *Order) (
	res *Response,
	refund *Refund,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/refunds", orderID)

	res, err = ors.client.post(ctx, u, order, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &refund); err != nil {
		return
	}

	return
}

// ListOrderRefunds retrieve all order refunds.
//
// See https://docs.mollie.com/reference/v2/orders-api/list-order-refunds
func (ors *OrdersService) ListOrderRefunds(ctx context.Context, orderID string, opts *OrderListRefundOptions) (
	res *Response,
	orderListRefund *OrderListRefund,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/refunds", orderID)

	res, err = ors.client.get(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &orderListRefund); err != nil {
		return
	}

	return
}

// ManageOrderLines allows to update, cancel, or add one or more order lines.
//
// See: https://docs.mollie.com/reference/v2/orders-api/manage-order-lines
func (ors *OrdersService) ManageOrderLines(ctx context.Context, orderID string, operations *OrderLineOperations) (
	res *Response,
	order *Order,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/lines", orderID)

	res, err = ors.client.patch(ctx, u, operations, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}
