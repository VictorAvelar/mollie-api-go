package mollie

// GiftCardIssuer type describes issuers supported
// by mollie.
type GiftCardIssuer string

// Supported gift card issuers.
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

// GiftCardIssuerStatus describes the status of a gift
// card issuer in your account.
type GiftCardIssuerStatus string

// Valid issuer statuses.
const (
	PendingIssuer GiftCardIssuerStatus = "pending-issuer"
	EnabledIssuer GiftCardIssuerStatus = "enabled"
)

// GiftCardEnabled describes the response of a gift card
// issuer enable operation.
type GiftCardEnabled struct {
	Resource    string               `json:"resource,omitempty"`
	ID          GiftCardIssuer       `json:"id,omitempty"`
	Description string               `json:"description,omitempty"`
	Status      GiftCardIssuerStatus `json:"status,omitempty"`
	Links       GiftCardLinks        `json:"_links,omitempty"`
}

// GiftCardLinks are links embedded when a gift card is enabled.
type GiftCardLinks struct {
	Self          *URL `json:"self,omitempty"`
	Documentation *URL `json:"documentation,omitempty"`
}
