package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type paymentMethodsServiceSuite struct{ suite.Suite }

func (ms *paymentMethodsServiceSuite) SetupSuite() { setEnv() }

func (ms *paymentMethodsServiceSuite) TearDownSuite() { unsetEnv() }

func (ms *paymentMethodsServiceSuite) TestMethodsService_List() {
	type args struct {
		ctx     context.Context
		options *PaymentMethodsListOptions
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
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				testQuery(ms.T(), r, "testmode=true")

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
				&PaymentMethodsListOptions{
					AmountCurrency: "EUR",
					AmountValue:    "100.00",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				testQuery(ms.T(), r, "amount%5Bcurrency%5D=EUR&amount%5Bvalue%5D=100.00&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		ms.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/methods", c.handler)

			res, m, err := tClient.PaymentMethods.List(c.args.ctx, c.args.options)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&PaymentMethodsList{}, m)
				ms.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ms *paymentMethodsServiceSuite) TestMethodsService_All() {
	type args struct {
		ctx     context.Context
		options *PaymentMethodsListOptions
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
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				testQuery(ms.T(), r, "testmode=true")

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
				&PaymentMethodsListOptions{
					AmountCurrency: "EUR",
					AmountValue:    "100.00",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				testQuery(ms.T(), r, "amount%5Bcurrency%5D=EUR&amount%5Bvalue%5D=100.00&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		ms.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/methods/all", c.handler)

			res, m, err := tClient.PaymentMethods.All(c.args.ctx, c.args.options)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&PaymentMethodsList{}, m)
				ms.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ms *paymentMethodsServiceSuite) TestMethodsService_Get() {
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
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				testQuery(ms.T(), r, "testmode=true")

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
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				testQuery(ms.T(), r, "locale=ca_ES&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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
		ms.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/methods/%s", c.args.method), c.handler)

			res, m, err := tClient.PaymentMethods.Get(c.args.ctx, c.args.method, c.args.options)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&PaymentMethodDetails{}, m)
				ms.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestMethodsService(t *testing.T) {
	suite.Run(t, new(paymentMethodsServiceSuite))
}
