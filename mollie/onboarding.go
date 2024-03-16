package mollie

import (
	"context"
	"encoding/json"
	"time"
)

const onboardingURLPath = "v2/onboarding/me"

// OnboardingStatus describes status of the organization’s onboarding process.
type OnboardingStatus string

// Possible status values.
const (
	NeedsDataOnboardingStatus OnboardingStatus = "needs-data"
	InReviewOnboardingStatus  OnboardingStatus = "in-review"
	CompletedOnboardingStatus OnboardingStatus = "completed"
)

// OnboardingLinks contains URL objects relevant to the onboarding status.
type OnboardingLinks struct {
	Self          *URL `json:"self,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
	Organization  *URL `json:"organization,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Onboarding data for an organization.
type Onboarding struct {
	CanReceivePayments    bool             `json:"canReceivePayments,omitempty"`
	CanReceiveSettlements bool             `json:"canReceiveSettlements,omitempty"`
	Resource              string           `json:"reference,omitempty"`
	Name                  string           `json:"name,omitempty"`
	SignedUpAt            *time.Time       `json:"signedUpAt,omitempty"`
	Status                OnboardingStatus `json:"status,omitempty"`
	Links                 OnboardingLinks  `json:"_links,omitempty"`
}

// OnboardingData request possible values.
//
// Please note that even though all parameters are optional,
// at least one of them needs to be provided in the request.
//
// Information that the merchant has entered in their dashboard will not be overwritten.

// OnboardingDataOrganization contains data of the organization you want to provide.
type OnboardingDataOrganization struct {
	Name               string   `json:"name,omitempty"`
	RegistrationNumber string   `json:"registrationNumber,omitempty"`
	VatNumber          string   `json:"vatNumber,omitempty"`
	VatRegulation      string   `json:"vatRegulation,omitempty"`
	Address            *Address `json:"address,omitempty"`
}

// OnboardingDataProfile contains data of the payment profile you want to provide.
type OnboardingDataProfile struct {
	Name             string           `json:"name,omitempty"`
	URL              string           `json:"url,omitempty"`
	Email            string           `json:"email,omitempty"`
	Description      string           `json:"description,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	BusinessCategory BusinessCategory `json:"businessCategory,omitempty"`
}

// Full onboarding data to be submitted.
type OnboardingData struct {
	Organization OnboardingDataOrganization `json:"organization,omitempty"`
	Profile      OnboardingDataProfile      `json:"profile,omitempty"`
}

// OnboardingService operates over the onboarding API.
type OnboardingService service

// GetOnboardingStatus gets the status of onboarding of the authenticated organization.
//
// See: https://docs.mollie.com/reference/v2/onboarding-api/get-onboarding-status
func (os *OnboardingService) GetOnboardingStatus(ctx context.Context) (res *Response, o *Onboarding, err error) {
	res, err = os.client.get(ctx, onboardingURLPath, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &o); err != nil {
		return
	}

	return
}

// SubmitOnboardingData sends data that will be prefilled in the merchant’s onboarding.
// Please note that the data you submit will only be processed when the onboarding status is needs-data.
//
// This endpoint has been deprecated. It will be supported for the foreseeable future, but new implementations should
// use the Create client link endpoint to create new clients and submit their organization’s details in one go.
//
// See: https://docs.mollie.com/reference/v2/onboarding-api/submit-onboarding-data
func (os *OnboardingService) SubmitOnboardingData(ctx context.Context, d *OnboardingData) (res *Response, err error) {
	res, err = os.client.post(ctx, onboardingURLPath, d, nil)
	if err != nil {
		return
	}

	return
}
