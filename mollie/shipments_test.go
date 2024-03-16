package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestShipmentsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx      context.Context
		order    string
		shipment string
	}
	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get shipment works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
			},
		},
		{
			"get shipment, an error is returned from the server",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get shipment, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get shipment, invalid url when building request",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments/%s", c.args.order, c.args.shipment), c.handler)

			res, m, err := tClient.Shipments.Get(c.args.ctx, c.args.order, c.args.shipment)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Shipment{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestShipmentsService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx      context.Context
		order    string
		shipment string
	}
	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"list shipment works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				w.Write([]byte(testdata.ListShipmentsResponse))
			},
		},
		{
			"list shipment, an error is returned from the server",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"list shipment, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list shipment, invalid url when building request",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments", c.args.order), c.handler)

			res, m, err := tClient.Shipments.List(c.args.ctx, c.args.order)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &ShipmentsList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestShipmentsService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx      context.Context
		order    string
		shipment CreateShipment
	}
	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"create shipment works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				CreateShipment{
					Tracking: &ShipmentTracking{
						Carrier: "fedex",
						Code:    "3wmsgCJN4U",
						URL:     "https://www.example.fedex.com/apps/fedextrack/?tracknumbers=3wmsgCJN4U",
					},
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
			},
		},
		{
			"create shipment using access tokens works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				CreateShipment{
					Tracking: &ShipmentTracking{
						Carrier: "fedex",
						Code:    "3wmsgCJN4U",
						URL:     "https://www.example.fedex.com/apps/fedextrack/?tracknumbers=3wmsgCJN4U",
					},
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
			},
		},
		{
			"create shipment, an error is returned from the server",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				CreateShipment{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"create shipment, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				CreateShipment{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create shipment, invalid url when building request",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				CreateShipment{},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments", c.args.order), c.handler)

			res, m, err := tClient.Shipments.Create(c.args.ctx, c.args.order, c.args.shipment)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Shipment{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestShipmentsService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx      context.Context
		order    string
		shipment string
		st       UpdateShipment
	}
	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"update shipment works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
				UpdateShipment{
					Tracking: &ShipmentTracking{
						Carrier: "dhl",
						Code:    "3wmsgCJN4U",
						URL:     "https://www.example.dhl.com/apps/dhltrack/?tracknumbers=3wmsgCJN4U",
					},
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "PATCH")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
			},
		},
		{
			"update shipment using access tokens works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
				UpdateShipment{
					Tracking: &ShipmentTracking{
						Carrier: "dhl",
						Code:    "3wmsgCJN4U",
						URL:     "https://www.example.dhl.com/apps/dhltrack/?tracknumbers=3wmsgCJN4U",
					},
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "PATCH")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetShipmentsResponse))
			},
		},
		{
			"update shipment, an error is returned from the server",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
				UpdateShipment{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"update shipment, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
				UpdateShipment{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update shipment, invalid url when building request",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				"shp_3wmsgCJN4U",
				UpdateShipment{},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments/%s", c.args.order, c.args.shipment), c.handler)

			res, m, err := tClient.Shipments.Update(c.args.ctx, c.args.order, c.args.shipment, c.args.st)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Shipment{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
