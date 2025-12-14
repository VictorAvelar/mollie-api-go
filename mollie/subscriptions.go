package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

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

// CreateSubscription contains the fields that are required to create a subscription.
type CreateSubscription struct {
	Times       int           `json:"times,omitempty"`
	Interval    string        `json:"interval,omitempty"`
	Description string        `json:"description,omitempty"`
	MandateID   string        `json:"mandateId,omitempty"`
	WebhookURL  string        `json:"webhookUrl,omitempty"`
	Amount      *Amount       `json:"amount,omitempty"`
	StartDate   *ShortDate    `json:"startDate,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
	Metadata    any           `json:"metadata,omitempty"`
	SubscriptionAccessTokenFields
}

// SubscriptionAccessTokenFields contains the fields that are available when using an access token.
type SubscriptionAccessTokenFields struct {
	Testmode       bool            `json:"testmode,omitempty"`
	ProfileID      string          `json:"profileId,omitempty"`
	ApplicationFee *ApplicationFee `json:"applicationFee,omitempty"`
}

// UpdateSubscription contains the fields that are required to create a subscription.
type UpdateSubscription struct {
	Times       int           `json:"times,omitempty"`
	Interval    string        `json:"interval,omitempty"`
	Description string        `json:"description,omitempty"`
	MandateID   string        `json:"mandateId,omitempty"`
	WebhookURL  string        `json:"webhookUrl,omitempty"`
	Amount      *Amount       `json:"amount,omitempty"`
	StartDate   *ShortDate    `json:"startDate,omitempty"`
	Method      PaymentMethod `json:"method,omitempty"`
	Metadata    any           `json:"metadata,omitempty"`
	SubscriptionAccessTokenFields
}

// SubscriptionLinks contains several URL objects relevant to the subscription.
type SubscriptionLinks struct {
	Self          *URL `json:"self,omitempty"`
	Customer      *URL `json:"customer,omitempty"`
	Profile       *URL `json:"profile,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Subscription contains information about a customer subscription.
type Subscription struct {
	Times           int                `json:"times,omitempty"`
	TimesRemaining  int                `json:"timesRemaining,omitempty"`
	Resource        string             `json:"resource,omitempty"`
	ID              string             `json:"id,omitempty"`
	Interval        string             `json:"interval,omitempty"`
	Description     string             `json:"description,omitempty"`
	MandateID       string             `json:"mandateId,omitempty"`
	WebhookURL      string             `json:"webhookUrl,omitempty"`
	Amount          *Amount            `json:"amount,omitempty"`
	ApplicationFee  *ApplicationFee    `json:"applicationFee,omitempty"`
	StartDate       *ShortDate         `json:"startDate,omitempty"`
	NextPaymentDate *ShortDate         `json:"nextPaymentDate,omitempty"`
	CreatedAT       *time.Time         `json:"createdAt,omitempty"`
	CanceledAt      *time.Time         `json:"canceledAt,omitempty"`
	Mode            Mode               `json:"mode,omitempty"`
	Status          SubscriptionStatus `json:"status,omitempty"`
	Method          PaymentMethod      `json:"method,omitempty"`
	Metadata        any                `json:"metadata,omitempty"`
	Links           SubscriptionLinks  `json:"_links,omitempty"`
}

// SubscriptionsList describes the response for subscription list endpoints.
type SubscriptionsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Subscriptions []*Subscription
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// ListSubscriptionsOptions holds query string parameters valid for subscription lists.
type ListSubscriptionsOptions struct {
	Testmode  bool   `url:"testmode,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	From      string `url:"from,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}

// SubscriptionsService operates over subscriptions resource.
type SubscriptionsService service

// Get retrieves a customer's subscription
//
// See: https://docs.mollie.com/reference/get-subscription
func (ss *SubscriptionsService) Get(ctx context.Context, customer, subscription string) (
	res *Response,
	s *Subscription,
	err error,
) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", customer, subscription)

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
// See: https://docs.mollie.com/reference/create-subscription
func (ss *SubscriptionsService) Create(ctx context.Context, customer string, sc CreateSubscription) (
	res *Response,
	s *Subscription,
	err error,
) {
	uri := fmt.Sprintf("v2/customers/%s/subscriptions", customer)

	if ss.client.HasAccessToken() && ss.client.config.testing {
		sc.Testmode = true
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
// See: https://docs.mollie.com/reference/update-subscription
func (ss *SubscriptionsService) Update(ctx context.Context, customer, subscription string, sc UpdateSubscription) (
	res *Response,
	s *Subscription,
	err error,
) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", customer, subscription)

	res, err = ss.client.patch(ctx, u, sc)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}

	return
}

// Cancel cancels a subscription.
//
// See: https://docs.mollie.com/reference/cancel-subscription
func (ss *SubscriptionsService) Cancel(ctx context.Context, customer, subscription string) (
	res *Response,
	s *Subscription,
	err error,
) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", customer, subscription)

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
// See: https://docs.mollie.com/reference/list-all-subscriptions
func (ss *SubscriptionsService) All(ctx context.Context, opts *ListSubscriptionsOptions) (
	res *Response,
	sl *SubscriptionsList,
	err error,
) {
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
// See: https://docs.mollie.com/reference/list-subscriptions
func (ss *SubscriptionsService) List(ctx context.Context, customer string, opts *ListSubscriptionsOptions) (
	res *Response,
	sl *SubscriptionsList,
	err error,
) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions", customer)

	res, err = ss.list(ctx, u, opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sl); err != nil {
		return
	}

	return
}

// ListPayments retrieves all payments of a specific subscriptions of a customer
//
// See: https://docs.mollie.com/reference/list-subscription-payments
func (ss *SubscriptionsService) ListPayments(
	ctx context.Context,
	customer, subscription string,
	opts *ListSubscriptionsOptions,
) (
	res *Response,
	sl *PaymentList,
	err error,
) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s/payments", customer, subscription)

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
