package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// DelayedRoutingDestinationType indicates the destination of the delayed routing.
type DelayedRoutingDestinationType string

// Possible values for DelayedRoutingDestination.
const (
	DelayedRoutingDestinationOrganization DelayedRoutingDestinationType = "organization"
)

// DelayedRoutingDestination represents the destination for delayed routing.
type DelayedRoutingDestination struct {
	Type           DelayedRoutingDestinationType `json:"type,omitempty"`
	OrganizationID string                        `json:"organizationId,omitempty"`
}

// Route represents a delayed routing associated with a payment.
type Route struct {
	Resource    string                    `json:"resource,omitempty"`
	ID          string                    `json:"id,omitempty"`
	PaymentID   string                    `json:"paymentId,omitempty"`
	Description string                    `json:"description,omitempty"`
	Amount      Amount                    `json:"amount,omitempty"`
	Destination DelayedRoutingDestination `json:"destination,omitempty"`
	Links       RouteLinks                `json:"_links,omitempty"`
	CreatedAt   *time.Time                `json:"createdAt,omitempty"`
}

// RouteLinks represents the links related to a delayed routing.
type RouteLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// PaymentRoutesList represents a list of delayed routings for a payment.
type PaymentRoutesList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Routes []Route `json:"routes,omitempty"`
	} `json:"_embedded,omitempty"`
	Links RouteLinks `json:"_links,omitempty"`
}

// CreateDelayedRouting represents the payload to create a delayed routing.
type CreateDelayedRouting struct {
	TestMode    bool                      `json:"testmode,omitempty"`
	Description string                    `json:"description,omitempty"`
	Amount      Amount                    `json:"amount,omitempty"`
	Destination DelayedRoutingDestination `json:"destination,omitempty"`
}

// DelayedRoutingService handles delayed routing related operations.
type DelayedRoutingService service

// Create creates a new delayed routing for a payment.
//
// See: https://docs.mollie.com/reference/payment-create-route
func (s *DelayedRoutingService) Create(ctx context.Context, payment string, dr CreateDelayedRouting) (
	res *Response,
	r *Route,
	err error,
) {
	if s.client.HasAccessToken() && s.client.config.testing {
		dr.TestMode = true
	}

	u := fmt.Sprintf("/v2/payments/%s/routes", payment)

	res, err = s.client.post(ctx, u, dr, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &r); err != nil {
		return
	}

	return
}

// List retrieves all delayed routings for a specific payment.
//
// See: https://docs.mollie.com/reference/payment-list-routes
func (s *DelayedRoutingService) List(ctx context.Context, payment string) (
	res *Response,
	prl *PaymentRoutesList,
	err error,
) {
	u := fmt.Sprintf("/v2/payments/%s/routes", payment)

	res, err = s.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &prl); err != nil {
		return
	}

	return
}
