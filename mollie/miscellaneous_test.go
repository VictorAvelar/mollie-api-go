package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestMiscellaneousService_ApplePaymentSession(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/wallets/applepay/sessions", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.ApplePaySessionResponse))
	})

	apple, err := tClient.Miscellaneous.ApplePaymentSession(nil, &ApplePaymentSessionRequest{})
	if err != nil {
		t.Fail()
	}

	if apple == nil {
		t.Error("nil session received")
	}
}

func TestMiscellaneousService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/wallets", errorHandler)

	tests := forceMiscellaneousErrors()

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func TestMiscellaneousService_JsonDecodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/wallets/", encodingHandler)

	tests := forceMiscellaneousErrors()

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestMiscellaneousService_HTTPRequestErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/wallets/", errorHandler)

	tests := forceMiscellaneousErrors()
	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func forceMiscellaneousErrors() []error {
	_, aperr := tClient.Miscellaneous.ApplePaymentSession(nil, nil)
	return []error{aperr}
}
