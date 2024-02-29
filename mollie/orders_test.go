package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestOrdersService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "profileId=pfl_1236h213bv1")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s", c.args.order), c.handler)

			res, m, err := tClient.Orders.Get(c.args.ctx, c.args.order, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Order{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "profileId=pfl_1236h213bv1")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/orders", c.handler)

			res, m, err := tClient.Orders.List(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &OrderList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")

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
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/orders", c.handler)

			res, m, err := tClient.Orders.Create(c.args.ctx, c.args.order, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Order{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "PATCH")

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
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "PATCH")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s", c.args.order.ID), c.handler)

			res, m, err := tClient.Orders.Update(c.args.ctx, c.args.order.ID, c.args.order)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Order{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_Cancel(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "DELETE")

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
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "DELETE")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s", c.args.order.ID), c.handler)

			res, m, err := tClient.Orders.Cancel(c.args.ctx, c.args.order.ID)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Order{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_UpdateOrderLine(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "PATCH")

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
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "PATCH")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/lines/%s", c.args.order, c.args.line.ID), c.handler)

			res, m, err := tClient.Orders.UpdateOrderLine(c.args.ctx, c.args.order, c.args.line.ID, c.args.line)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Order{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_CancelOrderLine(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "DELETE")

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
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "DELETE")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/lines", c.args.order), c.handler)

			l := []OrderLine{
				c.args.line,
			}

			res, err := tClient.Orders.CancelOrderLines(c.args.ctx, c.args.order, l)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_CreateOrderPayment(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")

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
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/payments", c.args.order), c.handler)

			res, m, err := tClient.Orders.CreateOrderPayment(c.args.ctx, c.args.order, c.args.payment)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Payment{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_CreateOrderRefund(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")

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
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/refunds", c.args.order.ID), c.handler)

			res, m, err := tClient.Orders.CreateOrderRefund(c.args.ctx, c.args.order.ID, c.args.order)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Refund{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_ListOrderRefund(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=100")

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
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "GET")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/refunds", c.args.order.ID), c.handler)

			res, m, err := tClient.Orders.ListOrderRefunds(c.args.ctx, c.args.order.ID, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &OrderListRefund{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService_ManageOrderLines(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		order      string
		operations *OrderLineOperations
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
			"manage order line works as expected.",
			args{
				context.Background(),
				"ord_pbjz8x",
				&OrderLineOperations{
					Operations: []*OrderLineChangeInstruction{
						{
							Operation: AddOrderLine,
							Data: &OrderLineOperationData{
								ID:   "odl_1.1l9vx0",
								Name: "new order line",
							},
						},
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
				_, _ = w.Write([]byte(testdata.ManageOrderLinesResponse))
			},
		},
		{
			"update order lines works as expected.",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				&OrderLineOperations{
					Operations: []*OrderLineChangeInstruction{
						{
							Operation: AddOrderLine,
							Data: &OrderLineOperationData{
								ID:   "odl_1.1l9vx0",
								Name: "new order line",
							},
						},
					},
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "PATCH")

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
				&OrderLineOperations{
					Operations: []*OrderLineChangeInstruction{
						{
							Operation: AddOrderLine,
							Data: &OrderLineOperationData{
								ID:   "odl_1.1l9vx0",
								Name: "new order line",
							},
						},
					},
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"update order lines, an error occurs when parsing json",
			args{
				context.Background(),
				"ord_kEn1PlbGa",
				&OrderLineOperations{
					Operations: []*OrderLineChangeInstruction{
						{
							Operation: AddOrderLine,
							Data: &OrderLineOperationData{
								ID:   "odl_1.1l9vx0",
								Name: "new order line",
							},
						},
					},
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
				&OrderLineOperations{
					Operations: []*OrderLineChangeInstruction{
						{
							Operation: AddOrderLine,
							Data: &OrderLineOperationData{
								ID:   "odl_1.1l9vx0",
								Name: "new order line",
							},
						},
					},
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/lines", c.args.order), c.handler)

			res, m, err := tClient.Orders.ManageOrderLines(c.args.ctx, c.args.order, c.args.operations)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Order{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
