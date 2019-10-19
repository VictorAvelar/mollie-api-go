package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ShipmentsService operates on shipments endpoints.
type ShipmentsService service

// Shipment contains information about a user service/product delivery and
// is used in the figurative sense here.
// It can also mean that a service was provided or digital content was delivered.
type Shipment struct {
	Resource  string           `json:"resource,omitempty"`
	ID        string           `json:"id,omitempty"`
	OrderID   string           `json:"orderId,omitempty"`
	CreatedAt *time.Time       `json:"createdAt,omitempty"`
	Tracking  ShipmentTracking `json:"tracking,omitempty"`
	Lines     []OrderLines     `json:"lines,omitempty"`
	Links     ShipmentLinks    `json:"_links,omitempty"`
}

// ShipmentTracking contains shipment tracking details.
type ShipmentTracking struct {
	Carrier string `json:"carrier,omitempty"`
	Code    string `json:"code,omitempty"`
	URL     string `json:"url,omitempty"`
}

// ShipmentLinks contains URL objects with shipment relevant
// information for the user.
type ShipmentLinks struct {
	Self          URL `json:"self,omitempty"`
	Order         URL `json:"order,omitempty"`
	Documentation URL `json:"documentation,omitempty"`
}

// Get retrieves a single shipment and the order lines shipped by a shipment’s ID.
//
// See: https://docs.mollie.com/reference/v2/shipments-api/get-shipment#
func (ss *ShipmentsService) Get(oID string, sID string) (s *Shipment, err error) {
	u := fmt.Sprintf("v2/orders/%s/shipments/%s", oID, sID)
	req, err := ss.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}

// CreateShipmentRequest defines information required to create a new shipment
type CreateShipmentRequest struct {
	Lines    []OrderLines     `json:"lines,omitempty"`
	Tracking ShipmentTracking `json:"tracking,omitempty"`
}

// Create can be used to ship order lines.
//
// See: https://docs.mollie.com/reference/v2/shipments-api/create-shipment
func (ss *ShipmentsService) Create(oID string, cs CreateShipmentRequest) (s *Shipment, err error) {
	u := fmt.Sprintf("v2/orders/%s/shipments", oID)
	req, err := ss.client.NewAPIRequest(http.MethodPost, u, cs)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}

// ShipmentsList describes how a list of payments will be retrieved by Mollie.
type ShipmentsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Shipments []Shipment
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// List retrieves all shipments for an order.
//
// See: https://docs.mollie.com/reference/v2/shipments-api/list-shipments
func (ss *ShipmentsService) List(oID string) (sl *ShipmentsList, err error) {
	u := fmt.Sprintf("v2/orders/%s/shipments", oID)
	req, err := ss.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sl); err != nil {
		return
	}
	return
}

// Update can be used to update the tracking information of a shipment
//
// See: https://docs.mollie.com/reference/v2/shipments-api/update-shipment
func (ss *ShipmentsService) Update(oID string, sID string, st ShipmentTracking) (s *Shipment, err error) {
	u := fmt.Sprintf("v2/orders/%s/shipments/%s", oID, sID)
	req, err := ss.client.NewAPIRequest(http.MethodPatch, u, st)
	if err != nil {
		return
	}

	res, err := ss.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}
	return
}
