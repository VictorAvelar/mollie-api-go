package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// SubscriptionsService operates over subscriptions resource.
type SubscriptionsService service

// SubscriptionStatus contains references to valid subscription statuses.
type SubscriptionStatus string

// Available subscription statuses.
const (
	SubscriptionStatusPending   SubscriptionStatus = "pending"
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusCanceled  SubscriptionStatus = "canceled"
	SubscriptionStatusSuspended SubscriptionStatus = "suspended"
	SubscriptionStatusCompleted SubscriptionStatus = "completed"
)

// SubscriptionLinks contains several URL objects relevant to the subscription.
type SubscriptionLinks struct {
	Self          *URL `json:"self,omitempty"`
	Customer      *URL `json:"customer,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Subscription contains information about a customer subscription.
type Subscription struct {
	Resource        string                 `json:"resource,omitempty"`
	ID              string                 `json:"id,omitempty"`
	MandateID       string                 `json:"mandateId,omitempty"`
	Mode            Mode                   `json:"mode,omitempty"`
	CreatedAT       *time.Time             `json:"createdAt,omitempty"`
	Status          SubscriptionStatus     `json:"status,omitempty"`
	Amount          *Amount                `json:"amount,omitempty"`
	Times           int                    `json:"times,omitempty"`
	TimesRemaining  int                    `json:"timesRemaining,omitempty"`
	Interval        string                 `json:"interval,omitempty"`
	StartDate       *ShortDate             `json:"startDate,omitempty"`
	NextPaymentDate *ShortDate             `json:"nextPaymentDate,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Method          PaymentMethod          `json:"method,omitempty"`
	CanceledAt      *time.Time             `json:"canceledAt,omitempty"`
	WebhookURL      string                 `json:"webhookUrl,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	ApplicationFee  *ApplicationFee        `json:"applicationFee,omitempty"`
	TestMode        bool                   `json:"testmode,omitempty"`
	Links           SubscriptionLinks      `json:"_links,omitempty"`
}

// SubscriptionList describes the response for subscription list endpoints.
type SubscriptionList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Subscriptions []*Subscription
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// SubscriptionListOptions holds query string parameters valid for subscription lists.
type SubscriptionListOptions struct {
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}

// Get retrieves a customer's subscription
//
// See: https://docs.mollie.com/reference/v2/subscriptions-api/get-subscription
func (ss *SubscriptionsService) Get(ctx context.Context, cID, sID string) (res *Response, s *Subscription, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", cID, sID)

	res, err = ss.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}

// Create stores a new subscription for a given customer
//
// See: https://docs.mollie.com/reference/v2/subscriptions-api/create-subscription
func (ss *SubscriptionsService) Create(ctx context.Context, cID string, sc *Subscription) (res *Response, s *Subscription, err error) {
	uri := fmt.Sprintf("v2/customers/%s/subscriptions", cID)

	if ss.client.HasAccessToken() && ss.client.config.testing {
		sc.TestMode = true
	}

	res, err = ss.client.post(ctx, uri, sc, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}

// Update changes fields on a subscription object
//
// See: https://docs.mollie.com/reference/v2/subscriptions-api/update-subscription
func (ss *SubscriptionsService) Update(ctx context.Context, cID, sID string, sc *Subscription) (res *Response, s *Subscription, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", cID, sID)

	res, err = ss.client.patch(ctx, u, sc, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}

// Delete cancels a subscription
//
// See: https://docs.mollie.com/reference/v2/subscriptions-api/cancel-subscription
func (ss *SubscriptionsService) Delete(ctx context.Context, cID, sID string) (res *Response, s *Subscription, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", cID, sID)

	res, err = ss.client.delete(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}

// All retrieves all subscriptions, ordered from newest to oldest.
// By using an API key all the subscriptions created with the current website profile will be returned.
// In the case of an OAuth Access Token relies the website profile on the profileId field
//
// See: https://docs.mollie.com/reference/v2/subscriptions-api/list-all-subscriptions
func (ss *SubscriptionsService) All(ctx context.Context, opts *SubscriptionListOptions) (res *Response, sl *SubscriptionList, err error) {
	u := "v2/subscriptions"

	res, err = ss.list(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sl); err != nil {
		return
	}
	return
}

// List retrieves all subscriptions of a customer
//
// See: https://docs.mollie.com/reference/v2/subscriptions-api/list-subscriptions
func (ss *SubscriptionsService) List(ctx context.Context, cID string, opts *SubscriptionListOptions) (res *Response, sl *SubscriptionList, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions", cID)

	res, err = ss.list(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sl); err != nil {
		return
	}
	return
}

// GetPayments retrieves all payments of a specific subscriptions of a customer
//
// See: https://docs.mollie.com/reference/v2/subscriptions-api/list-subscriptions-payments
func (ss *SubscriptionsService) GetPayments(ctx context.Context, cID, sID string, opts *SubscriptionListOptions) (res *Response, sl *PaymentList, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s/payments", cID, sID)

	res, err = ss.list(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sl); err != nil {
		return
	}
	return
}

func (ss *SubscriptionsService) list(ctx context.Context, uri string, opts interface{}) (r *Response, err error) {
	r, err = ss.client.get(ctx, uri, opts)
	if err != nil {
		return
	}

	return
}
