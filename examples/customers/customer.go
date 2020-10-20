// Example app to show the following Mollie features:
// - Get customer
// - Update customer
// - Delete customer
// - List customers//
// - Create customer payment
// - List customer payments
package main

import (
	"fmt"
	"github.com/VictorAvelar/mollie-api-go/mollie"
	"log"
	"os"
)

func main() {
	fmt.Printf("Connecting to Mollie....\n")
	m := CreateMollie()

	//create customer
	cs := CreateCustomer(m)

	// get customer
	cs = GetCustomer(cs, m)

	//update customer
	UpdateCustomer(cs, m)

	//list all customers
	ListAllCustomers(m)

	//create payment for a single customer
	CreateCustomerPayment(cs, m)

	//get all payments for a particular customer
	GetCustomerPayments(cs, m)

	os.Exit(0)
}

//GetCustomerPayments extract all payments for a particular customer from mollie
func GetCustomerPayments(cs *mollie.Customer, m *mollie.Client) {
	// List customer payment
	fmt.Printf("Listing customer ID %s payments\n", cs.ID)
	pl, err := m.Customers.GetPayments(cs.ID, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("....customer with id %s has %d payments", cs.ID, pl.Count)
}

//CreateCustomerPayment create a single payment for a particular customer from mollie
func CreateCustomerPayment(cs *mollie.Customer, m *mollie.Client)  {
	// Create customer payment
	fmt.Printf("Creating payment for customer  with id %s \n", cs.ID)
	var bt = mollie.BankTransfer

	var p = mollie.Payment{
		Amount: &mollie.Amount{
			Currency: "EUR",
			Value:    "10.12",
		},
		Description: "Order #12345",
		Method:      &bt,
		RedirectURL: "https://webshop.example.org/order/12345/",
		WebhookURL:  "https://webshop.example.org/payments/webhook/",
	}
	pp, err := m.Customers.CreatePayment(cs.ID, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("....created new payment with the following info - ID %s, currency : %s, "+
		"value : %s ",
		pp.ID, pp.Amount.Currency, pp.Amount.Value)
}

//ListAllCustomers list all customers from mollie
func ListAllCustomers(m *mollie.Client) {
	fmt.Printf("List all customers....\n")
	output, err := m.Customers.List(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("....total number customer - %d\n", output.Count)
}

//UpdateCustomer update email of a particular customer from mollie
func UpdateCustomer(cs *mollie.Customer, m *mollie.Client) {
	fmt.Printf("Updating customer with id %s\n", cs.ID)
	var uc = mollie.Customer{
		Email: "new_emailkingkong@jameskong.com",
	}
	cs, err := m.Customers.Update(cs.ID, uc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("....existing customer ID %s email has been updated to %s", cs.ID, cs.Email)
}

//GetCustomer extract a particular customer from mollie
func GetCustomer(cs *mollie.Customer, m *mollie.Client) *mollie.Customer {
	//2. Get customer
	fmt.Printf("Retrieving customer information with id %s\n", cs.ID)
	cs, err := m.Customers.Get(cs.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("....existing customer - name %s, ID %s, email %s\n", cs.Name, cs.ID, cs.Email)
	return cs
}

//CreateCustomer create a single customer from mollie
func CreateCustomer(m *mollie.Client) *mollie.Customer {
	var c = mollie.Customer{
		Name:  "testkingkong",
		Email: "jameskong@jameskong.com",
	}
	fmt.Printf("Creating new customer with the following info - name : %s , email : %s\n", c.Name,
		c.Email)
	cs, err := m.Customers.Create(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("....new customer created has ID %s\n", cs.ID)
	return cs
}

// CreateMollie create a mollie connection
func CreateMollie() *mollie.Client {
	config := mollie.NewConfig(false, mollie.APITokenEnv)
	m, err := mollie.NewClient(nil, config)
	if err != nil {
		log.Fatal(err)
	}
	return m
}
