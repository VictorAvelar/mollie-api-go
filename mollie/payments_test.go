package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPaymentsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
					Include: []IncludeValue{IncludeQrCode},
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "include=details.qrCode")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s", c.args.payment), c.handler)

			res, m, err := tClient.Payments.Get(c.args.ctx, c.args.payment, c.args.options)
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

func TestPaymentsService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		options *ListPaymentsOptions
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
				&ListPaymentsOptions{
					From: "tr_12o93213",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "from=tr_12o93213")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/payments", c.handler)

			res, m, err := tClient.Payments.List(c.args.ctx, c.args.options)
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

func TestPaymentsService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment CreatePayment
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
				CreatePayment{
					BillingEmail: "test@example.com",
				},
				&PaymentOptions{
					Include: []IncludeValue{},
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
				_, _ = w.Write([]byte(testdata.GetPaymentResponse))
			},
		},
		{
			"create payments with access token works as expected.",
			args{
				context.Background(),
				CreatePayment{
					BillingEmail: "test@example.com",
				},
				&PaymentOptions{
					Include: []IncludeValue{},
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
				_, _ = w.Write([]byte(testdata.GetPaymentResponse))
			},
		},
		{
			"create payments, an error is returned from the server",
			args{
				context.Background(),
				CreatePayment{},
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"create payments, an error occurs when parsing json",
			args{
				context.Background(),
				CreatePayment{},
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
				CreatePayment{},
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
			tMux.HandleFunc("/v2/payments", c.handler)

			res, m, err := tClient.Payments.Create(c.args.ctx, c.args.payment, c.args.options)
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

func TestPaymentsService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		id      string
		payment UpdatePayment
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
				UpdatePayment{
					BillingEmail: "test@example.com",
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
				_, _ = w.Write([]byte(testdata.UpdatePaymentResponse))
			},
		},
		{
			"update payments using access token works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				UpdatePayment{
					BillingEmail: "test@example.com",
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_example_token")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_example_token")
				testMethod(t, r, "PATCH")
				testQuery(t, r, "testmode=true")

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
				UpdatePayment{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"update payments, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				UpdatePayment{},
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
				UpdatePayment{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s", c.args.id), c.handler)

			res, m, err := tClient.Payments.Update(c.args.ctx, c.args.id, c.args.payment)
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

func TestPaymentsService_Cancel(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "DELETE")

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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s", c.args.id), c.handler)

			res, m, err := tClient.Payments.Cancel(c.args.ctx, c.args.id)
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
