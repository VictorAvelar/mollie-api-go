package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/VictorAvelar/mollie-api-go/v2/mollie/connect"
)

func main() {
	c := oauth2.Config{
		ClientID:     "your client id",
		ClientSecret: "your client secret",
		Endpoint:     connect.Endpoint,
	}
	secured := c.Client(context.TODO(), nil)

	config := mollie.NewConfig(true, "")

	client, err := mollie.NewClient(secured, config)
	if err != nil {
		log.Fatal(err)
	}

	output, err := client.Methods.List(nil)
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

	fmt.Printf("For detailed information see: %s", output.Links.Documentation.Href)
	os.Exit(0)
}
