package mollie

// GiftCardIssuer type describes issuers supported
// by mollie.
type GiftCardIssuer string

// sportenfitcadeau, sustainablefashion, travelcheq, vvvgiftcard, vvvdinercheque,
// vvvlekkerweg, webshopgiftcard, yourgift

// Supported gift card issuers
const (
	DecaudeuKaart              GiftCardIssuer = "decadeaukaart"
	Dinercadeau                               = "dinercadeau"
	Fashioncheque                             = "fashioncheque"
	Festivalcadeau                            = "festivalcadeau"
	Good4fun                                  = "good4fun"
	Kunstencultuurcadeaukaart                 = "kunstencultuurcadeaukaart"
	Nationalebioscoopbon                      = "nationalebioscoopbon"
	Nationaleentertainmentcard                = "nationaleentertainmentcard"
	Nationalegolfbon                          = "nationalegolfbon"
	Ohmygood                                  = "ohmygood"
	Podiumcadeaukaart                         = "podiumcadeaukaart"
	Reiscadeau                                = "reiscadeau"
	Restaurantcadeau                          = "restaurantcadeau"
	Sportenfitcadeau                          = "sportenfitcadeau"
	Sustainablefashion                        = "sustainablefashion"
	Travelcheq                                = "travelcheq"
	Vvvgiftcard                               = "vvvgiftcard"
	Vvvdinercheque                            = "vvvdinercheque"
	Vvvlekkerweg                              = "vvvlekkerweg"
	Webshopgiftcard                           = "webshopgiftcard"
	Yourgift                                  = "yourgift"
)

// GiftCardIssuerStatus describes the status of a gift
// card issuer in your account.
type GiftCardIssuerStatus string

// Valid issuer statuses
const (
	PendingIssuer GiftCardIssuerStatus = "pending-issuer"
	EnabledIssuer                      = "enabled"
)

// GiftCardEnabled describes the response of a gift card
// issuer enable operation.
type GiftCardEnabled struct {
	Resource    string               `json:"resource,omitempty"`
	ID          string               `json:"id,omitempty"`
	Description string               `json:"description,omitempty"`
	Status      GiftCardIssuerStatus `json:"status,omitempty"`
	Links       GiftCardLinks        `json:"_links,omitempty"`
}

// GiftCardLinks are links embeded when a gift card is enabled.
type GiftCardLinks struct {
	Self          URL `json:"self,omitempty"`
	Documentation URL `json:"documentation,omitempty"`
}
