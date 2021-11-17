package mollie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestShipmentsService_Get(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_kEn1PlbGa"
	shipmentID := "shp_3wmsgCJN4U"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments/%s", orderID, shipmentID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
	})

	shipment, err := tClient.Shipments.Get(orderID, shipmentID)
	if err != nil {
		t.Error(err)
	}

	if shipment.ID != shipmentID {
		t.Errorf("unexpected response, got: %v want: %v", shipment.ID, shipmentID)
	}
}

func TestShipmentsService_Create(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_kEn1PlbGa"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments", orderID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPost)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
	})

	csr := CreateShipmentRequest{
		Lines:    []OrderLine{},
		Tracking: ShipmentTracking{},
	}

	shipment, err := tClient.Shipments.Create(orderID, csr)
	if err != nil {
		t.Error(err)
	}

	if shipment.OrderID != orderID {
		t.Errorf("unexpected response, got: %v want: %v", shipment.OrderID, orderID)
	}
}

func TestShipmentsService_Create_AccessTokens(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("access_token")

	orderID := "ord_kEn1PlbGa"

	tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments", orderID), func(rw http.ResponseWriter, r *http.Request) {
		var ship Shipment
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&ship); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(ship)
	})

	csr := CreateShipmentRequest{
		Lines:    []OrderLine{},
		Tracking: ShipmentTracking{},
	}

	shipment, err := tClient.Shipments.Create(orderID, csr)
	if err != nil {
		t.Error(err)
	}

	if shipment.TestMode != true {
		t.Fatal("testmode flag is not set for access tokens")
	}
}

func TestShipmentsService_List(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_kEn1PlbGa"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments", orderID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
	})

	list, err := tClient.Shipments.List(orderID)
	if err != nil {
		t.Error(err)
	}

	if list.Count != len(list.Embedded.Shipments) {
		t.Error("the response doesn't match the reference")
	}
}

func TestShipmentsService_Update(t *testing.T) {
	setup()
	defer teardown()

	orderID := "ord_kEn1PlbGa"
	shipmentID := "shp_3wmsgCJN4U"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments/%s", orderID, shipmentID), func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPatch)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
	})

	shipment, err := tClient.Shipments.Update(orderID, shipmentID, ShipmentTracking{
		Carrier: "PostNL",
	})
	if err != nil {
		t.Error(err)
	}

	if shipment.ID != shipmentID {
		t.Errorf("unexpected response, got: %v want: %v", shipment.ID, shipmentID)
	}
}

func TestShipmentsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/orders/ord_kEn1PlbGa/shipments/", errorHandler)

	tests := forceShipmentsErrors()

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func TestShipmentsService_JsonDecodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/orders/ord_kEn1PlbGa/shipments/", encodingHandler)

	tests := forceShipmentsErrors()

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestShipmentsService_HTTPRequestErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/orders/ord_kEn1PlbGa/shipments/", errorHandler)

	tests := forceShipmentsErrors()
	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func forceShipmentsErrors() []error {
	_, gerr := tClient.Shipments.Get("ord_kEn1PlbGa", "as6da7d8sa9d")
	_, lerr := tClient.Shipments.List("ord_kEn1PlbGa")
	_, cerr := tClient.Shipments.Create("ord_kEn1PlbGa", CreateShipmentRequest{})
	_, uerr := tClient.Shipments.Update("ord_kEn1PlbGa", "as5das67d9s", ShipmentTracking{})

	return []error{gerr, lerr, cerr, uerr}
}
