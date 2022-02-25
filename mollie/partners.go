package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// PartnerClient describes a partner client.
type PartnerClient struct {
	Resource              string             `json:"resource,omitempty"`
	ID                    string             `json:"id,omitempty"`
	OrganizationCreatedAt *time.Time         `json:"organizationCreatedAt,omitempty"`
	Commission            Commission         `json:"commission,omitempty"`
	Links                 PartnerClientLinks `json:"_links,omitempty"`
}

// PartnerClientList describes a list of partner clients.
type PartnerClientList struct {
	Count          int `json:"count,omitempty"`
	PartnerClients struct {
		Clients []*PartnerClient `json:"clients,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// PartnerClientLinks contains URL objects relevant to the client.
type PartnerClientLinks struct {
	Self          *URL `json:"self,omitempty"`
	Organization  *URL `json:"organization,omitempty"`
	Onboarding    *URL `json:"onboarding,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Commission describes a partner take from any
// operation on Mollie's API.
type Commission struct {
	Count       int     `json:"count,omitempty"`
	TotalAmount *Amount `json:"totalAmount,omitempty"`
}

// GetPartnerClientOptions contains valid query parameters for the get clients endpoint.
type GetPartnerClientOptions struct {
	Embed string `url:"embed,omitempty"`
}

// ListPartnerClientsOptions contains valid query parameters for the list clients endpoint.
type ListPartnerClientsOptions struct {
	From  int `url:"from,omitempty"`
	Limit int `url:"limit,omitempty"`
	Year  int `url:"year,omitempty"`
	Month int `url:"month,omitempty"`
}

// PartnerService operates over the partners API.
type PartnerService service

// List retrieves all clients.
//
// See: https://docs.mollie.com/reference/v2/partners-api/list-clients
func (ps *PartnerService) List(ctx context.Context, opts *ListPartnerClientsOptions) (res *Response, pc *PartnerClientList, err error) {
	res, err = ps.client.get(ctx, "v2/clients", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pc); err != nil {
		return
	}

	return
}

// Get retrieves a single client, linked to your partner account, by its ID.
//
// See: https://docs.mollie.com/reference/v2/partners-api/get-client
func (ps *PartnerService) Get(ctx context.Context, id string, opts *GetPartnerClientOptions) (res *Response, pc *PartnerClient, err error) {
	res, err = ps.client.get(ctx, fmt.Sprintf("v2/clients/%s", id), opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pc); err != nil {
		return
	}

	return
}
