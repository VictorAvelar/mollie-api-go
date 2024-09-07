package mollie

import (
	"context"
	"encoding/json"
)

// Wallet describes the wallet types that Mollie supports.
type Wallet string

// Available wallet types.
const (
	ApplePayWallet Wallet = "applepay"
)

// WalletsService operates over the resources described
// in Mollie's wallets API endpoints section.
type WalletsService service

// ApplePaymentSession contains information about an Apple pay session.
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

// ApplePaymentSessionRequest contains the body parameters for requesting
// a valid PaymentSession from Apple.
type ApplePaymentSessionRequest struct {
	Domain        string `json:"domain,omitempty"`
	ValidationURL string `json:"validationUrl,omitempty"`
}

// ApplePaymentSession returns an Apple Payment Session object valid for one transaction.
//
// See: https://docs.mollie.com/reference/request-apple-pay-payment-session
func (ms *WalletsService) ApplePaymentSession(ctx context.Context, asr *ApplePaymentSessionRequest) (
	res *Response,
	aps *ApplePaymentSession,
	err error,
) {
	u := "v2/wallets/applepay/sessions"

	res, err = ms.client.post(ctx, u, asr, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &aps); err != nil {
		return
	}

	return
}
