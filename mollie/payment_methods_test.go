package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestMethodsService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		options *ListPaymentMethodsOptions
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
			"list methods works as expected.",
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
				_, _ = w.Write([]byte(testdata.ListMethodsResponse))
			},
		},
		{
			"list methods with options works as expected.",
			args{
				context.Background(),
				&ListPaymentMethodsOptions{
					Amount: &Amount{
						Value:    "100.00",
						Currency: "EUR",
					},
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "amount%5Bcurrency%5D=EUR&amount%5Bvalue%5D=100.00")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListMethodsResponse))
			},
		},
		{
			"list methods, an error is returned from the server",
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
			"list methods, an error occurs when parsing json",
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
			"list methods, invalid url when building request",
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
			tMux.HandleFunc("/v2/methods", c.handler)

			res, m, err := tClient.PaymentMethods.List(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentMethodsList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestMethodsService_All(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		options *ListPaymentMethodsOptions
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
			"list methods works as expected.",
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
				_, _ = w.Write([]byte(testdata.ListMethodsResponse))
			},
		},
		{
			"list methods with options works as expected.",
			args{
				context.Background(),
				&ListPaymentMethodsOptions{
					Amount: &Amount{
						Value:    "100.00",
						Currency: "EUR",
					},
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "amount%5Bcurrency%5D=EUR&amount%5Bvalue%5D=100.00")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListMethodsResponse))
			},
		},
		{
			"list methods, an error is returned from the server",
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
			"list methods, an error occurs when parsing json",
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
			"list methods, invalid url when building request",
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
			tMux.HandleFunc("/v2/methods/all", c.handler)

			res, m, err := tClient.PaymentMethods.All(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentMethodsList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestMethodsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		options *PaymentMethodOptions
		method  PaymentMethod
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
			"get methods works as expected.",
			args{
				context.Background(),
				nil,
				PayPal,
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
				_, _ = w.Write([]byte(testdata.ListMethodsResponse))
			},
		},
		{
			"get methods with options works as expected.",
			args{
				context.Background(),
				&PaymentMethodOptions{Locale: Catalan},
				PayPal,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "locale=ca_ES")

				fmt.Println(r.Context())

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListMethodsResponse))
			},
		},
		{
			"get methods, an error is returned from the server",
			args{
				context.Background(),
				nil,
				PayPal,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get methods, an error occurs when parsing json",
			args{
				context.Background(),
				nil,
				PayPal,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get methods, invalid url when building request",
			args{
				context.Background(),
				nil,
				PayPal,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/methods/%s", c.args.method), c.handler)

			res, m, err := tClient.PaymentMethods.Get(c.args.ctx, c.args.method, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentMethodDetails{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
