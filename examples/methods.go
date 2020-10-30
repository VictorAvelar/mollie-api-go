package examples

import (
	"fmt"
	"log"
	"os"

	"github.com/VictorAvelar/mollie-api-go/mollie"
)


func init() {
	_ = os.Setenv(mollie.APITokenEnv, "YOUR_API_TOKEN")

	config := mollie.NewConfig(true, mollie.APITokenEnv)
	client, _ = mollie.NewClient(nil, config)
}

// ListPaymentMethods code sample for getting a list of enabled payment methods on the website 
func ListPaymentMethods() {

	options := &mollie.MethodsOptions{Resource: "orders"}
	list, err := client.Methods.List(options)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list)
}

// ListAllPaymentMethods code sample for getting a list all payment methods offered by Mollie
func ListAllPaymentMethods() {

	options := &mollie.MethodsOptions{Locale: "en_US"}
	list, err := client.Methods.All(options)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list)
}

// GetPaymentMethod code sample for getting an enabled payment method on the website 
func GetPaymentMethod() {
	
	options := &mollie.MethodsOptions{Include: "issuers"}
	paymentMethod, err := client.Methods.Get("ideal", options)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(paymentMethod)
}
