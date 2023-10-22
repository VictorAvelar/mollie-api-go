package mollie

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/suite"
)

type invoiceServiceSuite struct{ suite.Suite }

func (is *invoiceServiceSuite) SetupSuite() { setEnv() }

func (is *invoiceServiceSuite) TearDownSuite() { unsetEnv() }

func (is *invoiceServiceSuite) TestInvoicesService_Get() {
	type args struct {
		ctx     context.Context
		invoice string
		options *InvoicesListOptions
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
			"get invoice works as expecter",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
				nil,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(is.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(is.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.GetInvoiceResponse))
			},
		},
		{
			"get invoice, an error is returned from the server",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get invoice, an error occurs when parsing json",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get invoice, invalid url when building request",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
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
		is.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/invoices/%s", c.args.invoice), c.handler)

			res, i, err := tClient.Invoices.Get(c.args.ctx, c.args.invoice)
			if c.wantErr {
				is.NotNil(err)
				is.EqualError(err, c.err.Error())
			} else {
				is.Nil(err)
				is.IsType(&Invoice{}, i)
				is.Same(c.args.ctx, res.Request.Context())
				is.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (is *invoiceServiceSuite) TestInvoicesService_List() {
	type args struct {
		ctx     context.Context
		invoice string
		options *InvoicesListOptions
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
			"list invoices works as expecter",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
				nil,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(is.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(is.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.ListInvoicesResponse))
			},
		},
		{
			"list invoices with options works as expecter",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
				&InvoicesListOptions{
					Year: strconv.Itoa(time.Now().Year()),
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(is.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(is.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.ListInvoicesResponse))
			},
		},
		{
			"list invoices, an error is returned from the server",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"list invoices, an error occurs when parsing json",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list invoices, invalid url when building request",
			args{
				context.Background(),
				"inv_xBEbP9rvAq",
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
		is.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/invoices", c.handler)

			res, i, err := tClient.Invoices.List(c.args.ctx, c.args.options)
			if c.wantErr {
				is.NotNil(err)
				is.EqualError(err, c.err.Error())
			} else {
				is.Nil(err)
				is.IsType(&InvoicesList{}, i)
				is.Same(c.args.ctx, res.Request.Context())
				is.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestInvoicesService(t *testing.T) {
	suite.Run(t, new(invoiceServiceSuite))
}
