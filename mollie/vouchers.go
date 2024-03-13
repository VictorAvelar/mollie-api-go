package mollie

// VoucherIssuer represents the issuer of a voucher.
type VoucherIssuer string

// List of known voucher issuers.
const (
	EdenredBelgiumCadeauVoucher     VoucherIssuer = "edenred-belgium-cadeau"
	EdenredBelgiumEcoVoucher        VoucherIssuer = "edenred-belgium-eco"
	EdenredBelgiumMealVoucher       VoucherIssuer = "edenred-belgium-meal"
	EdenredBelgiumSportsVoucher     VoucherIssuer = "edenred-belgium-sports"
	EdenredBelgiumAdditionalVoucher VoucherIssuer = "edenred-belgium-additional"
	EdenredBelgiumConsumeVoucher    VoucherIssuer = "edenred-belgium-consume"
	MonizzeCadeauVoucher            VoucherIssuer = "monizze-cadeau"
	MonizzeEcoVoucher               VoucherIssuer = "monizze-eco"
	MonizzeMealVoucher              VoucherIssuer = "monizze-meal"
	PluxeeCadeauVoucher             VoucherIssuer = "sodexo-cadeau"
	PluxeeEcoVoucher                VoucherIssuer = "sodexo-ecopass"
	PluxeeLunchVoucher              VoucherIssuer = "sodexo-lunchpass"
)

// VoucherEnabled describes the response of a voucher enable operation.
type VoucherEnabled struct {
	ID          string            `json:"id,omitempty"`
	Description string            `json:"description,omitempty"`
	Status      IssuerStatus      `json:"status,omitempty"`
	Contractor  VoucherContractor `json:"contractor,omitempty"`
	Links       VoucherLinks      `json:"_links,omitempty"`
}

// VoucherLinks are links embedded when a voucher is enabled.
type VoucherLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}

// VoucherContractor represents a contractor for a voucher.
type VoucherContractor struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ContractorID string `json:"contractorId,omitempty"`
}
