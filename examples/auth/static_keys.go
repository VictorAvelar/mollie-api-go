package main

import (
	"fmt"
	"log"
	"os"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
)

func main() {
	/*
		Create a base config object, the first parameter is a boolean
		to indicate if the client should be build for testing or for
		live requests.
		The second parameter is name of the environment variable from
		which the client should parse the token.

		By default the client will attempt to parse an API token from
		MOLLIE_API_TOKEN, but if you are using a diff variable name,
		you can pass it as second parameter.
	*/
	config := mollie.NewConfig(false, mollie.APITokenEnv)

	/*
		Then we build our Mollie API client, in this case we will
		pass nil as we are not using a prebuilt http client and our
		recently created configuration.

		Now after checking for errors we are ready to make a test call
		to Mollie's API.
	*/
	m, err := mollie.NewClient(nil, config)
	if err != nil {
		log.Fatal(err)
	}

	output, err := m.Methods.All(nil)
	if err != nil {
		log.Fatal(err)
	}

	/*
		Now you should see a list of the enabled payment methods for your
		account.
	*/
	fmt.Printf("Receive %d payment methods\n", output.Count)
	for _, o := range output.Embedded.Methods {
		fmt.Printf("- Name: %s\n", o.Description)
	}

	fmt.Printf("For detailed information see: %s\n", output.Links.Docs.Href)
	os.Exit(0)
}
