package mollie

import (
	"context"
	"encoding/json"
	"fmt"
)

// PermissionGrant defines supported permissions.
type PermissionGrant string

// Available permission grants.
const (
	PaymentsRead       PermissionGrant = "payments.read"
	PaymentsWrite      PermissionGrant = "payments.write"
	RefundsRead        PermissionGrant = "refunds.read"
	RefundsWrite       PermissionGrant = "refunds.write"
	CustomersRead      PermissionGrant = "customers.read"
	CustomersWrite     PermissionGrant = "customers.write"
	MandatesRead       PermissionGrant = "mandates.read"
	MandatesWrite      PermissionGrant = "mandates.write"
	SubscriptionsRead  PermissionGrant = "subscriptions.read"
	SubscriptionsWrite PermissionGrant = "subscriptions.write"
	ProfilesRead       PermissionGrant = "profiles.read"
	ProfilesWrite      PermissionGrant = "profiles.write"
	InvoicesRead       PermissionGrant = "invoices.read"
	SettlementsRead    PermissionGrant = "settlements.read"
	OrdersRead         PermissionGrant = "orders.read"
	OrdersWrite        PermissionGrant = "orders.write"
	ShipmentsRead      PermissionGrant = "shipments.read"
	ShipmentsWrite     PermissionGrant = "shipments.write"
	OrganizationsRead  PermissionGrant = "organizations.read"
	OrganizationsWrite PermissionGrant = "organizations.write"
	OnboardingRead     PermissionGrant = "onboarding.read"
	OnboardingWrite    PermissionGrant = "onbording.write"
	PaymentLinksRead   PermissionGrant = "payment-links.read"
	PaymentLinksWrite  PermissionGrant = "payment-links.write"
	BalancesRead       PermissionGrant = "balances.read"
	TerminalsRead      PermissionGrant = "terminals.read"
	TerminalsWrite     PermissionGrant = "terminals.write"
)

// Permission represents an action that
// can be performed by any API actor.
type Permission struct {
	Granted     bool            `json:"granted,omitempty"`
	Resource    string          `json:"resource,omitempty"`
	Description string          `json:"description,omitempty"`
	ID          PermissionGrant `json:"id,omitempty"`
	Links       PermissionLinks `json:"_links,omitempty"`
}

// PermissionLinks contains URL objects that make
// reference to an http address related to permissions.
type PermissionLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// PermissionsList lists all the permissions given to an
// API actor.
type PermissionsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Permissions []*Permission `json:"permissions,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PermissionLinks `json:"_links,omitempty"`
}

// PermissionsService operates over permission resources.
type PermissionsService service

// Get returns a permission by its id.
//
// See: https://docs.mollie.com/reference/get-permission
func (ps *PermissionsService) Get(ctx context.Context, id PermissionGrant) (res *Response, p *Permission, err error) {
	res, err = ps.client.get(ctx, fmt.Sprintf("v2/permissions/%s", id), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}

	return
}

// List retrieves all permissions available with the current app access token.
// The list is not paginated.
//
// See: https://docs.mollie.com/reference/list-permissions
func (ps *PermissionsService) List(ctx context.Context) (res *Response, pl *PermissionsList, err error) {
	res, err = ps.client.get(ctx, "v2/permissions", nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}

	return
}
