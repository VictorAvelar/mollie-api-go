package mollie

import (
	"context"
	"encoding/json"
	"fmt"
)

// PaymentMethodStatus tels the status that the method is in.
// Possible values: activated pending-boarding pending-review
// pending-external rejected.
type PaymentMethodStatus string

// Available payment method statuses.
const (
	PaymentMethodActivated       PaymentMethodStatus = "activated"
	PaymentMethodPendingBoarding PaymentMethodStatus = "pending-boarding"
	PaymentMethodPendingReview   PaymentMethodStatus = "pending-review"
	PaymentMethodPendingExternal PaymentMethodStatus = "pending-external"
	PaymentMethodRejected        PaymentMethodStatus = "pending-rejected"
)

// PaymentMethodDetails describes a single method with details.
type PaymentMethodDetails struct {
	Resource      string                  `json:"resource,omitempty"`
	ID            string                  `json:"id,omitempty"`
	Description   string                  `json:"description,omitempty"`
	MinimumAmount *Amount                 `json:"minimumAmount,omitempty"`
	MaximumAmount *Amount                 `json:"maximumAmount,omitempty"`
	Image         *Image                  `json:"image,omitempty"`
	Pricing       []*PaymentMethodPricing `json:"pricing,omitempty"`
	Issuers       []*PaymentMethodIssuer  `json:"issuers,omitempty"`
	Status        *PaymentMethodStatus    `json:"status,omitempty"`
	Links         MethodsLinks            `json:"_links,omitempty"`
}

// MethodsLinks describes links attached to methods service responses.
type MethodsLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// PaymentMethodPricing contains information about commissions and fees
// applicable to a payment method.
type PaymentMethodPricing struct {
	Description string    `json:"description,omitempty"`
	Variable    string    `json:"variable,omitempty"`
	Fixed       *Amount   `json:"fixed,omitempty"`
	FeeRegion   FeeRegion `json:"feeRegion,omitempty"`
}

// PaymentMethodIssuer available for the payment method
// (for iDEAL, KBC/CBC payment button, gift cards, or meal vouchers).
type PaymentMethodIssuer struct {
	Resource string `json:"resource,omitempty"`
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Image    *Image `json:"image,omitempty"`
}

// PaymentMethodsList describes a list of paginated payment methods.
type PaymentMethodsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Methods []*PaymentMethodDetails
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// PaymentMethodOptions are applicable query string parameters to get methods
// from mollie's API.
type PaymentMethodOptions struct {
	Locale    Locale         `url:"locale,omitempty"`
	Currency  string         `url:"currency,omitempty"`
	ProfileID string         `url:"profileId,omitempty"`
	Include   []IncludeValue `url:"include,omitempty"`
}

// PaymentMethodsListOptions are applicable query string parameters to list methods
// from mollie's API.
//
// It contains list specific options and embeds GetMethodOptions.
type PaymentMethodsListOptions struct {
	PaymentMethodOptions
	Resource            string                              `url:"resource,omitempty"`
	BillingCountry      string                              `url:"billingCountry,omitempty"`
	Amount              *Amount                             `url:"amount,omitempty"`
	IncludeWallets      []Wallet                            `url:"includeWallets,omitempty"`
	OrderLineCategories []OrderLineOperationProductCategory `url:"orderLineCategories,omitempty"`
	Locale              Locale                              `url:"locale,omitempty"`
	SequenceType        SequenceType                        `url:"sequenceType,omitempty"`
}

// PaymentMethodsService operates on methods endpoints.
type PaymentMethodsService service

// Get returns information about the payment method specified by id,
// it also receives a pointer to the method options containing applicable
// query string parameters.
//
// See: https://docs.mollie.com/reference/v2/methods-api/get-method
func (ms *PaymentMethodsService) Get(ctx context.Context, id PaymentMethod, options *PaymentMethodOptions) (
	res *Response,
	pmd *PaymentMethodDetails,
	err error,
) {
	u := fmt.Sprintf("v2/methods/%s", id)

	res, err = ms.client.get(ctx, u, options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pmd); err != nil {
		return
	}

	return
}

// All retrieves all the payment methods enabled for your account/organization.
//
// See: https://docs.mollie.com/reference/v2/methods-api/list-all-methods
func (ms *PaymentMethodsService) All(ctx context.Context, options *PaymentMethodsListOptions) (
	res *Response,
	pm *PaymentMethodsList,
	err error,
) {
	return ms.list(ctx, "v2/methods/all", options)
}

// List retrieves all enabled payment methods.
//
// The results are not paginated.
//
// See: https://docs.mollie.com/reference/v2/methods-api/list-methods
func (ms *PaymentMethodsService) List(ctx context.Context, options *PaymentMethodsListOptions) (
	res *Response,
	pm *PaymentMethodsList,
	err error,
) {
	return ms.list(ctx, "v2/methods", options)
}

func (ms *PaymentMethodsService) list(ctx context.Context, uri string, options interface{}) (
	res *Response,
	pm *PaymentMethodsList,
	err error,
) {
	res, err = ms.client.get(ctx, uri, options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pm); err != nil {
		return
	}

	return
}
