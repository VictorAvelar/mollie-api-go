package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// SalesInvoiceStatus represents the status for the invoice to end up in.
type SalesInvoiceStatus string

// Possible values for SalesInvoiceStatus.
const (
	// DraftSalesInvoiceStatus indicates that the invoice is not paid or not sent and can be updated after creation.
	DraftSalesInvoiceStatus SalesInvoiceStatus = "draft"
	// IssuedSalesInvoiceStatus sends it to the recipient so they may then pay through our payment system.
	IssuedSalesInvoiceStatus SalesInvoiceStatus = "issued"
	// PaidSalesInvoiceStatus marks the invoice as paid.
	PaidSalesInvoiceStatus SalesInvoiceStatus = "paid"
)

// SalesInvoiceSource is the way through which the invoice is to be set to paid.
type SalesInvoiceSource string

// Possible values for SalesInvoiceSource.
const (
	ManualSalesInvoiceSource      SalesInvoiceSource = "manual"
	PaymentLinkSalesInvoiceSource SalesInvoiceSource = "payment_link"
	PaymentSalesInvoiceSource     SalesInvoiceSource = "payment"
)

// SalesInvoiceVATScheme represents the VAT scheme for the sales invoice.
type SalesInvoiceVATScheme string

// Possible values for SalesInvoiceVATScheme.
const (
	StandardSalesInvoiceVATScheme SalesInvoiceVATScheme = "standard"
	// OneStopShopSalesInvoiceVATScheme indicates that the One Stop Shop (OSS) VAT scheme is used.
	// You must be enrolled with One Stop Shop enabled to use it.
	OneStopShopSalesInvoiceVATScheme SalesInvoiceVATScheme = "one-stop-shop"
)

// SalesInvoiceVATMode represents the VAT mode for the sales invoice.
type SalesInvoiceVATMode string

// Possible values for SalesInvoiceVATMode.
const (
	// VAT is added on top of the prices.
	ExclusiveSalesInvoiceVATMode SalesInvoiceVATMode = "exclusive"
	// VAT is included in the prices.
	InclusiveSalesInvoiceVATMode SalesInvoiceVATMode = "inclusive"
)

// SalesInvoicePaymentTerm represents the payment term for the sales invoice.
// Defaults to 30 days.
type SalesInvoicePaymentTerm string

// Possible values for SalesInvoicePaymentTerm.
const (
	PaymentTerm7Days  SalesInvoicePaymentTerm = "7 days"
	PaymentTerm14Days SalesInvoicePaymentTerm = "14 days"
	PaymentTerm30Days SalesInvoicePaymentTerm = "30 days"
	PaymentTerm45Days SalesInvoicePaymentTerm = "45 days"
	PaymentTerm60Days SalesInvoicePaymentTerm = "60 days"
	PaymentTerm90Days SalesInvoicePaymentTerm = "90 days"
)

// SalesInvoiceRecipientType represents the type of recipient for the sales invoice.
type SalesInvoiceRecipientType string

// Possible values for SalesInvoiceRecipientType.
const (
	ConsumerSalesInvoiceRecipientType SalesInvoiceRecipientType = "consumer"
	BusinessSalesInvoiceRecipientType SalesInvoiceRecipientType = "business"
)

// SalesInvoiceDiscountType represents the type of discount applied to the sales invoice.
type SalesInvoiceDiscountType string

// Possible values for SalesInvoiceDiscountType.
const (
	FixedAmountSalesInvoiceDiscountType SalesInvoiceDiscountType = "amount"
	PercentageSalesInvoiceDiscountType  SalesInvoiceDiscountType = "percentage"
)

// SalesInvoiceDiscount represents the discount to be applied to the line item.
type SalesInvoiceDiscount struct {
	Type  SalesInvoiceDiscountType `json:"type,omitempty"`
	Value string                   `json:"value,omitempty"`
}

// SalesInvoicePaymentDetails contains details about how the sales invoice was or will be paid.
type SalesInvoicePaymentDetails struct {
	SourceReference string             `json:"sourceReference,omitempty"`
	Source          SalesInvoiceSource `json:"source,omitempty"`
}

// SalesInvoiceEmailDetails contains details about the email sent with the sales invoice.
type SalesInvoiceEmailDetails struct {
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
}

// SalesInvoiceRecipient represents the recipient of the sales invoice.
type SalesInvoiceRecipient struct {
	Address
	Type               SalesInvoiceRecipientType `json:"type,omitempty"`
	Title              string                    `json:"title,omitempty"`
	Email              string                    `json:"email,omitempty"`
	Phone              string                    `json:"phone,omitempty"`
	OrganizationName   string                    `json:"organizationName,omitempty"`
	OrganizationNumber string                    `json:"organizationNumber,omitempty"`
	VATNumber          string                    `json:"vatNumber,omitempty"`
	Locale             Locale                    `json:"locale,omitempty"`
}

// SalesInvoiceLineItem represents a line item on the sales invoice.
type SalesInvoiceLineItem struct {
	Description string                `json:"description"`
	Quantity    int                   `json:"quantity"`
	VATRate     string                `json:"vatRate"`
	UnitPrice   Amount                `json:"unitPrice"`
	Discount    *SalesInvoiceDiscount `json:"discount,omitempty"`
}

// SalesInvoiceLinks contains links related to the sales invoice.
type SalesInvoiceLinks struct {
	Self           *URL `json:"self,omitempty"`
	InvoicePayment *URL `json:"invoicePayment,omitempty"`
	PDF            *URL `json:"pdfLink,omitempty"`
	Documentation  *URL `json:"documentation,omitempty"`
	Next           *URL `json:"next,omitempty"`
	Previous       *URL `json:"previous,omitempty"`
}

// CreateSalesInvoice represents the payload to create a sales invoice.
type CreateSalesInvoice struct {
	TestMode            bool                        `json:"testmode,omitempty"`
	ProfileID           string                      `json:"profileId,omitempty"`
	CustomerID          string                      `json:"customerId,omitempty"`
	MandateID           string                      `json:"mandateId,omitempty"`
	RecipientIdentifier string                      `json:"recipientIdentifier,omitempty"`
	Memo                string                      `json:"memo,omitempty"`
	Metadata            map[string]string           `json:"metadata,omitempty"`
	Status              SalesInvoiceStatus          `json:"status,omitempty"`
	VATScheme           SalesInvoiceVATScheme       `json:"vatScheme,omitempty"`
	VATMode             SalesInvoiceVATMode         `json:"vatMode,omitempty"`
	PaymentTerm         SalesInvoicePaymentTerm     `json:"paymentTerm,omitempty"`
	PaymentDetails      *SalesInvoicePaymentDetails `json:"paymentDetails,omitempty"`
	EmailDetails        SalesInvoiceEmailDetails    `json:"emailDetails,omitempty"`
	Recipient           SalesInvoiceRecipient       `json:"recipient,omitempty"`
	Lines               []SalesInvoiceLineItem      `json:"lines,omitempty"`
	Discount            *SalesInvoiceDiscount       `json:"discount,omitempty"`
}

// UpdateSalesInvoice represents the payload to update a sales invoice.
type UpdateSalesInvoice struct {
	TestMode            bool                     `json:"testmode,omitempty"`
	Memo                string                   `json:"memo,omitempty"`
	RecipientIdentifier string                   `json:"recipientIdentifier,omitempty"`
	Status              SalesInvoiceStatus       `json:"status,omitempty"`
	PaymentTerm         SalesInvoicePaymentTerm  `json:"paymentTerm,omitempty"`
	EmailDetails        SalesInvoiceEmailDetails `json:"emailDetails,omitempty"`
	Recipient           SalesInvoiceRecipient    `json:"recipient,omitempty"`
	Lines               []SalesInvoiceLineItem   `json:"lines,omitempty"`
	Discount            *SalesInvoiceDiscount    `json:"discount,omitempty"`
}

// SalesInvoice represents a sales invoice resource.
type SalesInvoice struct {
	Resource                 string                       `json:"resource,omitempty"`
	ID                       string                       `json:"id,omitempty"`
	ProfileID                string                       `json:"profileId,omitempty"`
	Currency                 string                       `json:"currency,omitempty"`
	InvoiceNumber            string                       `json:"invoiceNumber,omitempty"`
	Memo                     string                       `json:"memo,omitempty"`
	CustomerID               string                       `json:"customerId,omitempty"`
	MandateID                string                       `json:"mandateId,omitempty"`
	RecipientIdentifier      string                       `json:"recipientIdentifier,omitempty"`
	Metadata                 map[string]string            `json:"metadata,omitempty"`
	Mode                     Mode                         `json:"mode,omitempty"`
	AmountDue                Amount                       `json:"amountDue,omitempty"`
	SubtotalAmount           Amount                       `json:"subtotalAmount,omitempty"`
	TotalAmount              Amount                       `json:"totalAmount,omitempty"`
	TotalVATAmount           Amount                       `json:"totalVatAmount,omitempty"`
	DiscountedSubtotalAmount Amount                       `json:"discountedSubtotalAmount,omitempty"`
	Status                   SalesInvoiceStatus           `json:"status,omitempty"`
	VATScheme                SalesInvoiceVATScheme        `json:"vatScheme,omitempty"`
	VATMode                  SalesInvoiceVATMode          `json:"vatMode,omitempty"`
	PaymentTerm              SalesInvoicePaymentTerm      `json:"paymentTerm,omitempty"`
	PaymentDetails           []SalesInvoicePaymentDetails `json:"paymentDetails,omitempty"`
	EmailDetails             SalesInvoiceEmailDetails     `json:"emailDetails,omitempty"`
	Recipient                SalesInvoiceRecipient        `json:"recipient,omitempty"`
	Links                    SalesInvoiceLinks            `json:"_links,omitempty"`
	Lines                    []SalesInvoiceLineItem       `json:"lines,omitempty"`
	Discount                 *SalesInvoiceDiscount        `json:"discount,omitempty"`
	CreatedAt                *time.Time                   `json:"createdAt,omitempty"`
	IssuedAt                 *time.Time                   `json:"issuedAt,omitempty"`
	PaidAt                   *time.Time                   `json:"paidAt,omitempty"`
	DueAt                    *time.Time                   `json:"dueAt,omitempty"`
}

// SalesInvoiceList represents a paginated list of sales invoices.
type SalesInvoiceList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		SalesInvoices []SalesInvoice `json:"sales_invoices,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// ListSalesInvoicesOptions specifies the optional parameters to the List method.
type ListSalesInvoicesOptions struct {
	From     string `url:"from,omitempty"`
	Limit    int    `url:"limit,omitempty"`
	TestMode bool   `url:"testmode,omitempty"`
}

// SalesInvoicesService handles API operations for sales invoices.
type SalesInvoicesService service

// List retrieves a list of sales invoices.
//
// See: https://docs.mollie.com/reference/list-sales-invoices
func (s *SalesInvoicesService) List(ctx context.Context, opts *ListSalesInvoicesOptions) (
	res *Response,
	sil *SalesInvoiceList,
	err error,
) {
	res, err = s.client.get(ctx, "/v2/sales-invoices", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sil); err != nil {
		return
	}

	return
}

// Get retrieves a single sales invoice by its ID.
//
// See: https://docs.mollie.com/reference/get-sales-invoice
func (s *SalesInvoicesService) Get(ctx context.Context, salesInvoice string) (
	res *Response,
	si *SalesInvoice,
	err error,
) {
	u := fmt.Sprintf("/v2/sales-invoices/%s", salesInvoice)

	res, err = s.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &si); err != nil {
		return
	}

	return
}

// Create creates a new sales invoice.
//
// See: https://docs.mollie.com/reference/create-sales-invoice
func (s *SalesInvoicesService) Create(ctx context.Context, csi CreateSalesInvoice) (
	res *Response,
	si *SalesInvoice,
	err error,
) {
	if s.client.HasAccessToken() && s.client.config.testing {
		csi.TestMode = true
	}

	res, err = s.client.post(ctx, "/v2/sales-invoices", csi, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &si); err != nil {
		return
	}

	return
}

// Update updates an existing sales invoice.
//
// See: https://docs.mollie.com/reference/update-sales-invoice
func (s *SalesInvoicesService) Update(ctx context.Context, salesInvoice string, usi UpdateSalesInvoice) (
	res *Response,
	si *SalesInvoice,
	err error,
) {
	u := fmt.Sprintf("/v2/sales-invoices/%s", salesInvoice)

	if s.client.HasAccessToken() && s.client.config.testing {
		usi.TestMode = true
	}

	res, err = s.client.patch(ctx, u, usi)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &si); err != nil {
		return
	}

	return
}

// Delete deletes a sales invoice by its ID.
//
// See: https://docs.mollie.com/reference/delete-sales-invoice
func (s *SalesInvoicesService) Delete(ctx context.Context, salesInvoice string) (res *Response, err error) {
	u := fmt.Sprintf("/v2/sales-invoices/%s", salesInvoice)

	res, err = s.client.delete(ctx, u, nil)
	if err != nil {
		return
	}

	return
}
