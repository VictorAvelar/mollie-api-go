package mollie

import (
	"encoding/json"
	"fmt"
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
	bts, err := json.Marshal(d.Time.Format("2006-01-02"))
	if err != nil {
		return nil, fmt.Errorf("json_marshall_error: %w", err)
	}

	return bts, nil
}

// UnmarshalJSON overrides the default unmarshal action
// for the Date struct, as we need links to be pointers to the time.Time struct.
func (d *ShortDate) UnmarshalJSON(b []byte) error {
	s := string(b)

	s = strings.Trim(s, "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("time_parse_error: %w", err)
	}

	d.Time = date

	return nil
}

// Locale represents a country and language in ISO-15897 format.
type Locale string

// Mollie supported locales.
const (
	English       Locale = "en_US"
	EnglishGB     Locale = "en_GB"
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
	Self          *URL `json:"self,omitempty"`
	Previous      *URL `json:"previous,omitempty"`
	Next          *URL `json:"next,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// BusinessCategory specifies an industry or category.
type BusinessCategory string

// Available business categoreis.
const (
	PetShopsPetFoodAndSupplies                              BusinessCategory = "PET_SHOPS"
	VeterinaryServices                                      BusinessCategory = "VETERINARY_SERVICES"
	ACAndHeatingContractors                                 BusinessCategory = "AC_AND_HEATING_CONTRACTORS"
	CarpentryContractors                                    BusinessCategory = "CARPENTRY_CONTRACTORS"
	ElectricalContractors                                   BusinessCategory = "ELECTRICAL_CONTRACTORS"
	EquipmentToolsOrFurnitureRentalOrLeasing                BusinessCategory = "EQUIPMENT_TOOLS_FURNITURE_RENTAL_LEASING"
	GeneralContractors                                      BusinessCategory = "GENERAL_CONTRACTORS"
	SpecialTradeContractors                                 BusinessCategory = "SPECIAL_TRADE_CONTRACTORS"
	CharityAndDonations                                     BusinessCategory = "CHARITY_AND_DONATIONS"
	FundraisingCrowdfundingAndSocialServiceOrganizations    BusinessCategory = "FUNDRAISING_CROWDFUNDING_SOCIAL_SERVICE"
	Apps                                                    BusinessCategory = "APPS"
	BooksMediaMoviesMusic                                   BusinessCategory = "BOOKS_MEDIA_MOVIES_MUSIC"
	Games                                                   BusinessCategory = "GAMES"
	SoftwareAndSubscriptions                                BusinessCategory = "SOFTWARE_AND_SUBSCRIPTIONS"
	ChildCareServices                                       BusinessCategory = "CHILD_CARE_SERVICES"
	CollegesOrUniversities                                  BusinessCategory = "COLLEGES_UNIVERSITIES"
	ElementaryOrSecondarySchools                            BusinessCategory = "ELEMENTARY_SECONDARY_SCHOOLS"
	OtherEducationalServices                                BusinessCategory = "OTHER_EDUCATIONAL_SERVICES"
	VocationalSchoolsOrTradeSchools                         BusinessCategory = "VOCATIONAL_SCHOOLS_TRADE_SCHOOLS"
	AmusementParksCircusesCarnivalsAndFortuneTellers        BusinessCategory = "AMUSEMENT_PARKS"
	EventTicketing                                          BusinessCategory = "EVENT_TICKETING"
	GamingEstablishmentsBilliardsPoolBowlingArcades         BusinessCategory = "GAMING_ESTABLISHMENTS"
	MovieTheatres                                           BusinessCategory = "MOVIE_THEATRES"
	MusiciansBandsOrOrchestras                              BusinessCategory = "MUSICIANS_BANDS_ORCHESTRAS"
	OnlineGambling                                          BusinessCategory = "ONLINE_GAMBLING"
	OtherEntertainmentAndRecreation                         BusinessCategory = "OTHER_ENTERTAINMENT_RECREATION"
	SportingAndRecreationalCamps                            BusinessCategory = "SPORTING_RECREATIONAL_CAMPS"
	SportsForecastingOrPredictionServices                   BusinessCategory = "SPORTS_FORECASTING"
	CreditCounsellingOrCreditRepair                         BusinessCategory = "CREDIT_COUNSELLING_REPAIR"
	DigitalWallets                                          BusinessCategory = "DIGITAL_WALLETS"
	InvestmentServices                                      BusinessCategory = "INVESTMENT_SERVICES"
	MoneyServicesOrTransmission                             BusinessCategory = "MONEY_SERVICES"
	MortgagesInsurancesLoansAndFinancialAdvice              BusinessCategory = "MORTGAGES_INSURANCES_LOANS_FINANCIAL_ADVICE"
	SecurityBrokersOrDealers                                BusinessCategory = "SECURITY_BROKERS_DEALERS"
	TrustOffices                                            BusinessCategory = "TRUST_OFFICES"
	VirtualCurrenciesAndCryptoCurrencies                    BusinessCategory = "VIRTUAL_CRYPTO_CURRENCIES"
	CaterersPrepareAndDelivery                              BusinessCategory = "CATERERS"
	FastFoodRestaurants                                     BusinessCategory = "FAST_FOOD_RESTAURANTS"
	GroceryStoresSupermarketsAndFoodProductStores           BusinessCategory = "FOOD_PRODUCT_STORES"
	RestaurantsNightlifeAndOtherOnPremiseConsumption        BusinessCategory = "RESTAURANTS_NIGHTLIFE"
	BoatRentalsAndLeasing                                   BusinessCategory = "BOAT_RENTALS_LEASING"
	CruiseLines                                             BusinessCategory = "CRUISE_LINES"
	HotelsMotelsResortsInnsAndOtherLodgingAndHospitality    BusinessCategory = "LODGING"
	PropertyRentalsOrCamping                                BusinessCategory = "PROPERTY_RENTALS_CAMPING"
	Marketplaces                                            BusinessCategory = "MARKETPLACES"
	DentalLabAndOrOphthalmicEquipmentAndSupplies            BusinessCategory = "DENTAL_EQUIPMENT_SUPPLIES"
	DentistsAndOrthodontists                                BusinessCategory = "DENTISTS_ORTHODONTISTS"
	DoctorsPhysiciansAndOtherMedicalServices                BusinessCategory = "MEDICAL_SERVICES"
	DrugStoresPharmaciesAndPrescriptionMedicine             BusinessCategory = "DRUG_PHARMACIES_PRESCRIPTION"
	MedicalDevices                                          BusinessCategory = "MEDICAL_DEVICES"
	MedicalOrganizations                                    BusinessCategory = "MEDICAL_ORGANIZATIONS"
	MentalHealthServices                                    BusinessCategory = "MENTAL_HEALTH_SERVICES"
	NursingOrPersonalCareFacilitiesAndAssistedLiving        BusinessCategory = "NURSING"
	OpticiansAndEyeglasses                                  BusinessCategory = "OPTICIANS_EYEGLASSES"
	CivicFraternalOrSocialAssociations                      BusinessCategory = "SOCIAL_ASSOCIATIONS"
	GymsMembershipFeeBasedSports                            BusinessCategory = "MEMBERSHIP_FEE_BASED_SPORTS"
	OtherMembershipOrganizations                            BusinessCategory = "OTHER_MEMBERSHIP_ORGANIZATIONS"
	AdultContentOrServices                                  BusinessCategory = "ADULT_CONTENT_SERVICES"
	CounselingServices                                      BusinessCategory = "COUNSELING_SERVICES"
	DatingServices                                          BusinessCategory = "DATING_SERVICES"
	HealthAndBeautySpas                                     BusinessCategory = "HEALTH_BEAUTY_SPAS"
	LandscapingServices                                     BusinessCategory = "LANDSCAPING_SERVICES"
	LaundryOrDrycleaningServices                            BusinessCategory = "LAUNDRY_DRYCLEANING_SERVICES"
	MassageParlours                                         BusinessCategory = "MASSAGE_PARLOURS"
	OtherPersonalServices                                   BusinessCategory = "OTHER_PERSONAL_SERVICES"
	PhotographyStudios                                      BusinessCategory = "PHOTOGRAPHY_STUDIOS"
	SalonsOrBarbers                                         BusinessCategory = "SALONS_BARBERS"
	PoliticalParties                                        BusinessCategory = "POLITICAL_PARTIES"
	AccountingAuditingBookkeepingAndTaxPreparationServices  BusinessCategory = "ACCOUNTING_AUDITING_BOOKKEEPING_TAX_PREPARATION_SERVICES"
	AdvertisingServices                                     BusinessCategory = "ADVERTISING_SERVICES"
	CleaningAndMaintenanceJanitorialServices                BusinessCategory = "CLEANING_MAINTENANCE_JANITORIAL_SERVICES"
	ComputerRepair                                          BusinessCategory = "COMPUTER_REPAIR"
	Consultancy                                             BusinessCategory = "CONSULTANCY"
	DetectiveOrProtectiveAgenciesSecurityServices           BusinessCategory = "SECURITY_SERVICES"
	DirectMarketing                                         BusinessCategory = "DIRECT_MARKETING"
	FuneralServicesAndCrematories                           BusinessCategory = "FUNERAL_SERVICES"
	GovernmentServices                                      BusinessCategory = "GOVERNMENT_SERVICES"
	HostingAndVpnServices                                   BusinessCategory = "HOSTING_VPN_SERVICES"
	IndustrialSuppliesNotElsewhereClassified                BusinessCategory = "INDUSTRIAL_SUPPLIES_NOT_ELSEWHERE_CLASSIFIED"
	LegalServicesAndAttorneys                               BusinessCategory = "LEGAL_SERVICES_ATTORNEYS"
	MotionPictureOrVideoTapeProductionAndOrDistribution     BusinessCategory = "MOTION_PICTURES_DISTRIBUTION"
	OtherBusinessServices                                   BusinessCategory = "OTHER_BUSINESS_SERVICES"
	PrintingAndPublishing                                   BusinessCategory = "PRINTING_PUBLISHING"
	RealEstateAgents                                        BusinessCategory = "REAL_ESTATE_AGENTS"
	SanitationPolishingAndSpecialtyCleaning                 BusinessCategory = "SANITATION_POLISHING_SPECIALTY_CLEANING"
	StationeryOrOfficeSupplies                              BusinessCategory = "OFFICE_SUPPLIES"
	TestingLaboratoriesNotMedical                           BusinessCategory = "TESTING_LABORATORIES_NOT_MEDICAL"
	TrainingAndCoaching                                     BusinessCategory = "TRAINING_AND_COACHING"
	Utilities                                               BusinessCategory = "UTILITIES"
	ReligiousOrganizations                                  BusinessCategory = "RELIGIOUS_ORGANIZATIONS"
	SportsClothingShoesAndAccessories                       BusinessCategory = "CLOTHING_SHOES_ACCESSORIES"
	ArtDealersGalleriesCommercialPhotographyAndGraphics     BusinessCategory = "COMMERCIAL_ART"
	BeautyProducts                                          BusinessCategory = "BEAUTY_PRODUCTS"
	BooksPeriodicalsAndNewspapers                           BusinessCategory = "BOOKS_PERIODICALS_NEWSPAPERS"
	BuildingHomeImprovementAndEquipment                     BusinessCategory = "HOME_IMPROVEMENT"
	CardsGiftsNoveltyAndSouvenirShops                       BusinessCategory = "GIFTS_SHOPS"
	CbdOrMarijuanaRelatedProducts                           BusinessCategory = "CBD_MARIJUANA_PRODUCTS"
	CoffeeShopsOrGrowShops                                  BusinessCategory = "COFFEE_SHOPS"
	ConvenienceStoresSpecialtyMarketsHealthFoodStores       BusinessCategory = "CONVENIENCE_STORES"
	CreditsVouchersGiftCardsForNonFinancialInstitutions     BusinessCategory = "GIFT_CARDS"
	EroticToys                                              BusinessCategory = "EROTIC_TOYS"
	FloristsFloristSupplier                                 BusinessCategory = "FLORISTS"
	FuelDealers                                             BusinessCategory = "FUEL_DEALERS"
	FurnitureHomeFurnishingsAndEquipmentStores              BusinessCategory = "FURNITURE_FURNISHINGS_EQUIPMENT_STORES"
	GameToyAndHobbyShops                                    BusinessCategory = "GAME_TOY_HOBBY_SHOPS"
	GardenAndOutdoorEquipment                               BusinessCategory = "OUTDOOR_EQUIPMENT"
	HomeElectronicsAndPersonalComputers                     BusinessCategory = "HOME_ELECTRONICS"
	HouseholdApplianceStores                                BusinessCategory = "HOUSEHOLD_APPLIANCE_STORES"
	JewelryWatchClockAndSilverwareStoresUnder1000Euro       BusinessCategory = "JEWELRY_WATCH_CLOCK_AND_SILVERWARE_STORES_UNDER_1000"
	MusicStoresInstrumentsAndRecords                        BusinessCategory = "MUSIC_STORES"
	OtherMerchandise                                        BusinessCategory = "OTHER_MERCHANDISE"
	PackageStoresBeerWineAndLiquor                          BusinessCategory = "LIQUOR_STORES"
	PaidTelevisionOrRadioServicesCableOrSatellite           BusinessCategory = "PAID_TELEVISION_RADIO"
	PreciousStonesMetalsWatchesAndJewelryOver1000Euro       BusinessCategory = "PRECIOUS_STONES_METALS_JEWELRY_OVER_1000"
	RepairShopsAndRelatedServicesNotElsewhereClassified     BusinessCategory = "REPAIR_SHOPS"
	SecondHandOrUsedMerchandiseStores                       BusinessCategory = "SECOND_HAND_STORES"
	SportingGoodsStoresMiscellaneousAndSpecialtyRetailShops BusinessCategory = "SPORTING_GOODS_SPECIALTY_RETAIL_SHOPS"
	SupplementsNutritionVitaminStores                       BusinessCategory = "SUPPLEMENTS_STORES"
	TelecomEquipment                                        BusinessCategory = "TELECOM_EQUIPMENT"
	TelecomServices                                         BusinessCategory = "TELECOM_SERVICES"
	TobaccoCigarsECigarettesAndRelatedProducts              BusinessCategory = "TOBACCO_PRODUCTS"
	TradersInDiamonds                                       BusinessCategory = "TRADERS_DIAMONDS"
	TradersInGold                                           BusinessCategory = "TRADERS_GOLD"
	WeaponsOrAmmunition                                     BusinessCategory = "WEAPONS_AMMUNITION"
	CommuterTransportation                                  BusinessCategory = "COMMUTER_TRANSPORTATION"
	CourierServicesAndFreightForwarders                     BusinessCategory = "COURIER_SERVICES"
	OtherTransportationServices                             BusinessCategory = "OTHER_TRANSPORTATION_SERVICES"
	TaxisLimosAndRidesharing                                BusinessCategory = "RIDESHARING"
	TravelAgenciesTourOperatorsAndOtherTravalServices       BusinessCategory = "TRAVEL_SERVICES"
	AutomotivePartsAndAccessories                           BusinessCategory = "AUTOMOTIVE_PARTS_ACCESSORIES"
	AutoAndTruckSalesAndServiceDealersAndLeasingCompanies   BusinessCategory = "CAR_TRUCK_COMPANIES"
	AutomotiveServices                                      BusinessCategory = "AUTOMOTIVE_SERVICES"
	BicyclePartsShopsAndService                             BusinessCategory = "BICYCLE_PARTS_SHOPS_SERVICE"
	CarBoatCamperMobileHomeDealer                           BusinessCategory = "CAR_BOAT_CAMPER_MOBILE_HOME_DEALER"
	CarRentals                                              BusinessCategory = "CAR_RENTALS"
	MotorcyclePartsShopsAndDealers                          BusinessCategory = "MOTORCYCLE_PARTS_SHOPS_AND_DEALERS"
)

// Mode contains information about the creation environment.
type Mode string

// Valid modes.
const (
	LiveMode Mode = "live"
	TestMode Mode = "test"
)

// EmbedValue describes the valid value of embed query string.
type EmbedValue string

// Valid Embed query string value.
const (
	EmbedPayment     EmbedValue = "payments"
	EmbedRefund      EmbedValue = "refunds"
	EmbedShipments   EmbedValue = "shipments"
	EmbedChargebacks EmbedValue = "chargebacks"
)

// Rate describes service rates, further divided into fixed and percentage costs.
type Rate struct {
	Fixed    *Amount `json:"fixed,omitempty"`
	Variable string  `json:"variable,omitempty"`
}

// Image describes a generic image resource retrieved by Mollie.
type Image struct {
	Size1x string `json:"size1X,omitempty"`
	Size2X string `json:"size2X,omitempty"`
	Svg    string `json:"svg,omitempty"`
}
