package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v2/testdata"
)

func TestSubscriptionsService_Get(t *testing.T) {
	setup()
	defer teardown()

	cID := "cst_stTC2WHAuS"
	sID := "sub_rVKGtNd6s3"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions/%s", cID, sID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
	})

	sub, err := tClient.Subscriptions.Get(cID, sID)
	if err != nil {
		t.Error(err)
	}

	if sub.ID != sID {
		t.Errorf("unexpected response: got %v, want %v", sub.ID, sID)
	}
}

func TestSubscriptionsService_Create(t *testing.T) {
	setup()
	defer teardown()

	cID := "cst_stTC2WHAuS"
	sID := "sub_rVKGtNd6s3"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions", cID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
	})

	s := Subscription{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "11.99",
		},
		Interval: "12 months",
	}

	sub, err := tClient.Subscriptions.Create(cID, &s)
	if err != nil {
		t.Error(err)
	}

	if sub.ID != sID {
		t.Errorf("unexpected response: got %v, want %v", sub.ID, sID)
	}
}

func TestSubscriptionsService_Create_AccessTokens(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("access_token")

	cID := "cst_stTC2WHAuS"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions", cID)

	tMux.HandleFunc(u, func(rw http.ResponseWriter, r *http.Request) {
		var ship Shipment
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&ship); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(ship)
	})

	s := Subscription{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "11.99",
		},
		Interval: "12 months",
	}

	sub, err := tClient.Subscriptions.Create(cID, &s)
	if err != nil {
		t.Error(err)
	}

	if sub.TestMode != true {
		t.Fatal("testmode flag is not set for access tokens")
	}
}

func TestSubscriptionsService_Update(t *testing.T) {
	setup()
	defer teardown()

	cID := "cst_stTC2WHAuS"
	sID := "sub_rVKGtNd6s3"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions/%s", cID, sID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPatch)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
	})

	s := Subscription{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "11.99",
		},
		Interval: "12 months",
	}

	sub, err := tClient.Subscriptions.Update(cID, sID, &s)
	if err != nil {
		t.Error(err)
	}

	if sub.ID != sID {
		t.Errorf("unexpected response: got %v, want %v", sub.ID, sID)
	}
}

func TestSubscriptionsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	cID := "cst_stTC2WHAuS"
	sID := "sub_rVKGtNd6s3"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions/%s", cID, sID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodDelete)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
	})

	sub, err := tClient.Subscriptions.Delete(cID, sID)
	if err != nil {
		t.Error(err)
	}

	if sub.ID != sID {
		t.Errorf("unexpected response: got %v, want %v", sub.ID, sID)
	}
}

func TestSubscriptionsService_List(t *testing.T) {
	setup()
	defer teardown()

	cID := "cst_stTC2WHAuS"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions", cID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListAllSubscriptionsResponse))
	})

	sub, err := tClient.Subscriptions.List(cID, &SubscriptionListOptions{Limit: 6})
	if err != nil {
		t.Error(err)
	}

	if sub.Count != 3 {
		t.Errorf("unexpected response: got %v, want 3", sub.Count)
	}
}

func TestSubscriptionsService_All(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListAllSubscriptionsResponse))
	})

	sub, err := tClient.Subscriptions.All(&SubscriptionListOptions{Limit: 6})
	if err != nil {
		t.Error(err)
	}

	if sub.Count != 3 {
		t.Errorf("unexpected response: got %v, want 3", sub.Count)
	}
}

func TestSubscriptionsService_GetPayments(t *testing.T) {
	setup()
	defer teardown()
	cID := "cst_stTC2WHAuS"
	sID := "sub_rVKGtNd6s3"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions/%s/payments", cID, sID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
	})
	sub, err := tClient.Subscriptions.GetPayments(cID, sID, &SubscriptionListOptions{Limit: 6})
	if err != nil {
		t.Error(err)
	}
	if sub.Count != 5 {
		t.Errorf("unexpected response: got %v, want 5", sub.Count)
	}
}

func TestSubscriptionsService_HTTPErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/", errorHandler)
	tests := forceSubscriptionsErrors()

	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func TestSubscriptionsService_JsonEncodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/", encodingHandler)

	tests := forceSubscriptionsErrors()

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestSubscriptionsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/", errorHandler)

	tests := forceSubscriptionsErrors()

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func forceSubscriptionsErrors() []error {
	_, gerr := tClient.Subscriptions.Get("6s7d89a", "sd9af8ds")
	_, cerr := tClient.Subscriptions.Create("6s7d89a", &Subscription{Mode: LiveMode})
	_, uerr := tClient.Subscriptions.Update("6s7d89a", "0re9w0e", &Subscription{Mode: LiveMode})
	_, derr := tClient.Subscriptions.Delete("6s7d89a", "0re9w0e")
	_, lerr := tClient.Subscriptions.List("6s7d89a", nil)
	_, aerr := tClient.Subscriptions.All(nil)
	_, perr := tClient.Subscriptions.GetPayments("6s7d89a", "0re9w0e", nil)

	return []error{gerr, cerr, uerr, derr, lerr, aerr, perr}
}
