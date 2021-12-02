package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type ordersServiceSuite struct{ suite.Suite }

func (os *ordersServiceSuite) SetupSuite() { setEnv() }

func (os *ordersServiceSuite) TearDownSuite() { unsetEnv() }

func (os *ordersServiceSuite) TestOrdersService_Get() {
	type args struct {
		ctx     context.Context
		order   string
		options *OrderOptions
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
			"get orders works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				&OrderOptions{
					ProfileID: "pfl_1236h213bv1",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")
				testQuery(os.T(), r, "profileId=pfl_1236h213bv1&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOrderResponse))
			},
		},
		{
			"get orders, an error is returned from the server",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"get orders, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get orders, invalid url when building request",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				nil,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s", c.args.order), c.handler)

			m, err := tClient.Orders.Get(c.args.ctx, c.args.order, c.args.options)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Order{}, m)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_List() {
	type args struct {
		ctx     context.Context
		order   string
		options *OrderListOptions
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
			"list orders works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				&OrderListOptions{
					ProfileID: "pfl_1236h213bv1",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")
				testQuery(os.T(), r, "profileId=pfl_1236h213bv1&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOrderResponse))
			},
		},
		{
			"list orders, an error is returned from the server",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"list orders, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list orders, invalid url when building request",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				nil,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/orders", c.handler)

			m, err := tClient.Orders.List(c.args.ctx, c.args.options)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&OrderList{}, m)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_Create() {
	type args struct {
		ctx     context.Context
		order   Order
		options *OrderOptions
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
			"create orders works as expected.",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				&OrderOptions{},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "POST")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOrderResponse))
			},
		},
		{
			"create orders works as expected.",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				&OrderOptions{},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(os.T(), r, "POST")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOrderResponse))
			},
		},
		{
			"create orders, an error is returned from the server",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"create orders, an error occurs when parsing json",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create orders, invalid url when building request",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/orders", c.handler)

			m, err := tClient.Orders.Create(c.args.ctx, c.args.order, c.args.options)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Order{}, m)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_Update() {
	type args struct {
		ctx     context.Context
		order   Order
		options *OrderOptions
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
			"update orders works as expected.",
			args{
				context.Background(),
				Order{
					Method: PayPal,
					ID:     "ord_kEn1PlbGa",
				},
				&OrderOptions{},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "PATCH")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.UpdateOrderResponse))
			},
		},
		{
			"update orders works as expected.",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				&OrderOptions{},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(os.T(), r, "PATCH")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.UpdateOrderResponse))
			},
		},
		{
			"update orders, an error is returned from the server",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"update orders, an error occurs when parsing json",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update orders, invalid url when building request",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s", c.args.order.ID), c.handler)

			m, err := tClient.Orders.Update(c.args.ctx, c.args.order.ID, c.args.order)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Order{}, m)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_Cancel() {
	type args struct {
		ctx     context.Context
		order   Order
		options *OrderOptions
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
			"cancel orders works as expected.",
			args{
				context.Background(),
				Order{
					Method: PayPal,
					ID:     "ord_kEn1PlbGa",
				},
				&OrderOptions{},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "DELETE")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CancelOrderResponse))
			},
		},
		{
			"cancel orders works as expected.",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				&OrderOptions{},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(os.T(), r, "DELETE")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CancelOrderResponse))
			},
		},
		{
			"cancel orders, an error is returned from the server",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"cancel orders, an error occurs when parsing json",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"cancel orders, invalid url when building request",
			args{
				context.Background(),
				Order{
					Method: PayPal,
				},
				nil,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s", c.args.order.ID), c.handler)

			m, err := tClient.Orders.Cancel(c.args.ctx, c.args.order.ID)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Order{}, m)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_UpdateOrderLine() {
	type args struct {
		ctx   context.Context
		order string
		line  OrderLine
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
			"update orders works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "PATCH")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.UpdateOrderlineResponse))
			},
		},
		{
			"update order lines works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(os.T(), r, "PATCH")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.UpdateOrderlineResponse))
			},
		},
		{
			"update order lines, an error is returned from the server",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
				},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"update order lines, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update order lines, invalid url when building request",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/lines/%s", c.args.order, c.args.line.ID), c.handler)

			m, err := tClient.Orders.UpdateOrderLine(c.args.ctx, c.args.order, c.args.line.ID, c.args.line)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Order{}, m)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_CancelOrderLine() {
	type args struct {
		ctx   context.Context
		order string
		line  OrderLine
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
			"update orders works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "DELETE")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.UpdateOrderlineResponse))
			},
		},
		{
			"update order lines works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(os.T(), r, "DELETE")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.UpdateOrderlineResponse))
			},
		},
		{
			"update order lines, an error is returned from the server",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
				},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"update order lines, invalid url when building request",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				OrderLine{
					ID:     "odl_dgtxyl",
					Status: OrderLinePaid,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/lines", c.args.order), c.handler)

			l := []OrderLine{
				c.args.line,
			}

			err := tClient.Orders.CancelOrderLines(c.args.ctx, c.args.order, l)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_CreateOrderPayment() {
	type args struct {
		ctx     context.Context
		order   string
		payment *OrderPayment
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
			"create order payments works as expected.",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				&OrderPayment{
					Method: PayPal,
					Issuer: "tr_asdajnasd",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "POST")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CreateOrderPaymentResponse))
			},
		},
		{
			"create order payments works as expected.",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				&OrderPayment{
					Method: PayPal,
					Issuer: "tr_asdajnasd",
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(os.T(), r, "POST")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CreateOrderPaymentResponse))
			},
		},
		{
			"create order payments, an error is returned from the server",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"create order payments, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create order payments, invalid url when building request",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				nil,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/payments", c.args.order), c.handler)

			m, err := tClient.Orders.CreateOrderPayment(c.args.ctx, c.args.order, c.args.payment)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Payment{}, m)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_CreateOrderRefund() {
	type args struct {
		ctx   context.Context
		order *Order
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
			"create order refund works as expected.",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "POST")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CreateOrderRefundResponse))
			},
		},
		{
			"create order refund works as expected.",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(os.T(), r, "POST")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CreateOrderRefundResponse))
			},
		},
		{
			"create order refund, an error is returned from the server",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"create order refund, an error occurs when parsing json",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create order refund, invalid url when building request",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/refunds", c.args.order.ID), c.handler)

			m, err := tClient.Orders.CreateOrderRefund(c.args.ctx, c.args.order.ID, c.args.order)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(Refund{}, m)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_ListOrderRefund() {
	type args struct {
		ctx     context.Context
		order   *Order
		options *OrderListRefundOptions
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
			"list order refunds works as expected.",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
				&OrderListRefundOptions{Limit: 100},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")
				testQuery(os.T(), r, "limit=100&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CreateOrderRefundResponse))
			},
		},
		{
			"list order refunds works as expected.",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
				nil,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(os.T(), r, "GET")
				testQuery(os.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CreateOrderRefundResponse))
			},
		},
		{
			"list order refunds, an error is returned from the server",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"list order refunds, an error occurs when parsing json",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list order refunds, invalid url when building request",
			args{
				context.Background(),
				&Order{ID: "ord_8wmqcHMN4U"},
				nil,
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/refunds", c.args.order.ID), c.handler)

			m, err := tClient.Orders.ListOrderRefunds(c.args.ctx, c.args.order.ID, c.args.options)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(OrderListRefund{}, m)
			}
		})
	}
}
func TestOrdersService(t *testing.T) {
	suite.Run(t, new(ordersServiceSuite))
}
