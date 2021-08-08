package main

import (
	"fmt"
	"os"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
)

const (
	MyLocale = "es_ES"
	VAT      = "19.00"
)

var (
	client *mollie.Client
)

func init() {
	config := mollie.NewConfig(false, mollie.APITokenEnv)
	m, err := mollie.NewClient(nil, config)
	if err != nil {
		print(err)
		os.Exit(1)
	}

	client = m
}

func main() {
	orders, err := client.Orders.List(nil)
	if err != nil {
		print(err)
		os.Exit(1)
	}

	for _, o := range orders.Embedded.Orders {
		fmt.Printf("%+v\n", o)
	}
}
