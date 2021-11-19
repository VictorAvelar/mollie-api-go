package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestPermissionsService_Get(t *testing.T) {
	setup()
	defer teardown()
	id := string(PaymentsRead)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/permissions/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetPermissionsResponse))
	})

	p, err := tClient.Permissions.Get(nil, id)
	if err != nil {
		t.Error(err)
	}

	if p.ID != PaymentsRead {
		t.Errorf("the response content doesn't match expectations")
	}
}

func TestPermissionsService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/permissions", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPermissionsResponse))
	})

	pl, err := tClient.Permissions.List(nil)
	if err != nil {
		t.Error(err)
	}

	if pl.Count != 15 {
		t.Errorf("the response content doesn't match expectations")
	}
}

func TestPermissionsService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/permissions", errorHandler)

	_, gerr := tClient.Permissions.Get(nil, "payments.read")
	_, lerr := tClient.Permissions.List(nil)

	tests := []error{lerr, gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestPermissionsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/v2/permissions", errorHandler)

	_, gerr := tClient.Permissions.Get(nil, "payments.read")
	_, lerr := tClient.Permissions.List(nil)

	tests := []error{lerr, gerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Fail()
		}
	}
}

func TestPermissionsService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/permissions/", encodingHandler)

	_, gerr := tClient.Permissions.Get(nil, "payments.read")
	_, lerr := tClient.Permissions.List(nil)

	tests := []error{lerr, gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}
