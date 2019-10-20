package mollie

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/testdata"
)

func TestCapturesService_Get(t *testing.T) {
	setup()
	defer teardown()

	pID := "tr_WDqYK6vllg"
	cID := "cpt_4qqhO89gsT"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/captures/%s", pID, cID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetCaptureResponse))
	})

	capture, err := tClient.Captures.Get(pID, cID)
	if err != nil {
		t.Error(err)
	}

	if capture.ID != cID {
		t.Errorf("unexpected response: got %v, want %v", capture.ID, cID)
	}
}

func TestCapturesService_List(t *testing.T) {
	setup()
	defer teardown()

	pID := "tr_WDqYK6vllg"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/captures", pID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListCapturesResponse))
	})

	captures, err := tClient.Captures.List(pID)
	if err != nil {
		t.Error(err)
	}

	if len(captures.Embedded.Captures) < 1 {
		t.Errorf("empty response with %d captures, expecting 1", len(captures.Embedded.Captures))
	}
}

func TestCapturesService_HTTPRequestErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/payments/09d87sd8a9d/captures/", errorHandler)

	tests := forceCapturesErrors()

	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func TestCapturesService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/payments/09d87sd8a9d/captures/", errorHandler)

	tests := forceCapturesErrors()

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func TestCapturesService_JsonEncodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/payments/09d87sd8a9d/captures/", encodingHandler)

	tests := forceCapturesErrors()

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func forceCapturesErrors() []error {
	_, lerr := tClient.Captures.List("09d87sd8a9d")
	_, gerr := tClient.Captures.Get("09d87sd8a9d", "4asd5ad6")

	return []error{lerr, gerr}
}
