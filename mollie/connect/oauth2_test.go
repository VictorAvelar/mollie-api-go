package connect

import (
	"context"
	"fmt"
	"testing"

	"golang.org/x/oauth2"
)

func Test_Endpoint(t *testing.T) {
	if Endpoint.AuthURL != AuthURL {
		t.Fail()
	}

	if Endpoint.AuthStyle != oauth2.AuthStyleAutoDetect {
		t.Fail()
	}

	if Endpoint.TokenURL != TokensURL {
		t.Fail()
	}
}

func ExampleEndpoint() {
	config := oauth2.Config{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		Endpoint:     Endpoint,
	}

	client := config.Client(context.Background(), nil)

	fmt.Print(client)
}
