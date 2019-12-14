package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/testdata"
)

func TestPaymentsService_Get(t *testing.T) {
	setup()
	defer teardown()
	id := "tr_WDqYK6vllg"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetPaymentResponse))
	})

	opt := &PaymentOptions{
		Include: "details.qrCode",
	}

	res, err := tClient.Payments.Get(id, opt)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}
}

func TestPaymentsService_Create(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.GetPaymentResponse))
	})

	p := Payment{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "10.00",
		},
		Description: "Order #12345",
	}

	res, err := tClient.Payments.Create(p)
	if err != nil {
		t.Fatal(err)
	}

	if res.Description != p.Description {
		t.Errorf("mismatching info. want %v, got %v", p.Description, res.Description)
	}
}

func TestPaymentsService_Update(t *testing.T) {
	setup()
	defer teardown()
	id := "tr_7UhSN1zuXS"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "PATCH")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.UpdatePaymentResponse))
	})

	p := Payment{
		Description: "alter description",
	}

	res, err := tClient.Payments.Update(id, p)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}
}

func TestPaymentsService_Cancel(t *testing.T) {
	setup()
	defer teardown()
	id := "tr_WDqYK6vllg"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "DELETE")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.CancelPaymentResponse))
	})

	res, err := tClient.Payments.Cancel(id)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}
}

func TestPaymentsService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
	})

	res, err := tClient.Payments.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Count)
	}
}

func TestPaymentsService_ListWithOptions(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
	})

	options := &ListPaymentOptions{
		ProfileID: "pfl_QkEhN94Ba",
	}

	res, err := tClient.Payments.List(options)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Count)
	}
}

func TestPaymentsService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/payments/", errorHandler)

	p := Payment{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "10.00",
		},
		Description: "Order #12345",
	}

	_, cerr := tClient.Payments.Create(p)
	_, rerr := tClient.Payments.List(nil)
	_, uerr := tClient.Payments.Update("1212", p)
	_, derr := tClient.Payments.Cancel("1212")
	_, gerr := tClient.Payments.Get("1212", nil)

	tests := []error{cerr, rerr, uerr, derr, gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestPaymentsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/v2/payments/", errorHandler)

	p := Payment{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "10.00",
		},
		Description: "Order #12345",
	}

	_, cerr := tClient.Payments.Create(p)
	_, rerr := tClient.Payments.List(nil)
	_, uerr := tClient.Payments.Update("1212", p)
	_, derr := tClient.Payments.Cancel("1212")
	_, gerr := tClient.Payments.Get("1212", nil)

	tests := []error{cerr, rerr, uerr, derr, gerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Fail()
		}
	}
}

func TestPaymentsService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/payments/", encodingHandler)

	p := Payment{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "10.00",
		},
		Description: "Order #12345",
	}

	_, cerr := tClient.Payments.Create(p)
	_, rerr := tClient.Payments.List(nil)
	_, uerr := tClient.Payments.Update("1212", p)
	_, derr := tClient.Payments.Cancel("1212")
	_, gerr := tClient.Payments.Get("1212", nil)

	tests := []error{cerr, rerr, uerr, derr, gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestPaymentFailedResponseAvailable(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/payments/", unprocessableEntityHandler)

	_, err := tClient.Payments.Create(Payment{})

	if err == nil {
		t.Error("expected error and got nil")
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func encodingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{hello: [{},]}`))
}

func unprocessableEntityHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	_, _ = w.Write([]byte(testdata.CreateOrderPaymentResponseFailed))
}
