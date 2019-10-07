package mollie

import (
	"github.com/VictorAvelar/mollie-api-go/testdata"
	"net/http"
	"testing"
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
