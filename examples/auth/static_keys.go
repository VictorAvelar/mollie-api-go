package main

import (
	"fmt"
	"log"
	"os"

	"github.com/VictorAvelar/mollie-api-go/mollie"
)

// Available authentication options using static keys
const (
	APIKey string = "your_api_key_here"
)

func main() {
	/*
		To build our config, we need to tell our client if its
		purpose is to be used for testing or live requests, and
		then if using static tokens we will pass our key as
		second parameter, this will attach our token to every
		outgoing requests as a Bearer token.
	*/
	config := mollie.NewConfig(true, APIKey)

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

	output, err := m.Methods.List(nil)
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

	fmt.Printf("For detailed information see: %s", output.Links.Docs.Href)
	os.Exit(0)
}
