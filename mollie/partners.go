package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
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

// Get retrieves all clients.
//
// See: https://docs.mollie.com/reference/v2/partners-api/list-clients
func (ps *PartnerService) List(opts *ListPartnerClientsOptions) (pc *PartnerClientList, err error) {
	u := "v2/clients"
	if opts != nil {
		v, _ := query.Values(opts)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	req, err := ps.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pc); err != nil {
		return
	}
	return
}

// Get retrieves a single client by its ID.
//
// See: https://docs.mollie.com/reference/v2/partners-api/get-client
func (ps *PartnerService) Get(id string, opts *GetPartnerClientOptions) (pc *PartnerClient, err error) {
	u := fmt.Sprintf("v2/clients/%s", id)
	if opts != nil {
		v, _ := query.Values(opts)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	req, err := ps.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pc); err != nil {
		return
	}
	return
}
