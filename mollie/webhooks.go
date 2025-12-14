package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// WebhookEventType represents the type of event that triggers a webhook.
type WebhookEventType string

// List of possible webhook event types.
const (
	PaymentLinkPaidWebhookEvent           WebhookEventType = "payment-link.paid"
	BalanceTransactionCreatedWebhookEvent WebhookEventType = "balance-transaction.created"
	SalesInvoiceCreatedWebhookEvent       WebhookEventType = "sales-invoice.created"
	SalesInvoiceIssuedWebhookEvent        WebhookEventType = "sales-invoice.issued"
	SalesInvoiceCanceledWebhookEvent      WebhookEventType = "sales-invoice.canceled"
	SalesInvoicePaidWebhookEvent          WebhookEventType = "sales-invoice.paid"
	AllWebhookEvents                      WebhookEventType = "*"
)

// WebhookStatus represents the status of a webhook.
type WebhookStatus string

// List of possible webhook statuses.
const (
	WebhookStatusEnabled  WebhookStatus = "enabled"
	WebhookStatusDisabled WebhookStatus = "disabled"
	WebhookStatusBlocked  WebhookStatus = "blocked"
	WebhookStatusDeleted  WebhookStatus = "deleted"
)

// CreateWebhook represents the payload to create a new webhook.
type CreateWebhook struct {
	TestMode   bool               `json:"testmode,omitempty"`
	Name       string             `json:"name"`
	URL        string             `json:"url"`
	EventTypes []WebhookEventType `json:"eventTypes"`
}

// UpdateWebhook represents the payload to update an existing webhook.
type UpdateWebhook struct {
	TestMode   bool               `json:"testmode,omitempty"`
	Name       string             `json:"name,omitempty"`
	URL        string             `json:"url,omitempty"`
	EventTypes []WebhookEventType `json:"eventTypes,omitempty"`
}

// DeleteWebhook represents the payload to delete a webhook.
type DeleteWebhook struct {
	TestMode bool `json:"testmode,omitempty"`
}

// TestWebhook represents the payload to ping a webhook.
type TestWebhook struct {
	TestMode bool `json:"testmode,omitempty"`
}

// Webhook describes a webhook object registered in the Mollie system.
type Webhook struct {
	Resource      string             `json:"resource,omitempty"`
	ID            string             `json:"id,omitempty"`
	URL           string             `json:"url,omitempty"`
	ProfileID     string             `json:"profileId,omitempty"`
	Name          string             `json:"name,omitempty"`
	WebhookSecret string             `json:"webhookSecret,omitempty"`
	Mode          Mode               `json:"mode,omitempty"`
	Status        WebhookStatus      `json:"status,omitempty"`
	Links         WebhookLinks       `json:"_links,omitempty"`
	EventTypes    []WebhookEventType `json:"eventTypes,omitempty"`
	CreatedAt     *time.Time         `json:"createdAt,omitempty"`
}

// WebhookLinks represents the links related to a webhook.
type WebhookLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// WebhooksListOptions represents the query parameters for listing webhooks.
type WebhooksListOptions struct {
	From      string             `url:"from,omitempty"`
	Limit     int                `url:"limit,omitempty"`
	Sort      SortDirection      `url:"sort,omitempty"`
	EventType []WebhookEventType `url:"eventType,omitempty"`
}

// WebhookList represents a paginated list of webhooks.
type WebhookList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Webhooks []*Webhook `json:"webhooks,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// WebhookService handles webhook-related operations.
type WebhookService service

// Create creates a new webhook.
//
// See: https://docs.mollie.com/reference/create-webhook
func (s *WebhookService) Create(ctx context.Context, wh CreateWebhook) (
	res *Response,
	w *Webhook,
	err error,
) {
	if s.client.HasAccessToken() && s.client.config.testing {
		wh.TestMode = true
	}

	res, err = s.client.post(ctx, "/v2/webhooks", wh, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &w); err != nil {
		return
	}

	return
}

// Get retrieves a webhook by its ID.
//
// See: https://docs.mollie.com/reference/get-webhook
func (s *WebhookService) Get(ctx context.Context, webhook string) (
	res *Response,
	w *Webhook,
	err error,
) {
	res, err = s.client.get(ctx, fmt.Sprintf("/v2/webhooks/%s", webhook), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &w); err != nil {
		return
	}

	return
}

// Update modifies an existing webhook.
//
// See: https://docs.mollie.com/reference/update-webhook
func (s *WebhookService) Update(ctx context.Context, webhook string, uw UpdateWebhook) (
	res *Response,
	w *Webhook,
	err error,
) {
	if s.client.HasAccessToken() && s.client.config.testing {
		uw.TestMode = true
	}

	res, err = s.client.patch(ctx, fmt.Sprintf("/v2/webhooks/%s", webhook), uw)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &w); err != nil {
		return
	}

	return
}

// List retrieves a list of webhooks with optional filtering.
//
// See: https://docs.mollie.com/reference/list-webhooks
func (s *WebhookService) List(ctx context.Context, options *WebhooksListOptions) (
	res *Response,
	wl *WebhookList,
	err error,
) {
	res, err = s.client.get(ctx, "/v2/webhooks", options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &wl); err != nil {
		return
	}

	return
}

// Delete removes a webhook by its ID.
//
// See: https://docs.mollie.com/reference/delete-webhook
func (s *WebhookService) Delete(ctx context.Context, webhook string) (
	res *Response,
	err error,
) {
	var dw DeleteWebhook
	if s.client.HasAccessToken() && s.client.config.testing {
		dw = DeleteWebhook{
			TestMode: true,
		}
	}

	res, err = s.client.delete(ctx, fmt.Sprintf("/v2/webhooks/%s", webhook), dw)
	if err != nil {
		return
	}

	return
}

// Test sends a test ping to the specified webhook.
//
// See: https://docs.mollie.com/reference/test-webhook
func (s *WebhookService) Test(ctx context.Context, webhook string) (
	res *Response,
	err error,
) {
	var tw TestWebhook
	if s.client.HasAccessToken() && s.client.config.testing {
		tw = TestWebhook{
			TestMode: true,
		}
	}

	res, err = s.client.post(ctx, fmt.Sprintf("/v2/webhooks/%s/ping", webhook), tw, nil)
	if err != nil {
		return
	}

	return
}
