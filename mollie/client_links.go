package mollie

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
)

// CreateClientLink contains information to link a new organization to an
// OAuth application.
type CreateClientLink struct {
	Owner              Owner    `json:"owner,omitempty"`
	Name               string   `json:"name,omitempty"`
	Address            *Address `json:"address,omitempty"`
	RegistrationNumber string   `json:"registrationNumber,omitempty"`
	VATNumber          string   `json:"vatNumber,omitempty"`
}

// ClientLinkLinks describes all the possible links to be returned with
// a client links response object.
type ClientLinkLinks struct {
	ClientLink    *URL `json:"clientLink,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// ClientLink object with redirect target.
type ClientLink struct {
	ID       string          `json:"id,omitempty"`
	Resource string          `json:"resource,omitempty"`
	Links    ClientLinkLinks `json:"_links,omitempty"`
}

// ClientLinksService interacts with the Client Links API to create
// new organizations for your customers.
type ClientLinksService service

// Create a client link based on the provided CreateClientLink values.
//
// See: https://docs.mollie.com/reference/v2/client-links-api/create-client-link
func (cls *ClientLinksService) Create(ctx context.Context, cd CreateClientLink) (
	res *Response,
	cl *ClientLink,
	err error,
) {
	res, err = cls.client.post(ctx, "v2/client-links", cd, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &cl); err != nil {
		return
	}

	return
}

// ApprovalPromptAction represents possible actions to be performed
// once the client link is created and redirected to the dashboard.
type ApprovalPromptAction string

// Possible approval prompt actions.
const (
	ForceApproval ApprovalPromptAction = "force"
	AutoApproval  ApprovalPromptAction = "auto"
)

// ClientLinkAuthorizeOptions subset of the parameters allowed for the Authorize endpoint.
type ClientLinkAuthorizeOptions struct {
	ClientID       string               `url:"clientId,omitempty"`
	State          string               `url:"state,omitempty"`
	Scope          []PermissionGrant    `del:"+"                        url:"scope,omitempty"`
	ApprovalPrompt ApprovalPromptAction `url:"approvalPrompt,omitempty"`
}

// GetFinalClientLink returns the final client link URI with the provided options.
func (cls *ClientLinksService) GetFinalClientLink(
	ctx context.Context,
	clientLink string,
	options *ClientLinkAuthorizeOptions,
) (
	clientLinkURI string,
) {
	if options != nil {
		v, _ := query.Values(options)
		clientLinkURI = fmt.Sprintf("%s?%s", clientLink, v.Encode())
	}

	return
}
