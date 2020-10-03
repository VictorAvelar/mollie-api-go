package examples

import (
	"fmt"
	"log"
	"os"

	"github.com/VictorAvelar/mollie-api-go/mollie"
)

var (
	config       *mollie.Config
	client       *mollie.Client
	paymentID    string = "tr_WDqYK6vllg"
	chargebackID string = "chb_n9z0tp"
)

// Sets the mollie api token on an environment variable.
func init() {
	_ = os.Setenv(mollie.APITokenEnv, "YOUR_API_TOKEN")
	config = mollie.NewConfig(true, mollie.APITokenEnv)
	client, _ = mollie.NewClient(nil, config)
}

// GetSingleGhargeback code samples shows how  a payment chargeback for the matching
// payment and chargeback id should be retrieved.
func GetSingleGhargeback() {
	chargeback, err := client.Chargebacks.Get(paymentID, chargebackID, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", chargeback)
}

// ListAllPaymentRelatedChargebacks code sample shows how to retrieve all received
// chargebacks for a payment.
func ListAllPaymentRelatedChargebacks() {
	list, err := client.Chargebacks.ListForPayment(paymentID, nil)

	if err != nil {
		log.Fatal(err)
	}

	for _, c := range list.Embedded.Chargebacks {
		fmt.Printf("%+v\n", c)
	}
}

// ListAllPaymentsForAccount code sample shows how to retrieve all received
// chargebacks for your account.
func ListAllPaymentsForAccount() {
	list, err := client.Chargebacks.List(nil)

	if err != nil {
		log.Fatal(err)
	}

	for _, c := range list.Embedded.Chargebacks {
		fmt.Printf("%+v\n", c)
	}
}

// ListAllPaymentsWithOptions code sample shows how to retrieve all received
// chargebacks for your account.
func ListAllPaymentsWithOptions() {
	list, err := client.Chargebacks.List(&mollie.ListChargebackOptions{
		Include: "payments", // will include the payments these chargebacks were issued for.
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, c := range list.Embedded.Chargebacks {
		fmt.Printf("%+v\n", c)
	}
}
