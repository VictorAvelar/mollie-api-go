package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CustomersService operates over the customer resource.
type CustomersService service

// CustomerLinks contains the HAL resources for a customer response.
type CustomerLinks struct {
	Self          *URL `json:"self,omitempty"`
	Mandates      *URL `json:"mandates,omitempty"`
	Subscriptions *URL `json:"subscriptions,omitempty"`
	Payments      *URL `json:"payments,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
	Dashboard     *URL `json:"dashboard,omitempty"`
}

// Customer represents buyers.
type Customer struct {
	Resource  string        `json:"resource,omitempty"`
	ID        string        `json:"id,omitempty"`
	Mode      Mode          `json:"mode,omitempty"`
	Name      string        `json:"name,omitempty"`
	Email     string        `json:"email,omitempty"`
	Locale    Locale        `json:"locale,omitempty"`
	Metadata  interface{}   `json:"metadata,omitempty"`
	CreatedAt *time.Time    `json:"createdAt,omitempty"`
	Links     CustomerLinks `json:"_links,omitempty"`
}

// CustomersListOptions contains valid query parameters for the list customers endpoint.
type CustomersListOptions struct {
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

// Get finds a customer by its ID.
//
// See: https://docs.mollie.com/reference/v2/customers-api/get-customer
func (cs *CustomersService) Get(ctx context.Context, id string) (res *Response, c *Customer, err error) {
	u := fmt.Sprintf("v2/customers/%s", id)

	res, err = cs.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &c); err != nil {
		return
	}

	return
}

// Create creates a simple minimal representation of a customer in the Mollie API
// to use for the Mollie Checkout and Recurring features.
//
// See: https://docs.mollie.com/reference/v2/customers-api/create-customer
func (cs *CustomersService) Create(ctx context.Context, c Customer) (res *Response, cc *Customer, err error) {
	res, err = cs.client.post(ctx, "v2/customers", c, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cc); err != nil {
		return
	}

	return
}

// Update an existing customer.
//
// See: https://docs.mollie.com/reference/v2/customers-api/update-customer
func (cs *CustomersService) Update(ctx context.Context, id string, c Customer) (
	res *Response,
	cc *Customer,
	err error,
) {
	u := fmt.Sprintf("v2/customers/%s", id)

	res, err = cs.client.patch(ctx, u, c, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cc); err != nil {
		return
	}

	return
}

// Delete a customer.
//
// All mandates and subscriptions created for this customer will be canceled as well.
//
// See: https://docs.mollie.com/reference/v2/customers-api/delete-customer
func (cs *CustomersService) Delete(ctx context.Context, id string) (res *Response, err error) {
	u := fmt.Sprintf("v2/customers/%s", id)

	res, err = cs.client.delete(ctx, u, nil)
	if err != nil {
		return
	}

	return
}

// List retrieves all customers created.
//
// See: https://docs.mollie.com/reference/v2/customers-api/list-customers
func (cs *CustomersService) List(ctx context.Context, options *CustomersListOptions) (
	res *Response,
	cl *CustomersList,
	err error,
) {
	res, err = cs.list(ctx, "v2/customers", options)
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
func (cs *CustomersService) GetPayments(ctx context.Context, id string, options *CustomersListOptions) (
	res *Response,
	pl *PaymentList,
	err error,
) {
	u := fmt.Sprintf("v2/customers/%s/payments", id)

	res, err = cs.list(ctx, u, options)
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
func (cs *CustomersService) CreatePayment(ctx context.Context, id string, p CreatePayment) (
	res *Response,
	pp *Payment,
	err error,
) {
	u := fmt.Sprintf("v2/customers/%s/payments", id)

	res, err = cs.client.post(ctx, u, p, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &pp); err != nil {
		return
	}

	return
}

func (cs *CustomersService) list(ctx context.Context, uri string, options interface{}) (r *Response, err error) {
	r, err = cs.client.get(ctx, uri, options)
	if err != nil {
		return
	}

	return
}
