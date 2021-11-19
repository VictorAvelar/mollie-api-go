package mollie

import (
	"context"
	"encoding/json"
	"net/http"
)

// MiscellaneousService operates over the resources described
// in Mollie's miscellaneous API endpoints section
type MiscellaneousService service

// ApplePaymentSession contains information about an Apple pay session
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
// See: https://docs.mollie.com/reference/v2/wallets-api/request-apple-pay-payment-session
func (ms *MiscellaneousService) ApplePaymentSession(ctx context.Context, asr *ApplePaymentSessionRequest) (aps *ApplePaymentSession, err error) {
	u := "v2/wallets/applepay/sessions"
	req, err := ms.client.NewAPIRequest(ctx, http.MethodPost, u, asr)
	if err != nil {
		return
	}

	res, err := ms.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &aps); err != nil {
		return
	}
	return
}
