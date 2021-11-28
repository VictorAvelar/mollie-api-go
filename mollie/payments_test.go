package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type paymentsServiceSuite struct{ suite.Suite }

func (ps *paymentsServiceSuite) SetupSuite() { setEnv() }

func (ps *paymentsServiceSuite) TearDownSuite() { unsetEnv() }

func (ps *paymentsServiceSuite) TestPaymentsService_Get() {
	type args struct {
		ctx     context.Context
		payment string
		options *PaymentOptions
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
			"get payments works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				&PaymentOptions{
					Include: "settlements",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "include=settlements&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPaymentResponse))
			},
		},
		{
			"get payments, an error is returned from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"get payments, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get payments, invalid url when building request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s", c.args.payment), c.handler)

			m, err := tClient.Payments.Get(c.args.ctx, c.args.payment, c.args.options)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(Payment{}, m)
			}
		})
	}
}

func (ps *paymentsServiceSuite) TestPaymentsService_List() {
	type args struct {
		ctx     context.Context
		options *ListPaymentOptions
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
			"get payments works as expected.",
			args{
				context.Background(),
				&ListPaymentOptions{
					From: "tr_12o93213",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "from=tr_12o93213&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
			},
		},
		{
			"get payments, an error is returned from the server",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"get payments, an error occurs when parsing json",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get payments, invalid url when building request",
			args{
				context.Background(),
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/payments", c.handler)

			m, err := tClient.Payments.List(c.args.ctx, c.args.options)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(PaymentList{}, m)
			}
		})
	}
}

func (ps *paymentsServiceSuite) TestPaymentsService_Create() {
	type args struct {
		ctx     context.Context
		payment Payment
		options *PaymentOptions
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
			"create payments works as expected.",
			args{
				context.Background(),
				Payment{
					BillingEmail: "test@example.com",
				},
				&PaymentOptions{
					Include: "settlements",
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_example_token")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_example_token")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "include=settlements&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPaymentResponse))
			},
		},
		{
			"create payments with access token works as expected.",
			args{
				context.Background(),
				Payment{
					BillingEmail: "test@example.com",
				},
				&PaymentOptions{
					Include: "settlements",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "include=settlements&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPaymentResponse))
			},
		},
		{
			"create payments, an error is returned from the server",
			args{
				context.Background(),
				Payment{},
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"create payments, an error occurs when parsing json",
			args{
				context.Background(),
				Payment{},
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create payments, invalid url when building request",
			args{
				context.Background(),
				Payment{},
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/payments", c.handler)

			m, err := tClient.Payments.Create(c.args.ctx, c.args.payment, c.args.options)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(Payment{}, m)
			}
		})
	}
}

func (ps *paymentsServiceSuite) TestPaymentsService_Update() {
	type args struct {
		ctx     context.Context
		id      string
		payment Payment
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
			"update payments works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Payment{
					BillingEmail: "test@example.com",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "PATCH")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.UpdatePaymentResponse))
			},
		},
		{
			"update payments using access token works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Payment{
					BillingEmail: "test@example.com",
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_example_token")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_example_token")
				testMethod(ps.T(), r, "PATCH")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.UpdatePaymentResponse))
			},
		},
		{
			"update payments, an error is returned from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Payment{},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"update payments, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Payment{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update payments, invalid url when building request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Payment{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s", c.args.id), c.handler)

			m, err := tClient.Payments.Update(c.args.ctx, c.args.id, c.args.payment)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(Payment{}, m)
			}
		})
	}
}

func (ps *paymentsServiceSuite) TestPaymentsService_Cancel() {
	type args struct {
		ctx context.Context
		id  string
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
			"get payments works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "DELETE")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPaymentResponse))
			},
		},
		{
			"get payments, an error is returned from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"get payments, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get payments, invalid url when building request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s", c.args.id), c.handler)

			m, err := tClient.Payments.Cancel(c.args.ctx, c.args.id)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(Payment{}, m)
			}
		})
	}
}

func TestPaymentsService(t *testing.T) {
	suite.Run(t, new(paymentsServiceSuite))
}
