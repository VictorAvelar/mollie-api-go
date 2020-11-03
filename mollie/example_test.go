package mollie

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie/connect"
)

// ----- mollie examples -----
func ExampleNewClient() {
	// using api token
	apiTokenClient, err := NewClient(nil, &Config{
		testing: false,
		auth:    APITokenEnv,
	})
	if err != nil {
		return
	}

	// using an organization token
	orgTokenClient, err := NewClient(nil, &Config{
		testing: false,
		auth:    OrgTokenEnv,
	})
	if err != nil {
		return
	}

	// using mollie connect
	oauthConfig := oauth2.Config{
		ClientID:     "your client id",
		ClientSecret: "your client secret",
		Endpoint:     connect.Endpoint,
	}
	secured := oauthConfig.Client(context.TODO(), nil)

	config := NewConfig(true, "")

	oauthClient, err := NewClient(secured, config)
	if err != nil {
		return
	}

	fmt.Printf("%v\n", apiTokenClient)
	fmt.Printf("%v\n", orgTokenClient)
	fmt.Printf("%v\n", oauthClient)
}
