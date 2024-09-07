package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CreateShipment contains information required to create a new shipment.
type CreateShipment struct {
	Lines    []*OrderLine      `json:"lines,omitempty"`
	Tracking *ShipmentTracking `json:"tracking,omitempty"`
	ShipmentAccessTokenFields
}

// ShipmentAccessTokenFields describes the fields available when using an access token.
type ShipmentAccessTokenFields struct {
	Testmode bool `json:"testmode,omitempty"`
}

// Shipment contains information about a user service/product delivery and
// is used in the figurative sense here.
// It can also mean that a service was provided or digital content was delivered.
type Shipment struct {
	Resource  string            `json:"resource,omitempty"`
	ID        string            `json:"id,omitempty"`
	OrderID   string            `json:"orderId,omitempty"`
	CreatedAt *time.Time        `json:"createdAt,omitempty"`
	Tracking  *ShipmentTracking `json:"tracking,omitempty"`
	Lines     []*OrderLine      `json:"lines,omitempty"`
	Links     ShipmentLinks     `json:"_links,omitempty"`
	ShipmentAccessTokenFields
}

// UpdateShipment contains information required to update a shipment.
type UpdateShipment struct {
	Tracking *ShipmentTracking `json:"tracking,omitempty"`
	ShipmentAccessTokenFields
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
	Self          *URL `json:"self,omitempty"`
	Order         *URL `json:"order,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// ShipmentsList describes how a list of payments will be retrieved by Mollie.
type ShipmentsList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Shipments []Shipment
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// ShipmentsService operates on shipments endpoints.
type ShipmentsService service

// Get retrieves a single shipment and the order lines shipped by a shipment’s ID.
//
// See: https://docs.mollie.com/reference/get-shipment#
func (ss *ShipmentsService) Get(ctx context.Context, order string, shipment string) (
	res *Response,
	s *Shipment,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/shipments/%s", order, shipment)

	res, err = ss.client.get(ctx, u, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}

	return
}

// Create can be used to ship order lines.
//
// See: https://docs.mollie.com/reference/create-shipment
func (ss *ShipmentsService) Create(ctx context.Context, order string, cs CreateShipment) (
	res *Response,
	s *Shipment,
	err error,
) {
	uri := fmt.Sprintf("v2/orders/%s/shipments", order)

	if ss.client.HasAccessToken() && ss.client.config.testing {
		cs.Testmode = true
	}

	res, err = ss.client.post(ctx, uri, cs, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}

	return
}

// List retrieves all shipments for an order.
//
// See: https://docs.mollie.com/reference/list-shipments
func (ss *ShipmentsService) List(ctx context.Context, order string) (res *Response, sl *ShipmentsList, err error) {
	u := fmt.Sprintf("v2/orders/%s/shipments", order)

	res, err = ss.client.get(ctx, u, nil)
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
// See: https://docs.mollie.com/reference/update-shipment
func (ss *ShipmentsService) Update(ctx context.Context, order string, shipment string, us UpdateShipment) (
	res *Response,
	s *Shipment,
	err error,
) {
	u := fmt.Sprintf("v2/orders/%s/shipments/%s", order, shipment)

	res, err = ss.client.patch(ctx, u, us)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &s); err != nil {
		return
	}

	return
}
