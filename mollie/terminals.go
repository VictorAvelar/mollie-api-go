package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// TerminalsService operates over terminals resource.
type TerminalsService service

// TerminalStatus is the status of the terminal, which is a read-only value determined by Mollie.
type TerminalStatus string

// Possible terminal statuses.
const (
	TerminalPending  TerminalStatus = "pending"
	TerminalActive   TerminalStatus = "active"
	TerminalInactive TerminalStatus = "inactive"
)

// Terminal symbolizes a physical device to receive payments.
type Terminal struct {
	ID           string          `json:"id,omitempty"`
	Resource     string          `json:"resource,omitempty"`
	ProfileID    string          `json:"profileID,omitempty"`
	Status       TerminalStatus  `json:"status,omitempty"`
	Brand        string          `json:"brand,omitempty"`
	Model        string          `json:"model,omitempty"`
	SerialNumber string          `json:"serialNumber,omitempty"`
	Currency     string          `json:"currency,omitempty"`
	Description  string          `json:"description,omitempty"`
	CreatedAt    *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time      `json:"updatedAt,omitempty"`
	Links        PaginationLinks `json:"_links,omitempty"`
}

// Get terminal retrieves a single terminal object by its terminal ID.
func (ts *TerminalsService) Get(ctx context.Context, id string) (res *Response, t *Terminal, err error) {
	res, err = ts.client.get(ctx, fmt.Sprintf("v2/terminals/%s", id), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &t); err != nil {
		return
	}

	return
}

// TerminalListOptions holds query string parameters valid for terminals lists.
//
// ProfileID and TestMode are valid only when using access tokens.
type TerminalListOptions struct {
	From      string `url:"from,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	ProfileID string `url:"profileID,omitempty"`
	TestMode  bool   `url:"testMode,omitempty"`
}

// TerminalList describes the response for terminals list endpoints.
type TerminalList struct {
	Count    int `json:"count,omitempty"`
	Embedded struct {
		Terminals []*Terminal `json:"terminals,omitempty"`
	} `json:"_embedded,omitempty"`
	Links PaginationLinks `json:"_links,omitempty"`
}

// List retrieves a list of terminals symbolizing the physical devices to receive payments.
func (ts *TerminalsService) List(ctx context.Context, options *TerminalListOptions) (
	res *Response,
	tl *TerminalList,
	err error,
) {
	res, err = ts.client.get(ctx, "v2/terminals", options)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &tl); err != nil {
		return
	}

	return
}
