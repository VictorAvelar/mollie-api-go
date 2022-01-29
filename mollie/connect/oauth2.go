package connect

import (
	"golang.org/x/oauth2"
)

// Mollie's Oauth2 server URLs.
const (
	authURL   = "https://www.mollie.com/oauth2/authorize"
	tokensURL = "https://api.mollie.com/oauth2/tokens"
)

// OauthEndpoint is Mollies's OAuth 2.0 endpoint.
func OauthEndpoint() *oauth2.Endpoint {
	return &oauth2.Endpoint{
		AuthURL:   authURL,
		TokenURL:  tokensURL,
		AuthStyle: 0,
	}
}
