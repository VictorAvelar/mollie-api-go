package mollie

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/VictorAvelar/mollie-api-go/mollie/connect"
)

// ----- mollie examples -----
func ExampleAPITokenClient() {
	c, err := NewClient(nil, &Config{
		testing: false,
		auth:    APITokenEnv,
	})
	if err != nil {
		return
	}
	fmt.Printf("%v", c)
}

func ExampleOrganizationTokenClient() {
	c, err := NewClient(nil, &Config{
		testing: false,
		auth:    OrgTokenEnv,
	})
	if err != nil {
		return
	}
	fmt.Printf("%v", c)
}

func ExampleOAuth2Client() {
	c := oauth2.Config{
		ClientID:     "your client id",
		ClientSecret: "your client secret",
		Endpoint:     connect.Endpoint,
	}
	secured := c.Client(context.TODO(), nil)

	config := NewConfig(true, "")

	client, err := NewClient(secured, config)
	if err != nil {
		return
	}

	fmt.Printf("%v", client)
}
