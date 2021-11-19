package mollie

import (
	"context"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestPaymentLinkService_Get(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payment-links/test_id", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetPaymentLinkResponse))
	})

	v, err := tClient.PaymentLinks.Get(context.TODO(), "test_id")
	if err != nil {
		t.Fatal(err)
	}

	if v.Amount.Value != "24.95" {
		t.Error("unexpected response")
	}
}

func TestPaymentLinkService_Create(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payment-links", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.CreatePaymentLinkResponse))
	})

	v, err := tClient.PaymentLinks.Create(
		context.TODO(),
		PaymentLink{
			Description: "payment_test_desc",
		},
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if v.ID != "pl_4Y0eZitmBnQ6IDoMqZQKh" {
		t.Fail()
	}
}

func TestPaymentLinkService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payment-links", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPaymentLinksResponse))
	})

	v, err := tClient.PaymentLinks.List(
		context.TODO(),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if v.Count != 1 {
		t.Fail()
	}
}
