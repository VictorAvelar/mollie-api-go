package mollie

// ApplicationFee allows you to split a payment between a platform and connected merchant accounts.
type ApplicationFee struct {
	Amount      *Amount `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
}
