package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type CustomersService service

type CustomerReqOptions struct {
	Testmode bool `json:"testmode,omitempty"`
}

type Customer struct {
	Resource  string        `json:"resource,omitempty"`
	ID        string        `json:"id,omitempty"`
	Mode      Mode          `json:"mode,omitempty"`
	Name      string        `json:"name,omitempty"`
	Email     string        `json:"email,omitempty"`
	Metadata  interface{}   `json:"metadata,omitempty"`
	Locale    *Locale       `json:"locale,omitempty"`
	CreatedAt *time.Time    `json:"createdAt,omitempty"`
	Links     *PaymentLinks `json:"_links,omitempty"`
}

// Get retrieve a single customer by its ID.
func (sv *CustomersService) Get(ID string) (c Customer, err error) {
	path := fmt.Sprintf("/v2/customers/%s", ID)

	req, err := sv.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return
	}

	res, err := sv.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &c); err != nil {
		return
	}
	return
}

// CreateCustomerReq represents the body of CreateCustomer endpoint
// All attributes are optional
type CreateCustomerReq struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Locale   *Locale     `json:"locale"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// Create creates a simple minimal representation of a customer
// in the Mollie API to use for the Mollie Checkout and Recurring features.
// These customers will appear in your Mollie Dashboard where you can manage
// their details, and also see their payments and subscriptions.
func (sv *CustomersService) Create(reqBody CreateCustomerReq) (c Customer, err error) {
	path := "/v2/customers"

	req, err := sv.client.NewAPIRequest(http.MethodPost, path, reqBody)
	if err != nil {
		return
	}

	res, err := sv.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &c); err != nil {
		return
	}

	return
}

// UpdateCutomerReq represents the body of the update request.
// Its structure is currently identical to CreateCustomerReq
type UpdateCustomerReq CreateCustomerReq

// Update updates an existing customer
func (sv *CustomersService) Update(ID string, reqBody UpdateCustomerReq) (c Customer, err error) {
	path := fmt.Sprintf("/v2/customers/%s", ID)

	req, err := sv.client.NewAPIRequest(http.MethodPatch, path, reqBody)
	if err != nil {
		return
	}

	res, err := sv.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &c); err != nil {
		return
	}

	return
}

// Delete retrieve a single customer by its ID.
func (sv *CustomersService) Delete(ID string) (err error) {
	path := fmt.Sprintf("/v2/customers/%s", ID)

	req, err := sv.client.NewAPIRequest(http.MethodDelete, path, nil)
	if err != nil {
		return
	}

	_, err = sv.client.Do(req)

	return
}

// ListCustomerOptions represents the available options
// From offsets the result set to the customer with this ID.
//   The customer with this ID is included in the result set as well.
// Limit limits the number of customers to return (with a maximum of 250).
type ListCustomerOptions struct {
	From  string `json:"list,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

//
type CustomersList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Customers []Customer
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// List retrieves all customers created.
func (sv *CustomersService) List(options *ListCustomerOptions) (cl CustomersList, err error) {
	path := "v2/customers"
	if options != nil {
		v, _ := query.Values(options)
		path = fmt.Sprintf("%s?%s", path, v.Encode())
	}
	req, err := sv.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return
	}
	res, err := sv.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}

type CreateCustomerPaymentReq struct {
	Payment
	SequenceType string `json:"sequenceType,omitempty"`
	RedirectUrl  string `json:"redirectUrl, omitempty"`
}

type CustomerPaymentOptions struct {
	ProfileID string `json:"profileID"`
}

var (
	requiredCreateParam = "parameter required for creating a payment: %+v"
)

// CreatePayment creates a payment for the customer
func (sv *CustomersService) CreatePayment(reqBody *CreateCustomerPaymentReq, options *CustomerPaymentOptions) (p Payment, err error) {
	path := fmt.Sprintf("v2/customers/%s/payments")
	if options.ProfileID == "" {
		return p, fmt.Errorf(requiredCreateParam, options.ProfileID)
	}
	if options != nil {
		v, _ := query.Values(options)
		path = fmt.Sprintf("%s?%s", path, v.Encode())
	}
	if reqBody.Payment.Amount == nil {
		return p, fmt.Errorf(requiredCreateParam, reqBody.Payment.Amount)
	}
	if reqBody.Payment.Description == "" {
		return p, fmt.Errorf(requiredCreateParam, reqBody.Payment.Description)
	}
	req, err := sv.client.NewAPIRequest(http.MethodPost, path, reqBody)
	if err != nil {
		return
	}

	res, err := sv.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}
	return
}

func (sv *CustomersService) ListPayments(ID string, options *CustomerPaymentOptions) (pl PaymentsList, err error) {
	path := fmt.Sprintf("v2/customers/%s/payments")
	if options.ProfileID == "" {
		return lp, fmt.Errorf(requiredCreateParam, options.ProfileID)
	}
	if options != nil {
		v, _ := query.Values(options)
		path = fmt.Sprintf("%s?%s", path, v.Encode())
	}
	req, err := sv.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return
	}
	res, err := sv.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}
