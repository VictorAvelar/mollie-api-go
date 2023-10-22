package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/suite"
)

type shipmentsServiceSuite struct{ suite.Suite }

func (ps *shipmentsServiceSuite) SetupSuite() { setEnv() }

func (ps *shipmentsServiceSuite) TearDownSuite() { unsetEnv() }

func (ps *shipmentsServiceSuite) TestShipmentsService_Get() {
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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")

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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments/%s", c.args.order, c.args.shipment), c.handler)

			res, m, err := tClient.Shipments.Get(c.args.ctx, c.args.order, c.args.shipment)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Shipment{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *shipmentsServiceSuite) TestShipmentsService_List() {
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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")

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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments", c.args.order), c.handler)

			res, m, err := tClient.Shipments.List(c.args.ctx, c.args.order)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&ShipmentsList{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *shipmentsServiceSuite) TestShipmentsService_Create() {
	type args struct {
		ctx      context.Context
		order    string
		shipment CreateShipmentRequest
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
				CreateShipmentRequest{
					TestMode: true,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "POST")

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
				CreateShipmentRequest{
					TestMode: true,
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "testmode=true")

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
				CreateShipmentRequest{
					TestMode: true,
				},
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
				CreateShipmentRequest{
					TestMode: true,
				},
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
				CreateShipmentRequest{
					TestMode: true,
				},
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments", c.args.order), c.handler)

			res, m, err := tClient.Shipments.Create(c.args.ctx, c.args.order, c.args.shipment)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Shipment{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *shipmentsServiceSuite) TestShipmentsService_Update() {
	type args struct {
		ctx      context.Context
		order    string
		shipment string
		st       ShipmentTracking
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
				ShipmentTracking{
					Carrier: "fedex",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "PATCH")

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
				ShipmentTracking{
					Carrier: "fedex",
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(ps.T(), r, "PATCH")
				testQuery(ps.T(), r, "testmode=true")

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
				ShipmentTracking{
					Carrier: "fedex",
				},
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
				ShipmentTracking{
					Carrier: "fedex",
				},
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
				ShipmentTracking{
					Carrier: "fedex",
				},
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/shipments/%s", c.args.order, c.args.shipment), c.handler)

			res, m, err := tClient.Shipments.Update(c.args.ctx, c.args.order, c.args.shipment, c.args.st)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Shipment{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestShipmentsService(t *testing.T) {
	suite.Run(t, new(shipmentsServiceSuite))
}
