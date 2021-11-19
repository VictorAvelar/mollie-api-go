package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// ProfileStatus determines whether the profile is able to receive live payments
type ProfileStatus string

// Possible profile statuses
const (
	StatusUnverified ProfileStatus = "unverified"
	StatusVerified   ProfileStatus = "verified"
	StatusBlocked    ProfileStatus = "blocked"
)

// Profile will usually reflect the trademark or brand name
// of the profile’s website or application.
type Profile struct {
	ID           string       `json:"id,omitempty"`
	CategoryCode CategoryCode `json:"categoryCode,omitempty"`
	CreatedAt    *time.Time   `json:"createdAt,omitempty"`
	Email        string       `json:"email,omitempty"`
	Mode         Mode         `json:"mode,omitempty"`
	Name         string       `json:"name,omitempty"`
	Phone        PhoneNumber  `json:"phone,omitempty"`
	Resource     string       `json:"resource,omitempty"`
	Review       struct {
		Status string `json:"status,omitempty"`
	} `json:"review,omitempty"`
	Status  ProfileStatus `json:"status,omitempty"`
	Website string        `json:"website,omitempty"`
	Links   ProfileLinks  `json:"_links,omitempty"`
}

// ProfileLinks contains URL's to relevant information related to
// a profile.
type ProfileLinks struct {
	Self               *URL `json:"self,omitempty"`
	Chargebacks        *URL `json:"chargebacks,omitempty"`
	Methods            *URL `json:"methods,omitempty"`
	Refunds            *URL `json:"refunds,omitempty"`
	CheckoutPreviewURL *URL `json:"checkoutPreviewUrl,omitempty"`
	Documentation      *URL `json:"documentation,omitempty"`
	Dashboard          *URL `json:"dashboard,omitempty"`
}

// ProfileListOptions are optional query string parameters for the list profiles request
type ProfileListOptions struct {
	From  string `url:"from,omitempty"`
	Limit uint   `url:"limit,omitempty"`
}

// ProfileList contains a list of profiles for your account.
type ProfileList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Profiles []*Profile `json:"profiles,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// ProfilesService operates over profile resource
type ProfilesService service

// List returns all the profiles for the authenticated account
func (ps *ProfilesService) List(ctx context.Context, options *ProfileListOptions) (pl *ProfileList, err error) {
	u := "v2/profiles"
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}
	req, err := ps.client.NewAPIRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}

// Get retrieves the a profile by ID.
func (ps *ProfilesService) Get(ctx context.Context, id string) (p *Profile, err error) {
	return ps.get(ctx, id)
}

// Current returns the profile belonging to the API key.
// This method only works when using API keys.
func (ps *ProfilesService) Current(ctx context.Context) (p *Profile, err error) {
	return ps.get(ctx, "me")
}

func (ps *ProfilesService) get(ctx context.Context, id string) (p *Profile, err error) {
	u := fmt.Sprintf("v2/profiles/%s", id)
	req, err := ps.client.NewAPIRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}
	return
}

// Create stores a new profile in your Mollie account.
func (ps *ProfilesService) Create(ctx context.Context, np *Profile) (p *Profile, err error) {
	req, err := ps.client.NewAPIRequest(ctx, http.MethodPost, "v2/profiles", np)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}
	return
}

// Update allows you to perform mutations on a profile
func (ps *ProfilesService) Update(ctx context.Context, id string, up *Profile) (p *Profile, err error) {
	u := fmt.Sprintf("v2/profiles/%s", id)
	req, err := ps.client.NewAPIRequest(ctx, http.MethodPatch, u, up)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}
	return
}

// Delete  enables profile deletions, rendering the profile unavailable
// for further API calls and transactions.
func (ps *ProfilesService) Delete(ctx context.Context, id string) (err error) {
	u := fmt.Sprintf("v2/profiles/%s", id)
	req, err := ps.client.NewAPIRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return
	}
	_, err = ps.client.Do(req)
	if err != nil {
		return
	}
	return
}

// EnablePaymentMethod enables a payment method on a specific or authenticated profile.
// If you're using API tokens for authentication, pass "me" as id.
func (ps *ProfilesService) EnablePaymentMethod(ctx context.Context, id string, pm PaymentMethod) (pmi *PaymentMethodInfo, err error) {
	u := fmt.Sprintf("v2/profiles/%s/methods/%s", id, pm)
	req, err := ps.client.NewAPIRequest(ctx, http.MethodPost, u, nil)
	if err != nil {
		return
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pmi); err != nil {
		return
	}
	return
}

// DisablePaymentMethod disables a payment method on a specific or authenticated profile.
// If you're using API tokens for authentication, pass "me" as id.
func (ps *ProfilesService) DisablePaymentMethod(ctx context.Context, id string, pm PaymentMethod) (err error) {
	u := fmt.Sprintf("v2/profiles/%s/methods/%s", id, pm)
	req, err := ps.client.NewAPIRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return
	}
	_, err = ps.client.Do(req)
	if err != nil {
		return
	}
	return
}

// EnableGiftCardIssuer activates the requested giftcard issuer for the provided
// profile id.
//
// See: https://docs.mollie.com/reference/v2/profiles-api/enable-gift-card-issuer
func (ps *ProfilesService) EnableGiftCardIssuer(ctx context.Context, profileID string, issuer GiftCardIssuer) (gc *GiftCardEnabled, err error) {
	res, err := ps.toggleGiftCardIssuerStatus(ctx, profileID, http.MethodPost, issuer)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &gc); err != nil {
		return
	}
	return
}

// DisableGiftCardIssuer deactivates the requested giftcard issuer for the provided
// profile id.
//
// See: https://docs.mollie.com/reference/v2/profiles-api/disable-gift-card-issuer
func (ps *ProfilesService) DisableGiftCardIssuer(ctx context.Context, profileID string, issuer GiftCardIssuer) (err error) {
	_, err = ps.toggleGiftCardIssuerStatus(ctx, profileID, http.MethodDelete, issuer)
	if err != nil {
		return
	}
	return
}

// EnableGiftCardIssuerForCurrent activates the specified issuer for the
// curent profile (token owner).
//
// See: https://docs.mollie.com/reference/v2/profiles-api/enable-gift-card-issuer
func (ps *ProfilesService) EnableGiftCardIssuerForCurrent(ctx context.Context, issuer GiftCardIssuer) (gc *GiftCardEnabled, err error) {
	res, err := ps.toggleGiftCardIssuerStatus(ctx, "me", http.MethodPost, issuer)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &gc); err != nil {
		return
	}
	return
}

// DisableGiftCardIssuerForCurrent deactivates the specified issuer for the
// curent profile (token owner).
//
// See: https://docs.mollie.com/reference/v2/profiles-api/disable-gift-card-issuer
func (ps *ProfilesService) DisableGiftCardIssuerForCurrent(ctx context.Context, issuer GiftCardIssuer) (err error) {
	_, err = ps.toggleGiftCardIssuerStatus(ctx, "me", http.MethodDelete, issuer)
	if err != nil {
		return
	}
	return
}

func (ps *ProfilesService) toggleGiftCardIssuerStatus(ctx context.Context, profileID string, method string, issuer GiftCardIssuer) (r *Response, err error) {
	u := fmt.Sprintf("v2/profiles/%s/methods/giftcard/issuers/%s", profileID, issuer)
	req, err := ps.client.NewAPIRequest(ctx, method, u, nil)
	if err != nil {
		return
	}

	r, err = ps.client.Do(req)
	if err != nil {
		return
	}
	return
}
