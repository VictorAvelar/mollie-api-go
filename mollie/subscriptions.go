package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// SubscriptionsService operates over subscriptions resource
type SubscriptionsService service

// SubscriptionStatus contains references to valid subscription statuses
type SubscriptionStatus string

// Available subscription statuses
const (
	SubscriptionStatusPending   SubscriptionStatus = "pending"
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusCanceled  SubscriptionStatus = "canceled"
	SubscriptionStatusSuspended SubscriptionStatus = "suspended"
	SubscriptionStatusCompleted SubscriptionStatus = "completed"
)

// SubscriptionLinks contains several URL objects relevant to the subscription
type SubscriptionLinks struct {
	Self          *URL `json:"self,omitempty"`
	Customer      *URL `json:"customer,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Subscription contains information about a customer subscription
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

// SubscriptionList describes the response for subscription list endpoints
type SubscriptionList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Subscriptions []*Subscription
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// SubscriptionListOptions holds query string parameters valid for subscription lists
type SubscriptionListOptions struct {
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}

// Get retrieves a customer's subscription
//
// See: https://docs.mollie.com/reference/v2/subscriptions-api/get-subscription
func (ss *SubscriptionsService) Get(cID, sID string) (s *Subscription, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", cID, sID)
	req, err := ss.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
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
func (ss *SubscriptionsService) Create(cID string, sc *Subscription) (s *Subscription, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions", cID)

	if ss.client.HasAccessToken() && ss.client.config.testing {
		sc.TestMode = true
	}

	req, err := ss.client.NewAPIRequest(http.MethodPost, u, sc)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
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
func (ss *SubscriptionsService) Update(cID, sID string, sc *Subscription) (s *Subscription, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", cID, sID)

	req, err := ss.client.NewAPIRequest(http.MethodPatch, u, sc)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
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
func (ss *SubscriptionsService) Delete(cID, sID string) (s *Subscription, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s", cID, sID)
	req, err := ss.client.NewAPIRequest(http.MethodDelete, u, nil)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
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
func (ss *SubscriptionsService) All(options *SubscriptionListOptions) (sl *SubscriptionList, err error) {
	u := "v2/subscriptions"

	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	res, err := ss.list(u)
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
func (ss *SubscriptionsService) List(cID string, options *SubscriptionListOptions) (sl *SubscriptionList, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions", cID)

	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	res, err := ss.list(u)
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
func (ss *SubscriptionsService) GetPayments(cID, sID string, options *SubscriptionListOptions) (sl *PaymentList, err error) {
	u := fmt.Sprintf("v2/customers/%s/subscriptions/%s/payments", cID, sID)

	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	res, err := ss.list(u)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sl); err != nil {
		return
	}
	return
}

func (ss *SubscriptionsService) list(uri string) (r *Response, err error) {
	req, err := ss.client.NewAPIRequest(http.MethodGet, uri, nil)
	if err != nil {
		return
	}

	r, err = ss.client.Do(req)
	if err != nil {
		return
	}
	return
}
