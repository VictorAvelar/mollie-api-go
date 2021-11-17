package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestMethodsService_ListWithQueryOptionsAmountCurrency(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/methods", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		if r.URL.RawQuery != "amount%5Bcurrency%5D=USD&amount%5Bvalue%5D=100.00" {
			t.Fatal(r.URL.RawQuery)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListMethodsResponse))
	})

	opts := &MethodsOptions{
		AmountCurrency: "USD",
		AmountValue:    "100.00",
	}

	res, err := tClient.Methods.List(opts)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count <= 0 {
		t.Error("expecting methods and got 0")
	}
}

func TestMethodsService_ListWithQueryOptionsAll(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/methods", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		if r.URL.RawQuery != "amount%5Bcurrency%5D=USD&amount%5Bvalue%5D=100.00&billingCountry=DE&includeWallets=applepay&locale=de_DE&resource=orders&sequenceType=first" {
			t.Fatal(r.URL.RawQuery)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListMethodsResponse))
	})

	opts := &MethodsOptions{
		AmountCurrency: "USD",
		AmountValue:    "100.00",
		Resource:       "orders",
		SequenceType:   FirstSequence,
		Locale:         German,
		BillingCountry: "DE",
		IncludeWallets: "applepay",
	}

	res, err := tClient.Methods.List(opts)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count <= 0 {
		t.Error("expecting methods and got 0")
	}
}

func TestMethodsService_Get(t *testing.T) {
	setup()
	defer teardown()
	id := "ideal"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/methods/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetMethodResponse))
	})

	res, err := tClient.Methods.Get(id, nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}

	// 404 when not found
	mo := &MethodsOptions{
		Locale: English,
	}
	_, err2 := tClient.Methods.Get("sofort", mo)
	if err2 == nil {
		t.Fatal(err)
	} else if !strings.Contains(err2.Error(), "Not Found") {
		t.Fail()
	}
}

func TestMethodsService_All(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/methods/all", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListMethodsResponse))
	})

	res, err := tClient.Methods.All(&MethodsOptions{
		Locale: Dutch,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 13, res.Count)
	}
}

func TestMethodsService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/methods", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListMethodsResponse))
	})

	res, err := tClient.Methods.List(&MethodsOptions{
		SequenceType: FirstSequence,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 13, res.Count)
	}
}

func TestMethodsService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/methods/", errorHandler)

	_, lerr := tClient.Methods.List(nil)
	_, aerr := tClient.Methods.All(nil)
	_, gerr := tClient.Methods.Get("ideal", nil)

	tests := []error{lerr, aerr, gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestMethodsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/v2/methods/", errorHandler)

	_, rerr := tClient.Methods.List(nil)
	_, derr := tClient.Methods.All(nil)
	_, gerr := tClient.Methods.Get("1212", nil)

	tests := []error{rerr, derr, gerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Fail()
		}
	}
}

func TestMethodsService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/methods/", encodingHandler)

	_, cerr := tClient.Methods.All(nil)
	_, rerr := tClient.Methods.List(nil)
	_, uerr := tClient.Methods.Get("1212", nil)

	tests := []error{cerr, rerr, uerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}
