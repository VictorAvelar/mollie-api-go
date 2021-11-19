package mollie

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestSettlementsService_Get(t *testing.T) {
	setup()
	defer teardown()

	cID := "stl_jDk30akdN"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc(fmt.Sprintf("/v2/settlements/%s", cID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetSettlementsResponse))
	})

	settlement, err := tClient.Settlements.Get(context.TODO(), cID)
	if err != nil {
		t.Error(err)
	}

	if settlement.ID != cID {
		t.Errorf("unexpected response: got: %v, want %v", settlement.ID, cID)
	}
}

func TestSettlementsService_Next(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/settlements/next", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetNextSettlementResponse))
	})

	settlement, err := tClient.Settlements.Next(context.TODO())
	if err != nil {
		t.Error(err)
	}

	if settlement.ID != "next" {
		t.Errorf("unexpected response: got: %v, want 'next'", settlement.ID)
	}
}

func TestSettlementsService_Open(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/settlements/open", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetOpenSettlementResponse))
	})

	settlement, err := tClient.Settlements.Open(context.TODO())
	if err != nil {
		t.Error(err)
	}

	if settlement.ID != "open" {
		t.Errorf("unexpected response: got: %v, want 'open'", settlement.ID)
	}
}

func TestSettlementsService_List(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/settlements", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListSettlementsResponse))
	})

	settlement, err := tClient.Settlements.List(context.TODO(), &SettlementsListOptions{
		Limit: 10,
	},
	)
	if err != nil {
		t.Error(err)
	}

	if settlement.Count != 5 {
		t.Errorf("unexpected response: got: %v, want 5", settlement.Count)
	}
}

func TestSettlementsService_GetPayments(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/settlements/stl_jDk30akdN/payments", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
	})

	list, err := tClient.Settlements.GetPayments(context.TODO(), "stl_jDk30akdN", nil)
	if err != nil {
		t.Error(err)
	}

	if list.Count != 5 {
		t.Errorf("unexpected response: got: %v, want 5", list.Count)
	}
}

func TestSettlementsService_GetCaptures(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/settlements/stl_jDk30akdN/captures", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListCapturesResponse))
	})

	list, err := tClient.Settlements.GetCaptures(context.TODO(), "stl_jDk30akdN", nil)
	if err != nil {
		t.Error(err)
	}

	if list.Count != 1 {
		t.Errorf("unexpected response: got: %v, want 1", list.Count)
	}
}

func TestSettlementsService_GetChargebacks(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/settlements/stl_jDk30akdN/chargebacks", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
	})

	list, err := tClient.Settlements.GetChargebacks(context.TODO(), "stl_jDk30akdN", nil)
	if err != nil {
		t.Error(err)
	}

	if list.Count != 3 {
		t.Errorf("unexpected response: got: %v, want 3", list.Count)
	}
}

func TestSettlementsService_GetRefunds(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/settlements/stl_jDk30akdN/refunds", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListOrderRefundResponse))
	})

	list, err := tClient.Settlements.GetRefunds(context.TODO(), "stl_jDk30akdN", nil)
	if err != nil {
		t.Error(err)
	}

	if list.Count != 1 {
		t.Errorf("unexpected response: got: %v, want 1", list.Count)
	}
}

func TestSettlementsService_HTTPRequestErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/settlements/", errorHandler)

	tests := forceSettlementsErrors()

	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func TestSettlementsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/settlements/", errorHandler)

	tests := forceSettlementsErrors()

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func TestSettlementsService_JsonEncodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/settlements/", encodingHandler)

	tests := forceSettlementsErrors()

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func forceSettlementsErrors() []error {
	_, gerr := tClient.Settlements.Get(context.TODO(), "10298372891")
	_, nerr := tClient.Settlements.Next(context.TODO())
	_, oerr := tClient.Settlements.Open(context.TODO())
	_, lerr := tClient.Settlements.List(context.TODO(), nil)
	_, perr := tClient.Settlements.GetPayments(context.TODO(), "stl_jDk30akdN", nil)
	_, cerr := tClient.Settlements.GetCaptures(context.TODO(), "stl_jDk30akdN", nil)
	_, cberr := tClient.Settlements.GetChargebacks(context.TODO(), "stl_jDk30akdN", nil)
	_, rerr := tClient.Settlements.GetRefunds(context.TODO(), "stl_jDk30akdN", nil)

	return []error{gerr, nerr, oerr, lerr, perr, cerr, cberr, rerr}
}
