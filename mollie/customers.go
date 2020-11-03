package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// CustomersService operates over the customer resource
type CustomersService service

// CustomerLinks contains the HAL resources for a customer response
type CustomerLinks struct {
	Self          URL `json:"self,omitempty"`
	Mandates      URL `json:"mandates,omitempty"`
	Subscriptions URL `json:"subscriptions,omitempty"`
	Payments      URL `json:"payments,omitempty"`
	Documentation URL `json:"documentation,omitempty"`
	Dashboard     URL `json:"dashboard,omitempty"`
}

// Customer represents buyers
type Customer struct {
	Resource  string                 `json:"resource,omitempty"`
	ID        string                 `json:"id,omitempty"`
	Mode      Mode                   `json:"mode,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Email     string                 `json:"email,omitempty"`
	Locale    Locale                 `json:"locale,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"createdAt,omitempty"`
	Links     CustomerLinks          `json:"_links,omitempty"`
}

// ListCustomersOptions contains valid query parameters for the list customers endpoint.
type ListCustomersOptions struct {
	From         string       `url:"from,omitempty"`
	Limit        int          `url:"limit,omitempty"`
	ProfileID    string       `url:"profileId,omitempty"`
	SequenceType SequenceType `url:"sequenceType,omitempty"`
	RedirectURL  string       `url:"redirectUrl,omitempty"`
}

// CustomersList contains a embedded list of customers
// wrapped in a standard Mollie paginated response.
type CustomersList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Customers []Customer `json:"customers,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}

// Get finds a customer by its ID
//
// See: https://docs.mollie.com/reference/v2/customers-api/get-customer
func (cs *CustomersService) Get(id string) (c *Customer, err error) {
	u := fmt.Sprintf("v2/customers/%s", id)
	req, err := cs.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &c); err != nil {
		return
	}
	return
}

// Create creates a simple minimal representation of a customer in the Mollie API
// to use for the Mollie Checkout and Recurring features
//
// See: https://docs.mollie.com/reference/v2/customers-api/create-customer
func (cs *CustomersService) Create(c Customer) (cc *Customer, err error) {
	req, err := cs.client.NewAPIRequest(http.MethodPost, "v2/customers", c)
	if err != nil {
		return
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cc); err != nil {
		return
	}
	return
}

// Update mutates an existing customer
//
// See: https://docs.mollie.com/reference/v2/customers-api/update-customer
func (cs *CustomersService) Update(id string, c Customer) (cc *Customer, err error) {
	u := fmt.Sprintf("v2/customers/%s", id)
	req, err := cs.client.NewAPIRequest(http.MethodPatch, u, c)
	if err != nil {
		return
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cc); err != nil {
		return
	}
	return
}

// Delete a customer.
// All mandates and subscriptions created for this customer will be canceled as well.
//
// See: https://docs.mollie.com/reference/v2/customers-api/delete-customer
func (cs *CustomersService) Delete(id string) (err error) {
	u := fmt.Sprintf("v2/customers/%s", id)
	req, err := cs.client.NewAPIRequest(http.MethodDelete, u, nil)
	if err != nil {
		return
	}

	_, err = cs.client.Do(req)
	if err != nil {
		return
	}
	return
}

// List retrieves all customers created.
//
// See: https://docs.mollie.com/reference/v2/customers-api/list-customers
func (cs *CustomersService) List(options *ListCustomersOptions) (cl *CustomersList, err error) {
	u := "v2/customers"
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	res, err := cs.list(u)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}
	return
}

// GetPayments retrieves all payments linked to the customer.
//
// See: https://docs.mollie.com/reference/v2/customers-api/list-customer-payments
func (cs *CustomersService) GetPayments(id string, options *ListCustomersOptions) (pl *PaymentList, err error) {
	u := fmt.Sprintf("v2/customers/%s/payments", id)
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	res, err := cs.list(u)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pl); err != nil {
		return
	}
	return
}

// CreatePayment creates a payment for the customer.
//
// See: https://docs.mollie.com/reference/v2/customers-api/create-customer-payment
func (cs *CustomersService) CreatePayment(id string, p Payment) (pp *Payment, err error) {
	u := fmt.Sprintf("v2/customers/%s/payments", id)
	req, err := cs.client.NewAPIRequest(http.MethodPost, u, p)
	if err != nil {
		return
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pp); err != nil {
		return
	}
	return
}

func (cs *CustomersService) list(uri string) (r *Response, err error) {
	req, err := cs.client.NewAPIRequest(http.MethodGet, uri, nil)
	if err != nil {
		return
	}

	r, err = cs.client.Do(req)
	if err != nil {
		return
	}
	return
}
