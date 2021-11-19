package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Mandate allow you to charge a customer’s credit card or bank account recurrently.
type Mandate struct {
	ID               string         `json:"id,omitempty"`
	Resource         string         `json:"resource,omitempty"`
	Method           PaymentMethod  `json:"method,omitempty"`
	ConsumerName     string         `json:"consumerName,omitempty"`
	ConsumerAccount  string         `json:"consumerAccount,omitempty"`
	ConsumerBic      string         `json:"consumerBic,omitempty"`
	SignatureDate    *ShortDate     `json:"signatureDate,omitempty"`
	MandateReference string         `json:"mandateReference,omitempty"`
	Mode             Mode           `json:"mode,omitempty"`
	Status           MandateStatus  `json:"status,omitempty"`
	CreatedAt        *time.Time     `json:"createdAt,omitempty"`
	Details          MandateDetails `json:"details,omitempty"`
	Links            MandateLinks   `json:"_links,omitempty"`
}

// MandateDetails are possible values inside the mandate.details field
type MandateDetails struct {
	ConsumerName    string     `json:"consumerName,omitempty"`
	ConsumerAccount string     `json:"consumerAccount,omitempty"`
	ConsumerBic     string     `json:"consumerBic,omitempty"`
	CardHolder      string     `json:"cardHolder,omitempty"`
	CardNumber      string     `json:"cardNumber,omitempty"`
	CardLabel       CardLabel  `json:"cardLabel,omitempty"`
	CardFingerprint string     `json:"cardFingerprint,omitempty"`
	CardExpiryDate  *ShortDate `json:"cardExpiryDate,omitempty"`
}

// MandateStatus for the Mandate object
type MandateStatus string

// Valid mandate statuses
const (
	PendingMandate MandateStatus = "pending"
	ValidMandate   MandateStatus = "valid"
	InvalidMandate MandateStatus = "invalid"
)

// CardLabel Note that not all labels can be processed through Mollie.
type CardLabel string

// Available card labels
const (
	AmericaExpress CardLabel = "American Express"
	CartaSi        CardLabel = "Carta Si"
	CarteBleue     CardLabel = "Carte Bleue"
	Dankort        CardLabel = "Dankort"
	DinersClub     CardLabel = "Diners Club"
	Discover       CardLabel = "Discover"
	JCB            CardLabel = "JCB"
	Laser          CardLabel = "Laser"
	Maestro        CardLabel = "Maestro"
	Mastercard     CardLabel = "Mastercard"
	Unionpay       CardLabel = "Unionpay"
	Visa           CardLabel = "Visa"
	Empty          CardLabel = "null"
)

// MandatesService operates over customer mandates endpoints.
type MandatesService service

// MandateLinks response objects
type MandateLinks struct {
	Self          *URL `json:"self,omitempty"`
	Customer      *URL `json:"customer,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// ListMandatesOptions contains valid query parameters
// to filter the List mandates actions.
// From is a mandate id to offset from (inclusive)
// Limit is the max number of mandates to retrieve
type ListMandatesOptions struct {
	From  string `url:"from,omitempty"`
	Limit int    `url:"limit,omitempty"`
}

// MandateList describes how a list of mandates will be retrieved by Mollie.
type MandateList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Mandates []Mandate
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// Create a mandate for a specific customer.
// Mandates allow you to charge a customer’s credit card or bank account recurrently.
//
// See: https://docs.mollie.com/reference/v2/mandates-api/create-mandate
func (ms *MandatesService) Create(ctx context.Context, cID string, mandate Mandate) (mr *Mandate, err error) {
	u := fmt.Sprintf("v2/customers/%s/mandates", cID)
	req, err := ms.client.NewAPIRequest(ctx, http.MethodPost, u, mandate)
	if err != nil {
		return
	}

	res, err := ms.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &mr); err != nil {
		return
	}

	return
}

// Get retrieves a mandate by its ID and its customer’s ID.
// The mandate will either contain IBAN or credit card details,
// depending on the type of mandate.
//
// See: https://docs.mollie.com/reference/v2/mandates-api/get-mandate
func (ms *MandatesService) Get(ctx context.Context, cID, mID string) (mr *Mandate, err error) {
	u := fmt.Sprintf("v2/customers/%s/mandates/%s", cID, mID)
	req, err := ms.client.NewAPIRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ms.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &mr); err != nil {
		return
	}

	return
}

// Revoke a customer’s mandate.
// You will no longer be able to charge the consumer’s bank account or credit card with this mandate and all connected subscriptions will be canceled.
//
// See: https://docs.mollie.com/reference/v2/mandates-api/revoke-mandate
func (ms *MandatesService) Revoke(ctx context.Context, cID, mID string) (err error) {
	u := fmt.Sprintf("v2/customers/%s/mandates/%s", cID, mID)
	req, err := ms.client.NewAPIRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return
	}

	_, err = ms.client.Do(req)
	if err != nil {
		return
	}

	return
}

// List retrieves all mandates for the given customerId,
// ordered from newest to oldest.
//
// See: https://docs.mollie.com/reference/v2/mandates-api/list-mandates
func (ms *MandatesService) List(ctx context.Context, cID string, opt *ListMandatesOptions) (ml MandateList, err error) {
	u := fmt.Sprintf("v2/customers/%s/mandates", cID)
	if opt != nil {
		v, _ := query.Values(opt)
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}

	req, err := ms.client.NewAPIRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ms.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &ml); err != nil {
		return
	}

	return
}
