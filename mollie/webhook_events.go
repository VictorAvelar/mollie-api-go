package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// WebhookEventService handles webhook event API calls.
type WebhookEventService service

// WebhookEntityLinks represents the links related to the entity received with a webhook event.
type WebhookEntityLinks struct {
	Self        *URL `json:"self,omitempty"`
	PaymentLink *URL `json:"payment,omitempty"`
}

// WebhookEntity represents the entity received with a webhook event.
type WebhookEntity struct {
	Archived        bool                `json:"archived,omitempty"`
	Reusable        bool                `json:"reusable,omitempty"`
	CustomerID      string              `json:"customerId,omitempty"`
	ID              string              `json:"id,omitempty"`
	Description     string              `json:"description,omitempty"`
	ProfileID       string              `json:"profileId,omitempty"`
	RedirectURL     string              `json:"redirectUrl,omitempty"`
	Resource        string              `json:"resource,omitempty"`
	WebhookURL      string              `json:"webhookUrl,omitempty"`
	Amount          Amount              `json:"amount,omitempty"`
	ApplicationFee  ApplicationFee      `json:"applicationFee,omitempty"`
	MinimumAmount   Amount              `json:"minimumAmount,omitempty"`
	Mode            Mode                `json:"mode,omitempty"`
	SequenceType    SequenceType        `json:"sequenceType,omitempty"`
	AllowedMethods  []PaymentMethod     `json:"allowedMethods,omitempty"`
	Lines           []WebhookEntityLine `json:"lines,omitempty"`
	Links           WebhookEntityLinks  `json:"_links,omitempty"`
	BillingAddress  *Address            `json:"billingAddress,omitempty"`
	ShippingAddress *Address            `json:"shippingAddress,omitempty"`
	CreatedAt       *time.Time          `json:"createdAt,omitempty"`
	ExpiresAt       *time.Time          `json:"expiresAt,omitempty"`
	PaidAt          *time.Time          `json:"paidAt,omitempty"`
}

// WebhookEntityLine represents a line item in the entity received with a webhook event.
type WebhookEntityLine struct {
	Quantity       int               `json:"quantity,omitempty"`
	Description    string            `json:"description,omitempty"`
	ImageURL       string            `json:"imageUrl,omitempty"`
	ProductURL     string            `json:"productUrl,omitempty"`
	QuantityUnit   string            `json:"quantityUnit,omitempty"`
	SKU            string            `json:"sku,omitempty"`
	VATRate        string            `json:"vatRate,omitempty"`
	DiscountAmount Amount            `json:"discountAmount,omitempty"`
	Type           ProductKind       `json:"type,omitempty"`
	TotalAmount    Amount            `json:"totalAmount,omitempty"`
	UnitPrice      Amount            `json:"unitPrice,omitempty"`
	VATAmount      Amount            `json:"vatAmount,omitempty"`
	Categories     []VoucherCategory `json:"categories,omitempty"`
}

// WebhookEventLinks represents the links related to a webhook event.
type WebhookEventLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	Entity        *URL `json:"entity,omitempty"`
}

// WebhookEvent represents a webhook event received from Mollie.
type WebhookEvent struct {
	Resource  string            `json:"resource,omitempty"`
	ID        string            `json:"id,omitempty"`
	Type      string            `json:"type,omitempty"`
	EntityID  string            `json:"entityId,omitempty"`
	Embedded  WebhookEntity     `json:"_embedded,omitempty"`
	Links     WebhookEventLinks `json:"_links,omitempty"`
	CreatedAt *time.Time        `json:"createdAt,omitempty"`
}

// Get retrieves a webhook event by its ID.
//
// See: https://docs.mollie.com/reference/get-webhook-event
func (s *WebhookEventService) Get(ctx context.Context, webhook string) (res *Response, we *WebhookEvent, err error) {
	res, err = s.client.get(ctx, fmt.Sprintf("/v2/events/%s", webhook), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &we); err != nil {
		return
	}

	return
}
