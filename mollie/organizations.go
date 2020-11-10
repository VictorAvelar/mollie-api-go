package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Organization describes an organization detail
type Organization struct {
	Resource           string            `json:"resource,omitempty"`
	ID                 string            `json:"id,omitempty"`
	Name               string            `json:"name,omitempty"`
	Email              string            `json:"email,omitempty"`
	Locale             string            `json:"locale,omitempty"`
	Address            *Address          `json:"address,omitempty"`
	RegistrationNumber string            `json:"registrationNumber,omitempty"`
	VatNumber          string            `json:"vatNumber,omitempty"`
	VatRegulation      string            `json:"vatRegulation,omitempty"`
	Links              OrganizationLinks `json:"_links,omitempty"`
}

// OrganizationLinks describes all the possible links to be returned with
// a organization object.
type OrganizationLinks struct {
	Self          *URL `json:"self,omitempty"`
	Chargebacks   *URL `json:"chargebacks,omitempty"`
	Customers     *URL `json:"customers,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
	Invoices      *URL `json:"invoices,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Profiles      *URL `json:"profiles,omitempty"`
	Refunds       *URL `json:"refunds,omitempty"`
	Settlements   *URL `json:"settlements,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// OrganizationsService instance operates over organization resources
type OrganizationsService service

// Get retrieve an organization by its id.
func (os *OrganizationsService) Get(id string) (o *Organization, err error) {
	return os.get(fmt.Sprintf("v2/organizations/%s", id))
}

// GetCurrent retrieve the currently authenticated organization
func (os *OrganizationsService) GetCurrent() (o *Organization, err error) {
	return os.get("v2/organizations/me")
}

func (os *OrganizationsService) get(uri string) (o *Organization, err error) {
	req, err := os.client.NewAPIRequest(http.MethodGet, uri, nil)
	if err != nil {
		return
	}
	res, err := os.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &o); err != nil {
		return
	}
	return
}
