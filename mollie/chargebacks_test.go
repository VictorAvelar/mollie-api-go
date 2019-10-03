package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/testdata"
)

func TestChargebacksService_Get(t *testing.T) {
	setup()
	defer teardown()
	paymentID := "tr_WDqYK6vllg"
	chargebackID := "chb_n9z0tp"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+paymentID+"/chargebacks/"+chargebackID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetChargebackResponse))
	})

	opt := &ChargebackOptions{
		Include: "details.qrCode",
	}

	res, err := tClient.Chargebacks.Get(paymentID, chargebackID, opt)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != chargebackID {
		t.Errorf("mismatching info. want %v, got %v", chargebackID, res.ID)
	}
}

func TestChargebacksService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/chargebacks", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
	})

	res, err := tClient.Chargebacks.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Count)
	}
}

func TestChargebacksService_ListForPayment(t *testing.T) {
	setup()
	defer teardown()
	paymentID := "tr_WDqYK6vllg"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+paymentID+"/chargebacks", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
	})

	res, err := tClient.Chargebacks.ListForPayment(paymentID, nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Count)
	}
}

func TestChargebacksService_ListWithOptions(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/chargebacks", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
	})

	options := &ListChargebackOptions{
		ProfileID: "pfl_QkEhN94Ba",
	}

	res, err := tClient.Chargebacks.List(options)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Count)
	}
}

func TestChargebacksService_ListForPaymentWithOptions(t *testing.T) {
	setup()
	defer teardown()
	paymentID := "tr_WDqYK6vllg"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+paymentID+"/chargebacks", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
	})

	options := &ListChargebackOptions{
		ProfileID: "pfl_QkEhN94Ba",
	}

	res, err := tClient.Chargebacks.ListForPayment(paymentID, options)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Count)
	}
}

func TestChargebacksService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/chargebacks/", errorHandler)

	_, rerr := tClient.Chargebacks.List(nil)
	_, lerr := tClient.Chargebacks.ListForPayment("1212", nil)
	_, gerr := tClient.Chargebacks.Get("1212", "3232", nil)

	tests := []error{rerr, gerr, lerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestChargebacksService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/v2/chargebacks/", errorHandler)

	_, rerr := tClient.Chargebacks.List(nil)
	_, lerr := tClient.Chargebacks.ListForPayment("1212", nil)
	_, gerr := tClient.Chargebacks.Get("1212", "3232", nil)

	tests := []error{rerr, gerr, lerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Fail()
		}
	}
}

func TestChargebacksService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/payments/1212/chargebacks/", encodingHandler)

	_, rerr := tClient.Chargebacks.ListForPayment("1212", nil)
	_, gerr := tClient.Chargebacks.Get("1212", "3232", nil)

	tests := []error{rerr, gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestChargebacksService_EncodingResponseErrors_List(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/chargebacks/", encodingHandler)

	_, rerr := tClient.Chargebacks.List(nil)

	tests := []error{rerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}
