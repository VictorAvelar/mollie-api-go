package mollie

// FeeRegion contains the fee region for the payment.
type FeeRegion string

// Valid Fee regions
const (
	AmericanExpress FeeRegion = "american-express"
	CarteBancaire   FeeRegion = "carte-bancaire"
	IntraEU         FeeRegion = "intra-eu"
	MaestroRegion   FeeRegion = "maestro"
	Other           FeeRegion = "other"
)

type FailureReason string

const (
	ReasonInvalidCardNumber     FailureReason = "invalid_card_number"
	ReasonInvalidCCV            FailureReason = "invalid_ccv"
	ReasonInvalidCardHolderName FailureReason = "invalid_card_holder_name"
	ReasonCardExpired           FailureReason = "card_expired"
	ReasonInvalidCardType       FailureReason = "invalid_card_type"
	ReasonRefusedByIssuer       FailureReason = "refused_by_issuer"
	ReasonInsufficientFunds     FailureReason = "insufficient_funds"
	ReasonInactiveCard          FailureReason = "inactive_card"
	ReasonUnknown               FailureReason = "unknown_reason"
	ReasonPossibleFraud         FailureReason = "possible_fraud"
)

// UsedGiftCard describes a used gift card.
type UsedGiftCard struct {
	Issuer        string  `json:"issuer,omitempty"`
	Amount        *Amount `json:"amount,omitempty"`
	VoucherNumber string  `json:"voucherNumber,omitempty"`
}

// BankTransferDetails contains details for the specified payment method
type PaymentDetails struct {
	BankAccount        string         `json:"bankAccount,omitempty"`
	BankBIC            string         `json:"bankBic,omitempty"`
	BankName           string         `json:"bankName,omitempty"`
	BankReason         string         `json:"bankReason,omitempty"`
	BatchReference     string         `json:"batchReference,omitempty"`
	BillingEmail       string         `json:"billingEmail,omitempty"`
	CardAudience       string         `json:"cardAudience,omitempty"`
	CardCountryCode    string         `json:"cardCountryCode,omitempty"`
	CardFingerPrint    string         `json:"cardFingerPrint,omitempty"`
	CardHolder         string         `json:"cardHolder,omitempty"`
	CardLabel          string         `json:"cardLabel,omitempty"`
	CardNumber         string         `json:"cardNumber,omitempty"`
	CardSecurity       string         `json:"cardSecurity,omitempty"`
	ConsumerAccount    string         `json:"consumerAccount,omitempty"`
	ConsumerBIC        string         `json:"consumerBic,omitempty"`
	ConsumerName       string         `json:"consumerName,omitempty"`
	ConsumerReference  string         `json:"consumerReference,omitempty"`
	CreditorIdentifier string         `json:"creditorIdentifier,omitempty"`
	DueDate            *ShortDate     `json:"dueDate,omitempty"`
	EndToEndIdentifier string         `json:"endToEndIdentifier,omitempty"`
	FailureReason      *FailureReason `json:"failureReason,omitempty"`
	FeeRegion          *FeeRegion     `json:"feeRegion,omitempty"`
	FileReference      string         `json:"fileReference,omitempty"`
	GiftCards          []UsedGiftCard `json:"giftCards,omitempty"`
	MandateReference   string         `json:"mandateReference,omitempty"`
	PayPalReference    string         `json:"payPalReference,omitempty"`
	QRCode             *QRCode        `json:"qrCode,omitempty"`
	RemainderAmount    *Amount        `json:"remainderAmount,omitempty"`
	RemainderMethod    *PaymentMethod `json:"remainderMethod,omitempty"`
	SignatureDate      *ShortDate     `json:"signatureDate,omitempty"`
	TransferReference  string         `json:"transferReference,omitempty"`
	VoucherNumber      string         `json:"voucherNumber,omitempty"`
	Wallet             string         `json:"wallet,omitempty"`
	Links              struct {
		Status    URL `json:"status,omitempty"`
		PayOnline URL `json:"payOnline,omitempty"`
	} `json:"_links,omitempty"`
}
