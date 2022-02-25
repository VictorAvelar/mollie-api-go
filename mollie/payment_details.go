package mollie

// FeeRegion contains the fee region for the payment.
type FeeRegion string

// Valid Fee regions.
const (
	AmericanExpress FeeRegion = "american-express"
	CarteBancaire   FeeRegion = "carte-bancaire"
	IntraEU         FeeRegion = "intra-eu"
	MaestroRegion   FeeRegion = "maestro"
	Other           FeeRegion = "other"
)

// FailureReason provides a description on a failed payment.
type FailureReason string

// Possible failure reasons.
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

// EligibilityReasons for paypal seller protection.
type EligibilityReasons string

// Available eligibility reasons for Paypal sellers protection.
const (
	Eligible                        EligibilityReasons = "Eligible"
	Ineligible                      EligibilityReasons = "Ineligible"
	INROnly                         EligibilityReasons = "Partially Eligible - INR Only"
	AuthOnly                        EligibilityReasons = "Partially Eligible - Unauth Only"
	PartiallyEligible               EligibilityReasons = "PartiallyEligible"
	EligibilityNone                 EligibilityReasons = "None"
	ActiveFraudControlUnAuthPremium EligibilityReasons = "Active Fraud Control - Unauth Premium Eligible"
)

// UsedGiftCard describes a used gift card.
type UsedGiftCard struct {
	Issuer        string  `json:"issuer,omitempty"`
	Amount        *Amount `json:"amount,omitempty"`
	VoucherNumber string  `json:"voucherNumber,omitempty"`
}

// PaymentDetails contains details for the specified payment method.
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

// PaymentDetailsAddress identify both the address and the person the payment is shipped to.
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
