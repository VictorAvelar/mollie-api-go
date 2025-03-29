package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPaymentLinkService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx         context.Context
		paymentLink string
		opts        *PaymentLinkOptions
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
			"get payment links works as expected.",
			args{
				context.Background(),
				"pl_ka21123129",
				nil,
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
				_, _ = w.Write([]byte(testdata.GetPaymentLinkResponse))
			},
		},
		{
			"get payment links, an error is returned from the server",
			args{
				context.Background(),
				"pl_ka21123129",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get payment links, an error occurs when parsing json",
			args{
				context.Background(),
				"pl_ka21123129",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get payment links, invalid url when building request",
			args{
				context.Background(),
				"pl_ka21123129",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payment-links/%s", c.args.paymentLink), c.handler)

			res, m, err := tClient.PaymentLinks.Get(c.args.ctx, c.args.paymentLink)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentLink{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestPaymentLinkService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx         context.Context
		paymentLink PaymentLink
		options     *PaymentLinkOptions
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
			"create payment links works as expected.",
			args{
				context.Background(),
				PaymentLink{
					Amount: Amount{Value: "20.00", Currency: "EUR"},
				},
				&PaymentLinkOptions{ProfileID: "prf_12312312"},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")
				testQuery(t, r, "profileId=prf_12312312")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.CreatePaymentLinkResponse))
			},
		},
		{
			"create payment links, an error is returned from the server",
			args{
				context.Background(),
				PaymentLink{},
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"create payment links, an error occurs when parsing json",
			args{
				context.Background(),
				PaymentLink{},
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create payment links, invalid url when building request",
			args{
				context.Background(),
				PaymentLink{},
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
			tMux.HandleFunc("/v2/payment-links", c.handler)

			res, m, err := tClient.PaymentLinks.Create(c.args.ctx, c.args.paymentLink, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentLink{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestPaymentLinkService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx  context.Context
		opts *PaymentLinkOptions
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
			"list payment links works as expected.",
			args{
				context.Background(),
				nil,
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
				_, _ = w.Write([]byte(testdata.ListPaymentLinksResponse))
			},
		},
		{
			"list payment links with options works as expected.",
			args{
				context.Background(),
				&PaymentLinkOptions{
					ProfileID: "pfl_11211",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "profileId=pfl_11211")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListPaymentLinksResponse))
			},
		},
		{
			"list payment links, an error is returned from the server",
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
			"list payment links, an error occurs when parsing json",
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
			"list payment links, invalid url when building request",
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
			tMux.HandleFunc("/v2/payment-links", c.handler)

			res, m, err := tClient.PaymentLinks.List(c.args.ctx, c.args.opts)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentLinksList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestPaymentLinkService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx         context.Context
		paymentLink string
		pl          UpdatePaymentLinks
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
			"update payment links works as expected.",
			args{
				context.Background(),
				"pl_ka21123129",
				UpdatePaymentLinks{
					Archived: true,
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
				_, _ = w.Write([]byte(testdata.UpdatePaymentLinksResponse))
			},
		},
		{
			"update payment links, an error is returned from the server",
			args{
				context.Background(),
				"pl_ka21123129",
				UpdatePaymentLinks{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"update payment links, an error occurs when parsing json",
			args{
				context.Background(),
				"pl_ka21123129",
				UpdatePaymentLinks{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update payment links, invalid url when building request",
			args{
				context.Background(),
				"pl_ka21123129",
				UpdatePaymentLinks{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payment-links/%s", c.args.paymentLink), c.handler)

			res, m, err := tClient.PaymentLinks.Update(c.args.ctx, c.args.paymentLink, c.args.pl)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentLink{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestPaymentLinkService_Delete(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx         context.Context
		paymentLink string
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
			"delete payment links works as expected.",
			args{
				context.Background(),
				"pl_ka21123129",
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
				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"delete payment links, an error is returned from the server",
			args{
				context.Background(),
				"pl_ka21123129",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"delete payment links, invalid url when building request",
			args{
				context.Background(),
				"pl_ka21123129",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payment-links/%s", c.args.paymentLink), c.handler)

			res, err := tClient.PaymentLinks.Delete(c.args.ctx, c.args.paymentLink)
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

func TestPaymentLinkService_Payments(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx         context.Context
		paymentLink string
		opts        *PaymentLinkPaymentsListOptions
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
			"list payment link payments works as expected.",
			args{
				context.Background(),
				"pl_ka21123129",
				nil,
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
				_, _ = w.Write([]byte(testdata.ListPaymentLinkPaymentsResponse))
			},
		},
		{
			"list payment link payments with options works as expected.",
			args{
				context.Background(),
				"pl_ka21123129",
				&PaymentLinkPaymentsListOptions{Limit: 10},
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
				_, _ = w.Write([]byte(testdata.ListPaymentLinkPaymentsResponse))
			},
		},
		{
			"list payment link payments, an error is returned from the server",
			args{
				context.Background(),
				"pl_ka21123129",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"list payment link payments, an error occurs when parsing json",
			args{
				context.Background(),
				"pl_ka21123129",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list payment link payments, invalid url when building request",
			args{
				context.Background(),
				"pl_ka21123129",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payment-links/%s/payments", c.args.paymentLink), c.handler)

			res, m, err := tClient.PaymentLinks.Payments(c.args.ctx, c.args.paymentLink, c.args.opts)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentLinkPaymentsList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
