package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// PaymentMethodInfo describes a single method with details.
type PaymentMethodInfo struct {
	Resource      string                 `json:"resource,omitempty"`
	ID            string                 `json:"id,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MinimumAmount *Amount                `json:"minimumAmount,omitempty"`
	MaximumAmount *Amount                `json:"maximumAmount,omitempty"`
	Image         Image                  `json:"image,omitempty"`
	Pricing       []PaymentMethodPricing `json:"pricing,omitempty"`
	Links         MethodsLinks           `json:"_links,omitempty"`
}

// MethodsLinks describes links attached to methods service responses.
type MethodsLinks struct {
	Self URL `json:"self,omitempty"`
	Docs URL `json:"documentation,omitempty"`
}

// Image describes a generic image resource retrieved by Mollie.
type Image struct {
	Size1x string `json:"size1X,omitempty"`
	Size2X string `json:"size2X,omitempty"`
	Svg    string `json:"svg,omitempty"`
}

// PaymentMethodPricing contains information about commissions and fees
// applicable to a payment method.
type PaymentMethodPricing struct {
	Description string  `json:"description,omitempty"`
	Fixed       *Amount `json:"fixed,omitempty"`
	Variable    string  `json:"variable,omitempty"`
}

// ListMethods describes a list of paginated payment methods.
type ListMethods struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Methods []PaymentMethodInfo
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// MethodsOptions are applicable query string parameters to methods
// service endpoints.
type MethodsOptions struct {
	Locale    Locale `url:"locale,omitempty"`
	Currency  string `url:"currency,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
	Include   string `url:"include,omitempty"`
	// Use for List method only
	SequenceType   SequenceType `url:"sequenceType,omitempty"`
	Amount         Amount       `url:"amount,omitempty"`
	Resource       string       `url:"resource,omitempty"`
	BillingCountry string       `url:"billingCountry,omitempty"`
	IncludeWallets string       `url:"includeWallets,omitempty"`
}

// MethodsService operates on methods endpoints
type MethodsService service

// Get returns information about the payment method specified by id,
// it also receives a pointer to the method options containing applicable
// query string parameters
//
// See: https://docs.mollie.com/reference/v2/methods-api/get-method
func (ms *MethodsService) Get(id string, options *MethodsOptions) (pmi *PaymentMethodInfo, err error) {
	u := fmt.Sprintf("v2/methods/%s", id)
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}
	req, err := ms.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}
	res, err := ms.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pmi); err != nil {
		return
	}
	return
}

// All retrieves all the payment methods enabled for your account/organization
//
// See: https://docs.mollie.com/reference/v2/methods-api/list-all-methods
func (ms *MethodsService) All(options *MethodsOptions) (pm *ListMethods, err error) {
	u := "v2/methods/all"
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	return ms.list(u)
}

// List retrieves all enabled payment methods.
// The results are not paginated.
//
// See: https://docs.mollie.com/reference/v2/methods-api/list-methods
func (ms *MethodsService) List(options *MethodsOptions) (pm *ListMethods, err error) {
	u := "v2/methods"
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	return ms.list(u)
}

func (ms *MethodsService) list(uri string) (pm *ListMethods, err error) {
	req, err := ms.client.NewAPIRequest(http.MethodGet, uri, nil)
	if err != nil {
		return
	}
	res, err := ms.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pm); err != nil {
		return
	}
	return
}
