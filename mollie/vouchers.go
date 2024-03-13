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
