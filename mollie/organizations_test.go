package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v2/testdata"
)

func TestOrganizationsService_Get(t *testing.T) {
	setup()
	defer teardown()
	id := "org_12345678"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/organizations/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetOrganizationResponse))
	})

	res, err := tClient.Organizations.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}
}

func TestOrganizationsService_GetCurrent(t *testing.T) {
	setup()
	defer teardown()
	id := "org_12345678"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/organizations/me", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetCurrentOrganizationResponse))
	})

	res, err := tClient.Organizations.GetCurrent()
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}
}

func TestOrganizationsService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/organizations/", errorHandler)

	_, gerr := tClient.Organizations.Get("org_12345678")
	_, gcerr := tClient.Organizations.GetCurrent()

	tests := []error{gerr, gcerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestOrganizationsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/v2/organizations/", errorHandler)

	_, gerr := tClient.Organizations.Get("org_12345678")
	_, gcerr := tClient.Organizations.GetCurrent()

	tests := []error{gerr, gcerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Fail()
		}
	}
}

func TestOrganizationsService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/organizations/", encodingHandler)

	_, gerr := tClient.Organizations.Get("org_12345678")
	_, gcerr := tClient.Organizations.GetCurrent()

	tests := []error{gerr, gcerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}
