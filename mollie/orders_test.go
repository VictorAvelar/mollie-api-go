package mollie

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/testdata"
)

func TestOrdersService_Get(t *testing.T) {
	setup()
	defer teardown()

	ordID := "ord_kEn1PlbGa"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+ordID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetOrderResponse))
	})

	opt := &OrderOptions{
		Embed: []EmbedValue{
			EmbedPayment,
			EmbedRefund,
		},
	}

	res, err := tClient.Orders.Get(ordID, opt)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != ordID {
		t.Errorf("mismatching info. want %v got %v", ordID, res.ID)
	}
}

func TestOrdersService_Create(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateOrderResponse))
	})

	var order Orders
	err := json.Unmarshal([]byte(testdata.CreateOrderRequest), &order)
	if err != nil {
		t.Error(err)
	}

	opt := &OrderOptions{
		ProfileID: "pfl_3RkSN1zuPE",
	}

	res, err := tClient.Orders.Create(order, opt)
	if err != nil {
		t.Error(err)
	}

	if res.ID != "ord_pbjz8x" {
		t.Errorf("mismatching info. want %v got %v", "ord_pbjz8x", res.ID)
	}
}

func TestOrdersService_Update(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_kEn1PlbGa"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPatch)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.UpdateOrderResponse))
	})

	var order Orders
	if err := json.Unmarshal([]byte(testdata.UpdateOrderRequest), &order); err != nil {
		t.Error(err)
	}

	res, err := tClient.Orders.Update(orderID, order)
	if err != nil {
		t.Error(err)
	}

	if res.ID != orderID {
		t.Errorf("mismatching info. want %v got %v", orderID, res.ID)
	}
}

func TestOrdersService_Cancel(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_8wmqcHMN4U"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodDelete)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.CancelOrderResponse))
	})

	res, err := tClient.Orders.Cancel(orderID)
	if err != nil {
		t.Error(err)
	}

	if res.ID != orderID {
		t.Errorf("mismatching info. want %v got %v", orderID, res.ID)
	}
}

func TestOrdersService_List(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.OrderListResponse))
	})

	opt := &OrderListOptions{
		ProfileID: "pfl_3RkSN1zuPE",
	}

	res, err := tClient.Orders.List(opt)
	if err != nil {
		t.Error(err)
	}

	if res.Count != 3 {
		t.Errorf("mismatching info. want %v got %v", 3, res.Count)
	}
}

func TestOrdersService_UpdateOrderline(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_pbjz8x"
	orderlineID := "odl_dgtxyl"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID+"/lines/"+orderlineID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPatch)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.UpdateOrderlineResponse))
	})

	var orderline OrderLines
	if err := json.Unmarshal([]byte(testdata.UpdateOrderlineRequest), &orderline); err != nil {
		t.Error(err)
	}

	res, err := tClient.Orders.UpdateOrderline(orderID, orderlineID, orderline)
	if err != nil {
		t.Error(err)
	}

	if res.ID != orderID {
		t.Errorf("mismatching info. want %v got %v", orderID, res.ID)
	}
}

func TestOrdersService_CancelOrderLines(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_8wmqcHMN4U"
	isError := false

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID+"/lines", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodDelete)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		decoder := json.NewDecoder(r.Body)

		var ordLines Orders
		err := decoder.Decode(&ordLines)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		for _, el := range ordLines.Lines {
			if (*el).ID == "odl_dgtwkn" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				_, _ = w.Write([]byte(testdata.CancelOrderLinesResponseCancelReject))

				isError = true

				break
			}

			if (*el).Amount == nil || (*el).Amount.Value == "0.00" || (*el).Amount.Currency == "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				_, _ = w.Write([]byte(testdata.CancelOrderLinesResponseAmountRequired))

				isError = true

				break
			}
		}

		if !isError {
			w.WriteHeader(http.StatusNoContent)
		}

	})

	var ordLines Orders
	if err := json.Unmarshal([]byte(testdata.CancelOrderLinesRequest), &ordLines); err != nil {
		t.Error(err)
	}

	res, err := tClient.Orders.CancelOrderLine(orderID, &ordLines)
	if res == nil || err != nil {
		t.Fail()
	}
	t.Log(res)

	if err := json.Unmarshal([]byte(testdata.CancelOrderLinesRejectRequest), &ordLines); err != nil {
		t.Error(err)
	}

	res, err = tClient.Orders.CancelOrderLine(orderID, &ordLines)
	if res == nil || err != nil {
		t.Fail()
	}
	t.Log(res)
}

func TestOrdersService_CreatePayment(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_8wmqcHMN4U"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID+"/payments", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateOrderPaymentResponse))
	})

	var ordPay OrderPayment
	if err := json.Unmarshal([]byte(testdata.CreateOrderPaymentRequest), &ordPay); err != nil {
		t.Error(err)
	}

	res, errResp, err := tClient.Orders.CreateOrderPayment(orderID, &ordPay)
	if err != nil {
		t.Error(errResp)
	}

	if res.ID != "tr_WDqYK6vllg" {
		t.Errorf("mismatching info. want %v got %v", "tr_WDqYK6vllg", res.ID)
	}
}

func TestOrdersService_CreatePaymentFailed(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_stTC2WHAuS"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID+"/payments", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(testdata.CreateOrderPaymentResponseFailed))
	})

	var ordPay OrderPayment
	if err := json.Unmarshal([]byte(testdata.CreateOrderPaymentRequest), &ordPay); err != nil {
		t.Error(err)
	}

	_, errResp, err := tClient.Orders.CreateOrderPayment(orderID, &ordPay)
	if errResp == nil || err != nil {
		t.Fail()
	}

	t.Log(errResp)
}

func TestOrdersService_CreateOrderRefund(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_stTC2WHAuS"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID+"/refunds", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateOrderRefundResponse))
	})

	var order Orders
	if err := json.Unmarshal([]byte(testdata.CreateOrderRefundRequest), &order); err != nil {
		t.Error(err)
	}

	res, _, err := tClient.Orders.CreateOrderRefund(orderID, &order)
	if err != nil {
		t.Error(err)
	}

	if res.ID != "re_4qqhO89gsT" {
		t.Errorf("mismatching info. want %v got %v", "re_4qqhO89gsT", res.ID)
	}
}

func TestOrdersService_CreateOrderRefundFailed(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_stTC2WHAuS"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID+"/refunds", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(testdata.CreateOrderRefundResponseFailed))
	})

	var order Orders
	if err := json.Unmarshal([]byte(testdata.CreateOrderRefundRequest), &order); err != nil {
		t.Error(err)
	}

	_, errResp, err := tClient.Orders.CreateOrderRefund(orderID, &order)
	if errResp == nil || err != nil {
		t.Fail()
	}

	t.Log(errResp)
}

func TestOrdersService_ListOrderRefund(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_pbjz8x"

	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/orders/"+orderID+"/refunds", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListOrderRefundResponse))
	})

	opt := &OrderListRefundOptions{
		Embed: EmbedPayment,
	}

	res, err := tClient.Orders.ListOrderRefunds(orderID, opt)
	if err != nil {
		t.Error(err)
	}

	if res.Count != 1 {
		t.Errorf("mismatching info. want %v got %v", 1, res.Count)
	}
}

func TestOrdersService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()

	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/orders/", errorHandler)

	tests := forceOrdersErrors(true)

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func TestOrdersService_JsonDecodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/orders/", encodingHandler)

	tests := forceOrdersErrors(false)

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestOrdersService_HTTPRequestErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/orders/", errorHandler)

	tests := forceOrdersErrors(true)
	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func forceOrdersErrors(del bool) []error {
	id := "ord_8wmqcHMN4U"

	_, cerr := tClient.Orders.Create(Orders{}, nil)
	_, _, coperr := tClient.Orders.CreateOrderPayment(id, nil)
	_, _, corerr := tClient.Orders.CreateOrderRefund(id, nil)
	_, rerr := tClient.Orders.Get(id, nil)
	_, lerr := tClient.Orders.List(nil)
	_, lorerr := tClient.Orders.ListOrderRefunds(id, nil)
	_, uerr := tClient.Orders.Update(id, Orders{})
	_, uolerr := tClient.Orders.UpdateOrderline(id, "", OrderLines{})

	errs := []error{cerr, coperr, corerr, rerr, lerr, lorerr, uerr, uolerr}

	if del {
		_, cnlerr := tClient.Orders.Cancel(id)
		errs = append(errs, cnlerr)

		_, colerr := tClient.Orders.CancelOrderLine(id, nil)
		errs = append(errs, colerr)
	}

	return errs
}
