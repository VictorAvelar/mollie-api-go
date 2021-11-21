package mollie

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
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

	res, err := tClient.Payments.Get(context.TODO(), id, opt)
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

	res, err := tClient.Payments.Create(context.TODO(), p, nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.Description != p.Description {
		t.Errorf("mismatching info. want %v, got %v", p.Description, res.Description)
	}
}

func TestPaymentsService_Create_AccessTokens(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("access_token")

	tMux.HandleFunc("/v2/payments", func(rw http.ResponseWriter, r *http.Request) {
		var pay Payment
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&pay); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(pay)
	})

	p := Payment{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "10.00",
		},
		Description: "Order #12345",
	}

	payment, err := tClient.Payments.Create(context.TODO(), p, nil)
	if err != nil {
		t.Fatal(err)
	}

	if payment.TestMode != true {
		t.Error("testmode flag is not set for access tokens")
	}
}

func TestPaymentsService_Create_PaymentMethodFields(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("access_token")

	tMux.HandleFunc("/v2/payments", func(rw http.ResponseWriter, r *http.Request) {
		var pay Payment
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&pay); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if pay.Method != "ideal" && pay.Issuer == "" {
			t.Error("request payload is not properly encoded")
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(pay)
	})

	p := Payment{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "10.00",
		},
		Description: "Order #12345",
		Method:      IDeal,
		Issuer:      "ideal_INGBNL2A",
	}

	payment, err := tClient.Payments.Create(context.TODO(), p, nil)
	if err != nil {
		t.Fatal(err)
	}

	if payment.TestMode != true {
		t.Error("testmode flag is not set for access tokens")
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

	res, err := tClient.Payments.Update(context.TODO(), id, p)
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

	res, err := tClient.Payments.Cancel(context.TODO(), id)
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

	res, err := tClient.Payments.List(context.TODO(), nil)
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

	res, err := tClient.Payments.List(context.TODO(), options)
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

	_, cerr := tClient.Payments.Create(context.TODO(), p, nil)
	_, rerr := tClient.Payments.List(context.TODO(), nil)
	_, uerr := tClient.Payments.Update(context.TODO(), "1212", p)
	_, derr := tClient.Payments.Cancel(context.TODO(), "1212")
	_, gerr := tClient.Payments.Get(context.TODO(), "1212", nil)

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

	_, cerr := tClient.Payments.Create(context.TODO(), p, nil)
	_, rerr := tClient.Payments.List(context.TODO(), nil)
	_, uerr := tClient.Payments.Update(context.TODO(), "1212", p)
	_, derr := tClient.Payments.Cancel(context.TODO(), "1212")
	_, gerr := tClient.Payments.Get(context.TODO(), "1212", nil)

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

	_, cerr := tClient.Payments.Create(context.TODO(), p, nil)
	_, rerr := tClient.Payments.List(context.TODO(), nil)
	_, uerr := tClient.Payments.Update(context.TODO(), "1212", p)
	_, derr := tClient.Payments.Cancel(context.TODO(), "1212")
	_, gerr := tClient.Payments.Get(context.TODO(), "1212", nil)

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

	_, err := tClient.Payments.Create(context.TODO(), Payment{}, nil)

	if err == nil {
		t.Error("expected error and got nil")
	}
}

func unprocessableEntityHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	_, _ = w.Write([]byte(testdata.CreateOrderPaymentResponseFailed))
}
