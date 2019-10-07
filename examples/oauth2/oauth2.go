package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"

	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/VictorAvelar/mollie-api-go/mollie/connect"
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
