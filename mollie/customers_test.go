package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestCustomersService_Get(t *testing.T) {
	setup()
	defer teardown()

	id := "cst_kEn1PlbGa"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/customers/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetCustomerResponse))
	})

	c, err := tClient.Customers.Get(id)
	if err != nil {
		t.Error(err)
	}

	if c.ID != id {
		t.Errorf("unexpected response: got %s, want %s", c.ID, id)
	}
}

func TestCustomersService_Create(t *testing.T) {
	setup()
	defer teardown()

	id := "cst_8wmqcHMN4U"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/customers", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateCustomerResponse))
	})

	c, err := tClient.Customers.Create(Customer{Locale: German})
	if err != nil {
		t.Error(err)
	}

	if c.ID != id {
		t.Errorf("unexpected response: got %s, want %s", c.ID, id)
	}
}

func TestCustomersService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/customers", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListCustomersResponse))
	})

	c, err := tClient.Customers.List(&ListCustomersOptions{Limit: 10})
	if err != nil {
		t.Error(err)
	}

	if c.Count != 3 {
		t.Errorf("unexpected response: got %d, want 3", c.Count)
	}
}

func TestCustomersService_GetPayments(t *testing.T) {
	setup()
	defer teardown()
	id := "cst_kEn1PlbGa"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/customers/"+id+"/payments", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
	})

	p, err := tClient.Customers.GetPayments(id, &ListCustomersOptions{SequenceType: FirstSequence})
	if err != nil {
		t.Error(err)
	}

	if p.Count != 5 {
		t.Errorf("unexpected response: got %d, want 5", p.Count)
	}
}

func TestCustomersService_CreatePayment(t *testing.T) {
	setup()
	defer teardown()
	id := "cst_kEn1PlbGa"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/customers/"+id+"/payments", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetPaymentResponse))
	})

	p, err := tClient.Customers.CreatePayment(id, Payment{Mode: TestMode})
	if err != nil {
		t.Error(err)
	}

	if p.Resource != "payment" {
		t.Errorf("unexpected response: got %s, want payment", p.Resource)
	}
}

func TestCustomersService_Delete(t *testing.T) {
	setup()
	defer teardown()

	id := "cst_kEn1PlbGa"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/customers/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "DELETE")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	err := tClient.Customers.Delete(id)
	if err != nil {
		t.Error(err)
	}
}

func TestCustomersService_Update(t *testing.T) {
	setup()
	defer teardown()

	id := "cst_8wmqcHMN4U"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/customers/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "PATCH")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.UpdateCustomerResponse))
	})

	c, err := tClient.Customers.Update(id, Customer{
		Locale: French,
	})
	if err != nil {
		t.Error(err)
	}

	if c.ID != id {
		t.Errorf("unexpected response: got %s, want %s", c.ID, id)
	}
}

func TestCustomersService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/customers/", errorHandler)

	tests := forceCustomersErrors(true)

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func TestCustomersService_JsonDecodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/customers/", encodingHandler)

	tests := forceCustomersErrors(false)

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestCustomersService_HTTPRequestErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/customers/", errorHandler)

	tests := forceCustomersErrors(true)
	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func forceCustomersErrors(del bool) []error {
	id := "cst_kEn1PlbGa"
	_, lerr := tClient.Customers.List(nil)
	_, lperr := tClient.Customers.GetPayments(id, nil)
	_, cperr := tClient.Customers.CreatePayment(id, Payment{})
	_, gerr := tClient.Customers.Get(id)
	_, cerr := tClient.Customers.Create(Customer{})
	_, uerr := tClient.Customers.Update(id, Customer{})

	errs := []error{lerr, lperr, cperr, gerr, uerr, cerr}

	if del {
		derr := tClient.Customers.Delete(id)
		errs = append(errs, derr)
	}

	return errs
}
