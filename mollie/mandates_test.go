package mollie

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v2/testdata"
)

func TestMandatesService_Get(t *testing.T) {
	setup()
	defer teardown()

	mandateID := "mdt_h3gAaD5zP"
	customerID := "cst_4qqhO89gsT"
	tkn := "token_X12b31ggg23"

	_ = tClient.WithAuthenticationValue(tkn)

	tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/mandates/%s", customerID, mandateID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer "+tkn)
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetMandateResponse))
	})

	res, err := tClient.Mandates.Get(customerID, mandateID)

	if err != nil {
		t.Fatal(err)
	}

	if res.ID != mandateID {
		t.Errorf("mismatching mandate ID retrieved, want %s got %s", mandateID, res.ID)
	}
}

func TestMandatesService_Create(t *testing.T) {
	setup()
	defer teardown()

	customerID := "cst_4qqhO89gsT"
	tkn := "token_X12b31ggg23"

	_ = tClient.WithAuthenticationValue(tkn)
	tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/mandates", customerID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer "+tkn)
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateMandateResponse))
	})

	res, err := tClient.Mandates.Create(customerID, Mandate{})
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != "mdt_h3gAaD5zP" {
		t.Errorf("mismatching mandate ID, want %s got %s", "mdt_h3gAaD5zP", res.ID)
	}
}

func TestMandatesService_Revoke(t *testing.T) {
	setup()
	defer teardown()

	mandateID := "mdt_h3gAaD5zP"
	customerID := "cst_4qqhO89gsT"
	tkn := "token_X12b31ggg23"

	_ = tClient.WithAuthenticationValue(tkn)
	tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/mandates/%s", customerID, mandateID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer "+tkn)
		testMethod(t, r, "DELETE")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	err := tClient.Mandates.Revoke(customerID, mandateID)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMandatesService_List(t *testing.T) {
	setup()
	defer teardown()

	customerID := "cst_4qqhO89gsT"
	tkn := "token_X12b31ggg23"

	_ = tClient.WithAuthenticationValue(tkn)
	tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/mandates", customerID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer "+tkn)
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListMandatesResponse))
	})

	res, err := tClient.Mandates.List(customerID, &ListMandatesOptions{From: "mdt_h3gAaD5zP"})
	if err != nil {
		t.Fatal(err)
	}

	if res.Count != 2 {
		t.Errorf("unexpected response size, want 2 got %d", res.Count)
	}
}

func TestMandatesService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	customerID := "cst_4qqhO89gsT"
	tkn := "token_X12b31ggg23"

	_ = tClient.WithAuthenticationValue(tkn)

	tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/mandates/", customerID), encodingHandler)

	tests := forceMandatesErrors(false)

	for _, tt := range tests {
		if tt == nil {
			t.Errorf("got nil when expecting an error: %e", tt)
		}
	}
}

func TestMandatesService_NewHttpRequestErrorHandler(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/customers/cst_4qqhO89gsT/mandates/", errorHandler)

	tests := forceMandatesErrors(true)

	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func TestMandatesService_NewHTTPApiRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/customers/cst_4qqhO89gsT/mandates/", errorHandler)

	tests := forceMandatesErrors(true)

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func forceMandatesErrors(del bool) []error {
	mandateID := "mdt_h3gAaD5zP"
	customerID := "cst_4qqhO89gsT"
	_, gerr := tClient.Mandates.Get(customerID, mandateID)
	_, cerr := tClient.Mandates.Create(customerID, Mandate{})
	_, lerr := tClient.Mandates.List(customerID, nil)

	tests := []error{gerr, cerr, lerr}

	if del {
		rerr := tClient.Mandates.Revoke(customerID, mandateID)
		tests = append(tests, rerr)
	}

	return tests
}
