package examples

import (
	"fmt"
	"log"
	"os"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
)

var (
	customerID     string = "cst_qPgbV2VkwT"
	subscriptionID string = "sub_8EjeBVgtEn"
)

// Sets the mollie api token on an environment variable
func init() {
	_ = os.Setenv(mollie.APITokenEnv, "YOUR_API_TOKEN")
	_ = os.Setenv(mollie.OrgTokenEnv, "YOUR_ORG_TOKEN")

	config := mollie.NewConfig(true, mollie.APITokenEnv)
	client, _ = mollie.NewClient(nil, config)

}

//CreateSubscription code sample for creating a subscription
func CreateSubscription() {

	newSubscription := &mollie.Subscription{
		Amount:      &mollie.Amount{Currency: "EUR", Value: "25.00"},
		Times:       4,
		Interval:    "3 months",
		Description: "Quarterly payment",
		WebhookURL:  "https://webshop.example.org/subscriptions/webhook/",
	}

	sub, err := client.Subscriptions.Create(customerID, newSubscription)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sub)
}

//GetSubscription code sample for getting a specific subscription using a customer ID
func GetSubscription() {

	sub, err := client.Subscriptions.Get(customerID, subscriptionID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sub)
}

//UpdateSubscription code sample for updating a specific subscription
func UpdateSubscription() {

	updatedSubscription := &mollie.Subscription{
		Amount: &mollie.Amount{Currency: "EUR", Value: "10.00"},
	}

	sub, err := client.Subscriptions.Update(customerID, subscriptionID, updatedSubscription)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sub)

}

//CancelSubscription code sample for canceling a subscription
func CancelSubscription() {

	sub, err := client.Subscriptions.Delete(customerID, subscriptionID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sub)
}

//ListSubscriptions code sample for getting a list of subscriptions for a certain customer
func ListSubscriptions() {

	options := &mollie.SubscriptionListOptions{Limit: 5}

	list, err := client.Subscriptions.List(customerID, options)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list)
}

//ListAllSubscriptions code sample to retrieving all subscriptions for the current website profile
func ListAllSubscriptions() {

	listOptions := &mollie.SubscriptionListOptions{ProfileID: "pfl_K4zQxvyTVH"}
	all, err := client.Subscriptions.All(listOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(all)
}

//ListSubscriptionPayments code sample for retrieving all payments of a specific subscription for a certain customer
func ListSubscriptionPayments() {

	options := &mollie.SubscriptionListOptions{Limit: 5}

	list, err := client.Subscriptions.GetPayments(customerID, subscriptionID, options)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list.Embedded.Payments)
}
