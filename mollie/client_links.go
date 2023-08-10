package mollie

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
)

// ClientLinksService interacts with the Client Links API to create
// new organizations for your customers.
type ClientLinksService service

// ClientDetails contains information to link a new organization to an
// OAuth application.
type ClientDetails struct {
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

// CreateClientLink based on the provided ClientDetails.
//
// See: https://docs.mollie.com/reference/v2/client-links-api/create-client-link
func (cls *ClientLinksService) CreateClientLink(ctx context.Context, cd *ClientDetails) (
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

// ClientLinkFinalizeOptions subset of the parameters allowed for the Authorize endpoint.
type ClientLinkFinalizeOptions struct {
	ClientID       string `url:"clientID,omitempty"`
	State          string `url:"state,omitempty"`
	Scope          string `url:"scope,omitempty"`
	ApprovalPrompt string `url:"approvalPrompt,omitempty"`
}

func (cls *ClientLinksService) CreateFinalizeClientLink(
	ctx context.Context,
	clientLink string,
	options *ClientLinkFinalizeOptions,
) (
	clientLinkURI string,
) {
	if options != nil {
		v, _ := query.Values(options)
		clientLinkURI = fmt.Sprintf("%s?%s", clientLink, v.Encode())
	}

	return
}
