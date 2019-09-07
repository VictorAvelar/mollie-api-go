package mollie

import (
	"net/url"
	"strings"
	"time"
)

//Locale is a string representing the country and language
//in ISO 15897 format.
type Locale string

//PhoneNumber is a string in the E.164 format.
//For example, `+31208202070`.
type PhoneNumber string

//Amount contains a currency code compliant to the ISO 4217
//specification and a value with the exact amount in the
//given currency.
type Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

//Address contains information to provide a means of physically
//locating a place in the urban geographical space.
//
//For Mollie the country must follow an ISO 3166-1 alpha-2 format,
//postal codes must belong to the specified city and country.
//When providing an address object as parameter to a request,
//the following conditions must be met:
//
//	- If any of the fields is provided, all fields have to be provided
//	  with exception of the region field.
//	- If only the region field is given, one should provide all the other
//	  fields as per the previous condition.
//	- For certain PayPal payments the region field is required.
type Address struct {
	StreetAndNumber  string `json:"streetAndNumber"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Region           string `json:"region,omitempty"`
	StreetAdditional string `json:"streetAdditional,omitempty"`
}

//Date is represented on `YYYY-MM-DD` format for Mollie requests,
//in order to comply with this the time.Time struct is embedded
//on Date and custom marshal and unmarshal will be provided.
type Date struct {
	time.Time
}

//UnmarshalJSON overrides the default unmarshal action
//for the Date struct, as we need links to be pointers
//to the time.Time struct.
func (d *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

//QRCode represents an image of a QR code.
type QRCode struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Source int `json:"src"`
}

//Link for Mollie are commonly represented as objects
//with an href and type field.
type Link struct {
	Href        HalURL `json:"href"`
	ContentType string `json:"type"`
}

//HalURL has url.URL embedded to modify the default json
//marshal behaviour.
type HalURL struct {
	*url.URL
}

//UnmarshalJSON overrides the default unmarshal action
//for the HalURL struct, as we need links to be pointers
//to the url.URL struct.
func (hl *HalURL) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	uri, err := url.Parse(s)
	if err != nil {
		return err
	}
	hl.URL = uri
	return nil
}

//HAL is a map of Links that contains relevant resources
//according to the response.
type HAL map[string]Link

//ResponseErr describes the content of the response body
//when the http call to Mollie API is not successful.
type ResponseErr struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Field  string `json:"field,omitempty"`
	Links  HAL    `json:"_links"`
}
