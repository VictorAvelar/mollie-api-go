package mollie

import (
	"encoding/json"
	"strings"
	"time"
)

// Amount represents a currency and value pair.
type Amount struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

// Address provides a human friendly representation of a geographical space.
//
// When providing an address object as parameter to a request, the following conditions must be met:
//
// If any of the fields is provided, all fields have to be provided with exception of the region field.
// If only the region field is given, one should provide all the other fields as per the previous condition.
// For certain PayPal payments the region field is required.
type Address struct {
	StreetAndNumber  string `json:"streetAndNumber,omitempty"`
	StreetAdditional string `json:"streetAdditional,omitempty"`
	PostalCode       string `json:"postalCode,omitempty"`
	City             string `json:"city,omitempty"`
	Region           string `json:"region,omitempty"`
	Country          string `json:"country,omitempty"`
}

// ShortDate is a string representing a date in YYYY-MM-DD format.
type ShortDate struct {
	time.Time
}

// MarshalJSON overrides the default marshal action
// for the Date struct. Returns date as YYYY-MM-DD formatted string.
func (d *ShortDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

// UnmarshalJSON overrides the default unmarshal action
// for the Date struct, as we need links to be pointers to the time.Time struct.
func (d *ShortDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// Locale represents a country and language in ISO-15897 format.
type Locale string

// Mollie supported locales
const (
	English       Locale = "en_US"
	Dutch         Locale = "nl_NL"
	DutchBelgium  Locale = "nl_BE"
	French        Locale = "fr_FR"
	FrenchBelgium Locale = "fr_BE"
	German        Locale = "de_DE"
	GermanAustria Locale = "de_AT"
	GermanSwiss   Locale = "de_CH"
	Spanish       Locale = "es_ES"
	Catalan       Locale = "ca_ES"
	Portuguese    Locale = "pt_PT"
	Italian       Locale = "it_IT"
	Norwegian     Locale = "nb_NO"
	Swedish       Locale = "sv_SE"
	Finish        Locale = "fi_FI"
	Danish        Locale = "da_DK"
	Icelandic     Locale = "is_IS"
	Hungarian     Locale = "hu_HU"
	Polish        Locale = "pl_PL"
	Latvian       Locale = "lv_LV"
	Lithuanian    Locale = "lt_LT"
)

// PhoneNumber represents a phone number in the E.164 format.
type PhoneNumber string

// QRCode object represents an image of a QR code.
type QRCode struct {
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	Src    string `json:"src,omitempty"`
}

// URL in Mollie are commonly represented as objects with an href and type field.
type URL struct {
	Href string `json:"href,omitempty"`
	Type string `json:"type,omitempty"`
}

// PaginationLinks describes the hal component of paginated responses.
type PaginationLinks struct {
	Self          URL `json:"self,omitempty"`
	Previous      URL `json:"previous,omitempty"`
	Next          URL `json:"next,omitempty"`
	Documentation URL `json:"documentation,omitempty"`
}

// CategoryCode specifies an industry or category
type CategoryCode uint

// Available category codes
const (
	BookMagazinesAndNewspapers      CategoryCode = 5192
	GeneralMerchandise              CategoryCode = 5399
	FoodAndDrinks                   CategoryCode = 5499
	AutomotiveProducts              CategoryCode = 5533
	ChildrenProducts                CategoryCode = 5641
	ClothingAndShoes                CategoryCode = 5651
	ElectronicsComputersAndSoftware CategoryCode = 5732
	HostingOrVpnServices            CategoryCode = 5734
	Entertainment                   CategoryCode = 5735
	CreditsOrVouchersOrGiftCards    CategoryCode = 5815
	Alcohol                         CategoryCode = 5921
	JewelryAndAccessories           CategoryCode = 5944
	HealthAndBeautyProducts         CategoryCode = 5977
	FinancialServices               CategoryCode = 6012
	Consultancy                     CategoryCode = 7299
	TravelRentalAndTransportation   CategoryCode = 7999
	AdvisingOrCoachingOrTraining    CategoryCode = 8299
	CharityAndDonations             CategoryCode = 8398
	PoliticalParties                CategoryCode = 8699
	Others                          CategoryCode = 0
)

// Mode contains information about the creation environment.
type Mode string

// Valid modes
const (
	LiveMode Mode = "live"
	TestMode Mode = "test"
)

// EmbedValue describes the valid value of embed query string.
type EmbedValue string

// Valid Embed query string value.
const (
	EmbedPayment     EmbedValue = "payment"
	EmbedRefund      EmbedValue = "refund"
	EmbedShipments   EmbedValue = "shipments"
	EmbedChangebacks EmbedValue = "chanrgebacks"
)

// Rate describes service rates, further divided into fixed and percentage costs.
type Rate struct {
	Fixed    Amount `json:"fixed,omitempty"`
	Variable string `json:"variable,omitempty"`
}
