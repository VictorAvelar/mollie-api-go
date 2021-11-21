package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type methodsServiceSuite struct{ suite.Suite }

func (ms *methodsServiceSuite) SetupSuite() { setEnv() }

func (ms *methodsServiceSuite) TearDownSuite() { unsetEnv() }

func (ms *methodsServiceSuite) TestMethodsService_List() {
	type args struct {
		ctx     context.Context
		options *ListMethodsOptions
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
				&ListMethodsOptions{
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

			m, err := tClient.Methods.List(c.args.ctx, c.args.options)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&ListMethods{}, m)
			}
		})
	}
}

func (ms *methodsServiceSuite) TestMethodsService_All() {
	type args struct {
		ctx     context.Context
		options *ListMethodsOptions
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
				&ListMethodsOptions{
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

			m, err := tClient.Methods.All(c.args.ctx, c.args.options)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&ListMethods{}, m)
			}
		})
	}
}

func (ms *methodsServiceSuite) TestMethodsService_Get() {
	type args struct {
		ctx     context.Context
		options *GetMethodsOptions
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
				&GetMethodsOptions{Locale: Catalan},
				PayPal,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				testQuery(ms.T(), r, "locale=ca_ES&testmode=true")

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

			m, err := tClient.Methods.Get(c.args.ctx, string(c.args.method), c.args.options)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&PaymentMethodInfo{}, m)
			}
		})
	}
}

func TestMethodsService(t *testing.T) {
	suite.Run(t, new(methodsServiceSuite))
}
