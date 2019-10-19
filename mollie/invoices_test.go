package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/testdata"
)

func TestInvoicesService_Get(t *testing.T) {
	setup()
	defer teardown()
	id := "inv_xBEbP9rvAq"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/invoices/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetInvoiceResponse))
	})

	res, err := tClient.Invoices.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}
}

func TestInvoicesService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/invoices", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListInvoicesResponse))
	})

	res, err := tClient.Invoices.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Count)
	}
}

func TestInvoicesService_ListWithOptions(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/invoices", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListInvoicesResponse))
	})

	options := &ListInvoiceOptions{
		Reference: "2016.10000",
	}

	res, err := tClient.Invoices.List(options)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Count)
	}

	if res.Embedded.Invoices[0].Reference != "2016.10000" {
		t.Errorf("mismatching info. want %v, got %v", 0, res.Embedded.Invoices[0].Reference)
	}
}

func TestInvoicesService_HTTPErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/invoices/", errorHandler)

	tests := forceInvoicesErrors()
	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func TestInvoicesService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/invoices/", errorHandler)

	tests := forceInvoicesErrors()
	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func TestInvoicesService_JsonEncodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/invoices/", encodingHandler)

	tests := forceInvoicesErrors()
	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func forceInvoicesErrors() []error {
	_, gerr := tClient.Invoices.Get("g6d7f8gds76dfs78")
	_, lerr := tClient.Invoices.List(nil)

	return []error{gerr, lerr}
}
