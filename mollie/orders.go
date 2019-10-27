package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Orders explain the items that customers need to pay for.
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
	ConsumerDateOfBirth *ShortDate     `json:"consumerDateOfBirth,omitempty"`
	OrderNumber         string         `json:"orderNumber,omitempty"`
	ShippingAddress     *OrderAddress  `json:"shippingAddress,omitempty"`
	Locale              *Locale        `json:"locale,omitempty"`
	Metadata            interface{}    `json:"metadata,omitempty"`
	RedirectURL         string         `json:"redirectUrl,omitempty"`
	Lines               []*OrderLines  `json:"lines,omitempty"`
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
		Refunds  []Refund  `json:"refunds,omitempty"`
	} `json:"_embedded,omitempty"`
	Links        *OrderLinks   `json:"_links,omitempty"`
	OrderPayment *OrderPayment `json:"payment,omitempty"`
	Description  string        `json:"description,omitempty"`
}

// OrderPayment describes payment specific parameters that can be passed during order creation
type OrderPayment struct {
	ConsumerAccount   string          `json:"consumerAccount,omitempty"`
	CustomerID        string          `json:"customerId,omitempty"`
	CustomerReference string          `json:"customerReference,omitempty"`
	Issuer            string          `json:"issuer,omitempty"`
	MandateID         string          `json:"mandateId,omitempty"`
	SequenceType      *SequenceType   `json:"sequenceType,omitempty"`
	VoucherNumber     string          `json:"voucherNumber,omitempty"`
	VoucherPin        string          `json:"voucherPin,omitempty"`
	WebhookURL        string          `json:"webhookUrl,omitempty"`
	ApplicationFee    *ApplicationFee `json:"applicationFee,omitempty"`
	Method            *PaymentMethod  `json:"method,omitempty"`
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
	OrganizationName string       `json:"organizationName,omitempty"`
	Title            string       `json:"title,omitempty"`
	GivenName        string       `json:"givenName,omitempty"`
	FamilyName       string       `json:"familyName,omitempty"`
	Email            string       `json:"email,omitempty"`
	Phone            *PhoneNumber `json:"phone,omitempty"`
	StreetAndNumber  string       `json:"streetAndNumber,omitempty"`
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

// OrderListLinks describes an object with several URL objects
// relevant to the order.
// Every URL object will contain an href and a type field.
type OrderListLinks struct {
	Self          *URL `json:"self,omitempty"`
	Previous      *URL `json:"previous,omitempty"`
	Next          *URL `json:"next,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// OrderLines contain the actual things the customer bought.
type OrderLines struct {
	Resource           string           `json:"resource,omitempty"`
	ID                 string           `json:"id,omitempty"`
	OrderID            string           `json:"orderId,omitempty"`
	ProductType        *ProductType     `json:"type,omitempty"`
	Name               string           `json:"name,omitempty"`
	Amount             *Amount          `json:"amount,omitempty"`
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
	ImageURL           string           `json:"imageUrl,omitempty"`
	ProductURL         string           `json:"productUrl,omitempty"`
}

// OrderList for containing the response of list orders
type OrderList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Orders []Orders `json:"orders,omitempty"`
	} `json:"_embedded,omitempty"`
	Links OrderListLinks `json:"links,omitempty"`
}

// OrderListRefund for containing the response of list orders
type OrderListRefund struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Refunds []Refund `json:"refund,omitempty"`
	} `json:"_embedded,omitempty"`
	Links OrderListLinks `json:"links,omitempty"`
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
// See: https://docs.mollie.com/reference/v2/orders-api/get-order.
type OrderOptions struct {
	Embed     []EmbedValue `url:"embed,omitempty"`
	ProfileID string       `url:"profileId,omitempty"`
}

// OrderListOptions describes order endpoint valid query string parameters.
// See: https://docs.mollie.com/reference/v2/orders-api/list-orders.
type OrderListOptions struct {
	ProfileID string `url:"profileId,omitempty"`
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

// OrderListRefundOptions describes order endpoint valid query string parameters.
// See: https://docs.mollie.com/reference/v2/orders-api/list-orders.
type OrderListRefundOptions struct {
	From  string     `url:"from,omitempty"`
	Limit int        `url:"limit,omitempty"`
	Embed EmbedValue `url:"embed,omitempty"`
}

// OrdersService instance operates over refund resources.
type OrdersService service

// Get retrieve a single order by its ID.
// See https://docs.mollie.com/reference/v2/orders-api/get-order
func (ors *OrdersService) Get(orID string, opt *OrderOptions) (order Orders, err error) {
	u := fmt.Sprintf("v2/orders/%s", orID)
	if opt != nil {
		v, _ := query.Values(opt)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	req, err := ors.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// Create an order will automatically create the required payment to allow your customer to pay for the order.
// See https://docs.mollie.com/reference/v2/orders-api/create-order
func (ors *OrdersService) Create(ord Orders, opt *OrderOptions) (order Orders, err error) {
	u := fmt.Sprintf("v2/orders")
	if opt != nil {
		v, _ := query.Values(opt)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	req, err := ors.client.NewAPIRequest(http.MethodPost, u, ord)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// Update is used to update the billing and/or shipping address of an order.
// See https://docs.mollie.com/reference/v2/orders-api/update-order
func (ors *OrdersService) Update(orderID string, ord Orders) (order Orders, err error) {
	u := fmt.Sprintf("v2/orders/%s", orderID)

	req, err := ors.client.NewAPIRequest(http.MethodPatch, u, ord)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// Cancel try to cancel the order that fulfill certain requirements
// See https://docs.mollie.com/reference/v2/orders-api/cancel-order
func (ors *OrdersService) Cancel(orderID string) (order Orders, err error) {
	u := fmt.Sprintf("v2/orders/" + orderID)

	req, err := ors.client.NewAPIRequest(http.MethodDelete, u, nil)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// List is to retrieve all orders.
// See https://docs.mollie.com/reference/v2/orders-api/list-orders
func (ors *OrdersService) List(opt *OrderListOptions) (ordList OrderList, err error) {
	u := fmt.Sprintf("v2/orders")
	if opt != nil {
		v, _ := query.Values(opt)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	req, err := ors.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &ordList); err != nil {
		return
	}

	return
}

// UpdateOrderline can be used to update an order line.
// See https://docs.mollie.com/reference/v2/orders-api/update-orderline
func (ors *OrdersService) UpdateOrderline(orderID string, orderlineID string, orderline OrderLines) (order Orders, err error) {
	u := fmt.Sprintf("v2/orders/%s/lines/%s", orderID, orderlineID)

	req, err := ors.client.NewAPIRequest(http.MethodPatch, u, nil)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

// CancelOrderLine can be used to cancel one or more order lines
// that were previously authorized using a pay after delivery payment method.
// Use the Cancel Order API if you want to cancel the entire order or the remainder of the order.
// See https://docs.mollie.com/reference/v2/orders-api/cancel-order-lines
func (ors *OrdersService) CancelOrderLine(orderID string, orderlines *Orders) (errorResponse *ErrorResponse, err error) {
	u := fmt.Sprintf("v2/orders/%s/lines", orderID)

	req, err := ors.client.NewAPIRequest(http.MethodDelete, u, orderlines)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		if len(res.content) != 0 {
			err = json.Unmarshal(res.content, &errorResponse)
			if err != nil {
				return
			}
		}

		return
	}

	if res.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	return
}

// CreateOrderPayment can only be created while the status of the order is created,
// and when the status of the existing payment is either expired, canceled or failed.
// See https://docs.mollie.com/reference/v2/orders-api/create-order-payment
func (ors *OrdersService) CreateOrderPayment(orderID string, ordPay *OrderPayment) (payment *Payment, errorResponse *ErrorResponse, err error) {
	u := fmt.Sprintf("v2/orders/%s/payments", orderID)

	req, err := ors.client.NewAPIRequest(http.MethodPost, u, ordPay)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		if len(res.content) != 0 {
			err = json.Unmarshal(res.content, &errorResponse)
			if err != nil {
				return
			}
		}

		return
	}

	if err = json.Unmarshal(res.content, &payment); err != nil {
		return
	}

	return
}

// CreateOrderRefund using the Orders API, refunds should be made against the order.
// See https://docs.mollie.com/reference/v2/orders-api/create-order-refund
func (ors *OrdersService) CreateOrderRefund(orderID string, order *Orders) (refund Refund, errorResponse *ErrorResponse, err error) {
	u := fmt.Sprintf("v2/orders/%s/refunds", orderID)

	req, err := ors.client.NewAPIRequest(http.MethodPost, u, order)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		if len(res.content) != 0 {
			err = json.Unmarshal(res.content, &errorResponse)
			if err != nil {
				return
			}
		}

		return
	}

	if err = json.Unmarshal(res.content, &refund); err != nil {
		return
	}

	return
}

// ListOrderRefunds retrieve all order refunds.
// See https://docs.mollie.com/reference/v2/orders-api/list-order-refunds
func (ors *OrdersService) ListOrderRefunds(orderID string, opt *OrderListRefundOptions) (orderListRefund OrderListRefund, err error) {
	u := fmt.Sprintf("v2/orders/%s/refunds", orderID)
	if opt != nil {
		v, _ := query.Values(opt)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	req, err := ors.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ors.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &orderListRefund); err != nil {
		return
	}

	return
}
