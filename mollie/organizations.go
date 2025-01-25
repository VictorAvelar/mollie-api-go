package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Organization describes an organization detail.
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

// PartnerType alias for organization partner types.
type PartnerType string

// Available partner types.
const (
	PartnerTypeOauth      PartnerType = "oauth"
	PartnerTypeSignUpLink PartnerType = "signuplink"
	PartnerTypeUserAgent  PartnerType = "useragent"
)

// UserAgentToken are time limited valid access tokens.
type UserAgentToken struct {
	Token    string
	StartsAt *time.Time
	EndsAt   *time.Time
}

// OrganizationPartnerLinks is an object with several URL objects
// relevant to the partner resource.
type OrganizationPartnerLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	SignUpLink    *URL `json:"signuplink,omitempty"`
}

// OrganizationPartnerStatus response descriptor.
type OrganizationPartnerStatus struct {
	IsCommissionPartner            bool                     `json:"isCommissionPartner,omitempty"`
	PartnerContractUpdateAvailable bool                     `json:"partnerContractUpdate_available,omitempty"`
	Resource                       string                   `json:"resource,omitempty"`
	PartnerType                    PartnerType              `json:"partnerType,omitempty"`
	UserAgentTokens                []*UserAgentToken        `json:"userAgentTokens,omitempty"`
	PartnerContractSignedAt        *time.Time               `json:"partnerContractSignedAt,omitempty"`
	PartnerContractExpiresAt       *time.Time               `json:"partnerContractExpiresAt,omitempty"`
	Links                          OrganizationPartnerLinks `json:"_links,omitempty"`
}

// OrganizationsService instance operates over organization resources.
type OrganizationsService service

// Get retrieve an organization by its id.
func (os *OrganizationsService) Get(ctx context.Context, id string) (res *Response, o *Organization, err error) {
	return os.get(ctx, fmt.Sprintf("v2/organizations/%s", id))
}

// GetCurrent retrieve the currently authenticated organization.
func (os *OrganizationsService) GetCurrent(ctx context.Context) (res *Response, o *Organization, err error) {
	return os.get(ctx, "v2/organizations/me")
}

// GetPartnerStatus retrieves details about the partner status
// of the currently authenticated organization.
//
// See: https://docs.mollie.com/reference/get-partner-status
func (os *OrganizationsService) GetPartnerStatus(ctx context.Context) (
	res *Response,
	ops *OrganizationPartnerStatus,
	err error,
) {
	res, err = os.client.get(ctx, "v2/organizations/me/partner", nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &ops); err != nil {
		return
	}

	return
}

func (os *OrganizationsService) get(ctx context.Context, uri string) (res *Response, o *Organization, err error) {
	res, err = os.client.get(ctx, uri, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &o); err != nil {
		return
	}

	return
}
