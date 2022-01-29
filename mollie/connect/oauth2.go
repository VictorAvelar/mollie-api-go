package connect

import (
	"golang.org/x/oauth2"
)

// Mollie's Oauth2 server URLs.
const (
	AuthURL   = "https://www.mollie.com/oauth2/authorize"
	TokensURL = "https://api.mollie.com/oauth2/tokens"
)

// Endpoint is Mollies's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:   AuthURL,
	TokenURL:  TokensURL,
	AuthStyle: 0,
}
