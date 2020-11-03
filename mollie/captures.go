package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CapturesService operates over captures resource
type CapturesService service

// CaptureLinks contains relevant links for a capture object
type CaptureLinks struct {
	Self          URL `json:"self,omitempty"`
	Payment       URL `json:"payment,omitempty"`
	Shipment      URL `json:"shipment,omitempty"`
	Settlement    URL `json:"settlement,omitempty"`
	Documentation URL `json:"documentation,omitempty"`
}

// Capture describes a single capture
// Captures are used for payments that have the authorize-then-capture flow
type Capture struct {
	Resource         string       `json:"resource,omitempty"`
	ID               string       `json:"id,omitempty"`
	Mode             Mode         `json:"mode,omitempty"`
	Amount           Amount       `json:"amount,omitempty"`
	SettlementAmount Amount       `json:"settlementAmount,omitempty"`
	PaymentID        string       `json:"paymentId,omitempty"`
	ShipmentID       string       `json:"shipmentId,omitempty"`
	SettlementID     string       `json:"settlementId,omitempty"`
	CreatedAt        time.Time    `json:"createdAt,omitempty"`
	Links            CaptureLinks `json:"links,omitempty"`
}

// CapturesList describes a list of captures
type CapturesList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Captures []Capture
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// Get retrieves a single capture by its ID.
// Note the original payment’s ID is needed as well.
//
// See: https://docs.mollie.com/reference/v2/captures-api/get-capture
func (cs *CapturesService) Get(pID, cID string) (c *Capture, err error) {
	u := fmt.Sprintf("v2/payments/%s/captures/%s", pID, cID)
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

// List retrieves all captures for a certain payment
//
// See: https://docs.mollie.com/reference/v2/captures-api/list-captures
func (cs *CapturesService) List(pID string) (cl *CapturesList, err error) {
	u := fmt.Sprintf("v2/payments/%s/captures", pID)
	req, err := cs.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := cs.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}
	return
}
