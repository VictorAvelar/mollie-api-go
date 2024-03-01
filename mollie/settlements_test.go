package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestSettlementsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		settlement string
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
			"get settlement works as expected.",
			args{
				context.Background(),
				"stl_jDk30akdN",
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
				_, _ = w.Write([]byte(testdata.GetSettlementsResponse))
			},
		},
		{
			"get settlement, an error is returned from the server",
			args{
				context.Background(),
				"stl_jDk30akdN",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get settlement, an error occurs when parsing json",
			args{
				context.Background(),
				"stl_jDk30akdN",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get settlement, invalid url when building request",
			args{
				context.Background(),
				"stl_jDk30akdN",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/settlements/%s", c.args.settlement), c.handler)

			res, m, err := tClient.Settlements.Get(c.args.ctx, c.args.settlement)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Settlement{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSettlementsService_Next(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		settlement string
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
			"get next settlement works as expected.",
			args{
				context.Background(),
				"next",
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
				_, _ = w.Write([]byte(testdata.GetNextSettlementResponse))
			},
		},
		{
			"get next settlement, an error is returned from the server",
			args{
				context.Background(),
				"next",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get next settlement, an error occurs when parsing json",
			args{
				context.Background(),
				"next",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get next settlement, invalid url when building request",
			args{
				context.Background(),
				"next",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/settlements/%s", c.args.settlement), c.handler)

			res, m, err := tClient.Settlements.Next(c.args.ctx)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Settlement{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSettlementsService_Open(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		settlement string
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
			"get open settlement works as expected.",
			args{
				context.Background(),
				"open",
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
				_, _ = w.Write([]byte(testdata.GetOpenSettlementResponse))
			},
		},
		{
			"get open settlement, an error is returned from the server",
			args{
				context.Background(),
				"open",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get open settlement, an error occurs when parsing json",
			args{
				context.Background(),
				"open",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get open settlement, invalid url when building request",
			args{
				context.Background(),
				"open",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/settlements/%s", c.args.settlement), c.handler)

			res, m, err := tClient.Settlements.Open(c.args.ctx)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Settlement{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSettlementsService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		options *SettlementsListOptions
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
				&SettlementsListOptions{
					Limit: 40,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=40")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				w.Write([]byte(testdata.ListSettlementsResponse))
			},
		},
		{
			"list shipment, an error is returned from the server",
			args{
				context.Background(),
				&SettlementsListOptions{
					Limit: 40,
				},
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
				&SettlementsListOptions{
					Limit: 40,
				},
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
				&SettlementsListOptions{
					Limit: 40,
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
			tMux.HandleFunc("/v2/settlements", c.handler)

			res, m, err := tClient.Settlements.List(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &SettlementsList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSettlementsService_GetPayments(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		settlement string
		options    *SettlementsListOptions
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
			"get settlement payments works as expected.",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=10")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
			},
		},
		{
			"get settlement payments, an error is returned from the server",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get settlement payments, an error occurs when parsing json",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get settlement payments, invalid url when building request",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/settlements/%s/payments", c.args.settlement), c.handler)

			res, m, err := tClient.Settlements.GetPayments(c.args.ctx, c.args.settlement, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSettlementsService_GetCaptures(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		settlement string
		options    *SettlementsListOptions
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
			"get settlement captures works as expected.",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=10")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListCapturesResponse))
			},
		},
		{
			"get settlement captures, an error is returned from the server",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get settlement captures, an error occurs when parsing json",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get settlement captures, invalid url when building request",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/settlements/%s/captures", c.args.settlement), c.handler)

			res, m, err := tClient.Settlements.GetCaptures(c.args.ctx, c.args.settlement, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &CapturesList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSettlementsService_GetChargebacks(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		settlement string
		options    *SettlementsListOptions
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
			"get settlement chargebacks works as expected.",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=10")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
			},
		},
		{
			"get settlement chargebacks, an error is returned from the server",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get settlement chargebacks, an error occurs when parsing json",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get settlement chargebacks, invalid url when building request",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/settlements/%s/chargebacks", c.args.settlement), c.handler)

			res, m, err := tClient.Settlements.GetChargebacks(c.args.ctx, c.args.settlement, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &ChargebacksList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSettlementsService_GetRefunds(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		settlement string
		options    *SettlementsListOptions
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
			"get settlement refunds works as expected.",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=10")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetRefundListResponse))
			},
		},
		{
			"get settlement refunds, an error is returned from the server",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get settlement refunds, an error occurs when parsing json",
			args{
				context.Background(),
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
				},
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
				"stl_jDk30akdN",
				&SettlementsListOptions{
					Limit: 10,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/settlements/%s/refunds", c.args.settlement), c.handler)

			res, m, err := tClient.Settlements.GetRefunds(c.args.ctx, c.args.settlement, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &RefundList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
