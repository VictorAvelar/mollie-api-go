package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestRefundsService_CreatePaymentRefund(t *testing.T) {
	type args struct {
		ctx       context.Context
		paymentID string
		re        CreatePaymentRefund
		options   *PaymentRefundOptions
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
			"create payment refund works as expected",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				CreatePaymentRefund{
					Amount: &Amount{
						Value:    "5.00",
						Currency: "EUR",
					},
				},
				&PaymentRefundOptions{},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(testdata.CreatePaymentRefundResponse))
			},
		},
		{
			"create payment refund works as expected with access tokens",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				CreatePaymentRefund{
					Amount: &Amount{
						Value:    "5.00",
						Currency: "EUR",
					},
				},
				&PaymentRefundOptions{},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(testdata.CreatePaymentRefundResponse))
			},
		},
		{
			"create payment refund, an error is returned from the server",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				CreatePaymentRefund{},
				&PaymentRefundOptions{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"create payment refund, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				CreatePaymentRefund{},
				&PaymentRefundOptions{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get settlement refunds, invalid url when building request",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				CreatePaymentRefund{},
				&PaymentRefundOptions{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds", c.args.paymentID), c.handler)

			res, ref, err := tClient.Refunds.CreatePaymentRefund(
				c.args.ctx,
				c.args.paymentID,
				c.args.re,
				c.args.options,
			)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Refund{}, ref)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_GetPaymentRefund(t *testing.T) {
	type args struct {
		ctx       context.Context
		paymentID string
		refundID  string
		opts      *PaymentRefundOptions
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
			"get payment refund works as expected",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
				&PaymentRefundOptions{},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(testdata.GetPaymentRefundResponse))
			},
		},
		{
			"get payment refund works as expected with access tokens",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
				&PaymentRefundOptions{},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(testdata.GetPaymentRefundResponse))
			},
		},
		{
			"get payment refund, an error is returned from the server",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
				&PaymentRefundOptions{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get payment refund, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
				&PaymentRefundOptions{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get payment refund, invalid url when building request",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
				&PaymentRefundOptions{},
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
			tMux.HandleFunc(
				fmt.Sprintf("/v2/payments/%s/refunds/%s", c.args.paymentID, c.args.refundID),
				c.handler,
			)

			res, ref, err := tClient.Refunds.GetPaymentRefund(
				c.args.ctx,
				c.args.paymentID,
				c.args.refundID,
				c.args.opts,
			)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Refund{}, ref)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_ListPaymentRefunds(t *testing.T) {
	type args struct {
		ctx       context.Context
		paymentID string
		opts      *RefundsListOptions
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
			"list payment refunds works as expected",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				&RefundsListOptions{},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(testdata.ListPaymentRefundsResponse))
			},
		},
		{
			"list payment refunds works as expected with access tokens",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				&RefundsListOptions{},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(testdata.ListPaymentRefundsResponse))
			},
		},
		{
			"list payment refunds, an error is returned from the server",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				&RefundsListOptions{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"list payment refunds, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				&RefundsListOptions{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list payment refunds, invalid url when building request",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				&RefundsListOptions{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds", c.args.paymentID), c.handler)

			res, rl, err := tClient.Refunds.ListPaymentRefunds(
				c.args.ctx,
				c.args.paymentID,
				c.args.opts,
			)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &RefundsList{}, rl)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_CancelPaymentRefund(t *testing.T) {
	type args struct {
		ctx       context.Context
		paymentID string
		refundID  string
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
			"cancel payment refund works as expected",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"cancel payment refund works as expected with access tokens",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"cancel payment refund, an error is returned from the server",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"cancel payment refund, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"cancel payment refund, invalid url when building request",
			args{
				context.Background(),
				"tr_7UhSN1zuXS",
				"re_4qqhO89gsT",
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
			tMux.HandleFunc(
				fmt.Sprintf("/v2/payments/%s/refunds/%s", c.args.paymentID, c.args.refundID),
				c.handler,
			)

			res, err := tClient.Refunds.CancelPaymentRefund(
				c.args.ctx,
				c.args.paymentID,
				c.args.refundID,
			)
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

func TestRefundsService_CreateOrderRefund(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderID string
		r       CreateOrderRefund
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
			"create order refund works as expected",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				CreateOrderRefund{
					Lines: []*OrderRefundLine{
						{
							ID:       "odl_dgtxyl",
							Quantity: 1,
							Amount:   &Amount{Value: "5.00", Currency: "EUR"},
						},
					},
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(testdata.CreateOrderRefundResponse))
			},
		},
		{
			"create order refund works as expected with access tokens",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				CreateOrderRefund{
					Lines: []*OrderRefundLine{
						{
							ID:       "odl_dgtxyl",
							Quantity: 1,
							Amount:   &Amount{Value: "5.00", Currency: "EUR"},
						},
					},
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(testdata.CreateOrderRefundResponse))
			},
		},
		{
			"create order refund, an error is returned from the server",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				CreateOrderRefund{},
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
				"ord_8wmqcHMN4U",
				CreateOrderRefund{},
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
				"ord_8wmqcHMN4U",
				CreateOrderRefund{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/refunds", c.args.orderID), c.handler)

			res, ref, err := tClient.Refunds.CreateOrderRefund(
				c.args.ctx,
				c.args.orderID,
				c.args.r,
			)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Refund{}, ref)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_ListOrderRefunds(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderID string
		opts    *RefundsListOptions
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
			"list order refunds works as expected",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				&RefundsListOptions{},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(testdata.ListOrderRefundsResponse))
			},
		},
		{
			"list order refunds works as expected with access tokens",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				&RefundsListOptions{},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(testdata.ListOrderRefundsResponse))
			},
		},
		{
			"list order refunds, an error is returned from the server",
			args{
				context.Background(),
				"ord_8wmqcHMN4U",
				&RefundsListOptions{},
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
				"ord_8wmqcHMN4U",
				&RefundsListOptions{},
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
				"ord_8wmqcHMN4U",
				&RefundsListOptions{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/orders/%s/refunds", c.args.orderID), c.handler)

			res, rl, err := tClient.Refunds.ListOrderRefunds(
				c.args.ctx,
				c.args.orderID,
				c.args.opts,
			)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &RefundsList{}, rl)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_List(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts *RefundsListOptions
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
			"list refunds works as expected",
			args{
				context.Background(),
				&RefundsListOptions{},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(testdata.ListRefundsResponse))
			},
		},
		{
			"list refunds works as expected with access tokens",
			args{
				context.Background(),
				&RefundsListOptions{},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(testdata.ListRefundsResponse))
			},
		},
		{
			"list refunds, an error is returned from the server",
			args{
				context.Background(),
				&RefundsListOptions{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"list refunds, an error occurs when parsing json",
			args{
				context.Background(),
				&RefundsListOptions{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list refunds, invalid url when building request",
			args{
				context.Background(),
				&RefundsListOptions{},
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
			tMux.HandleFunc("/v2/refunds", c.handler)

			res, rl, err := tClient.Refunds.List(
				c.args.ctx,
				c.args.opts,
			)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error(), "error message does not match")
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &RefundsList{}, rl)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
