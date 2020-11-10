package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// InvoiceStatus status of the invoice
type InvoiceStatus string

// Valid status of the invoice
const (
	InvoiceStatusOpen    InvoiceStatus = "open"
	InvoiceStatusPaid    InvoiceStatus = "paid"
	InvoiceStatusOverdue InvoiceStatus = "overdue"
)

// Invoice describes an invoice details
type Invoice struct {
	Resource    string        `json:"resource,omitempty"`
	ID          string        `json:"id,omitempty"`
	Reference   string        `json:"reference,omitempty"`
	VatNumber   string        `json:"vatNumber,omitempty"`
	Status      InvoiceStatus `json:"status,omitempty"`
	IssuedAt    string        `json:"issuedAt,omitempty"`
	PaidAt      string        `json:"paidAt,omitempty"`
	DueAt       string        `json:"dueAt,omitempty"`
	NetAmount   *Amount       `json:"netAmount,omitempty"`
	VatAmount   *Amount       `json:"vatAmount,omitempty"`
	GrossAmount *Amount       `json:"grossAmount,omitempty"`
	Lines       []*LineItem   `json:"lines,omitempty"`
	Links       InvoiceLinks  `json:"_links,omitempty"`
}

// LineItem product details
type LineItem struct {
	Period        string  `json:"period,omitempty"`
	Description   string  `json:"description,omitempty"`
	Count         int64   `json:"count,omitempty"`
	VatPercentage float64 `json:"vatPercentage,omitempty"`
	Amount        *Amount `json:"amount,omitempty"`
}

// InvoiceLinks describes all the possible links to be returned with
// a invoice object.
type InvoiceLinks struct {
	Self          *URL `json:"self,omitempty"`
	PDF           *URL `json:"pdf,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// ListInvoiceOptions describes list invoices endpoint valid query string parameters.
type ListInvoiceOptions struct {
	Reference string `json:"reference,omitempty"`
	Year      string `json:"year,omitempty"`
	From      int64  `json:"from,omitempty"`
	Limit     int64  `json:"limit,omitempty"`
}

// InvoiceList describes how a list of invoices will be retrieved by Mollie.
type InvoiceList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Invoices []Invoice `json:"invoices"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// InvoicesService instance operates over invoice resources
type InvoicesService service

// Get retrieve details of an invoice, using the invoice’s identifier.
func (is *InvoicesService) Get(id string) (i Invoice, err error) {
	getURL := fmt.Sprintf("v2/invoices/%s", id)

	req, err := is.client.NewAPIRequest(http.MethodGet, getURL, nil)
	if err != nil {
		return
	}
	res, err := is.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &i); err != nil {
		return
	}
	return
}

// List retrieves a list of invoices associated with your account/organization.
func (is *InvoicesService) List(options *ListInvoiceOptions) (il InvoiceList, err error) {
	u := "v2/invoices"
	if options != nil {
		v, _ := query.Values(options)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}
	req, err := is.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}
	res, err := is.client.Do(req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.content, &il); err != nil {
		return
	}
	return
}
