package mollie

import (
	"encoding/json"
	"net/http"
	"time"
)

const onboardingTarget = "v2/onboarding/me"

// OnboardingStatus describes status of the organization’s onboarding process.
type OnboardingStatus string

// Possible status values.
const (
	NeedsDataOnboardingStatus OnboardingStatus = "needs-data"
	InReviewOnboardingStatus  OnboardingStatus = "in-review"
	CompletedOnboardingStatus OnboardingStatus = "completed"
)

// OnboardingService operates over the onboarding API.
type OnboardingService service

// OnboardingLinks contains URL objects relevant to the onboarding status.
type OnboardingLinks struct {
	Self          *URL `json:"self,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
	Organization  *URL `json:"organization,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Onboarding data for an organization.
type Onboarding struct {
	Resource              string           `json:"reference,omitempty"`
	Name                  string           `json:"name,omitempty"`
	SignedUpAt            *time.Time       `json:"signedUpAt,omitempty"`
	Status                OnboardingStatus `json:"status,omitempty"`
	CanReceivePayments    bool             `json:"canReceivePayments,omitempty"`
	CanReveiceSettlements bool             `json:"canReceiveSettlements,omitempty"`
	Links                 OnboardingLinks  `json:"_links,omitempty"`
}

// GetOnboardingStatus gets the status of onboarding of the authenticated organization.
//
// See: https://docs.mollie.com/reference/v2/onboarding-api/get-onboarding-status
func (os *OnboardingService) GetOnboardingStatus() (o *Onboarding, err error) {
	req, err := os.client.NewAPIRequest(http.MethodGet, onboardingTarget, nil)
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

// OnboardingData request possible values.
//
// Please note that even though all parameters are optional,
// at least one of them needs to be provided in the request.
//
// Information that the merchant has entered in their dashboard will not be overwritten.
type OnboardingData struct {
	Organization struct {
		Name               string   `json:"name,omitempty"`
		Address            *Address `json:"address,omitempty"`
		RegistrationNumber string   `json:"registrationNumber,omitempty"`
		VatNumber          string   `json:"vatNumber,omitempty"`
		VatRegulation      string   `json:"vatRegulation,omitempty"`
	} `json:"organization,omitempty"`
	Profile struct {
		Name         string       `json:"name,omitempty"`
		URL          string       `json:"url,omitempty"`
		Email        string       `json:"email,omitempty"`
		Description  string       `json:"description,omitempty"`
		Phone        string       `json:"phone,omitempty"`
		CategoryCode CategoryCode `json:"categoryCode,omitempty"`
	} `json:"profile,omitempty"`
}

// SubmitOnboardingData sends data that will be prefilled in the merchant’s onboarding.
// Please note that the data you submit will only be processed when the onboarding status is needs-data.
//
// See: https://docs.mollie.com/reference/v2/onboarding-api/submit-onboarding-data
func (os *OnboardingService) SubmitOnboardingData(d *OnboardingData) (err error) {
	req, err := os.client.NewAPIRequest(http.MethodPost, onboardingTarget, d)
	if err != nil {
		return
	}

	_, err = os.client.Do(req)
	if err != nil {
		return
	}

	return
}
