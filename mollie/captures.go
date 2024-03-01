package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CaptureMode describes the mode of a capture.
type CaptureMode string

// CaptureMode possible values.
const (
	AutomaticCapture CaptureMode = "automatic"
	ManualCapture    CaptureMode = "manual"
)

// CapturesService operates over captures resource.
type CapturesService service

// CaptureLinks contains relevant links for a capture object.
type CaptureLinks struct {
	Self          *URL `json:"self,omitempty"`
	Payment       *URL `json:"payment,omitempty"`
	Shipment      *URL `json:"shipment,omitempty"`
	Settlement    *URL `json:"settlement,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// Capture describes a single capture.
// Captures are used for payments that have the authorize-then-capture flow.
type Capture struct {
	Resource         string       `json:"resource,omitempty"`
	ID               string       `json:"id,omitempty"`
	Mode             Mode         `json:"mode,omitempty"`
	Amount           *Amount      `json:"amount,omitempty"`
	SettlementAmount *Amount      `json:"settlementAmount,omitempty"`
	PaymentID        string       `json:"paymentId,omitempty"`
	ShipmentID       string       `json:"shipmentId,omitempty"`
	SettlementID     string       `json:"settlementId,omitempty"`
	CreatedAt        *time.Time   `json:"createdAt,omitempty"`
	Links            CaptureLinks `json:"_links,omitempty"`
}

// CapturesList describes a list of captures.
type CapturesList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Captures []*Capture
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// Get retrieves a single capture by its ID.
// Note the original payment’s ID is needed as well.
//
// See: https://docs.mollie.com/reference/v2/captures-api/get-capture
func (cs *CapturesService) Get(ctx context.Context, payment, capture string) (res *Response, c *Capture, err error) {
	u := fmt.Sprintf("v2/payments/%s/captures/%s", payment, capture)

	res, err = cs.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &c); err != nil {
		return
	}

	return
}

// List retrieves all captures for a certain payment.
//
// See: https://docs.mollie.com/reference/v2/captures-api/list-captures
func (cs *CapturesService) List(ctx context.Context, payment string) (res *Response, cl *CapturesList, err error) {
	u := fmt.Sprintf("v2/payments/%s/captures", payment)

	res, err = cs.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}

	return
}
