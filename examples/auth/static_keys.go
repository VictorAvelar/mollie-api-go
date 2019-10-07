package main

import (
	"fmt"

	"github.com/VictorAvelar/mollie-api-go/mollie"
)

const (
	APIKey string = "your api key here"
	OrgKey string = "your org key here"
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
		return
	}

	output, err := m.Methods.List(nil)
	if err != nil {
		return
	}

	/*
		Now you should see a list of the enabled payment methods for your
		account.
	*/
	for _, o := range output.Embedded.Methods {
		fmt.Printf("Id: %s | Name: %s", o.ID, o.Description)
	}
}
