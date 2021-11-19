package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestPartnerService_Get(t *testing.T) {
	setup()
	defer teardown()
	id := "org_1337"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/clients/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetPartnerClientResponse))
	})

	res, err := tClient.Partners.Get(nil, id, nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}
}

func TestPartnerService_GetWithOptions(t *testing.T) {
	setup()
	defer teardown()
	id := "org_1337"
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/clients/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if !strings.Contains(r.URL.String(), "?embed=organization") {
			w.WriteHeader(http.StatusBadRequest)
		}
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetPartnerClientResponse))
	})

	res, err := tClient.Partners.Get(nil, id, &GetPartnerClientOptions{Embed: "organization"})
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id {
		t.Errorf("mismatching info. want %v, got %v", id, res.ID)
	}
}

func TestPartnerService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/clients", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPartnerClientsResponse))
	})

	res, err := tClient.Partners.List(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.PartnerClients.Clients) != res.Count {
		t.Error("mismatching info.")
	}
}

func TestPartnerService_ListWithOptions(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/clients", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if !strings.Contains(r.URL.String(), "?year=2020") {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPartnerClientsResponse))
	})

	res, err := tClient.Partners.List(nil, &ListPartnerClientsOptions{
		Year: 2020,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(res.PartnerClients.Clients) != res.Count {
		t.Error("mismatching info.")
	}
}

func TestPartnerService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/clients/", errorHandler)

	_, gerr := tClient.Partners.Get(nil, "org_1337", nil)
	_, rerr := tClient.Partners.List(nil, nil)

	tests := []error{rerr, gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestPartnerService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/v2/clients/", errorHandler)

	_, gerr := tClient.Partners.Get(nil, "org_1337", nil)
	_, rerr := tClient.Partners.List(nil, nil)

	tests := []error{rerr, gerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Fail()
		}
	}
}

func TestPartnerService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/clients/", encodingHandler)

	_, gerr := tClient.Partners.Get(nil, "org_1337", nil)
	_, rerr := tClient.Partners.List(nil, nil)

	tests := []error{rerr, gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}
