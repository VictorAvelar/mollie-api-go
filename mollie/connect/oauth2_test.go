package connect

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func Test_Endpoint(t *testing.T) {
	require.Equal(t, Endpoint.AuthURL, AuthURL)
	require.Equal(t, Endpoint.AuthStyle, oauth2.AuthStyleAutoDetect)
	require.Equal(t, Endpoint.TokenURL, TokensURL)
}
