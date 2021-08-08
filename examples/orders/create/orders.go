package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
	var order mollie.Order
	{
		addr := &mollie.OrderAddress{
			GivenName:       "John",
			FamilyName:      "Doe",
			Email:           "john.doe@example.com",
			StreetAndNumber: "C. d'Arístides Maillol, 12",
			PostalCode:      "08028",
			City:            "Barcelona",
			Country:         "ES",
		}
		order = mollie.Order{
			OrderNumber:     "myRandUniqueOrderNumber-1023128312",
			Locale:          MyLocale,
			BillingAddress:  addr,
			ShippingAddress: *addr, // if different just construct a new address object.
			RedirectURL:     "https://github.com/VictorAvelar/mollie-api-go",
		}

		product := createOrderLine(
			"Macbook PRO 2021 - M1 Spacegray",
			1,
			createAmount("1299.00", "EUR"),
		)

		order.Lines = []*mollie.OrderLine{product}
		order.Method = mollie.CreditCard
		order.Amount = product.TotalAmount
	}

	ord, err := client.Orders.Create(order, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Order created: Checkout at %s\n", ord.Links.Checkout.Href)
}

func createOrderLine(name string, quantity int, unitPrice *mollie.Amount) *mollie.OrderLine {
	return &mollie.OrderLine{
		Name:        name,
		Quantity:    quantity,                           // quantity
		VatRate:     VAT,                                // Tax rate as string of form 19.00 for 19%.
		VatAmount:   calculateVatAmount(VAT, unitPrice), // Calculate vat amount from gross unit price.
		UnitPrice:   unitPrice,                          // gross unit price.
		TotalAmount: unitPrice,                          // in case of single units the same as unit price, otherwise the sum of gross unit prices.
	}
}

func createAmount(value, currency string) *mollie.Amount {
	return &mollie.Amount{
		Value:    value,
		Currency: currency,
	}
}

func calculateVatAmount(rate string, rawPrice *mollie.Amount) (vat *mollie.Amount) {
	vatRateF := parseStrToFloat(VAT)
	rawPriceF := parseStrToFloat(rawPrice.Value)

	netPrice := rawPriceF / (float64(1) + (vatRateF / float64(100)))

	return &mollie.Amount{
		Value:    parseFloatToStr(rawPriceF - netPrice),
		Currency: rawPrice.Currency,
	}
}

func parseStrToFloat(value string) float64 {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("%+v", err)
	}

	return val
}

func parseFloatToStr(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
