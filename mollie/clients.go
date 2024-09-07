package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// LinkedClient describes a single client, linked to your partner account.
type LinkedClient struct {
	Resource              string            `json:"resource,omitempty"`
	ID                    string            `json:"id,omitempty"`
	OrganizationCreatedAt *time.Time        `json:"organizationCreatedAt,omitempty"`
	Links                 LinkedClientLinks `json:"_links,omitempty"`
}

// LinkedClientLinks contains URL objects relevant to the client.
type LinkedClientLinks struct {
	Self          *URL `json:"self,omitempty"`
	Organization  *URL `json:"organization,omitempty"`
	Onboarding    *URL `json:"onboarding,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// GetLinkedClientOptions contains valid query parameters for the get clients endpoint.
type GetLinkedClientOptions struct {
	Embed []EmbedValue `url:"embed,omitempty"`
}

// LinkedClientList describes a list of partner clients.
type LinkedClientList struct {
	Count          int `json:"count,omitempty"`
	PartnerClients struct {
		Clients []*LinkedClient `json:"clients,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// ListLinkedClientsOptions contains valid query parameters for the list clients endpoint.
type ListLinkedClientsOptions struct {
	Limit int          `url:"limit,omitempty"`
	From  string       `url:"from,omitempty"`
	Embed []EmbedValue `url:"embed,omitempty"`
}

// ClientsService operates over the partners API.
type ClientsService service

// List retrieves all clients.
//
// See: https://docs.mollie.com/reference/list-clients
func (ps *ClientsService) List(ctx context.Context, opts *ListLinkedClientsOptions) (
	res *Response,
	pc *LinkedClientList,
	err error,
) {
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
// See: https://docs.mollie.com/reference/get-client
func (ps *ClientsService) Get(ctx context.Context, id string, opts *GetLinkedClientOptions) (
	res *Response,
	pc *LinkedClient,
	err error,
) {
	res, err = ps.client.get(ctx, fmt.Sprintf("v2/clients/%s", id), opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pc); err != nil {
		return
	}

	return
}
