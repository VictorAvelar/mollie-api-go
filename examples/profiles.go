package examples

import (
	"fmt"
	"log"
	"os"

	"github.com/VictorAvelar/mollie-api-go/mollie"
)
var(
	profileID string = "pfl_jmWPWbFgeh"
)

// Sets the mollie api token on an environment variable.
func init() {
	_ = os.Setenv(mollie.APITokenEnv, "YOUR_API_TOKEN")
	_ = os.Setenv(mollie.OrgTokenEnv, "YOUR_ORG_TOKEN")

	config := mollie.NewConfig(true, mollie.OrgTokenEnv)
	client, _ = mollie.NewClient(nil, config)
	
}

//CreateProfile code sample for creating a new profile.
func CreateProfile() {

	newProfile := &mollie.Profile{
		Name:         "My website name",            //Brand name or trademark of the website profile
		Website:      "https://www.mywebsite.com",  //The profiles website URL (The URL must be compliant to RFC3986)
		Email:        "info@mywebsite.com",         //Email address for the profile
		Phone:        "+31208202071",               //Phone number for the profile
		CategoryCode: mollie.GeneralMerchandise,    //Optional value for the websites associated industry
		Mode:         mollie.TestMode,              //Mode: "Test" for testing or "Live" for the live use of the api
	}

	profile, err := client.Profiles.Create(newProfile) 

	if err != nil {
		log.Fatal(err) 
	}

	fmt.Println(profile)
}

//GetProfile code sample for getting a specific profile using a profile ID.
func GetProfile() {

	profile, err := client.Profiles.Get(profileID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profile)
}

//CurrentProfile code sample for getting the current profile belonging to the API key.
func CurrentProfile() {

	
	newConfig := mollie.NewConfig(true, mollie.APITokenEnv) //Requires a API key 
	newClient, _ := mollie.NewClient(nil,newConfig)
	
	profile, err := newClient.Profiles.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(profile)

}

//UpdateProfile code sample for updating a specific profile.
func UpdateProfile() {

	updatedProfile := &mollie.Profile{
		Name:    "My Website Name",
		Website: "http://www.mywebsite.com",
		Email:   "info@mywebsite.com",
	}

	profile, err := client.Profiles.Update(profileID, updatedProfile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profile)
}

//DeleteProfile code sample for deleting a specific profile using a profile ID.
func DeleteProfile() {

	err := client.Profiles.Delete(profileID)

	if err != nil {
		log.Fatal(err)
	}
}

//ListProfiles code sample for list all profiles.
func ListProfiles() {

	list, err := client.Profiles.List(nil) //Optional list options 
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < list.Count; i++ {
		profile := list.Embedded.Profiles[i]
		fmt.Println(profile)
	}
}

//EnablePaymentMethod code sample for enabling a payment method on a profile.
func EnablePaymentMethod() {

	paymentMethodInfo, err := client.Profiles.EnablePaymentMethod(profileID, mollie.IDeal)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(paymentMethodInfo)
}

//DisablePaymentMethod code sample for disabling a payment method on a profile.
func DisablePaymentMethod() {

	err := client.Profiles.DisablePaymentMethod(profileID, mollie.IDeal)

	if err != nil {
		log.Fatal(err)
	}

}
