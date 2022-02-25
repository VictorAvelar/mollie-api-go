package connect

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func Test_Endpoint(t *testing.T) {
	ept := OauthEndpoint()
	require.Equal(t, ept.AuthURL, authURL)
	require.Equal(t, ept.AuthStyle, oauth2.AuthStyleAutoDetect)
	require.Equal(t, ept.TokenURL, tokensURL)
}
