# mollie

--
 import "github.com/VictorAvelar/mollie-api-golang/v3/mollie"

Package mollie is a wrapper around Mollie's REST API.

See: <https://www.mollie.com/en/developers>

The Mollie API implements a Representational state transfer (REST) architecture.

Sounds technical, but it’s really quite easy. It mainly breaks down to
HTTP-methods GET, PATCH, POST and DELETE matching the operations to read,
update, create and delete.

REST also implies a nice and clean structure for URLs or endpoints. This means
you can reach any part of the Mollie API on <https://api.mollie.com/v2/> adding
the name of the resource you want to interact with.

## Usage

```go
const (
 BaseURL              string = "https://api.mollie.com/"
 AuthHeader           string = "Authorization"
 TokenType            string = "Bearer"
 APITokenEnv          string = "MOLLIE_API_TOKEN"
 OrgTokenEnv          string = "MOLLIE_ORG_TOKEN"
 RequestContentType   string = "application/json"
 IdempotencyKeyHeader string = "Idempotency-Key"
)
```

Constants holding values for client initialization and request instantiation.

#### func  CheckResponse

```go
func CheckResponse(r *Response) error
```

CheckResponse checks the API response for errors, and returns them if present. A
response is considered an error if it has a status code outside the 200 range.
API error responses are expected to have either no response body, or a JSON
response body.

#### type Address

```go
type Address struct {
 StreetAndNumber  string `json:"streetAndNumber,omitempty"`
 StreetAdditional string `json:"streetAdditional,omitempty"`
 PostalCode       string `json:"postalCode,omitempty"`
 City             string `json:"city,omitempty"`
 Region           string `json:"region,omitempty"`
 Country          string `json:"country,omitempty"`
}
```

Address provides a human friendly representation of a geographical space.

When providing an address object as parameter to a request, the following
conditions must be met:

If any of the fields is provided, all fields have to be provided with exception
of the region field. If only the region field is given, one should provide all
the other fields as per the previous condition. For certain PayPal payments the
region field is required.

#### type Amount

```go
type Amount struct {
 Currency string `json:"currency,omitempty"`
 Value    string `json:"value,omitempty"`
}
```

Amount represents a currency and value pair.

#### type ApplePaymentSession

```go
type ApplePaymentSession struct {
 EpochTimestamp    int    `json:"epochTimestamp,omitempty"`
 ExpiresAt         int    `json:"expiresAt,omitempty"`
 MerchantSessionID string `json:"merchantSessionIdentifier,omitempty"`
 Nonce             string `json:"nonce,omitempty"`
 MerchantID        string `json:"merchantIdentified,omitempty"`
 DomainName        string `json:"domainName,omitempty"`
 DisplayName       string `json:"displayName,omitempty"`
 Signature         string `json:"signature,omitempty"`
}
```

ApplePaymentSession contains information about an Apple pay session.

#### type ApplePaymentSessionRequest

```go
type ApplePaymentSessionRequest struct {
 Domain        string `json:"domain,omitempty"`
 ValidationURL string `json:"validationUrl,omitempty"`
}
```

ApplePaymentSessionRequest contains the body parameters for requesting a valid
PaymentSession from Apple.

#### type ApplicationFee

```go
type ApplicationFee struct {
 Amount      *Amount `json:"amount,omitempty"`
 Description string  `json:"description,omitempty"`
}
```

ApplicationFee allows you to split a payment between a platform and connected
merchant accounts.

#### type Balance

```go
type Balance struct {
 ID                  string               `json:"id,omitempty"`
 Resource            string               `json:"resource,omitempty"`
 Currency            string               `json:"currency,omitempty"`
 TransferReference   string               `json:"transferReference,omitempty"`
 Status              BalanceStatus        `json:"status,omitempty"`
 TransferFrequency   TransferFrequency    `json:"transferFrequency,omitempty"`
 TransferThreshold   *Amount              `json:"transferThreshold,omitempty"`
 AvailableAmount     *Amount              `json:"availableAmount,omitempty"`
 PendingAmount       *Amount              `json:"pendingAmount,omitempty"`
 TransferDestination *TransferDestination `json:"transferDestination,omitempty"`
 CreatedAt           *time.Time           `json:"createdAt,omitempty"`
 Links               BalanceLinks         `json:"_links,omitempty"`
}
```

Balance holds the payments processed with Mollie once fees have been deducted.

#### type BalanceAmount

```go
type BalanceAmount struct {
 Amount *Amount `json:"amount,omitempty"`
}
```

BalanceAmount wraps the std Amount type.

#### type BalanceGroupingFormat

```go
type BalanceGroupingFormat string
```

BalanceGroupingFormat defines a grouping mechanism for transactions included in
a balance report.

```go
const (
 StatusBalancesGrouping        BalanceGroupingFormat = "status-balances"
 TransactionCategoriesGrouping BalanceGroupingFormat = "transaction-categories"
)
```

Supported transaction grouping.

#### type BalanceLinks

```go
type BalanceLinks struct {
 Self          *URL `json:"self,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

BalanceLinks holds URL objects relevant to the balance.

#### type BalanceListOptions

```go
type BalanceListOptions struct {
 Currency string `url:"currency,omitempty"`
 From     string `url:"from,omitempty"`
 Limit    int    `url:"limit,omitempty"`
}
```

BalanceListOptions contains valid query parameters for the list balances
endpoint.

#### type BalanceReport

```go
type BalanceReport struct {
 Resource  string                       `json:"resource,omitempty"`
 BalanceID string                       `json:"balanceId,omitempty"`
 TimeZone  string                       `json:"timeZone,omitempty"`
 From      *ShortDate                   `json:"from,omitempty"`
 Until     *ShortDate                   `json:"until,omitempty"`
 Totals    *BalanceReportTotalsGrouping `json:"totals,omitempty"`
 Grouping  BalanceGroupingFormat        `json:"grouping,omitempty"`
 Links     BalanceReportLinks           `json:"_links,omitempty"`
}
```

BalanceReport contains the common fields between different balance grouping
options.

#### type BalanceReportDetail

```go
type BalanceReportDetail struct {
 Open                 *BalanceAmount `json:"open,omitempty"`
 Pending              *BalanceAmount `json:"pending,omitempty"`
 MovedToAvailable     *BalanceAmount `json:"movedToAvailable,omitempty"`
 ImmediatelyAvailable *BalanceAmount `json:"immediatelyAvailable,omitempty"`
 Close                *BalanceAmount `json:"close,omitempty"`
}
```

BalanceReportDetail contains the breakdown categories when grouping balance
transactions.

#### type BalanceReportLinks

```go
type BalanceReportLinks struct {
 Self          *URL `json:"self,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

BalanceReportLinks holds URL objects relevant to the balance report.

#### type BalanceReportOptions

```go
type BalanceReportOptions struct {
 Grouping BalanceGroupingFormat `url:"grouping,omitempty"`
 From     *ShortDate            `url:"from,omitempty"`
 Until    *ShortDate            `url:"until,omitempty"`
}
```

BalanceReportOptions contains valid query parameters for the list balances
endpoint.

#### type BalanceReportTotalsGrouping

```go
type BalanceReportTotalsGrouping struct {
 PendingBalance   *BalanceReportDetail `json:"pendingBalance,omitempty"`
 AvailableBalance *BalanceReportDetail `json:"availableBalance,omitempty"`
 Open             *BalanceReportDetail `json:"open,omitempty"`
 Payments         *BalanceReportDetail `json:"payments,omitempty"`
 Refunds          *BalanceReportDetail `json:"refunds,omitempty"`
 Chargebacks      *BalanceReportDetail `json:"chargebacks,omitempty"`
 Capital          *BalanceReportDetail `json:"capital,omitempty"`
 Transfers        *BalanceReportDetail `json:"transfers,omitempty"`
 FeePrePayments   *BalanceReportDetail `json:"fee-prepayments,omitempty"`
 Corrections      *BalanceReportDetail `json:"corrections,omitempty"`
 Close            *BalanceReportDetail `json:"close,omitempty"`
}
```

BalanceReportTotalsGrouping contains the per totals grouped balances for the
requested period.

#### type BalanceStatus

```go
type BalanceStatus string
```

BalanceStatus reflects whether a balance is operational or not.

```go
const (
 BalanceActive   BalanceStatus = "active"
 BalanceInactive BalanceStatus = "inactive"
)
```

Possible values for type BalanceStatus.

#### type BalanceTransaction

```go
type BalanceTransaction struct {
 Resource        string        `json:"resource,omitempty"`
 ID              string        `json:"id,omitempty"`
 TransactionType string        `json:"transactionType,omitempty"`
 ResultAmount    *Amount       `json:"resultAmount,omitempty"`
 InitialAmount   *Amount       `json:"initialAmount,omitempty"`
 Deductions      *Amount       `json:"deductions,omitempty"`
 CreatedAt       *time.Time    `json:"createdAt,omitempty"`
 Context         ContextValues `json:"context,omitempty"`
}
```

BalanceTransaction represents a the movement on your balance.

#### type BalanceTransactionsList

```go
type BalanceTransactionsList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  BalanceTransactions []*BalanceTransaction `json:"balance_transactions,omitempty"`
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

BalanceTransactionsList contains an array of embedded transactions.

#### type BalanceTransactionsListOptions

```go
type BalanceTransactionsListOptions struct {
 From  string `url:"from,omitempty"`
 Limit int    `url:"limit,omitempty"`
}
```

BalanceTransactionsListOptions are valid query parameters for list balance
transactions requests.

#### type BalancesList

```go
type BalancesList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Balances []*Balance `json:"balances,omitempty"`
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

BalancesList describes a list of captures.

#### type BalancesService

```go
type BalancesService service
```

BalancesService allows you to retrieve real-time as well as historical
information about your Mollie balance.

Works with Organization access tokens and App access tokens.

The API is in **BETA** so be careful and expect changes.

See: <https://docs.mollie.com/reference/v2/balances-api/overview>

#### func (*BalancesService) Get

```go
func (bs *BalancesService) Get(ctx context.Context, balance string) (res *Response, b *Balance, err error)
```

GetBalance retrieves a balance by its id.

See: <https://docs.mollie.com/reference/v2/balances-api/get-balance>

#### func (*BalancesService) GetPrimaryReport

```go
func (bs *BalancesService) GetPrimaryReport(ctx context.Context, options *BalanceReportOptions) (
 res *Response,
 br *BalanceReport,
 err error,
)
```

GetPrimaryReport returns the report for the primary balance.

See:
<https://docs.mollie.com/reference/v2/balances-api/get-primary-balance-report>

#### func (*BalancesService) GetPrimaryTransactionsList

```go
func (bs *BalancesService) GetPrimaryTransactionsList(ctx context.Context, options *BalanceTransactionsListOptions) (
 res *Response,
 btl *BalanceTransactionsList,
 err error,
)
```

GetPrimaryTransactionsList retrieves the list of movements (transactions) for
the primary balance of the account.

See:
<https://docs.mollie.com/reference/v2/balances-api/list-primary-balance-transactions>

#### func (*BalancesService) GetReport

```go
func (bs *BalancesService) GetReport(ctx context.Context, balance string, options *BalanceReportOptions) (
 res *Response,
 br *BalanceReport,
 err error,
)
```

GetReport returns the balance report for the specified balance id.

See: <https://docs.mollie.com/reference/v2/balances-api/get-balance-report>

#### func (*BalancesService) GetTransactionsList

```go
func (bs *BalancesService) GetTransactionsList(
 ctx context.Context,
 balance string,
 options *BalanceTransactionsListOptions,
) (
 res *Response,
 btl *BalanceTransactionsList,
 err error,
)
```

GetTransactionsList retrieves a list of movements (transactions) for the
specified balance.

See: <https://docs.mollie.com/reference/v2/balances-api/list-balance-transactions>

#### func (*BalancesService) List

```go
func (bs *BalancesService) List(ctx context.Context, options *BalanceListOptions) (
 res *Response,
 bl *BalancesList,
 err error,
)
```

List retrieves all the organization’s balances, including the primary balance,
ordered from newest to oldest.

See: <https://docs.mollie.com/reference/v2/balances-api/list-balances>

#### func (*BalancesService) Primary

```go
func (bs *BalancesService) Primary(ctx context.Context) (res *Response, b *Balance, err error)
```

Primary retrieves the primary balance. This is the balance of your account’s
primary currency, where all payments are settled to by default.

See: <https://docs.mollie.com/reference/v2/balances-api/get-primary-balance>

#### type BaseError

```go
type BaseError struct {
 Status int         `json:"status,omitempty"`
 Title  string      `json:"title,omitempty"`
 Detail string      `json:"detail,omitempty"`
 Field  string      `json:"field,omitempty"`
 Links  *ErrorLinks `json:"_links,omitempty"`
}
```

BaseError contains the general error structure returned by mollie.

#### func (*BaseError) Error

```go
func (be *BaseError) Error() string
```

Error interface compliance.

#### type Capture

```go
type Capture struct {
 Resource         string       `json:"resource,omitempty"`
 ID               string       `json:"id,omitempty"`
 Mode             Mode         `json:"mode,omitempty"`
 Amount           *Amount      `json:"amount,omitempty"`
 SettlementAmount *Amount      `json:"settlementAmount,omitempty"`
 PaymentID        string       `json:"paymentId,omitempty"`
 ShipmentID       string       `json:"shipmentId,omitempty"`
 SettlementID     string       `json:"settlementId,omitempty"`
 CreatedAt        *time.Time   `json:"createdAt,omitempty"`
 Links            CaptureLinks `json:"_links,omitempty"`
}
```

Capture describes a single capture. Captures are used for payments that have the
authorize-then-capture flow.

#### type CaptureLinks

```go
type CaptureLinks struct {
 Self          *URL `json:"self,omitempty"`
 Payment       *URL `json:"payment,omitempty"`
 Shipment      *URL `json:"shipment,omitempty"`
 Settlement    *URL `json:"settlement,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

CaptureLinks contains relevant links for a capture object.

#### type CapturesList

```go
type CapturesList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Captures []*Capture
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

CapturesList describes a list of captures.

#### type CapturesService

```go
type CapturesService service
```

CapturesService operates over captures resource.

#### func (*CapturesService) Get

```go
func (cs *CapturesService) Get(ctx context.Context, payment, capture string) (res *Response, c *Capture, err error)
```

Get retrieves a single capture by its ID. Note the original payment’s ID is
needed as well.

See: <https://docs.mollie.com/reference/v2/captures-api/get-capture>

#### func (*CapturesService) List

```go
func (cs *CapturesService) List(ctx context.Context, payment string) (res *Response, cl *CapturesList, err error)
```

List retrieves all captures for a certain payment.

See: <https://docs.mollie.com/reference/v2/captures-api/list-captures>

#### type CardLabel

```go
type CardLabel string
```

CardLabel Note that not all labels can be processed through Mollie.

```go
const (
 AmericaExpress CardLabel = "American Express"
 CartaSi        CardLabel = "Carta Si"
 CarteBleue     CardLabel = "Carte Bleue"
 Dankort        CardLabel = "Dankort"
 DinersClub     CardLabel = "Diners Club"
 Discover       CardLabel = "Discover"
 JCB            CardLabel = "JCB"
 Laser          CardLabel = "Laser"
 Maestro        CardLabel = "Maestro"
 Mastercard     CardLabel = "Mastercard"
 Unionpay       CardLabel = "Unionpay"
 Visa           CardLabel = "Visa"
 Empty          CardLabel = "null"
)
```

Available card labels.

#### type BusinessCategory

```go
type BusinessCategory string
```

BusinessCategory specifies an industry or category.

```go
const (
 PetShopsPetFoodAndSupplies                              BusinessCategory = "PET_SHOPS"
 VeterinaryServices                                      BusinessCategory = "VETERINARY_SERVICES"
 ACAndHeatingContractors                                 BusinessCategory = "AC_AND_HEATING_CONTRACTORS"
 CarpentryContractors                                    BusinessCategory = "CARPENTRY_CONTRACTORS"
 ElectricalContractors                                   BusinessCategory = "ELECTRICAL_CONTRACTORS"
 EquipmentToolsOrFurnitureRentalOrLeasing                BusinessCategory = "EQUIPMENT_TOOLS_FURNITURE_RENTAL_LEASING"
 GeneralContractors                                      BusinessCategory = "GENERAL_CONTRACTORS"
 SpecialTradeContractors                                 BusinessCategory = "SPECIAL_TRADE_CONTRACTORS"
 CharityAndDonationsCategory                             BusinessCategory = "CHARITY_AND_DONATIONS"
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
 PoliticalPartiesCategory                                BusinessCategory = "POLITICAL_PARTIES"
 AccountingAuditingBookkeepingAndTaxPreparationServices  BusinessCategory = "ACCOUNTING_AUDITING_BOOKKEEPING_TAX_PREPARATION_SERVICES"
 AdvertisingServices                                     BusinessCategory = "ADVERTISING_SERVICES"
 CleaningAndMaintenanceJanitorialServices                BusinessCategory = "CLEANING_MAINTENANCE_JANITORIAL_SERVICES"
 ComputerRepair                                          BusinessCategory = "COMPUTER_REPAIR"
 ConsultancyCategory                                     BusinessCategory = "CONSULTANCY"
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
```

Available category codes.

#### type Chargeback

```go
type Chargeback struct {
 Resource         string          `json:"resource,omitempty"`
 ID               string          `json:"id,omitempty"`
 Amount           *Amount         `json:"amount,omitempty"`
 SettlementAmount *Amount         `json:"settlementAmount,omitempty"`
 CreatedAt        *time.Time      `json:"createdAt,omitempty"`
 ReversedAt       *time.Time      `json:"reversedAt,omitempty"`
 PaymentID        string          `json:"paymentId,omitempty"`
 Links            ChargebackLinks `json:"_links,omitempty"`
}
```

Chargeback describes a forced transaction reversal initiated by the cardholder's
bank.

#### type ChargebackLinks

```go
type ChargebackLinks struct {
 Self          *URL `json:"self,omitempty"`
 Payment       *URL `json:"payment,omitempty"`
 Settlement    *URL `json:"settlement,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

ChargebackLinks describes all the possible links to be returned with a
chargeback object.

#### type ChargebackOptions

```go
type ChargebackOptions struct {
 Include string `url:"include,omitempty"`
 Embed   string `url:"embed,omitempty"`
}
```

ChargebackOptions describes chargeback endpoint valid query string parameters.

#### type ChargebacksList

```go
type ChargebacksList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Chargebacks []Chargeback
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

ChargebacksList describes how a list of chargebacks will be retrieved by Mollie.

#### type ChargebacksListOptions

```go
type ChargebacksListOptions struct {
 From      string `url:"from,omitempty"`
 Limit     int    `url:"limit,omitempty"`
 Include   string `url:"include,omitempty"`
 Embed     string `url:"embed,omitempty"`
 ProfileID string `url:"profileId,omitempty"`
}
```

ChargebacksListOptions describes list chargebacks endpoint valid query string
parameters.

#### type ChargebacksService

```go
type ChargebacksService service
```

ChargebacksService instance operates over chargeback resources.

#### func (*ChargebacksService) Get

```go
func (cs *ChargebacksService) Get(ctx context.Context, payment, chargeback string, opts *ChargebackOptions) (
 res *Response,
 p *Chargeback,
 err error,
)
```

Get retrieves a single chargeback by its ID. Note the original payment’s ID is
needed as well.

See: <https://docs.mollie.com/reference/v2/chargebacks-api/get-chargeback>

#### func (*ChargebacksService) List

```go
func (cs *ChargebacksService) List(ctx context.Context, options *ChargebacksListOptions) (
 res *Response,
 cl *ChargebacksList,
 err error,
)
```

List retrieves a list of chargebacks associated with your account/organization.

See: <https://docs.mollie.com/reference/v2/chargebacks-api/list-chargebacks>

#### func (*ChargebacksService) ListForPayment

```go
func (cs *ChargebacksService) ListForPayment(ctx context.Context, payment string, options *ChargebacksListOptions) (
 res *Response,
 cl *ChargebacksList,
 err error,
)
```

ListForPayment retrieves a list of chargebacks associated with a single payment.

See: <https://docs.mollie.com/reference/v2/chargebacks-api/list-chargebacks>

#### type Client

```go
type Client struct {
 BaseURL *url.URL

 // Services
 Payments       *PaymentsService
 Chargebacks    *ChargebacksService
 PaymentMethods *PaymentMethodsService
 Invoices       *InvoicesService
 Organizations  *OrganizationsService
 Profiles       *ProfilesService
 Refunds        *RefundsService
 Shipments      *ShipmentsService
 Orders         *OrdersService
 Settlements    *SettlementsService
 Captures       *CapturesService
 Subscriptions  *SubscriptionsService
 Customers      *CustomersService
 Miscellaneous  *MiscellaneousService
 Mandates       *MandatesService
 Permissions    *PermissionsService
 Onboarding     *OnboardingService
 PaymentLinks   *PaymentLinksService
 Partners       *PartnerService
 Balances       *BalancesService
}
```

Client manages communication with Mollie's API.

#### func  NewClient

```go
func NewClient(baseClient *http.Client, conf *Config) (mollie *Client, err error)
```

NewClient returns a new Mollie HTTP API client. You can pass a previously build
http client, if none is provided then http.DefaultClient will be used.

NewClient will lookup the environment for values to assign to the API token
(`MOLLIE_API_TOKEN`) and the Organization token (`MOLLIE_ORG_TOKEN`) according
to the provided Config object.

You can also set the token values programmatically by using the Client
WithAPIKey and WithOrganizationKey functions.

#### func (*Client) Do

```go
func (c *Client) Do(req *http.Request) (*Response, error)
```

Do sends an API request and returns the API response or returned as an error if
an API error has occurred.

#### func (*Client) HasAccessToken

```go
func (c *Client) HasAccessToken() bool
```

HasAccessToken will return true when the provided authentication token complies
with the access token REGEXP match check. This will enable TestMode inside the
request body.

See: <https://github.com/VictorAvelar/mollie-api-go/issues/123>

#### func (*Client) NewAPIRequest

```go
func (c *Client) NewAPIRequest(ctx context.Context, method string, uri string, body interface{}) (
 req *http.Request,
 err error,
)
```

NewAPIRequest is a wrapper around the http.NewRequest function.

It will setup the authentication headers/parameters according to the client
config.

#### func (*Client) SetIdempotencyKeyGenerator

```go
func (c *Client) SetIdempotencyKeyGenerator(kg idempotency.KeyGenerator)
```

SetIdempotencyKeyGenerator allows you to pass your own idempotency key
generator.

#### func (*Client) WithAuthenticationValue

```go
func (c *Client) WithAuthenticationValue(k string) error
```

WithAuthenticationValue offers a convenient setter for any of the valid
authentication tokens provided by Mollie.

Ideally your API key will be provided from and environment variable or a secret
management engine. This should only be used when environment variables are
"impossible" to be used.

#### type Commission

```go
type Commission struct {
 Count       int     `json:"count,omitempty"`
 TotalAmount *Amount `json:"totalAmount,omitempty"`
}
```

Commission describes a partner take from any operation on Mollie's API.

#### type Config

```go
type Config struct {
}
```

Config contains information that helps during the setup of a new Mollie client.

#### func  NewAPIConfig

```go
func NewAPIConfig(reqIdem bool) *Config
```

NewAPIConfig builds a configuration object with the following settings: tests
mode: disabled api token source: MOLLIE_API_TOKEN

it receives `reqIdem (boolean)` to enable the request idempotency feature.

#### func  NewAPITestingConfig

```go
func NewAPITestingConfig(reqIdem bool) *Config
```

NewAPITestingConfig builds a configuration object with the following settings:
tests mode: enabled api token source: MOLLIE_API_TOKEN

it receives `reqIdem (boolean)` to enable the request idempotency feature.

#### func  NewConfig

```go
func NewConfig(t bool, auth string) *Config
```

NewConfig builds a Mollie configuration object, it takes t to indicate if our
client is meant to create requests for testing, and auth to indicate the
authentication method we want to use.

#### func  NewOrgConfig

```go
func NewOrgConfig(reqIdem bool) *Config
```

NewOrgConfig builds a configuration object with the following settings: tests
mode: disabled Org token source: MOLLIE_ORG_TOKEN

it receives `reqIdem (boolean)` to enable the request idempotency feature.

#### func  NewOrgTestingConfig

```go
func NewOrgTestingConfig(reqIdem bool) *Config
```

NewOrgTestingConfig builds a configuration object with the following settings:
tests mode: enabled api token source: MOLLIE_ORG_TOKEN

it receives `reqIdem (boolean)` to enable the request idempotency feature.

#### func (*Config) SwitchAuthStrategy

```go
func (c *Config) SwitchAuthStrategy(auth string) string
```

SwitchAuthStrategy changes the environment variable used to fetch the auth
tokens.

Known values are: [MOLLIE_API_TOKEN,MOLLIE_ORG_TOKEN], if you use a custom
environment variable pass it as argument.

#### func (*Config) ToggleIdempotency

```go
func (c *Config) ToggleIdempotency() bool
```

ToggleIdempotency enables/disables the request idempotency feature in the
current Config.

#### func (*Config) ToggleTesting

```go
func (c *Config) ToggleTesting() bool
```

ToggleTesting enables/disables the test-mode in the current Config.

#### type ContextValue

```go
type ContextValue string
```

ContextValue represents a relevant value in the system associated with a
BalanceTransaction.

#### type ContextValues

```go
type ContextValues map[TransactionType]ContextValue
```

ContextValues is a map of TransactionType to ContextValue.

#### type CreateShipmentRequest

```go
type CreateShipmentRequest struct {
 Lines    []OrderLine      `json:"lines,omitempty"`
 Tracking ShipmentTracking `json:"tracking,omitempty"`
 TestMode bool             `json:"testmode,omitempty"`
}
```

CreateShipmentRequest defines information required to create a new shipment.

#### type Customer

```go
type Customer struct {
 Resource  string        `json:"resource,omitempty"`
 ID        string        `json:"id,omitempty"`
 Mode      Mode          `json:"mode,omitempty"`
 Name      string        `json:"name,omitempty"`
 Email     string        `json:"email,omitempty"`
 Locale    Locale        `json:"locale,omitempty"`
 Metadata  interface{}   `json:"metadata,omitempty"`
 CreatedAt *time.Time    `json:"createdAt,omitempty"`
 Links     CustomerLinks `json:"_links,omitempty"`
}
```

Customer represents buyers.

#### type CustomerLinks

```go
type CustomerLinks struct {
 Self          *URL `json:"self,omitempty"`
 Mandates      *URL `json:"mandates,omitempty"`
 Subscriptions *URL `json:"subscriptions,omitempty"`
 Payments      *URL `json:"payments,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
 Dashboard     *URL `json:"dashboard,omitempty"`
}
```

CustomerLinks contains the HAL resources for a customer response.

#### type CustomersList

```go
type CustomersList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Customers []Customer `json:"customers,omitempty"`
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"links,omitempty"`
}
```

CustomersList contains a embedded list of customers wrapped in a standard Mollie
paginated response.

#### type CustomersListOptions

```go
type CustomersListOptions struct {
 From         string       `url:"from,omitempty"`
 Limit        int          `url:"limit,omitempty"`
 ProfileID    string       `url:"profileId,omitempty"`
 SequenceType SequenceType `url:"sequenceType,omitempty"`
 RedirectURL  string       `url:"redirectUrl,omitempty"`
}
```

CustomersListOptions contains valid query parameters for the list customers
endpoint.

#### type CustomersService

```go
type CustomersService service
```

CustomersService operates over the customer resource.

#### func (*CustomersService) Create

```go
func (cs *CustomersService) Create(ctx context.Context, c Customer) (res *Response, cc *Customer, err error)
```

Create creates a simple minimal representation of a customer in the Mollie API
to use for the Mollie Checkout and Recurring features.

See: <https://docs.mollie.com/reference/v2/customers-api/create-customer>

#### func (*CustomersService) CreatePayment

```go
func (cs *CustomersService) CreatePayment(ctx context.Context, id string, p Payment) (
 res *Response,
 pp *Payment,
 err error,
)
```

CreatePayment creates a payment for the customer.

See: <https://docs.mollie.com/reference/v2/customers-api/create-customer-payment>

#### func (*CustomersService) Delete

```go
func (cs *CustomersService) Delete(ctx context.Context, id string) (res *Response, err error)
```

Delete a customer.

All mandates and subscriptions created for this customer will be canceled as
well.

See: <https://docs.mollie.com/reference/v2/customers-api/delete-customer>

#### func (*CustomersService) Get

```go
func (cs *CustomersService) Get(ctx context.Context, id string) (res *Response, c *Customer, err error)
```

Get finds a customer by its ID.

See: <https://docs.mollie.com/reference/v2/customers-api/get-customer>

#### func (*CustomersService) GetPayments

```go
func (cs *CustomersService) GetPayments(ctx context.Context, id string, options *CustomersListOptions) (
 res *Response,
 pl *PaymentList,
 err error,
)
```

GetPayments retrieves all payments linked to the customer.

See: <https://docs.mollie.com/reference/v2/customers-api/list-customer-payments>

#### func (*CustomersService) List

```go
func (cs *CustomersService) List(ctx context.Context, options *CustomersListOptions) (
 res *Response,
 cl *CustomersList,
 err error,
)
```

List retrieves all customers created.

See: <https://docs.mollie.com/reference/v2/customers-api/list-customers>

#### func (*CustomersService) Update

```go
func (cs *CustomersService) Update(ctx context.Context, id string, c Customer) (
 res *Response,
 cc *Customer,
 err error,
)
```

Update an existing customer.

See: <https://docs.mollie.com/reference/v2/customers-api/update-customer>

#### type EligibilityReasons

```go
type EligibilityReasons string
```

EligibilityReasons for paypal seller protection.

```go
const (
 Eligible                        EligibilityReasons = "Eligible"
 Ineligible                      EligibilityReasons = "Ineligible"
 INROnly                         EligibilityReasons = "Partially Eligible - INR Only"
 AuthOnly                        EligibilityReasons = "Partially Eligible - Unauth Only"
 PartiallyEligible               EligibilityReasons = "PartiallyEligible"
 EligibilityNone                 EligibilityReasons = "None"
 ActiveFraudControlUnAuthPremium EligibilityReasons = "Active Fraud Control - Unauth Premium Eligible"
)
```

Available eligibility reasons for Paypal sellers protection.

#### type EmbedValue

```go
type EmbedValue string
```

EmbedValue describes the valid value of embed query string.

```go
const (
 EmbedPayment     EmbedValue = "payments"
 EmbedRefund      EmbedValue = "refunds"
 EmbedShipments   EmbedValue = "shipments"
 EmbedChargebacks EmbedValue = "chargebacks"
)
```

Valid Embed query string value.

#### type ErrorLinks

```go
type ErrorLinks struct {
 Documentation *URL `json:"documentation,omitempty"`
}
```

ErrorLinks container references to common urls returned with errors.

#### type FailureReason

```go
type FailureReason string
```

FailureReason provides a description on a failed payment.

```go
const (
 ReasonAuthenticationAbandoned      FailureReason = "authentication_abandoned"
 ReasonAuthenticationUnavailableACS FailureReason = "authentication_unavailable_acs"
 ReasonInvalidCardNumber            FailureReason = "invalid_card_number"
 ReasonInvalidCCV                   FailureReason = "invalid_ccv"
 ReasonInvalidCardHolderName        FailureReason = "invalid_card_holder_name"
 ReasonCardExpired                  FailureReason = "card_expired"
 ReasonInvalidCardType              FailureReason = "invalid_card_type"
 ReasonRefusedByIssuer              FailureReason = "refused_by_issuer"
 ReasonInsufficientFunds            FailureReason = "insufficient_funds"
 ReasonInactiveCard                 FailureReason = "inactive_card"
 ReasonUnknown                      FailureReason = "unknown_reason"
 ReasonPossibleFraud                FailureReason = "possible_fraud"
)
```

Possible failure reasons.

#### type FeeRegion

```go
type FeeRegion string
```

FeeRegion contains the fee region for the payment.

```go
const (
 AmericanExpress FeeRegion = "american-express"
 CarteBancaire   FeeRegion = "carte-bancaire"
 IntraEU         FeeRegion = "intra-eu"
 MaestroRegion   FeeRegion = "maestro"
 Other           FeeRegion = "other"
)
```

Valid Fee regions.

#### type GetPartnerClientOptions

```go
type GetPartnerClientOptions struct {
 Embed string `url:"embed,omitempty"`
}
```

GetPartnerClientOptions contains valid query parameters for the get clients
endpoint.

#### type GiftCardEnabled

```go
type GiftCardEnabled struct {
 Resource    string               `json:"resource,omitempty"`
 ID          GiftCardIssuer       `json:"id,omitempty"`
 Description string               `json:"description,omitempty"`
 Status      GiftCardIssuerStatus `json:"status,omitempty"`
 Links       GiftCardLinks        `json:"_links,omitempty"`
}
```

GiftCardEnabled describes the response of a gift card issuer enable operation.

#### type GiftCardIssuer

```go
type GiftCardIssuer string
```

GiftCardIssuer type describes issuers supported by mollie.

```go
const (
 BloemenCadeuKaart          GiftCardIssuer = "bloemencadeaukaart"
 BloemPlantGiftCard         GiftCardIssuer = "bloemplantgiftcard"
 Boekenbon                  GiftCardIssuer = "boekenbon"
 DecaudeuKaart              GiftCardIssuer = "decadeaukaart"
 DelokaleDecauKaart         GiftCardIssuer = "delokalecadeaukaart"
 Dinercadeau                GiftCardIssuer = "dinercadeau"
 Doenkadotickets            GiftCardIssuer = "doenkadotickets"
 Fashioncheque              GiftCardIssuer = "fashioncheque"
 Festivalcadeau             GiftCardIssuer = "festivalcadeau"
 Good4fun                   GiftCardIssuer = "good4fun"
 Horseandgifts              GiftCardIssuer = "horseandgifts"
 HuistuinCadeauKaart        GiftCardIssuer = "huistuincadeaukaart"
 JewelCard                  GiftCardIssuer = "jewelcard"
 KlusCadeu                  GiftCardIssuer = "kluscadeau"
 Kunstencultuurcadeaukaart  GiftCardIssuer = "kunstencultuurcadeaukaart"
 Nationalebioscoopbon       GiftCardIssuer = "nationalebioscoopbon"
 Nationaleentertainmentcard GiftCardIssuer = "nationaleentertainmentcard"
 Nationalegolfbon           GiftCardIssuer = "nationalegolfbon"
 Ohmygood                   GiftCardIssuer = "ohmygood"
 Podiumcadeaukaart          GiftCardIssuer = "podiumcadeaukaart"
 Reiscadeau                 GiftCardIssuer = "reiscadeau"
 Restaurantcadeau           GiftCardIssuer = "restaurantcadeau"
 Shoesandsneakerscadeu      GiftCardIssuer = "shoesandsneakerscadeau"
 SodexoSportCulturePass     GiftCardIssuer = "sodexosportculturepass"
 Sportenfitcadeau           GiftCardIssuer = "sportenfitcadeau"
 Sustainablefashion         GiftCardIssuer = "sustainablefashion"
 Travelcheq                 GiftCardIssuer = "travelcheq"
 Vvvgiftcard                GiftCardIssuer = "vvvgiftcard"
 Vvvdinercheque             GiftCardIssuer = "vvvdinercheque"
 Vvvlekkerweg               GiftCardIssuer = "vvvlekkerweg"
 Webshopgiftcard            GiftCardIssuer = "webshopgiftcard"
 Wijncadeukaart             GiftCardIssuer = "wijncadeaukaart"
 Yourgift                   GiftCardIssuer = "yourgift"
)
```

Supported gift card issuers. #nosec G101 -- This are the brands issuing gift
cards.

#### type GiftCardIssuerStatus

```go
type GiftCardIssuerStatus string
```

GiftCardIssuerStatus describes the status of a gift card issuer in your account.

```go
const (
 PendingIssuer GiftCardIssuerStatus = "pending-issuer"
 EnabledIssuer GiftCardIssuerStatus = "enabled"
)
```

Valid issuer statuses.

#### type GiftCardLinks

```go
type GiftCardLinks struct {
 Self          *URL `json:"self,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

GiftCardLinks are links embedded when a gift card is enabled.

#### type Image

```go
type Image struct {
 Size1x string `json:"size1X,omitempty"`
 Size2X string `json:"size2X,omitempty"`
 Svg    string `json:"svg,omitempty"`
}
```

Image describes a generic image resource retrieved by Mollie.

#### type Invoice

```go
type Invoice struct {
 Resource    string        `json:"resource,omitempty"`
 ID          string        `json:"id,omitempty"`
 Reference   string        `json:"reference,omitempty"`
 VatNumber   string        `json:"vatNumber,omitempty"`
 Status      InvoiceStatus `json:"status,omitempty"`
 IssuedAt    string        `json:"issuedAt,omitempty"`
 PaidAt      string        `json:"paidAt,omitempty"`
 DueAt       string        `json:"dueAt,omitempty"`
 NetAmount   *Amount       `json:"netAmount,omitempty"`
 VatAmount   *Amount       `json:"vatAmount,omitempty"`
 GrossAmount *Amount       `json:"grossAmount,omitempty"`
 Lines       []*LineItem   `json:"lines,omitempty"`
 Links       InvoiceLinks  `json:"_links,omitempty"`
}
```

Invoice describes an invoice details.

#### type InvoiceLinks

```go
type InvoiceLinks struct {
 Self          *URL `json:"self,omitempty"`
 PDF           *URL `json:"pdf,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

InvoiceLinks describes all the possible links to be returned with a invoice
object.

#### type InvoiceStatus

```go
type InvoiceStatus string
```

InvoiceStatus status of the invoice.

```go
const (
 InvoiceStatusOpen    InvoiceStatus = "open"
 InvoiceStatusPaid    InvoiceStatus = "paid"
 InvoiceStatusOverdue InvoiceStatus = "overdue"
)
```

Valid status of the invoice.

#### type InvoicesList

```go
type InvoicesList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Invoices []Invoice `json:"invoices"`
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

InvoicesList describes how a list of invoices will be retrieved by Mollie.

#### type InvoicesListOptions

```go
type InvoicesListOptions struct {
 Reference string `url:"reference,omitempty"`
 Year      string `url:"year,omitempty"`
 From      string `url:"from,omitempty"`
 Limit     int64  `url:"limit,omitempty"`
}
```

InvoicesListOptions describes list invoices endpoint valid query string
parameters.

#### type InvoicesService

```go
type InvoicesService service
```

InvoicesService instance operates over invoice resources.

#### func (*InvoicesService) Get

```go
func (is *InvoicesService) Get(ctx context.Context, id string) (res *Response, i *Invoice, err error)
```

Get retrieve details of an invoice, using the invoice’s identifier.

#### func (*InvoicesService) List

```go
func (is *InvoicesService) List(ctx context.Context, options *InvoicesListOptions) (
 res *Response,
 il *InvoicesList,
 err error,
)
```

List retrieves a list of invoices associated with your account/organization.

#### type LineItem

```go
type LineItem struct {
 Period        string  `json:"period,omitempty"`
 Description   string  `json:"description,omitempty"`
 Count         int64   `json:"count,omitempty"`
 VatPercentage float64 `json:"vatPercentage,omitempty"`
 Amount        *Amount `json:"amount,omitempty"`
}
```

LineItem product details.

#### type ListPartnerClientsOptions

```go
type ListPartnerClientsOptions struct {
 From  int `url:"from,omitempty"`
 Limit int `url:"limit,omitempty"`
 Year  int `url:"year,omitempty"`
 Month int `url:"month,omitempty"`
}
```

ListPartnerClientsOptions contains valid query parameters for the list clients
endpoint.

#### type ListPaymentOptions

```go
type ListPaymentOptions struct {
 Limit     int    `url:"limit,omitempty"`
 Include   string `url:"include,omitempty"`
 Embed     string `url:"embed,omitempty"`
 ProfileID string `url:"profileId,omitempty"`
 From      string `url:"from,omitempty"`
}
```

ListPaymentOptions describes list payments endpoint valid query string
parameters.

#### type ListRefundOptions

```go
type ListRefundOptions struct {
 From      string     `url:"from,omitempty"`
 Limit     int        `url:"limit,omitempty"`
 ProfileID string     `url:"profileId,omitempty"`
 Embed     EmbedValue `url:"embed,omitempty"`
}
```

ListRefundOptions describes list refund endpoint valid query string parameters.

See: <https://docs.mollie.com/reference/v2/refunds-api/list-refunds>.

#### type Locale

```go
type Locale string
```

Locale represents a country and language in ISO-15897 format.

```go
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
```

Mollie supported locales.

#### type Mandate

```go
type Mandate struct {
 ID               string         `json:"id,omitempty"`
 Resource         string         `json:"resource,omitempty"`
 Method           PaymentMethod  `json:"method,omitempty"`
 ConsumerName     string         `json:"consumerName,omitempty"`
 ConsumerAccount  string         `json:"consumerAccount,omitempty"`
 ConsumerBic      string         `json:"consumerBic,omitempty"`
 SignatureDate    *ShortDate     `json:"signatureDate,omitempty"`
 MandateReference string         `json:"mandateReference,omitempty"`
 Mode             Mode           `json:"mode,omitempty"`
 Status           MandateStatus  `json:"status,omitempty"`
 CreatedAt        *time.Time     `json:"createdAt,omitempty"`
 Details          MandateDetails `json:"details,omitempty"`
 Links            MandateLinks   `json:"_links,omitempty"`
}
```

Mandate allow you to charge a customer’s credit card or bank account
recurrently.

#### type MandateDetails

```go
type MandateDetails struct {
 ConsumerName    string     `json:"consumerName,omitempty"`
 ConsumerAccount string     `json:"consumerAccount,omitempty"`
 ConsumerBic     string     `json:"consumerBic,omitempty"`
 CardHolder      string     `json:"cardHolder,omitempty"`
 CardNumber      string     `json:"cardNumber,omitempty"`
 CardLabel       CardLabel  `json:"cardLabel,omitempty"`
 CardFingerprint string     `json:"cardFingerprint,omitempty"`
 CardExpiryDate  *ShortDate `json:"cardExpiryDate,omitempty"`
}
```

MandateDetails are possible values inside the mandate.details field.

#### type MandateLinks

```go
type MandateLinks struct {
 Self          *URL `json:"self,omitempty"`
 Customer      *URL `json:"customer,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

MandateLinks response objects.

#### type MandateStatus

```go
type MandateStatus string
```

MandateStatus for the Mandate object.

```go
const (
 PendingMandate MandateStatus = "pending"
 ValidMandate   MandateStatus = "valid"
 InvalidMandate MandateStatus = "invalid"
)
```

Valid mandate statuses.

#### type MandatesList

```go
type MandatesList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Mandates []Mandate
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

MandatesList describes how a list of mandates will be retrieved by Mollie.

#### type MandatesListOptions

```go
type MandatesListOptions struct {
 From  string `url:"from,omitempty"`
 Limit int    `url:"limit,omitempty"`
}
```

MandatesListOptions contains valid query parameters to filter the List mandates
actions.

From is a mandate id to offset from (inclusive) Limit is the max number of
mandates to retrieve.

#### type MandatesService

```go
type MandatesService service
```

MandatesService operates over customer mandates endpoints.

#### func (*MandatesService) Create

```go
func (ms *MandatesService) Create(ctx context.Context, customer string, mandate Mandate) (
 res *Response,
 mr *Mandate,
 err error,
)
```

Create a mandate for a specific customer.

Mandates allow you to charge a customer’s credit card or bank account
recurrently.

See: <https://docs.mollie.com/reference/v2/mandates-api/create-mandate>

#### func (*MandatesService) Get

```go
func (ms *MandatesService) Get(ctx context.Context, customer, mandate string) (res *Response, mr *Mandate, err error)
```

Get retrieves a mandate by its ID and its customer’s ID. The mandate will either
contain IBAN or credit card details, depending on the type of mandate.

See: <https://docs.mollie.com/reference/v2/mandates-api/get-mandate>

#### func (*MandatesService) List

```go
func (ms *MandatesService) List(ctx context.Context, customer string, options *MandatesListOptions) (
 res *Response,
 ml *MandatesList,
 err error,
)
```

List retrieves all mandates for the given customerId, ordered from newest to
oldest.

See: <https://docs.mollie.com/reference/v2/mandates-api/list-mandates>

#### func (*MandatesService) Revoke

```go
func (ms *MandatesService) Revoke(ctx context.Context, customer, mandate string) (res *Response, err error)
```

Revoke a customer’s mandate.

You will no longer be able to charge the consumer’s bank account or credit card
with this mandate and all connected subscriptions will be canceled.

See: <https://docs.mollie.com/reference/v2/mandates-api/revoke-mandate>

#### type MethodsLinks

```go
type MethodsLinks struct {
 Self          *URL `json:"self,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

MethodsLinks describes links attached to methods service responses.

#### type MiscellaneousService

```go
type MiscellaneousService service
```

MiscellaneousService operates over the resources described in Mollie's
miscellaneous API endpoints section.

#### func (*MiscellaneousService) ApplePaymentSession

```go
func (ms *MiscellaneousService) ApplePaymentSession(ctx context.Context, asr *ApplePaymentSessionRequest) (
 res *Response,
 aps *ApplePaymentSession,
 err error,
)
```

ApplePaymentSession returns an Apple Payment Session object valid for one
transaction.

See:
<https://docs.mollie.com/reference/v2/wallets-api/request-apple-pay-payment-session>

#### type Mode

```go
type Mode string
```

Mode contains information about the creation environment.

```go
const (
 LiveMode Mode = "live"
 TestMode Mode = "test"
)
```

Valid modes.

#### type Onboarding

```go
type Onboarding struct {
 Resource              string           `json:"reference,omitempty"`
 Name                  string           `json:"name,omitempty"`
 SignedUpAt            *time.Time       `json:"signedUpAt,omitempty"`
 Status                OnboardingStatus `json:"status,omitempty"`
 CanReceivePayments    bool             `json:"canReceivePayments,omitempty"`
 CanReveiceSettlements bool             `json:"canReceiveSettlements,omitempty"`
 Links                 OnboardingLinks  `json:"_links,omitempty"`
}
```

Onboarding data for an organization.

#### type OnboardingData

```go
type OnboardingDataOrganization struct {
 Name               string   `json:"name,omitempty"`
 Address            *Address `json:"address,omitempty"`
 RegistrationNumber string   `json:"registrationNumber,omitempty"`
 VatNumber          string   `json:"vatNumber,omitempty"`
 VatRegulation      string   `json:"vatRegulation,omitempty"`
}
type OnboardingDataProfile struct {
 Name             string           `json:"name,omitempty"`
 URL              string           `json:"url,omitempty"`
 Email            string           `json:"email,omitempty"`
 Description      string           `json:"description,omitempty"`
 Phone            string           `json:"phone,omitempty"`
 BusinessCategory BusinessCategory `json:"businessCategory,omitempty"`
}
type OnboardingData struct {
 Organization OnboardingDataOrganization `json:"organization,omitempty"`
 Profile      OnboardingDataProfile      `json:"profile,omitempty"`
}
```

OnboardingData request possible values.

Please note that even though all parameters are optional, at least one of them
needs to be provided in the request.

Information that the merchant has entered in their dashboard will not be
overwritten.

#### type OnboardingLinks

```go
type OnboardingLinks struct {
 Self          *URL `json:"self,omitempty"`
 Dashboard     *URL `json:"dashboard,omitempty"`
 Organization  *URL `json:"organization,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

OnboardingLinks contains URL objects relevant to the onboarding status.

#### type OnboardingService

```go
type OnboardingService service
```

OnboardingService operates over the onboarding API.

#### func (*OnboardingService) GetOnboardingStatus

```go
func (os *OnboardingService) GetOnboardingStatus(ctx context.Context) (res *Response, o *Onboarding, err error)
```

GetOnboardingStatus gets the status of onboarding of the authenticated
organization.

See: <https://docs.mollie.com/reference/v2/onboarding-api/get-onboarding-status>

#### func (*OnboardingService) SubmitOnboardingData

```go
func (os *OnboardingService) SubmitOnboardingData(ctx context.Context, d *OnboardingData) (res *Response, err error)
```

SubmitOnboardingData sends data that will be prefilled in the merchant’s
onboarding. Please note that the data you submit will only be processed when the
onboarding status is needs-data.

See: <https://docs.mollie.com/reference/v2/onboarding-api/submit-onboarding-data>

#### type OnboardingStatus

```go
type OnboardingStatus string
```

OnboardingStatus describes status of the organization’s onboarding process.

```go
const (
 NeedsDataOnboardingStatus OnboardingStatus = "needs-data"
 InReviewOnboardingStatus  OnboardingStatus = "in-review"
 CompletedOnboardingStatus OnboardingStatus = "completed"
)
```

Possible status values.

#### type Order

```go
type Order struct {
 TestMode                                 bool          `json:"testmode,omitempty"`
 IsCancelable                             bool          `json:"isCancelable,omitempty"`
 ShopperCountryMustMatchTheBillingCountry bool          `json:"shopperCountryMustMatchTheBillingCountry,omitempty"`
 Resource                                 string        `json:"resource,omitempty"`
 ID                                       string        `json:"id,omitempty"`
 ProfileID                                string        `json:"profileId,omitempty"`
 OrderNumber                              string        `json:"orderNumber,omitempty"`
 RedirectURL                              string        `json:"redirectUrl,omitempty"`
 WebhookURL                               string        `json:"webhookUrl,omitempty"`
 Description                              string        `json:"description,omitempty"`
 Sku                                      string        `json:"sku,omitempty"`
 CancelURL                                string        `json:"cancelUrl,omitempty"`
 Metadata                                 interface{}   `json:"metadata,omitempty"`
 Mode                                     Mode          `json:"mode,omitempty"`
 Method                                   PaymentMethod `json:"method,omitempty"`
 Status                                   OrderStatus   `json:"status,omitempty"`
 Locale                                   Locale        `json:"locale,omitempty"`
 ShippingAddress                          OrderAddress  `json:"shippingAddress,omitempty"`
 Links                                    OrderLinks    `json:"_links,omitempty"`
 Amount                                   *Amount       `json:"amount,omitempty"`
 AmountCaptured                           *Amount       `json:"amountCaptured,omitempty"`
 AmountRefunded                           *Amount       `json:"amountRefunded,omitempty"`
 BillingAddress                           *OrderAddress `json:"billingAddress,omitempty"`
 ConsumerDateOfBirth                      *ShortDate    `json:"consumerDateOfBirth,omitempty"`
 CreatedAt                                *time.Time    `json:"createdAt,omitempty"`
 ExpiresAt                                *time.Time    `json:"expiresAt,omitempty"`
 ExpiredAt                                *time.Time    `json:"expiredAt,omitempty"`
 PaidAt                                   *time.Time    `json:"paidAt,omitempty"`
 AuthorizedAt                             *time.Time    `json:"authorizedAt,omitempty"`
 CanceledAt                               *time.Time    `json:"canceledAt,omitempty"`
 CompletedAt                              *time.Time    `json:"completedAt,omitempty"`
 OrderPayment                             *OrderPayment `json:"payment,omitempty"`
 Lines                                    []*OrderLine  `json:"lines,omitempty"`
 Embedded                                 struct {
  Payments []*Payment `json:"payments,omitempty"`
  Refunds  []*Refund  `json:"refunds,omitempty"`
 } `json:"_embedded,omitempty"`
}
```

Order explain the items that customers need to pay for.

#### type OrderAddress

```go
type OrderAddress struct {
 OrganizationName string      `json:"organizationName,omitempty"`
 Title            string      `json:"title,omitempty"`
 GivenName        string      `json:"givenName,omitempty"`
 FamilyName       string      `json:"familyName,omitempty"`
 Email            string      `json:"email,omitempty"`
 Phone            PhoneNumber `json:"phone,omitempty"`
 StreetAndNumber  string      `json:"streetAndNumber,omitempty"`
 StreetAdditional string      `json:"streetAdditional,omitempty"`
 PostalCode       string      `json:"postalCode,omitempty"`
 City             string      `json:"city,omitempty"`
 Region           string      `json:"region,omitempty"`
 Country          string      `json:"country,omitempty"`
}
```

OrderAddress identify both the address and the person the order is billed or
shipped to.

#### type OrderLine

```go
type OrderLine struct {
 Resource           string          `json:"resource,omitempty"`
 ID                 string          `json:"id,omitempty"`
 OrderID            string          `json:"orderId,omitempty"`
 ProductType        ProductType     `json:"type,omitempty"`
 Name               string          `json:"name,omitempty"`
 Amount             *Amount         `json:"amount,omitempty"`
 Status             OrderLineStatus `json:"status,omitempty"`
 IsCancelable       bool            `json:"isCancelable,omitempty"`
 Quantity           int             `json:"quantity,omitempty"`
 QuantityShipped    int             `json:"quantityShipped,omitempty"`
 AmountShipped      *Amount         `json:"amountShipped,omitempty"`
 QuantityRefunded   int             `json:"quantityRefunded,omitempty"`
 AmountRefunded     *Amount         `json:"amountRefunded,omitempty"`
 QuantityCanceled   int             `json:"quantityCanceled,omitempty"`
 AmountCanceled     *Amount         `json:"amountCanceled,omitempty"`
 ShippableQuantity  int             `json:"shippableQuantity,omitempty"`
 RefundableQuantity int             `json:"refundableQuantity,omitempty"`
 CancelableQuantity int             `json:"cancelableQuantity,omitempty"`
 UnitPrice          *Amount         `json:"unitPrice,omitempty"`
 DiscountAmount     *Amount         `json:"discountAmount,omitempty"`
 TotalAmount        *Amount         `json:"totalAmount,omitempty"`
 VatRate            string          `json:"vatRate,omitempty"`
 VatAmount          *Amount         `json:"vatAmount,omitempty"`
 SKU                string          `json:"sku,omitempty"`
 CreatedAt          *time.Time      `json:"createdAt,omitempty"`
 Links              OrderLineLinks  `json:"_links,omitempty"`
 ImageURL           string          `json:"imageUrl,omitempty"`
 ProductURL         string          `json:"productUrl,omitempty"`
 Metadata           interface{}     `json:"metadata,omitempty"`
}
```

OrderLine contain the actual things the customer bought.

#### type OrderLineLinks

```go
type OrderLineLinks struct {
 ProductURL *URL `json:"productUrl,omitempty"`
 ImageURL   *URL `json:"imageUrl,omitempty"`
}
```

OrderLineLinks describes object with several URL objects relevant to the order
line.

#### type OrderLineStatus

```go
type OrderLineStatus string
```

OrderLineStatus describes status of the order line.

```go
const (
 OrderLineCreated    OrderLineStatus = "created"
 OrderLineAuthorized OrderLineStatus = "authorized"
 OrderLinePaid       OrderLineStatus = "paid"
 OrderLineShipping   OrderLineStatus = "shipping"
 OrderLineCanceled   OrderLineStatus = "canceled"
 OrderLineCompleted  OrderLineStatus = "completed"
)
```

Valid order line status.

#### type OrderLinks

```go
type OrderLinks struct {
 Self          *URL `json:"self,omitempty"`
 Checkout      *URL `json:"checkout,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
 Dashboard     *URL `json:"dashboard,omitempty"`
}
```

OrderLinks describes an object with several URL objects relevant to the order.
Every URL object will contain an href and a type field.

#### type OrderList

```go
type OrderList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Orders []*Order `json:"orders,omitempty"`
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"links,omitempty"`
}
```

OrderList for containing the response of list orders.

#### type OrderListOptions

```go
type OrderListOptions struct {
 ProfileID string `url:"profileId,omitempty"`
 From      string `url:"from,omitempty"`
 Limit     int    `url:"limit,omitempty"`
}
```

OrderListOptions describes order endpoint valid query string parameters.

#### type OrderListRefund

```go
type OrderListRefund struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Refunds []*Refund `json:"refund,omitempty"`
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"links,omitempty"`
}
```

OrderListRefund for containing the response of list orders.

#### type OrderListRefundOptions

```go
type OrderListRefundOptions struct {
 From  string     `url:"from,omitempty"`
 Limit int        `url:"limit,omitempty"`
 Embed EmbedValue `url:"embed,omitempty"`
}
```

OrderListRefundOptions describes order endpoint valid query string parameters.

#### type OrderOptions

```go
type OrderOptions struct {
 Embed     []EmbedValue `url:"embed,omitempty"`
 ProfileID string       `url:"profileId,omitempty"`
}
```

OrderOptions describes order endpoint valid query string parameters.

#### type OrderPayment

```go
type OrderPayment struct {
 ConsumerAccount   string          `json:"consumerAccount,omitempty"`
 CustomerID        string          `json:"customerId,omitempty"`
 CustomerReference string          `json:"customerReference,omitempty"`
 Issuer            string          `json:"issuer,omitempty"`
 MandateID         string          `json:"mandateId,omitempty"`
 SequenceType      SequenceType    `json:"sequenceType,omitempty"`
 VoucherNumber     string          `json:"voucherNumber,omitempty"`
 VoucherPin        string          `json:"voucherPin,omitempty"`
 WebhookURL        string          `json:"webhookUrl,omitempty"`
 ApplicationFee    *ApplicationFee `json:"applicationFee,omitempty"`
 Method            PaymentMethod   `json:"method,omitempty"`
}
```

OrderPayment describes payment specific parameters that can be passed during
order creation.

#### type OrderStatus

```go
type OrderStatus string
```

OrderStatus describes the valid order status.

```go
const (
 Created    OrderStatus = "created"
 Paid       OrderStatus = "paid"
 Authorized OrderStatus = "authorized"
 Canceled   OrderStatus = "canceled"
 Shipping   OrderStatus = "shipping"
 Completed  OrderStatus = "completed"
 Expired    OrderStatus = "expired"
)
```

Valid order status.

#### type OrdersService

```go
type OrdersService service
```

OrdersService instance operates over refund resources.

#### func (*OrdersService) Cancel

```go
func (ors *OrdersService) Cancel(ctx context.Context, orderID string) (res *Response, order *Order, err error)
```

Cancel try to cancel the order that fulfill certain requirements.

See <https://docs.mollie.com/reference/v2/orders-api/cancel-order>

#### func (*OrdersService) CancelOrderLines

```go
func (ors *OrdersService) CancelOrderLines(ctx context.Context, orderID string, orderLines []OrderLine) (
 res *Response,
 err error,
)
```

CancelOrderLines can be used to cancel one or more order lines that were
previously authorized using a pay after delivery payment method. Use the Cancel
Order API if you want to cancel the entire order or the remainder of the order.

See <https://docs.mollie.com/reference/v2/orders-api/cancel-order-lines>

#### func (*OrdersService) Create

```go
func (ors *OrdersService) Create(ctx context.Context, ord Order, opts *OrderOptions) (
 res *Response,
 order *Order,
 err error,
)
```

Create an order will automatically create the required payment to allow your
customer to pay for the order.

See <https://docs.mollie.com/reference/v2/orders-api/create-order>

#### func (*OrdersService) CreateOrderPayment

```go
func (ors *OrdersService) CreateOrderPayment(ctx context.Context, orderID string, ordPay *OrderPayment) (
 res *Response,
 payment *Payment,
 err error,
)
```

CreateOrderPayment can only be created while the status of the order is created,
and when the status of the existing payment is either expired, canceled or
failed.

See <https://docs.mollie.com/reference/v2/orders-api/create-order-payment>

#### func (*OrdersService) CreateOrderRefund

```go
func (ors *OrdersService) CreateOrderRefund(ctx context.Context, orderID string, order *Order) (
 res *Response,
 refund *Refund,
 err error,
)
```

CreateOrderRefund using the Orders API, refunds should be made against the
order.

See <https://docs.mollie.com/reference/v2/orders-api/create-order-refund>

#### func (*OrdersService) Get

```go
func (ors *OrdersService) Get(ctx context.Context, orID string, opts *OrderOptions) (
 res *Response,
 order *Order,
 err error,
)
```

Get retrieve a single order by its ID.

See <https://docs.mollie.com/reference/v2/orders-api/get-order>

#### func (*OrdersService) List

```go
func (ors *OrdersService) List(ctx context.Context, opts *OrderListOptions) (
 res *Response,
 ordList *OrderList,
 err error,
)
```

List is to retrieve all orders.

See <https://docs.mollie.com/reference/v2/orders-api/list-orders>

#### func (*OrdersService) ListOrderRefunds

```go
func (ors *OrdersService) ListOrderRefunds(ctx context.Context, orderID string, opts *OrderListRefundOptions) (
 res *Response,
 orderListRefund *OrderListRefund,
 err error,
)
```

ListOrderRefunds retrieve all order refunds.

See <https://docs.mollie.com/reference/v2/orders-api/list-order-refunds>

#### func (*OrdersService) Update

```go
func (ors *OrdersService) Update(ctx context.Context, orderID string, ord Order) (
 res *Response,
 order *Order,
 err error,
)
```

Update is used to update the billing and/or shipping address of an order.

See <https://docs.mollie.com/reference/v2/orders-api/update-order>

#### func (*OrdersService) UpdateOrderLine

```go
func (ors *OrdersService) UpdateOrderLine(
 ctx context.Context,
 orderID string,
 orderLineID string,
 orderLine OrderLine) (
 res *Response,
 order *Order,
 err error,
)
```

UpdateOrderLine can be used to update an order line.

See <https://docs.mollie.com/reference/v2/orders-api/update-orderline>

#### type Organization

```go
type Organization struct {
 Resource           string            `json:"resource,omitempty"`
 ID                 string            `json:"id,omitempty"`
 Name               string            `json:"name,omitempty"`
 Email              string            `json:"email,omitempty"`
 Locale             string            `json:"locale,omitempty"`
 Address            *Address          `json:"address,omitempty"`
 RegistrationNumber string            `json:"registrationNumber,omitempty"`
 VatNumber          string            `json:"vatNumber,omitempty"`
 VatRegulation      string            `json:"vatRegulation,omitempty"`
 Links              OrganizationLinks `json:"_links,omitempty"`
}
```

Organization describes an organization detail.

#### type OrganizationLinks

```go
type OrganizationLinks struct {
 Self          *URL `json:"self,omitempty"`
 Chargebacks   *URL `json:"chargebacks,omitempty"`
 Customers     *URL `json:"customers,omitempty"`
 Dashboard     *URL `json:"dashboard,omitempty"`
 Invoices      *URL `json:"invoices,omitempty"`
 Payments      *URL `json:"payments,omitempty"`
 Profiles      *URL `json:"profiles,omitempty"`
 Refunds       *URL `json:"refunds,omitempty"`
 Settlements   *URL `json:"settlements,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

OrganizationLinks describes all the possible links to be returned with a
organization object.

#### type OrganizationPartnerLinks

```go
type OrganizationPartnerLinks struct {
 Self          *URL `json:"self,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
 SignUpLink    *URL `json:"signuplink,omitempty"`
}
```

OrganizationPartnerLinks is an object with several URL objects relevant to the
partner resource.

#### type OrganizationPartnerStatus

```go
type OrganizationPartnerStatus struct {
 IsCommissionPartner            bool                     `json:"isCommissionPartner,omitempty"`
 PartnerContractUpdateAvailable bool                     `json:"partnerContractUpdate_available,omitempty"`
 Resource                       string                   `json:"resource,omitempty"`
 PartnerType                    PartnerType              `json:"partnerType,omitempty"`
 UserAgentTokens                []*UserAgentToken        `json:"userAgentTokens,omitempty"`
 PartnerContractSignedAt        *time.Time               `json:"partnerContractSignedAt,omitempty"`
 Links                          OrganizationPartnerLinks `json:"_links,omitempty"`
}
```

OrganizationPartnerStatus response descriptor.

#### type OrganizationsService

```go
type OrganizationsService service
```

OrganizationsService instance operates over organization resources.

#### func (*OrganizationsService) Get

```go
func (os *OrganizationsService) Get(ctx context.Context, id string) (res *Response, o *Organization, err error)
```

Get retrieve an organization by its id.

#### func (*OrganizationsService) GetCurrent

```go
func (os *OrganizationsService) GetCurrent(ctx context.Context) (res *Response, o *Organization, err error)
```

GetCurrent retrieve the currently authenticated organization.

#### func (*OrganizationsService) GetPartnerStatus

```go
func (os *OrganizationsService) GetPartnerStatus(ctx context.Context) (
 res *Response,
 ops *OrganizationPartnerStatus,
 err error,
)
```

GetPartnerStatus retrieves details about the partner status of the currently
authenticated organization.

See: <https://docs.mollie.com/reference/v2/organizations-api/get-partner>

#### type PaginationLinks

```go
type PaginationLinks struct {
 Self          *URL `json:"self,omitempty"`
 Previous      *URL `json:"previous,omitempty"`
 Next          *URL `json:"next,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

PaginationLinks describes the hal component of paginated responses.

#### type PartnerClient

```go
type PartnerClient struct {
 Resource              string             `json:"resource,omitempty"`
 ID                    string             `json:"id,omitempty"`
 OrganizationCreatedAt *time.Time         `json:"organizationCreatedAt,omitempty"`
 Commission            Commission         `json:"commission,omitempty"`
 Links                 PartnerClientLinks `json:"_links,omitempty"`
}
```

PartnerClient describes a partner client.

#### type PartnerClientLinks

```go
type PartnerClientLinks struct {
 Self          *URL `json:"self,omitempty"`
 Organization  *URL `json:"organization,omitempty"`
 Onboarding    *URL `json:"onboarding,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

PartnerClientLinks contains URL objects relevant to the client.

#### type PartnerClientList

```go
type PartnerClientList struct {
 Count          int `json:"count,omitempty"`
 PartnerClients struct {
  Clients []*PartnerClient `json:"clients,omitempty"`
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

PartnerClientList describes a list of partner clients.

#### type PartnerService

```go
type PartnerService service
```

PartnerService operates over the partners API.

#### func (*PartnerService) Get

```go
func (ps *PartnerService) Get(ctx context.Context, id string, opts *GetPartnerClientOptions) (
 res *Response,
 pc *PartnerClient,
 err error,
)
```

Get retrieves a single client, linked to your partner account, by its ID.

See: <https://docs.mollie.com/reference/v2/partners-api/get-client>

#### func (*PartnerService) List

```go
func (ps *PartnerService) List(ctx context.Context, opts *ListPartnerClientsOptions) (
 res *Response,
 pc *PartnerClientList,
 err error,
)
```

List retrieves all clients.

See: <https://docs.mollie.com/reference/v2/partners-api/list-clients>

#### type PartnerType

```go
type PartnerType string
```

PartnerType alias for organization partner types.

```go
const (
 PartnerTypeOauth      PartnerType = "oauth"
 PartnerTypeSignUpLink PartnerType = "signuplink"
 PartnerTypeUserAgent  PartnerType = "useragent"
)
```

Available partner types.

#### type Payment

```go
type Payment struct {
 IsCancellable                   bool                   `json:"isCancellable,omitempty"`
 TestMode                        bool                   `json:"testmode,omitempty"`
 DigitalGoods                    bool                   `json:"digitalGoods,omitempty"`
 ApplePayPaymentToken            string                 `json:"applePayPaymentToken,omitempty"`
 BillingEmail                    string                 `json:"billingEmail,omitempty"`
 CardToken                       string                 `json:"cardToken,omitempty"`
 Issuer                          string                 `json:"issuer,omitempty"`
 VoucherNumber                   string                 `json:"voucherNumber,omitempty"`
 VoucherPin                      string                 `json:"voucherPin,omitempty"`
 ExtraMerchantData               string                 `json:"extraMerchantData,omitempty"`
 SessionID                       string                 `json:"sessionId,omitempty"`
 CustomerReference               string                 `json:"customerReference,omitempty"`
 ConsumerName                    string                 `json:"consumerName,omitempty"`
 ConsumerAccount                 string                 `json:"consumerAccount,omitempty"`
 WebhookURL                      string                 `json:"webhookUrl,omitempty"`
 Resource                        string                 `json:"resource,omitempty"`
 ID                              string                 `json:"id,omitempty"`
 MandateID                       string                 `json:"mandateId,omitempty"`
 OrderID                         string                 `json:"orderId,omitempty"`
 ProfileID                       string                 `json:"profileId,omitempty"`
 SettlementID                    string                 `json:"settlementId,omitempty"`
 CustomerID                      string                 `json:"customerId,omitempty"`
 Status                          string                 `json:"status,omitempty"`
 Description                     string                 `json:"description,omitempty"`
 RedirectURL                     string                 `json:"redirectUrl,omitempty"`
 CountryCode                     string                 `json:"countryCode,omitempty"`
 SubscriptionID                  string                 `json:"subscriptionId,omitempty"`
 CancelURL                       string                 `json:"cancelUrl,omitempty"`
 Metadata                        interface{}            `json:"metadata,omitempty"`
 Amount                          *Amount                `json:"amount,omitempty"`
 AmountRefunded                  *Amount                `json:"amountRefunded,omitempty"`
 AmountRemaining                 *Amount                `json:"amountRemaining,omitempty"`
 AmountCaptured                  *Amount                `json:"amountCaptured,omitempty"`
 AmountChargedBack               *Amount                `json:"amountChargedBack,omitempty"`
 SettlementAmount                *Amount                `json:"settlementAmount,omitempty"`
 ApplicationFee                  *ApplicationFee        `json:"applicationFee,omitempty"`
 Details                         *PaymentDetails        `json:"details,omitempty"`
 CreatedAt                       *time.Time             `json:"createdAt,omitempty"`
 AuthorizedAt                    *time.Time             `json:"authorizedAt,omitempty"`
 PaidAt                          *time.Time             `json:"paidAt,omitempty"`
 CanceledAt                      *time.Time             `json:"canceledAt,omitempty"`
 ExpiresAt                       *time.Time             `json:"expiresAt,omitempty"`
 ExpiredAt                       *time.Time             `json:"expiredAt,omitempty"`
 FailedAt                        *time.Time             `json:"failedAt,omitempty"`
 DueDate                         *ShortDate             `json:"dueDate,omitempty"`
 BillingAddress                  *Address               `json:"billingAddress,omitempty"`
 ShippingAddress                 *PaymentDetailsAddress `json:"shippingAddress,omitempty"`
 Mode                            Mode                   `json:"mode,omitempty"`
 Locale                          Locale                 `json:"locale,omitempty"`
 RestrictPaymentMethodsToCountry Locale                 `json:"restrictPaymentMethodsToCountry,omitempty"`
 Method                          PaymentMethod          `json:"method,omitempty"`
 Links                           PaymentLinks           `json:"_links,omitempty"`
 SequenceType                    SequenceType           `json:"sequenceType,omitempty"`
}
```

Payment describes a transaction between a customer and a merchant.

#### type PaymentDetails

```go
type PaymentDetails struct {
 PaypalDigitalGoods bool                   `json:"digitalGoods,omitempty"`
 BankAccount        string                 `json:"bankAccount,omitempty"`
 BankBIC            string                 `json:"bankBic,omitempty"`
 BankName           string                 `json:"bankName,omitempty"`
 BankReason         string                 `json:"bankReason,omitempty"`
 BatchReference     string                 `json:"batchReference,omitempty"`
 BillingEmail       string                 `json:"billingEmail,omitempty"`
 CardAudience       string                 `json:"cardAudience,omitempty"`
 CardCountryCode    string                 `json:"cardCountryCode,omitempty"`
 CardFingerPrint    string                 `json:"cardFingerPrint,omitempty"`
 CardHolder         string                 `json:"cardHolder,omitempty"`
 CardLabel          string                 `json:"cardLabel,omitempty"`
 CardNumber         string                 `json:"cardNumber,omitempty"`
 CardSecurity       string                 `json:"cardSecurity,omitempty"`
 ConsumerAccount    string                 `json:"consumerAccount,omitempty"`
 ConsumerBIC        string                 `json:"consumerBic,omitempty"`
 ConsumerName       string                 `json:"consumerName,omitempty"`
 ConsumerReference  string                 `json:"consumerReference,omitempty"`
 CreditorIdentifier string                 `json:"creditorIdentifier,omitempty"`
 EndToEndIdentifier string                 `json:"endToEndIdentifier,omitempty"`
 FileReference      string                 `json:"fileReference,omitempty"`
 MandateReference   string                 `json:"mandateReference,omitempty"`
 PaypalReference    string                 `json:"paypalReference,omitempty"`
 PaypalPayerID      string                 `json:"paypalPayerId,omitempty"`
 TransferReference  string                 `json:"transferReference,omitempty"`
 VoucherNumber      string                 `json:"voucherNumber,omitempty"`
 Wallet             string                 `json:"wallet,omitempty"`
 PaypalFee          Amount                 `json:"paypalFee,omitempty"`
 FailureReason      FailureReason          `json:"failureReason,omitempty"`
 FeeRegion          FeeRegion              `json:"feeRegion,omitempty"`
 RemainderMethod    PaymentMethod          `json:"remainderMethod,omitempty"`
 SellerProtection   EligibilityReasons     `json:"sellerProtection,omitempty"`
 DueDate            *ShortDate             `json:"dueDate,omitempty"`
 GiftCards          []*UsedGiftCard        `json:"giftCards,omitempty"`
 QRCode             *QRCode                `json:"qrCode,omitempty"`
 RemainderAmount    *Amount                `json:"remainderAmount,omitempty"`
 ShippingAddress    *PaymentDetailsAddress `json:"shippingAddress,omitempty"`
 SignatureDate      *ShortDate             `json:"signatureDate,omitempty"`
 Links              struct {
  Status    *URL `json:"status,omitempty"`
  PayOnline *URL `json:"payOnline,omitempty"`
 } `json:"_links,omitempty"`
}
```

PaymentDetails contains details for the specified payment method.

#### type PaymentDetailsAddress

```go
type PaymentDetailsAddress struct {
 OrganizationName string      `json:"organizationName,omitempty"`
 Title            string      `json:"title,omitempty"`
 GivenName        string      `json:"givenName,omitempty"`
 FamilyName       string      `json:"familyName,omitempty"`
 Email            string      `json:"email,omitempty"`
 Phone            PhoneNumber `json:"phone,omitempty"`
 StreetAndNumber  string      `json:"streetAndNumber,omitempty"`
 StreetAdditional string      `json:"streetAdditional,omitempty"`
 PostalCode       string      `json:"postalCode,omitempty"`
 City             string      `json:"city,omitempty"`
 Region           string      `json:"region,omitempty"`
 Country          string      `json:"country,omitempty"`
}
```

PaymentDetailsAddress identify both the address and the person the payment is
shipped to.

#### type PaymentLink

```go
type PaymentLink struct {
 ID          string           `json:"id,omitempty"`
 Resource    string           `json:"resource,omitempty"`
 Description string           `json:"description,omitempty"`
 ProfileID   string           `json:"profileId,omitempty"`
 RedirectURL string           `json:"redirectUrl,omitempty"`
 WebhookURL  string           `json:"webhookUrl,omitempty"`
 Mode        Mode             `json:"mode,omitempty"`
 Amount      Amount           `json:"amount,omitempty"`
 CreatedAt   *time.Time       `json:"createdAt,omitempty"`
 PaidAt      *time.Time       `json:"paidAt,omitempty"`
 UpdatedAt   *time.Time       `json:"updatedAt,omitempty"`
 ExpiresAt   *time.Time       `json:"expiresAt,omitempty"`
 Links       PaymentLinkLinks `json:"_links,omitempty"`
}
```

PaymentLink is a resource that can be shared with your customers and will
redirect them to them the payment page where they can complete the payment.

See: <https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link>

#### type PaymentLinkLinks

```go
type PaymentLinkLinks struct {
 Self          *URL `json:"self,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
 PaymentLink   *URL `json:"paymentLink,omitempty"`
 Next          *URL `json:"next,omitempty"`
 Previous      *URL `json:"previous,omitempty"`
}
```

PaymentLinkLinks describes all the possible links returned with a payment link
struct.

See: <https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link>

#### type PaymentLinkOptions

```go
type PaymentLinkOptions struct {
 ProfileID string `url:"profileId,omitempty"`
 From      string `url:"from,omitempty"`
 Limit     int    `url:"limit,omitempty"`
}
```

PaymentLinkOptions represents query string parameters to modify the payment
links requests.

#### type PaymentLinks

```go
type PaymentLinks struct {
 Self               *URL `json:"self,omitempty"`
 Checkout           *URL `json:"checkout,omitempty"`
 ChangePaymentState *URL `json:"changePaymentState,omitempty"`
 Refunds            *URL `json:"refunds,omitempty"`
 ChargeBacks        *URL `json:"chargebacks,omitempty"`
 Captures           *URL `json:"captures,omitempty"`
 Settlement         *URL `json:"settlement,omitempty"`
 Documentation      *URL `json:"documentation,omitempty"`
 Mandate            *URL `json:"mandate,omitempty"`
 Subscription       *URL `json:"subscription,omitempty"`
 Customer           *URL `json:"customer,omitempty"`
 Order              *URL `json:"order,omitempty"`
 Dashboard          *URL `json:"dashboard,omitempty"`
}
```

PaymentLinks describes all the possible links to be returned with a payment
object.

#### type PaymentLinksList

```go
type PaymentLinksList struct {
 Count    int              `json:"count,omitempty"`
 Links    PaymentLinkLinks `json:"_links,omitempty"`
 Embedded struct {
  PaymentLinks []*PaymentLink `json:"payment_links,omitempty"`
 } `json:"_embedded,omitempty"`
}
```

PaymentLinksList retrieves a list of payment links for the active profile or
account token owner.

#### type PaymentLinksService

```go
type PaymentLinksService service
```

PaymentLinksService operates over the payment link resource.

#### func (*PaymentLinksService) Create

```go
func (pls *PaymentLinksService) Create(ctx context.Context, p PaymentLink, opts *PaymentLinkOptions) (
 res *Response,
 np *PaymentLink,
 err error,
)
```

Create generates payment links that by default, unlike regular payments, do not
expire.

See: <https://docs.mollie.com/reference/v2/payment-links-api/create-payment-link>

#### func (*PaymentLinksService) Get

```go
func (pls *PaymentLinksService) Get(ctx context.Context, id string) (res *Response, pl *PaymentLink, err error)
```

Get retrieves a single payment link object by its id/token.

See: <https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link>

#### func (*PaymentLinksService) List

```go
func (pls *PaymentLinksService) List(ctx context.Context, opts *PaymentLinkOptions) (
 res *Response,
 pl *PaymentLinksList,
 err error,
)
```

List retrieves all payments links created with the current website profile,
ordered from newest to oldest.

See: <https://docs.mollie.com/reference/v2/payment-links-api/list-payment-links>

#### type PaymentList

```go
type PaymentList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Payments []Payment
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

PaymentList describes how a list of payments will be retrieved by Mollie.

#### type PaymentMethod

```go
type PaymentMethod string
```

PaymentMethod is a payment method supported by Mollie.

```go
const (
 Bancontact     PaymentMethod = "bancontact"
 BankTransfer   PaymentMethod = "banktransfer"
 Belfius        PaymentMethod = "belfius"
 CreditCard     PaymentMethod = "creditcard"
 DirectDebit    PaymentMethod = "directdebit"
 EPS            PaymentMethod = "eps"
 GiftCard       PaymentMethod = "giftcard"
 GiroPay        PaymentMethod = "giropay"
 IDeal          PaymentMethod = "ideal"
 KBC            PaymentMethod = "kbc"
 KlarnaPayLater PaymentMethod = "klarnapaylater"
 KlarnaSliceIt  PaymentMethod = "klarnasliceit"
 MyBank         PaymentMethod = "mybank"
 PayPal         PaymentMethod = "paypal"
 PaySafeCard    PaymentMethod = "paysafecard"
 PRZelewy24     PaymentMethod = "przelewy24"
 Sofort         PaymentMethod = "sofort"
)
```

Supported payment methods.

#### type PaymentMethodDetails

```go
type PaymentMethodDetails struct {
 Resource      string                  `json:"resource,omitempty"`
 ID            string                  `json:"id,omitempty"`
 Description   string                  `json:"description,omitempty"`
 MinimumAmount *Amount                 `json:"minimumAmount,omitempty"`
 MaximumAmount *Amount                 `json:"maximumAmount,omitempty"`
 Image         *Image                  `json:"image,omitempty"`
 Pricing       []*PaymentMethodPricing `json:"pricing,omitempty"`
 Issuers       []*PaymentMethodIssuer  `json:"issuers,omitempty"`
 Status        *PaymentMethodStatus    `json:"status,omitempty"`
 Links         MethodsLinks            `json:"_links,omitempty"`
}
```

PaymentMethodDetails describes a single method with details.

#### type PaymentMethodIssuer

```go
type PaymentMethodIssuer struct {
 Resource string `json:"resource,omitempty"`
 ID       string `json:"id,omitempty"`
 Name     string `json:"name,omitempty"`
 Image    Image  `json:"image,omitempty"`
}
```

PaymentMethodIssuer available for the payment method (for iDEAL, KBC/CBC payment
button, gift cards, or meal vouchers).

#### type PaymentMethodOptions

```go
type PaymentMethodOptions struct {
 Locale    Locale `url:"locale,omitempty"`
 Currency  string `url:"currency,omitempty"`
 ProfileID string `url:"profileId,omitempty"`
 Include   string `url:"include,omitempty"`
}
```

PaymentMethodOptions are applicable query string parameters to get methods from
mollie's API.

#### type PaymentMethodPricing

```go
type PaymentMethodPricing struct {
 Description string    `json:"description,omitempty"`
 Fixed       *Amount   `json:"fixed,omitempty"`
 Variable    string    `json:"variable,omitempty"`
 FeeRegion   FeeRegion `json:"feeRegion,omitempty"`
}
```

PaymentMethodPricing contains information about commissions and fees applicable
to a payment method.

#### type PaymentMethodStatus

```go
type PaymentMethodStatus string
```

PaymentMethodStatus tels the status that the method is in. Possible values:
activated pending-boarding pending-review pending-external rejected.

```go
const (
 PaymentMethodActivated       PaymentMethodStatus = "activated"
 PaymentMethodPendingBoarding PaymentMethodStatus = "pending-boarding"
 PaymentMethodPendingReview   PaymentMethodStatus = "pending-review"
 PaymentMethodPendingExternal PaymentMethodStatus = "pending-external"
 PaymentMethodRejected        PaymentMethodStatus = "pending-rejected"
)
```

Available payment method statuses.

#### type PaymentMethodsList

```go
type PaymentMethodsList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Methods []*PaymentMethodDetails
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

PaymentMethodsList describes a list of paginated payment methods.

#### type PaymentMethodsListOptions

```go
type PaymentMethodsListOptions struct {
 PaymentMethodOptions
 SequenceType   SequenceType `url:"sequenceType,omitempty"`
 AmountCurrency string       `url:"amount[currency],omitempty"`
 AmountValue    string       `url:"amount[value],omitempty"`
 Resource       string       `url:"resource,omitempty"`
 BillingCountry string       `url:"billingCountry,omitempty"`
 IncludeWallets string       `url:"includeWallets,omitempty"`
}
```

PaymentMethodsListOptions are applicable query string parameters to list methods
from mollie's API.

It contains list specific options and embeds GetMethodOptions.

#### type PaymentMethodsService

```go
type PaymentMethodsService service
```

PaymentMethodsService operates on methods endpoints.

#### func (*PaymentMethodsService) All

```go
func (ms *PaymentMethodsService) All(ctx context.Context, options *PaymentMethodsListOptions) (
 res *Response,
 pm *PaymentMethodsList,
 err error,
)
```

All retrieves all the payment methods enabled for your account/organization.

See: <https://docs.mollie.com/reference/v2/methods-api/list-all-methods>

#### func (*PaymentMethodsService) Get

```go
func (ms *PaymentMethodsService) Get(ctx context.Context, id PaymentMethod, options *PaymentMethodOptions) (
 res *Response,
 pmd *PaymentMethodDetails,
 err error,
)
```

Get returns information about the payment method specified by id, it also
receives a pointer to the method options containing applicable query string
parameters.

See: <https://docs.mollie.com/reference/v2/methods-api/get-method>

#### func (*PaymentMethodsService) List

```go
func (ms *PaymentMethodsService) List(ctx context.Context, options *PaymentMethodsListOptions) (
 res *Response,
 pm *PaymentMethodsList,
 err error,
)
```

List retrieves all enabled payment methods.

The results are not paginated.

See: <https://docs.mollie.com/reference/v2/methods-api/list-methods>

#### type PaymentOptions

```go
type PaymentOptions struct {
 Include string `url:"include,omitempty"`
 Embed   string `url:"embed,omitempty"`
}
```

PaymentOptions describes payments endpoint valid query string parameters.

See: <https://docs.mollie.com/reference/v2/payments-api/get-payment>

#### type PaymentsService

```go
type PaymentsService service
```

PaymentsService instance operates over payment resources.

#### func (*PaymentsService) Cancel

```go
func (ps *PaymentsService) Cancel(ctx context.Context, id string) (res *Response, p *Payment, err error)
```

Cancel removes a payment (if possible) from your Mollie account.

See: <https://docs.mollie.com/reference/v2/payments-api/cancel-payment>

#### func (*PaymentsService) Create

```go
func (ps *PaymentsService) Create(ctx context.Context, p Payment, opts *PaymentOptions) (
 res *Response,
 np *Payment,
 err error,
)
```

Create stores a new payment object attached to your Mollie account.

See: <https://docs.mollie.com/reference/v2/payments-api/create-payment#>

#### func (*PaymentsService) Get

```go
func (ps *PaymentsService) Get(ctx context.Context, id string, opts *PaymentOptions) (
 res *Response,
 p *Payment,
 err error,
)
```

Get retrieves a single payment object by its payment token.

#### func (*PaymentsService) List

```go
func (ps *PaymentsService) List(ctx context.Context, opts *ListPaymentOptions) (
 res *Response,
 pl *PaymentList,
 err error,
)
```

List retrieves a list of payments associated with your account/organization.

See: <https://docs.mollie.com/reference/v2/payments-api/list-payments>

#### func (*PaymentsService) Update

```go
func (ps *PaymentsService) Update(ctx context.Context, id string, up Payment) (res *Response, p *Payment, err error)
```

Update can be used to update some details of a created payment.

See: <https://docs.mollie.com/reference/v2/payments-api/update-payment#>

#### type Permission

```go
type Permission struct {
 Resource    string          `json:"resource,omitempty"`
 ID          PermissionGrant `json:"id,omitempty"`
 Description string          `json:"description,omitempty"`
 Granted     bool            `json:"granted,omitempty"`
 Links       PermissionLinks `json:"_links,omitempty"`
}
```

Permission represents an action that can be performed by any API actor.

#### type PermissionGrant

```go
type PermissionGrant string
```

PermissionGrant defines supported permissions.

```go
const (
 PaymentsRead       PermissionGrant = "payments.read"
 PaymentsWrite      PermissionGrant = "payments.write"
 RefundsRead        PermissionGrant = "refunds.read"
 RefundsWrite       PermissionGrant = "refunds.write"
 CustomersRead      PermissionGrant = "customers.read"
 CustomersWrite     PermissionGrant = "customers.write"
 MandatesRead       PermissionGrant = "mandates.read"
 MandatesWrite      PermissionGrant = "mandates.write"
 SubscriptionsRead  PermissionGrant = "subscriptions.read"
 SubscriptionsWrite PermissionGrant = "subscriptions.write"
 ProfilesRead       PermissionGrant = "profiles.read"
 ProfilesWrite      PermissionGrant = "profiles.write"
 InvoicesRead       PermissionGrant = "invoices.read"
 OrdersRead         PermissionGrant = "orders.read"
 OrdersWrite        PermissionGrant = "orders.write"
 ShipmentsRead      PermissionGrant = "shipments.read"
 ShipmentsWrite     PermissionGrant = "shipments.write"
 OrganizationsRead  PermissionGrant = "organizations.read"
 OrganizationsWrite PermissionGrant = "organizations.write"
 OnboardingRead     PermissionGrant = "onbording.read"
 OnboardingWrite    PermissionGrant = "onbording.write"
)
```

Available permission grants.

#### type PermissionLinks

```go
type PermissionLinks struct {
 Self          *URL `json:"self,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

PermissionLinks contains URL objects that make reference to an http address
related to permissions.

#### type PermissionsList

```go
type PermissionsList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Permissions []*Permission `json:"permissions,omitempty"`
 } `json:"_embedded,omitempty"`
 Links PermissionLinks `json:"_links,omitempty"`
}
```

PermissionsList lists all the permissions given to an API actor.

#### type PermissionsService

```go
type PermissionsService service
```

PermissionsService operates over permission resources.

#### func (*PermissionsService) Get

```go
func (ps *PermissionsService) Get(ctx context.Context, id PermissionGrant) (res *Response, p *Permission, err error)
```

Get returns a permission by its id.

See: <https://docs.mollie.com/reference/v2/permissions-api/get-permission>

#### func (*PermissionsService) List

```go
func (ps *PermissionsService) List(ctx context.Context) (res *Response, pl *PermissionsList, err error)
```

List retrieves all permissions available with the current app access token. The
list is not paginated.

See: <https://docs.mollie.com/reference/v2/permissions-api/list-permissions>

#### type PhoneNumber

```go
type PhoneNumber string
```

PhoneNumber represents a phone number in the E.164 format.

#### type ProductType

```go
type ProductType string
```

ProductType describes the type of product bought, for example, a physical or a
digital product.

```go
const (
 Physical        ProductType = "physical"
 Discount        ProductType = "discount"
 Digital         ProductType = "digital"
 ShippingFee     ProductType = "shipping_fee"
 StoreCredit     ProductType = "store_credit"
 GiftCardProduct ProductType = "gift_card"
 Surcharge       ProductType = "surcharge"
)
```

Valid product type.

#### type Profile

```go
type Profile struct {
 ID               string       `json:"id,omitempty"`
 BusinessCategory BusinessCategory `json:"businessCategory,omitempty"`
 CreatedAt        *time.Time   `json:"createdAt,omitempty"`
 Email            string       `json:"email,omitempty"`
 Mode             Mode         `json:"mode,omitempty"`
 Name             string       `json:"name,omitempty"`
 Phone            PhoneNumber  `json:"phone,omitempty"`
 Resource         string       `json:"resource,omitempty"`
 Review           struct {
  Status string `json:"status,omitempty"`
 } `json:"review,omitempty"`
 Status           ProfileStatus `json:"status,omitempty"`
 Website          string        `json:"website,omitempty"`
 Links            ProfileLinks  `json:"_links,omitempty"`
}
```

Profile will usually reflect the trademark or brand name of the profile’s
website or application.

#### type ProfileLinks

```go
type ProfileLinks struct {
 Self               *URL `json:"self,omitempty"`
 Chargebacks        *URL `json:"chargebacks,omitempty"`
 Methods            *URL `json:"methods,omitempty"`
 Refunds            *URL `json:"refunds,omitempty"`
 CheckoutPreviewURL *URL `json:"checkoutPreviewUrl,omitempty"`
 Documentation      *URL `json:"documentation,omitempty"`
 Dashboard          *URL `json:"dashboard,omitempty"`
}
```

ProfileLinks contains URL's to relevant information related to a profile.

#### type ProfileList

```go
type ProfileList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Profiles []*Profile `json:"profiles,omitempty"`
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

ProfileList contains a list of profiles for your account.

#### type ProfileListOptions

```go
type ProfileListOptions struct {
 From  string `url:"from,omitempty"`
 Limit uint   `url:"limit,omitempty"`
}
```

ProfileListOptions are optional query string parameters for the list profiles
request.

#### type ProfileStatus

```go
type ProfileStatus string
```

ProfileStatus determines whether the profile is able to receive live payments.

```go
const (
 StatusUnverified ProfileStatus = "unverified"
 StatusVerified   ProfileStatus = "verified"
 StatusBlocked    ProfileStatus = "blocked"
)
```

Possible profile statuses.

#### type ProfilesService

```go
type ProfilesService service
```

ProfilesService operates over profile resource.

#### func (*ProfilesService) Create

```go
func (ps *ProfilesService) Create(ctx context.Context, np *Profile) (res *Response, p *Profile, err error)
```

Create stores a new profile in your Mollie account.

#### func (*ProfilesService) Current

```go
func (ps *ProfilesService) Current(ctx context.Context) (res *Response, p *Profile, err error)
```

Current returns the profile belonging to the API key. This method only works
when using API keys.

#### func (*ProfilesService) Delete

```go
func (ps *ProfilesService) Delete(ctx context.Context, id string) (res *Response, err error)
```

Delete enables profile deletions, rendering the profile unavailable for further
API calls and transactions.

#### func (*ProfilesService) DisableGiftCardIssuer

```go
func (ps *ProfilesService) DisableGiftCardIssuer(ctx context.Context, profileID string, issuer GiftCardIssuer) (
 res *Response,
 err error,
)
```

DisableGiftCardIssuer deactivates the requested gift card issuer for the
provided profile id.

See: <https://docs.mollie.com/reference/v2/profiles-api/disable-gift-card-issuer>

#### func (*ProfilesService) DisableGiftCardIssuerForCurrent

```go
func (ps *ProfilesService) DisableGiftCardIssuerForCurrent(ctx context.Context, issuer GiftCardIssuer) (
 res *Response,
 err error,
)
```

DisableGiftCardIssuerForCurrent deactivates the specified issuer for the current
profile (token owner).

See: <https://docs.mollie.com/reference/v2/profiles-api/disable-gift-card-issuer>

#### func (*ProfilesService) DisablePaymentMethod

```go
func (ps *ProfilesService) DisablePaymentMethod(ctx context.Context, id string, pm PaymentMethod) (
 res *Response,
 err error,
)
```

DisablePaymentMethod disables a payment method on a specific or authenticated
profile. If you're using API tokens for authentication, pass "me" as id.

#### func (*ProfilesService) EnableGiftCardIssuer

```go
func (ps *ProfilesService) EnableGiftCardIssuer(ctx context.Context, profileID string, issuer GiftCardIssuer) (
 res *Response,
 gc *GiftCardEnabled,
 err error,
)
```

EnableGiftCardIssuer activates the requested gift card issuer for the provided
profile id.

See: <https://docs.mollie.com/reference/v2/profiles-api/enable-gift-card-issuer>

#### func (*ProfilesService) EnableGiftCardIssuerForCurrent

```go
func (ps *ProfilesService) EnableGiftCardIssuerForCurrent(ctx context.Context, issuer GiftCardIssuer) (
 res *Response,
 gc *GiftCardEnabled,
 err error,
)
```

EnableGiftCardIssuerForCurrent activates the specified issuer for the current
profile (token owner).

See: <https://docs.mollie.com/reference/v2/profiles-api/enable-gift-card-issuer>

#### func (*ProfilesService) EnablePaymentMethod

```go
func (ps *ProfilesService) EnablePaymentMethod(ctx context.Context, id string, pm PaymentMethod) (
 res *Response,
 pmi *PaymentMethodDetails,
 err error,
)
```

EnablePaymentMethod enables a payment method on a specific or authenticated
profile. If you're using API tokens for authentication, pass "me" as id.

#### func (*ProfilesService) Get

```go
func (ps *ProfilesService) Get(ctx context.Context, id string) (res *Response, p *Profile, err error)
```

Get retrieves the a profile by ID.

#### func (*ProfilesService) List

```go
func (ps *ProfilesService) List(ctx context.Context, opts *ProfileListOptions) (
 res *Response,
 pl *ProfileList,
 err error,
)
```

List returns all the profiles for the authenticated account.

#### func (*ProfilesService) Update

```go
func (ps *ProfilesService) Update(ctx context.Context, id string, up *Profile) (res *Response, p *Profile, err error)
```

Update allows you to perform mutations on a profile.

#### type QRCode

```go
type QRCode struct {
 Height int    `json:"height,omitempty"`
 Width  int    `json:"width,omitempty"`
 Src    string `json:"src,omitempty"`
}
```

QRCode object represents an image of a QR code.

#### type Rate

```go
type Rate struct {
 Fixed    *Amount `json:"fixed,omitempty"`
 Variable string  `json:"variable,omitempty"`
}
```

Rate describes service rates, further divided into fixed and percentage costs.

#### type Refund

```go
type Refund struct {
 Resource         string       `json:"resource,omitempty"`
 ID               string       `json:"id,omitempty"`
 Amount           *Amount      `json:"amount,omitempty"`
 SettlementID     string       `json:"settlementId,omitempty"`
 SettlementAmount *Amount      `json:"settlementAmount,omitempty"`
 Description      string       `json:"description,omitempty"`
 Metadata         interface{}  `json:"metadata,omitempty"`
 Status           RefundStatus `json:"status,omitempty"`
 Lines            []*OrderLine `json:"lines,omitempty"`
 PaymentID        string       `json:"paymentId,omitempty"`
 OrderID          string       `json:"orderId,omitempty"`
 CreatedAt        *time.Time   `json:"createdAt,omitempty"`
 TestMode         bool         `json:"testmode,omitempty"`
 Links            RefundLinks  `json:"_links,omitempty"`
}
```

Refund describe a refund for a certain payment.

#### type RefundLinks

```go
type RefundLinks struct {
 Self          *URL `json:"self,omitempty"`
 Payment       *URL `json:"payment,omitempty"`
 Settlement    *URL `json:"settlement,omitempty"`
 Order         *URL `json:"order,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

RefundLinks describes all the possible links to be returned with a Refund
object.

#### type RefundList

```go
type RefundList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Refunds []*Refund
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

RefundList describes how a list of refunds will be retrieved by Mollie.

#### type RefundOptions

```go
type RefundOptions struct {
 Embed EmbedValue `url:"embed,omitempty"`
}
```

RefundOptions describes refund endpoint valid query string parameters.

See: <https://docs.mollie.com/reference/v2/refunds-api/get-refund>.

#### type RefundStatus

```go
type RefundStatus string
```

RefundStatus describes the status of the refund.

```go
const (
 Queued     RefundStatus = "queued"
 Pending    RefundStatus = "pending"
 Processing RefundStatus = "processing"
 Refunded   RefundStatus = "refunded"
 Failed     RefundStatus = "failed"
)
```

Valid refund status.

#### type RefundsService

```go
type RefundsService service
```

RefundsService instance operates over refund resources.

#### func (*RefundsService) Cancel

```go
func (rs *RefundsService) Cancel(ctx context.Context, paymentID, refundID string) (res *Response, err error)
```

Cancel try to cancel the refund request. The refund can only be canceled while
the refund’s status is either queued or pending. See
<https://docs.mollie.com/reference/v2/refunds-api/cancel-refund>

#### func (*RefundsService) Create

```go
func (rs *RefundsService) Create(ctx context.Context, paymentID string, re Refund, options *RefundOptions) (
 res *Response,
 rf *Refund,
 err error,
)
```

Create a refund payment request.

See <https://docs.mollie.com/reference/v2/refunds-api/create-refund>.

#### func (*RefundsService) Get

```go
func (rs *RefundsService) Get(ctx context.Context, paymentID, refundID string, opts *RefundOptions) (
 res *Response,
 refund *Refund,
 err error,
)
```

Get retrieve a single refund by its ID.

If you do not know the original payment’s ID, you can use the List payment
refunds endpoint.

#### func (*RefundsService) ListRefund

```go
func (rs *RefundsService) ListRefund(ctx context.Context, opts *ListRefundOptions) (
 res *Response,
 rl *RefundList,
 err error,
)
```

ListRefund calls the top level <https://api.mollie.com/v2/refunds>.

See <https://docs.mollie.com/reference/v2/refunds-api/list-refunds>.

#### func (*RefundsService) ListRefundPayment

```go
func (rs *RefundsService) ListRefundPayment(ctx context.Context, paymentID string, opts *ListRefundOptions) (
 res *Response,
 rl *RefundList,
 err error,
)
```

ListRefundPayment calls the payment-specific
<https://api.mollie.com/v2/payments/*paymentId*/refunds>.

Only refunds for that specific payment are returned. See:
<https://docs.mollie.com/reference/v2/refunds-api/list-refunds>

#### type Response

```go
type Response struct {
 *http.Response
}
```

Response is a Mollie API response. This wraps the standard http.Response
returned from Mollie and provides convenient access to things like pagination
links.

#### type SequenceType

```go
type SequenceType string
```

SequenceType indicates which type of payment this is in a recurring sequence.

```go
const (
 OneOffSequence    SequenceType = "oneoff"
 FirstSequence     SequenceType = "first"
 RecurringSequence SequenceType = "recurring"
)
```

Valid sequence types.

#### type Settlement

```go
type Settlement struct {
 ID        string           `json:"id,omitempty"`
 Resource  string           `json:"resource,omitempty"`
 Reference string           `json:"reference,omitempty"`
 CreatedAt *time.Time       `json:"createdAt,omitempty"`
 SettledAt *time.Time       `json:"settledAt,omitempty"`
 Status    SettlementStatus `json:"status,omitempty"`
 Amount    *Amount          `json:"amount,omitempty"`
 Periods   SettlementObject `json:"periods,omitempty"`
 InvoiceID string           `json:"invoiceId,omitempty"`
 Links     SettlementLinks  `json:"_links,omitempty"`
}
```

Settlement contains successful payments, together with refunds, captures and
chargebacks into settlements.

#### type SettlementCosts

```go
type SettlementCosts struct {
 Description string        `json:"description,omitempty"`
 AmountNet   *Amount       `json:"amountNet,omitempty"`
 AmountVAT   *Amount       `json:"amountVat,omitempty"`
 AmountGross *Amount       `json:"amountGross,omitempty"`
 Count       int           `json:"count,omitempty"`
 Rate        *Rate         `json:"rate,omitempty"`
 Method      PaymentMethod `json:"method,omitempty"`
}
```

SettlementCosts contains information about costs related to a settlement.

#### type SettlementLinks

```go
type SettlementLinks struct {
 Self          *URL `json:"self,omitempty"`
 Payments      *URL `json:"payments,omitempty"`
 Refunds       *URL `json:"refunds,omitempty"`
 Chargebacks   *URL `json:"chargebacks,omitempty"`
 Captures      *URL `json:"captures,omitempty"`
 Invoice       *URL `json:"invoice,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

SettlementLinks is an object with several URL objects relevant to the
settlement.

#### type SettlementObject

```go
type SettlementObject map[string]map[string]SettlementPeriod
```

SettlementObject nests as describes for settlement periods.

#### type SettlementPeriod

```go
type SettlementPeriod struct {
 Revenue []*SettlementRevenue `json:"revenue,omitempty"`
 Costs   []*SettlementCosts   `json:"costs,omitempty"`
}
```

SettlementPeriod describe the settlement by month in full detail.

#### type SettlementRevenue

```go
type SettlementRevenue struct {
 Description string        `json:"description,omitempty"`
 AmountNet   *Amount       `json:"amountNet,omitempty"`
 AmountVAT   *Amount       `json:"amountVat,omitempty"`
 AmountGross *Amount       `json:"amountGross,omitempty"`
 Count       int           `json:"count,omitempty"`
 Method      PaymentMethod `json:"method,omitempty"`
}
```

SettlementRevenue objects contain the total revenue for each payment method
during this period.

#### type SettlementStatus

```go
type SettlementStatus string
```

SettlementStatus describes the status of the settlement.

```go
const (
 SettlementStatusOpen    SettlementStatus = "open"
 SettlementStatusPending SettlementStatus = "pending"
 SettlementStatusPaidOut SettlementStatus = "paidout"
 SettlementStatusFailed  SettlementStatus = "failed"
)
```

Available settlement statuses.

#### type SettlementsList

```go
type SettlementsList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Settlements []*Settlement
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

SettlementsList describes a list of settlements.

#### type SettlementsListOptions

```go
type SettlementsListOptions struct {
 From  *ShortDate `url:"from,omitempty"`
 Limit int        `url:"limit,omitempty"`
 Embed EmbedValue `url:"embed,omitempty"`
}
```

SettlementsListOptions contains query parameters for settlement lists.

#### type SettlementsService

```go
type SettlementsService service
```

SettlementsService operates over settlements resource.

#### func (*SettlementsService) Get

```go
func (ss *SettlementsService) Get(ctx context.Context, id string) (res *Response, s *Settlement, err error)
```

Get returns a settlement by its id or the bank reference id

See: <https://docs.mollie.com/reference/v2/settlements-api/get-settlement>

#### func (*SettlementsService) GetCaptures

```go
func (ss *SettlementsService) GetCaptures(ctx context.Context, id string, slo *SettlementsListOptions) (
 res *Response,
 cl *CapturesList,
 err error,
)
```

GetCaptures retrieves all captures included in a settlement.

See:
<https://docs.mollie.com/reference/v2/settlements-api/list-settlement-captures>

#### func (*SettlementsService) GetChargebacks

```go
func (ss *SettlementsService) GetChargebacks(ctx context.Context, id string, slo *SettlementsListOptions) (
 res *Response,
 cl *ChargebacksList,
 err error,
)
```

GetChargebacks retrieves all chargebacks included in a settlement.

See:
<https://docs.mollie.com/reference/v2/settlements-api/list-settlement-chargebacks>

#### func (*SettlementsService) GetPayments

```go
func (ss *SettlementsService) GetPayments(ctx context.Context, id string, slo *SettlementsListOptions) (
 res *Response,
 pl *PaymentList,
 err error,
)
```

GetPayments retrieves all payments included in a settlement.

See:
<https://docs.mollie.com/reference/v2/settlements-api/list-settlement-payments>

#### func (*SettlementsService) GetRefunds

```go
func (ss *SettlementsService) GetRefunds(ctx context.Context, id string, slo *SettlementsListOptions) (
 res *Response,
 rl *RefundList,
 err error,
)
```

GetRefunds retrieves all refunds included in a settlement.

See:
<https://docs.mollie.com/reference/v2/settlements-api/list-settlement-refunds>

#### func (*SettlementsService) List

```go
func (ss *SettlementsService) List(ctx context.Context, slo *SettlementsListOptions) (
 res *Response,
 sl *SettlementsList,
 err error,
)
```

List retrieves all settlements, ordered from new to old

See: <https://docs.mollie.com/reference/v2/settlements-api/list-settlements>

#### func (*SettlementsService) Next

```go
func (ss *SettlementsService) Next(ctx context.Context) (res *Response, s *Settlement, err error)
```

Next retrieves the details of the current settlement that has not yet been paid
out.

See: <https://docs.mollie.com/reference/v2/settlements-api/get-next-settlement>

#### func (*SettlementsService) Open

```go
func (ss *SettlementsService) Open(ctx context.Context) (res *Response, s *Settlement, err error)
```

Open retrieves the details of the open balance of the organization. This will
return a settlement object representing your organization’s balance.

See: <https://docs.mollie.com/reference/v2/settlements-api/get-open-settlement>

#### type Shipment

```go
type Shipment struct {
 Resource  string            `json:"resource,omitempty"`
 ID        string            `json:"id,omitempty"`
 OrderID   string            `json:"orderId,omitempty"`
 TestMode  bool              `json:"testmode,omitempty"`
 CreatedAt *time.Time        `json:"createdAt,omitempty"`
 Tracking  *ShipmentTracking `json:"tracking,omitempty"`
 Lines     []*OrderLine      `json:"lines,omitempty"`
 Links     ShipmentLinks     `json:"_links,omitempty"`
}
```

Shipment contains information about a user service/product delivery and is used
in the figurative sense here. It can also mean that a service was provided or
digital content was delivered.

#### type ShipmentLinks

```go
type ShipmentLinks struct {
 Self          *URL `json:"self,omitempty"`
 Order         *URL `json:"order,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

ShipmentLinks contains URL objects with shipment relevant information for the
user.

#### type ShipmentTracking

```go
type ShipmentTracking struct {
 Carrier string `json:"carrier,omitempty"`
 Code    string `json:"code,omitempty"`
 URL     string `json:"url,omitempty"`
}
```

ShipmentTracking contains shipment tracking details.

#### type ShipmentsList

```go
type ShipmentsList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Shipments []Shipment
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

ShipmentsList describes how a list of payments will be retrieved by Mollie.

#### type ShipmentsService

```go
type ShipmentsService service
```

ShipmentsService operates on shipments endpoints.

#### func (*ShipmentsService) Create

```go
func (ss *ShipmentsService) Create(ctx context.Context, oID string, cs CreateShipmentRequest) (
 res *Response,
 s *Shipment,
 err error,
)
```

Create can be used to ship order lines.

See: <https://docs.mollie.com/reference/v2/shipments-api/create-shipment>

#### func (*ShipmentsService) Get

```go
func (ss *ShipmentsService) Get(ctx context.Context, oID string, sID string) (res *Response, s *Shipment, err error)
```

Get retrieves a single shipment and the order lines shipped by a shipment’s ID.

See: <https://docs.mollie.com/reference/v2/shipments-api/get-shipment#>

#### func (*ShipmentsService) List

```go
func (ss *ShipmentsService) List(ctx context.Context, oID string) (res *Response, sl *ShipmentsList, err error)
```

List retrieves all shipments for an order.

See: <https://docs.mollie.com/reference/v2/shipments-api/list-shipments>

#### func (*ShipmentsService) Update

```go
func (ss *ShipmentsService) Update(ctx context.Context, oID string, sID string, st ShipmentTracking) (
 res *Response,
 s *Shipment,
 err error,
)
```

Update can be used to update the tracking information of a shipment

See: <https://docs.mollie.com/reference/v2/shipments-api/update-shipment>

#### type ShortDate

```go
type ShortDate struct {
 time.Time
}
```

ShortDate is a string representing a date in YYYY-MM-DD format.

#### func (*ShortDate) MarshalJSON

```go
func (d *ShortDate) MarshalJSON() ([]byte, error)
```

MarshalJSON overrides the default marshal action for the Date struct. Returns
date as YYYY-MM-DD formatted string.

#### func (*ShortDate) UnmarshalJSON

```go
func (d *ShortDate) UnmarshalJSON(b []byte) error
```

UnmarshalJSON overrides the default unmarshal action for the Date struct, as we
need links to be pointers to the time.Time struct.

#### type Subscription

```go
type Subscription struct {
 Resource        string             `json:"resource,omitempty"`
 ID              string             `json:"id,omitempty"`
 MandateID       string             `json:"mandateId,omitempty"`
 Mode            Mode               `json:"mode,omitempty"`
 CreatedAT       *time.Time         `json:"createdAt,omitempty"`
 Status          SubscriptionStatus `json:"status,omitempty"`
 Amount          *Amount            `json:"amount,omitempty"`
 Times           int                `json:"times,omitempty"`
 TimesRemaining  int                `json:"timesRemaining,omitempty"`
 Interval        string             `json:"interval,omitempty"`
 StartDate       *ShortDate         `json:"startDate,omitempty"`
 NextPaymentDate *ShortDate         `json:"nextPaymentDate,omitempty"`
 Description     string             `json:"description,omitempty"`
 Method          PaymentMethod      `json:"method,omitempty"`
 CanceledAt      *time.Time         `json:"canceledAt,omitempty"`
 WebhookURL      string             `json:"webhookUrl,omitempty"`
 Metadata        interface{}        `json:"metadata,omitempty"`
 ApplicationFee  *ApplicationFee    `json:"applicationFee,omitempty"`
 TestMode        bool               `json:"testmode,omitempty"`
 Links           SubscriptionLinks  `json:"_links,omitempty"`
}
```

Subscription contains information about a customer subscription.

#### type SubscriptionLinks

```go
type SubscriptionLinks struct {
 Self          *URL `json:"self,omitempty"`
 Customer      *URL `json:"customer,omitempty"`
 Payments      *URL `json:"payments,omitempty"`
 Documentation *URL `json:"documentation,omitempty"`
}
```

SubscriptionLinks contains several URL objects relevant to the subscription.

#### type SubscriptionList

```go
type SubscriptionList struct {
 Count    int `json:"count,omitempty"`
 Embedded struct {
  Subscriptions []*Subscription
 } `json:"_embedded,omitempty"`
 Links PaginationLinks `json:"_links,omitempty"`
}
```

SubscriptionList describes the response for subscription list endpoints.

#### type SubscriptionListOptions

```go
type SubscriptionListOptions struct {
 From      string `url:"from,omitempty"`
 Limit     int    `url:"limit,omitempty"`
 ProfileID string `url:"profileId,omitempty"`
}
```

SubscriptionListOptions holds query string parameters valid for subscription
lists.

#### type SubscriptionStatus

```go
type SubscriptionStatus string
```

SubscriptionStatus contains references to valid subscription statuses.

```go
const (
 SubscriptionStatusPending   SubscriptionStatus = "pending"
 SubscriptionStatusActive    SubscriptionStatus = "active"
 SubscriptionStatusCanceled  SubscriptionStatus = "canceled"
 SubscriptionStatusSuspended SubscriptionStatus = "suspended"
 SubscriptionStatusCompleted SubscriptionStatus = "completed"
)
```

Available subscription statuses.

#### type SubscriptionsService

```go
type SubscriptionsService service
```

SubscriptionsService operates over subscriptions resource.

#### func (*SubscriptionsService) All

```go
func (ss *SubscriptionsService) All(ctx context.Context, opts *SubscriptionListOptions) (
 res *Response,
 sl *SubscriptionList,
 err error,
)
```

All retrieves all subscriptions, ordered from newest to oldest. By using an API
key all the subscriptions created with the current website profile will be
returned. In the case of an OAuth Access Token relies the website profile on the
profileId field

See:
<https://docs.mollie.com/reference/v2/subscriptions-api/list-all-subscriptions>

#### func (*SubscriptionsService) Create

```go
func (ss *SubscriptionsService) Create(ctx context.Context, cID string, sc *Subscription) (
 res *Response,
 s *Subscription,
 err error,
)
```

Create stores a new subscription for a given customer

See: <https://docs.mollie.com/reference/v2/subscriptions-api/create-subscription>

#### func (*SubscriptionsService) Delete

```go
func (ss *SubscriptionsService) Delete(ctx context.Context, cID, sID string) (
 res *Response,
 s *Subscription,
 err error,
)
```

Delete cancels a subscription

See: <https://docs.mollie.com/reference/v2/subscriptions-api/cancel-subscription>

#### func (*SubscriptionsService) Get

```go
func (ss *SubscriptionsService) Get(ctx context.Context, cID, sID string) (res *Response, s *Subscription, err error)
```

Get retrieves a customer's subscription

See: <https://docs.mollie.com/reference/v2/subscriptions-api/get-subscription>

#### func (*SubscriptionsService) GetPayments

```go
func (ss *SubscriptionsService) GetPayments(ctx context.Context, cID, sID string, opts *SubscriptionListOptions) (
 res *Response,
 sl *PaymentList,
 err error,
)
```

GetPayments retrieves all payments of a specific subscriptions of a customer

See:
<https://docs.mollie.com/reference/v2/subscriptions-api/list-subscriptions-payments>

#### func (*SubscriptionsService) List

```go
func (ss *SubscriptionsService) List(ctx context.Context, cID string, opts *SubscriptionListOptions) (
 res *Response,
 sl *SubscriptionList,
 err error,
)
```

List retrieves all subscriptions of a customer

See: <https://docs.mollie.com/reference/v2/subscriptions-api/list-subscriptions>

#### func (*SubscriptionsService) Update

```go
func (ss *SubscriptionsService) Update(ctx context.Context, cID, sID string, sc *Subscription) (
 res *Response,
 s *Subscription,
 err error,
)
```

Update changes fields on a subscription object

See: <https://docs.mollie.com/reference/v2/subscriptions-api/update-subscription>

#### type Subtotal

```go
type Subtotal struct {
 TransactionType string      `json:"transactionType,omitempty"`
 Count           int         `json:"count,omitempty"`
 Amount          *Amount     `json:"amount,omitempty"`
 Subtotals       []*Subtotal `json:"subtotals,omitempty"`
}
```

Subtotal balance descriptor.

#### type TransactionType

```go
type TransactionType string
```

TransactionType specifies the reason for the movement.

```go
const (
 PaymentTransaction                     TransactionType = "payment"
 CaptureTransaction                     TransactionType = "capture"
 UnauthorizedDirectDebitTransaction     TransactionType = "unauthorized-direct-debit"
 FailedPaymentTransaction               TransactionType = "failed-payment"
 RefundTransaction                      TransactionType = "refund-transaction"
 ReturnedRefundTransaction              TransactionType = "returned-refund"
 ChargebackTransaction                  TransactionType = "chargeback"
 ChargebackReversalTransaction          TransactionType = "chargeback-reversal"
 OutgoingTransferTransaction            TransactionType = "outgoing-transfer"
 CanceledOutgoingTransfer               TransactionType = "canceled-outgoing-transfer"
 ReturnedTransferTransaction            TransactionType = "returned-transfer"
 InvoiceCompensationTransferTransaction TransactionType = "invoice-compensation"
 BalanceCorrectionTransaction           TransactionType = "balance-correction"
 ApplicationFeeTransaction              TransactionType = "application-fee"
 SplitPaymentTransaction                TransactionType = "split-payment"
 PlatformPaymentRefundTransaction       TransactionType = "platform-payment-refund"
 PlatformPaymentChargeback              TransactionType = "platform-payment-chargeback"
)
```

Known and supported transaction types.

#### type TransferDestination

```go
type TransferDestination struct {
 Type            string `json:"type,omitempty"`
 BankAccount     string `json:"bankAccount,omitempty"`
 BeneficiaryName string `json:"beneficiaryName,omitempty"`
}
```

TransferDestination where the available amount will be automatically
transferred.

#### type TransferFrequency

```go
type TransferFrequency string
```

TransferFrequency reflects the frequency at which the available amount on the
balance will be settled to the configured transfer destination.

```go
const (
 TransferDaily          TransferFrequency = "daily"
 TransferTwiceAWeek     TransferFrequency = "twice-a-week"
 TransferEveryMonday    TransferFrequency = "every-monday"
 TransferEveryTuesday   TransferFrequency = "every-tuesday"
 TransferEveryWednesday TransferFrequency = "every-wednesday"
 TransferEveryThursday  TransferFrequency = "every-thursday"
 TransferEveryFriday    TransferFrequency = "every-friday"
 TransferTwiceAMonth    TransferFrequency = "twice-a-month"
 TransferMonthly        TransferFrequency = "monthly"
 TransferNever          TransferFrequency = "never"
)
```

Possible values for type TransferFrequency.

#### type URL

```go
type URL struct {
 Href string `json:"href,omitempty"`
 Type string `json:"type,omitempty"`
}
```

URL in Mollie are commonly represented as objects with an href and type field.

#### type UsedGiftCard

```go
type UsedGiftCard struct {
 Issuer        string  `json:"issuer,omitempty"`
 Amount        *Amount `json:"amount,omitempty"`
 VoucherNumber string  `json:"voucherNumber,omitempty"`
}
```

UsedGiftCard describes a used gift card.

#### type UserAgentToken

```go
type UserAgentToken struct {
 Token    string
 StartsAt *time.Time
 EndsAt   *time.Time
}
```

UserAgentToken are time limited valid access tokens.
