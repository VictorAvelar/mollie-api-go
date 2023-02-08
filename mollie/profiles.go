package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ProfileStatus determines whether the profile is able to receive live payments.
type ProfileStatus string

// Possible profile statuses.
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

// ProfileListOptions are optional query string parameters for the list profiles request.
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

// ProfilesService operates over profile resource.
type ProfilesService service

// List returns all the profiles for the authenticated account.
func (ps *ProfilesService) List(ctx context.Context, opts *ProfileListOptions) (res *Response, pl *ProfileList, err error) {
	res, err = ps.client.get(ctx, "v2/profiles", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}

// Get retrieves the a profile by ID.
func (ps *ProfilesService) Get(ctx context.Context, id string) (res *Response, p *Profile, err error) {
	return ps.get(ctx, id)
}

// Current returns the profile belonging to the API key.
// This method only works when using API keys.
func (ps *ProfilesService) Current(ctx context.Context) (res *Response, p *Profile, err error) {
	return ps.get(ctx, "me")
}

func (ps *ProfilesService) get(ctx context.Context, id string) (res *Response, p *Profile, err error) {
	res, err = ps.client.get(ctx, fmt.Sprintf("v2/profiles/%s", id), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}
	return
}

// Create stores a new profile in your Mollie account.
func (ps *ProfilesService) Create(ctx context.Context, np *Profile) (res *Response, p *Profile, err error) {
	res, err = ps.client.post(ctx, "v2/profiles", np, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}
	return
}

// Update allows you to perform mutations on a profile.
func (ps *ProfilesService) Update(ctx context.Context, id string, up *Profile) (res *Response, p *Profile, err error) {
	res, err = ps.client.patch(ctx, fmt.Sprintf("v2/profiles/%s", id), up, nil)
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
func (ps *ProfilesService) Delete(ctx context.Context, id string) (res *Response, err error) {
	res, err = ps.client.delete(ctx, fmt.Sprintf("v2/profiles/%s", id), nil)
	if err != nil {
		return
	}
	return
}

// EnablePaymentMethod enables a payment method on a specific or authenticated profile.
// If you're using API tokens for authentication, pass "me" as id.
func (ps *ProfilesService) EnablePaymentMethod(ctx context.Context, id string, pm PaymentMethod) (res *Response, pmi *PaymentMethodDetails, err error) {
	res, err = ps.client.post(ctx, fmt.Sprintf("v2/profiles/%s/methods/%s", id, pm), nil, nil)
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
func (ps *ProfilesService) DisablePaymentMethod(ctx context.Context, id string, pm PaymentMethod) (res *Response, err error) {
	res, err = ps.client.delete(ctx, fmt.Sprintf("v2/profiles/%s/methods/%s", id, pm), nil)
	if err != nil {
		return
	}
	return
}

// EnableGiftCardIssuer activates the requested gift card issuer for the provided
// profile id.
//
// See: https://docs.mollie.com/reference/v2/profiles-api/enable-gift-card-issuer
func (ps *ProfilesService) EnableGiftCardIssuer(ctx context.Context, profileID string, issuer GiftCardIssuer) (res *Response, gc *GiftCardEnabled, err error) {
	res, err = ps.toggleGiftCardIssuerStatus(ctx, profileID, http.MethodPost, issuer)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &gc); err != nil {
		return
	}
	return
}

// DisableGiftCardIssuer deactivates the requested gift card issuer for the provided
// profile id.
//
// See: https://docs.mollie.com/reference/v2/profiles-api/disable-gift-card-issuer
func (ps *ProfilesService) DisableGiftCardIssuer(ctx context.Context, profileID string, issuer GiftCardIssuer) (res *Response, err error) {
	res, err = ps.toggleGiftCardIssuerStatus(ctx, profileID, http.MethodDelete, issuer)
	if err != nil {
		return
	}
	return
}

// EnableGiftCardIssuerForCurrent activates the specified issuer for the
// current profile (token owner).
//
// See: https://docs.mollie.com/reference/v2/profiles-api/enable-gift-card-issuer
func (ps *ProfilesService) EnableGiftCardIssuerForCurrent(ctx context.Context, issuer GiftCardIssuer) (res *Response, gc *GiftCardEnabled, err error) {
	res, err = ps.toggleGiftCardIssuerStatus(ctx, "me", http.MethodPost, issuer)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &gc); err != nil {
		return
	}
	return
}

// DisableGiftCardIssuerForCurrent deactivates the specified issuer for the
// current profile (token owner).
//
// See: https://docs.mollie.com/reference/v2/profiles-api/disable-gift-card-issuer
func (ps *ProfilesService) DisableGiftCardIssuerForCurrent(ctx context.Context, issuer GiftCardIssuer) (res *Response, err error) {
	res, err = ps.toggleGiftCardIssuerStatus(ctx, "me", http.MethodDelete, issuer)
	if err != nil {
		return
	}
	return
}

func (ps *ProfilesService) toggleGiftCardIssuerStatus(ctx context.Context, profileID string, method string, issuer GiftCardIssuer) (r *Response, err error) {
	u := fmt.Sprintf("v2/profiles/%s/methods/giftcard/issuers/%s", profileID, issuer)

	switch method {
	case http.MethodDelete:
		r, err = ps.client.delete(ctx, u, nil)
	case http.MethodPost:
		r, err = ps.client.post(ctx, u, nil, nil)
	}

	if err != nil {
		return
	}

	return
}
