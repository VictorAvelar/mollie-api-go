package mollie

import (
	"context"
	"encoding/json"
	"fmt"
)

// InvoiceStatus status of the invoice.
type InvoiceStatus string

// Valid status of the invoice.
const (
	InvoiceStatusOpen    InvoiceStatus = "open"
	InvoiceStatusPaid    InvoiceStatus = "paid"
	InvoiceStatusOverdue InvoiceStatus = "overdue"
)

// Invoice describes an invoice details.
type Invoice struct {
	Resource    string        `json:"resource,omitempty"`
	ID          string        `json:"id,omitempty"`
	Reference   string        `json:"reference,omitempty"`
	VatNumber   string        `json:"vatNumber,omitempty"`
	IssuedAt    string        `json:"issuedAt,omitempty"`
	PaidAt      string        `json:"paidAt,omitempty"`
	DueAt       string        `json:"dueAt,omitempty"`
	NetAmount   *Amount       `json:"netAmount,omitempty"`
	VatAmount   *Amount       `json:"vatAmount,omitempty"`
	GrossAmount *Amount       `json:"grossAmount,omitempty"`
	Lines       []*LineItem   `json:"lines,omitempty"`
	Status      InvoiceStatus `json:"status,omitempty"`
	Links       InvoiceLinks  `json:"_links,omitempty"`
}

// LineItem product details.
type LineItem struct {
	Count         int64   `json:"count,omitempty"`
	VatPercentage float64 `json:"vatPercentage,omitempty"`
	Period        string  `json:"period,omitempty"`
	Description   string  `json:"description,omitempty"`
	Amount        *Amount `json:"amount,omitempty"`
}

// InvoiceLinks describes all the possible links to be returned with
// a invoice object.
type InvoiceLinks struct {
	Self          *URL `json:"self,omitempty"`
	PDF           *URL `json:"pdf,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// ListInvoicesOptions describes list invoices endpoint valid query string parameters.
type ListInvoicesOptions struct {
	Limit     int64  `url:"limit,omitempty"`
	Reference string `url:"reference,omitempty"`
	Year      string `url:"year,omitempty"`
	From      string `url:"from,omitempty"`
}

// InvoicesList describes how a list of invoices will be retrieved by Mollie.
type InvoicesList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Invoices []*Invoice `json:"invoices"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// InvoicesService instance operates over invoice resources.
type InvoicesService service

// Get retrieve details of an invoice, using the invoice’s identifier.
func (is *InvoicesService) Get(ctx context.Context, id string) (res *Response, i *Invoice, err error) {
	u := fmt.Sprintf("v2/invoices/%s", id)

	res, err = is.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &i); err != nil {
		return
	}

	return
}

// List retrieves a list of invoices associated with your account/organization.
func (is *InvoicesService) List(ctx context.Context, options *ListInvoicesOptions) (
	res *Response,
	il *InvoicesList,
	err error,
) {
	res, err = is.client.get(ctx, "v2/invoices", options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &il); err != nil {
		return
	}

	return
}
