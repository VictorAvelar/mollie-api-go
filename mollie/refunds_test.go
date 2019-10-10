package mollie

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/testdata"
)

func TestRefundsService_Get(t *testing.T) {
	setup()
	defer teardown()

	paymentID := "tr_WDqYK6vllg"
	refundID := "re_4qqhO89gsT"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+paymentID+"/refunds/"+refundID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetRefundResponse))
	})

	opt := &RefundOptions{
		Embed:    EmbedPayment,
		TestMode: true,
	}

	res, err := tClient.Refunds.Get(paymentID, refundID, opt)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != refundID {
		t.Errorf("mismatching info. want %v got %v", refundID, res.ID)
	}
}

func TestRefundsService_Create(t *testing.T) {
	setup()
	defer teardown()

	paymentID := "tr_WDqYK6vllg"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+paymentID+"/refunds", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetRefundResponse))
	})

	opt := &RefundOptions{
		TestMode: true,
	}

	refund := &Refund{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "20",
		},
		Description: "Order #33",
	}

	res, err := tClient.Refunds.Create(paymentID, *refund, opt)
	if err != nil {
		t.Fatal(err)
	}

	if res.Description != refund.Description {
		t.Errorf("mismatching info. want %s got %s", refund.Description, res.Description)
	}
}

func TestRefundsService_CreateInvalidParams(t *testing.T) {
	setup()
	defer teardown()

	paymentID := "tr_WDqYK6vllg"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+paymentID+"/refunds", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetRefundResponse))
	})

	opt := &RefundOptions{
		TestMode: true,
	}

	refcurr := Refund{
		Amount: &Amount{
			Currency: "IDR",
		},
	}

	refval := Refund{
		Amount: &Amount{
			Value: "100000",
		},
	}

	_, errcurr := tClient.Refunds.Create(paymentID, refcurr, opt)
	_, errval := tClient.Refunds.Create(paymentID, refval, opt)

	tests := []error{errcurr, errval}

	for _, test := range tests {
		if test == nil {
			t.Fail()
		} else if !strings.Contains(test.Error(), "parameter required") {
			t.Errorf("unexpected error %v", test)
		}
	}
}

func TestRefundsService_Cancel(t *testing.T) {
	setup()
	defer teardown()

	paymentID := "tr_7UhSN1zuXS"
	refundID := "re_4qqhO89gsT"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+paymentID+"/refunds/"+refundID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodDelete)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	opt := &RefundOptions{
		TestMode: true,
	}

	err := tClient.Refunds.Cancel(paymentID, refundID, opt)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRefundsService_List(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/refunds", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		_, _ = w.Write([]byte(testdata.GetRefundListResponse))
	})

	opt := &ListRefundOptions{
		ProfileID: "pfl_3RkSN1zuPE",
		TestMode:  true,
	}

	res, err := tClient.Refunds.ListRefund(opt)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want >0 got %v", res.Count)
	}
}

func TestServiceRefunds_ListPayment(t *testing.T) {
	setup()
	defer teardown()

	paymentID := "tr_7UhSN1zuXS"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/payments/"+paymentID+"/refunds", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetRefundListResponse))
	})

	opt := &ListRefundOptions{
		From:  "re_4qqhO89gsT",
		Limit: "100",
	}

	res, err := tClient.Refunds.ListRefundPayment(paymentID, opt)
	if err != nil {
		t.Fatal(err)
	}

	if res.Count == 0 {
		t.Errorf("mismatching info. want >0 got %v", res.Count)
	}
}

func TestRefundsService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/refunds/", errorHandler)

	refund := Refund{
		Amount: &Amount{
			Currency: "IDR",
			Value:    "100000",
		},
		Description: "Order #12345",
	}

	_, cerr := tClient.Refunds.Create("123132", refund, nil)
	_, rerr1 := tClient.Refunds.ListRefund(nil)
	_, rerr2 := tClient.Refunds.ListRefundPayment("123132", nil)
	derr := tClient.Refunds.Cancel("12123", "123123", nil)
	_, rerr3 := tClient.Refunds.Get("1212", "12312", nil)

	tests := []error{cerr, rerr1, rerr2, derr, rerr3}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestRefundsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/v2/refunds/", errorHandler)

	refund := Refund{
		Amount: &Amount{
			Currency: "IDR",
			Value:    "100000",
		},
		Description: "Order #12345",
	}

	_, cerr := tClient.Refunds.Create("123132", refund, nil)
	_, rerr1 := tClient.Refunds.ListRefund(nil)
	_, rerr2 := tClient.Refunds.ListRefundPayment("123132", nil)
	derr := tClient.Refunds.Cancel("12123", "123123", nil)
	_, rerr3 := tClient.Refunds.Get("1212", "12312", nil)

	tests := []error{cerr, rerr1, rerr2, derr, rerr3}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Fail()
		}
	}
}

func TestRefundsService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()

	paymentID := "tr_7UhSN1zuXS"
	refundID := "re_4qqhO89gsT"

	tMux.HandleFunc("/v2/payments/"+paymentID+"/refunds", encodingHandler)
	tMux.HandleFunc("/v2/refunds/", encodingHandler)
	tMux.HandleFunc("/v2/payments/"+paymentID+"/refunds/"+refundID, encodingHandler)

	refund := Refund{
		Amount: &Amount{
			Currency: "IDR",
			Value:    "100000",
		},
		Description: "Order #12345",
	}

	_, cerr := tClient.Refunds.Create(paymentID, refund, nil)
	_, rerr1 := tClient.Refunds.ListRefund(nil)
	_, rerr2 := tClient.Refunds.ListRefundPayment(paymentID, nil)
	_, rerr3 := tClient.Refunds.Get(paymentID, refundID, nil)

	tests := []error{cerr, rerr1, rerr2, rerr3}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}
