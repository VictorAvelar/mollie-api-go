package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Local embed value type.
const (
	EmbedNestedPaymentDetails EmbedValue = "payments.details.remainderDetails"
)

// CreateOrder contains the parameters to create an order.
type CreateOrder struct {
	ShopperCountryMustMatchTheBillingCountry bool            `json:"shopperCountryMustMatchTheBillingCountry,omitempty"`
	OrderNumber                              string          `json:"orderNumber,omitempty"`
	RedirectURL                              string          `json:"redirectUrl,omitempty"`
	WebhookURL                               string          `json:"webhookUrl,omitempty"`
	CancelURL                                string          `json:"cancelUrl,omitempty"`
	Amount                                   *Amount         `json:"amount,omitempty"`
	BillingAddress                           *OrderAddress   `json:"billingAddress,omitempty"`
	ShippingAddress                          *OrderAddress   `json:"shippingAddress,omitempty"`
	ConsumerDateOfBirth                      *ShortDate      `json:"consumerDateOfBirth,omitempty"`
	Payment                                  *OrderPayment   `json:"payment,omitempty"`
	ExpiresAt                                *ShortDate      `json:"expiresAt,omitempty"`
	Lines                                    []OrderLine     `json:"lines,omitempty"`
	Method                                   []PaymentMethod `json:"method,omitempty"`
	Locale                                   Locale          `json:"locale,omitempty"`
	Metadata                                 any             `json:"metadata,omitempty"`
	OrderAccessTokenFields
}

// OrderAccessTokenFields contains the fields available to include in an order when using an access token.
type OrderAccessTokenFields struct {
	ProfileID string `json:"profileId,omitempty"`
	Testmode  bool   `json:"testmode,omitempty"`
}

// Order explain the items that customers need to pay for.
type Order struct {
	IsCancelable                             bool          `json:"isCancelable,omitempty"`
	ShopperCountryMustMatchTheBillingCountry bool          `json:"shopperCountryMustMatchTheBillingCountry,omitempty"`
	Resource                                 string        `json:"resource,omitempty"`
	ID                                       string        `json:"id,omitempty"`
	ProfileID                                string        `json:"profileId,omitempty"`
	RedirectURL                              string        `json:"redirectUrl,omitempty"`
	CancelURL                                string        `json:"cancelUrl,omitempty"`
	WebhookURL                               string        `json:"webhookUrl,omitempty"`
	OrderNumber                              string        `json:"orderNumber,omitempty"`
	Lines                                    []*OrderLine  `json:"lines,omitempty"`
	Amount                                   *Amount       `json:"amount,omitempty"`
	AmountCaptured                           *Amount       `json:"amountCaptured,omitempty"`
	AmountRefunded                           *Amount       `json:"amountRefunded,omitempty"`
	BillingAddress                           *OrderAddress `json:"billingAddress,omitempty"`
	ConsumerDateOfBirth                      *ShortDate    `json:"consumerDateOfBirth,omitempty"`
	ShippingAddress                          *OrderAddress `json:"shippingAddress,omitempty"`
	CreatedAt                                *time.Time    `json:"createdAt,omitempty"`
	ExpiresAt                                *time.Time    `json:"expiresAt,omitempty"`
	ExpiredAt                                *time.Time    `json:"expiredAt,omitempty"`
	PaidAt                                   *time.Time    `json:"paidAt,omitempty"`
	AuthorizedAt                             *time.Time    `json:"authorizedAt,omitempty"`
	CanceledAt                               *time.Time    `json:"canceledAt,omitempty"`
	CompletedAt                              *time.Time    `json:"completedAt,omitempty"`
	Method                                   PaymentMethod `json:"method,omitempty"`
	Mode                                     Mode          `json:"mode,omitempty"`
	Locale                                   Locale        `json:"locale,omitempty"`
	Status                                   OrderStatus   `json:"status,omitempty"`
	Links                                    OrderLinks    `json:"_links,omitempty"`
	Metadata                                 any           `json:"metadata,omitempty"`
	Embedded                                 struct {
		Payments []*Payment `json:"payments,omitempty"`
		Refunds  []*Refund  `json:"refunds,omitempty"`
	} `json:"_embedded,omitempty"`
}

// UpdateOrder contains the parameters to update an order.
type UpdateOrder struct {
	OrderNumber     string        `json:"orderNumber,omitempty"`
	RedirectURL     string        `json:"redirectUrl,omitempty"`
	CancelURL       string        `json:"cancelUrl,omitempty"`
	WebhookURL      string        `json:"webhookUrl,omitempty"`
	BillingAddress  *OrderAddress `json:"billingAddress,omitempty"`
	ShippingAddress *OrderAddress `json:"shippingAddress,omitempty"`
	OrderAccessTokenFields
}

// OrdersList for containing the response of list orders.
type OrdersList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Orders []*Order `json:"orders,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
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

// OrderPayment describes payment specific parameters that can be passed during order creation.
type OrderPayment struct {
	ApplePayPaymentToken string     `json:"applePayPaymentToken,omitempty"`
	CardToken            string     `json:"cardToken,omitempty"`
	ConsumerAccount      string     `json:"consumerAccount,omitempty"`
	CustomerID           string     `json:"customerId,omitempty"`
	CustomerReference    string     `json:"customerReference,omitempty"`
	ExtraMerchantData    string     `json:"extraMerchantData,omitempty"`
	Issuer               string     `json:"issuer,omitempty"`
	VoucherNumber        string     `json:"voucherNumber,omitempty"`
	VoucherPin           string     `json:"voucherPin,omitempty"`
	WebhookURL           string     `json:"webhookUrl,omitempty"`
	BillingEmail         string     `json:"billingEmail,omitempty"`
	SessionID            string     `json:"sessionId,omitempty"`
	TerminalID           string     `json:"terminalId,omitempty"`
	ConsumerName         string     `json:"consumerName,omitempty"`
	DueDate              *ShortDate `json:"dueDate,omitempty"`
	ShippingAddress      *Address   `json:"shippingAddress,omitempty"`
	BillingAddress       *Address   `json:"billingAddress,omitempty"`
	Company              *Company   `json:"company,omitempty"`
	// Only available when using access tokens.
	ApplicationFee *ApplicationFee `json:"applicationFee,omitempty"`
	SequenceType   SequenceType    `json:"sequenceType,omitempty"`
	Method         []PaymentMethod `json:"method,omitempty"`
}

// OrderAddress identify both the address and the person the order is billed or shipped to.
type OrderAddress struct {
	OrganizationName string      `json:"organizationName,omitempty"`
	Title            string      `json:"title,omitempty"`
	GivenName        string      `json:"givenName,omitempty"`
	FamilyName       string      `json:"familyName,omitempty"`
	Email            string      `json:"email,omitempty"`
	StreetAndNumber  string      `json:"streetAndNumber,omitempty"`
	StreetAdditional string      `json:"streetAdditional,omitempty"`
	PostalCode       string      `json:"postalCode,omitempty"`
	City             string      `json:"city,omitempty"`
	Region           string      `json:"region,omitempty"`
	Country          string      `json:"country,omitempty"`
	Phone            PhoneNumber `json:"phone,omitempty"`
}

// OrderLine contain the actual things the customer bought.
type OrderLine struct {
	IsCancelable       bool            `json:"isCancelable,omitempty"`
	Quantity           int             `json:"quantity,omitempty"`
	QuantityShipped    int             `json:"quantityShipped,omitempty"`
	QuantityRefunded   int             `json:"quantityRefunded,omitempty"`
	QuantityCanceled   int             `json:"quantityCanceled,omitempty"`
	ShippableQuantity  int             `json:"shippableQuantity,omitempty"`
	RefundableQuantity int             `json:"refundableQuantity,omitempty"`
	CancelableQuantity int             `json:"cancelableQuantity,omitempty"`
	Resource           string          `json:"resource,omitempty"`
	ID                 string          `json:"id,omitempty"`
	OrderID            string          `json:"orderId,omitempty"`
	Name               string          `json:"name,omitempty"`
	VatRate            string          `json:"vatRate,omitempty"`
	SKU                string          `json:"sku,omitempty"`
	AmountShipped      *Amount         `json:"amountShipped,omitempty"`
	AmountRefunded     *Amount         `json:"amountRefunded,omitempty"`
	AmountCanceled     *Amount         `json:"amountCanceled,omitempty"`
	UnitPrice          *Amount         `json:"unitPrice,omitempty"`
	DiscountAmount     *Amount         `json:"discountAmount,omitempty"`
	TotalAmount        *Amount         `json:"totalAmount,omitempty"`
	VatAmount          *Amount         `json:"vatAmount,omitempty"`
	CreatedAt          *time.Time      `json:"createdAt,omitempty"`
	ProductType        ProductKind     `json:"type,omitempty"`
	Status             OrderLineStatus `json:"status,omitempty"`
	Links              OrderLineLinks  `json:"_links,omitempty"`
}

// OrderLineLinks describes object with several URL objects relevant to the order line.
type OrderLineLinks struct {
	ProductURL *URL `json:"productUrl,omitempty"`
	ImageURL   *URL `json:"imageUrl,omitempty"`
}

// UpdateOrderLine contains the parameters to update an order line.
type UpdateOrderLine struct {
	Quantity       int     `json:"quantity,omitempty"`
	Name           string  `json:"name,omitempty"`
	ImageURL       string  `json:"imageUrl,omitempty"`
	ProductURL     string  `json:"productUrl,omitempty"`
	SKU            string  `json:"sku,omitempty"`
	VATRate        string  `json:"vatRate,omitempty"`
	UnitPrice      *Amount `json:"unitPrice,omitempty"`
	DiscountAmount *Amount `json:"discountAmount,omitempty"`
	TotalAmount    *Amount `json:"totalAmount,omitempty"`
	VATAmount      *Amount `json:"vatAmount,omitempty"`
	Metadata       any     `json:"metadata,omitempty"`
	OrderAccessTokenFields
}

// OrderLineOperation describes supported operations when managing order lines.
type OrderLineOperation string

// Supported order lines operation types.
const (
	AddOrderLineOperation    OrderLineOperation = "add"
	UpdateOrderLineOperation OrderLineOperation = "update"
	CancelOrderLineOperation OrderLineOperation = "cancel"
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
	Kind           ProductKind                       `json:"type,omitempty"`
	Category       OrderLineOperationProductCategory `json:"category,omitempty"`
	Amount         *Amount                           `json:"amount,omitempty"`
	UnitPrice      *Amount                           `json:"unitPrice,omitempty"`
	DiscountAmount *Amount                           `json:"discountAmount,omitempty"`
	VATAmount      *Amount                           `json:"vatAmount,omitempty"`
	TotalAmount    *Amount                           `json:"totalAmount,omitempty"`
	Metadata       any                               `json:"metadata,omitempty"`
	OrderAccessTokenFields
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

// OrderRefundsList for containing the response of list orders.
type OrderRefundsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Refunds []*Refund `json:"refund,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}

// ProductKind describes the type of product bought, for example, a physical or a digital product.
type ProductKind string

// Valid product type.
const (
	PhysicalProduct        ProductKind = "physical"
	DiscountProduct        ProductKind = "discount"
	DigitalProduct         ProductKind = "digital"
	ShippingFeeProduct     ProductKind = "shipping_fee"
	StoreCreditProduct     ProductKind = "store_credit"
	GiftCardProductProduct ProductKind = "gift_card"
	SurchargeProduct       ProductKind = "surcharge"
	TipProduct             ProductKind = "tip"
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

// OrderOptions describes order endpoint valid query string parameters.
type OrderOptions struct {
	ProfileID string       `url:"profileId,omitempty"`
	Embed     []EmbedValue `url:"embed,omitempty"`
}

// ListOrdersOptions describes order endpoint valid query string parameters.
type ListOrdersOptions struct {
	Limit     int    `url:"limit,omitempty"`
	From      string `url:"from,omitempty"`
	Sort      string `url:"sort,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}

// ListOrderRefundsOptions describes order endpoint valid query string parameters.
type ListOrderRefundsOptions struct {
	From  string     `url:"from,omitempty"`
	Limit int        `url:"limit,omitempty"`
	Embed EmbedValue `url:"embed,omitempty"`
}

// OrdersService instance operates over refund resources.
type OrdersService service

// Get retrieve a single order by its ID.
//
// See https://docs.mollie.com/reference/get-order
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
// See https://docs.mollie.com/reference/create-order
func (ors *OrdersService) Create(ctx context.Context, ord CreateOrder, opts *OrderOptions) (
	res *Response,
	order *Order,
	err error,
) {
	if ors.client.HasAccessToken() && ors.client.config.testing {
		ord.Testmode = true
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
// See https://docs.mollie.com/reference/update-order
func (ors *OrdersService) Update(ctx context.Context, orderID string, ord UpdateOrder) (
	res *Response,
	order *Order,
	err error,
) {
	if ors.client.HasAccessToken() && ors.client.config.testing {
		ord.Testmode = true
	}

	res, err = ors.client.patch(ctx, fmt.Sprintf("v2/orders/%s", orderID), ord)
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
// See https://docs.mollie.com/reference/cancel-order
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
// See https://docs.mollie.com/reference/list-orders
func (ors *OrdersService) List(ctx context.Context, opts *ListOrdersOptions) (
	res *Response,
	ordList *OrdersList,
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
// See https://docs.mollie.com/reference/update-order-line
func (ors *OrdersService) UpdateOrderLine(
	ctx context.Context,
	orderID string,
	orderLineID string,
	orderLine UpdateOrderLine) (
	res *Response,
	order *Order,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/lines/%s", orderID, orderLineID)

	if ors.client.HasAccessToken() && ors.client.config.testing {
		orderLine.Testmode = true
	}

	res, err = ors.client.patch(ctx, u, orderLine)
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
// See https://docs.mollie.com/reference/cancel-order-lines
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
// See https://docs.mollie.com/reference/create-order-payment
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
// See https://docs.mollie.com/reference/create-order-refund
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
// See https://docs.mollie.com/reference/list-order-refunds-1
func (ors *OrdersService) ListOrderRefunds(ctx context.Context, orderID string, opts *ListOrderRefundsOptions) (
	res *Response,
	orderListRefund *OrderRefundsList,
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
// See: https://docs.mollie.com/reference/manage-order-lines
func (ors *OrdersService) ManageOrderLines(ctx context.Context, orderID string, operations *OrderLineOperations) (
	res *Response,
	order *Order,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/lines", orderID)

	res, err = ors.client.patch(ctx, u, operations)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}
