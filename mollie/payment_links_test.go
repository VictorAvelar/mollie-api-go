package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type paymentLinksSuite struct{ suite.Suite }

func (ps *paymentLinksSuite) SetupSuite() { setEnv() }

func (ps *paymentLinksSuite) TearDownSuite() { unsetEnv() }

func (ps *paymentLinksSuite) TestPaymentLinkService_Get() {
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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/payment-links/%s", c.args.paymentLink), c.handler)

			m, err := tClient.PaymentLinks.Get(c.args.ctx, c.args.paymentLink)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&PaymentLink{}, m)
			}
		})
	}
}

func (ps *paymentLinksSuite) TestPaymentLinkService_Create() {
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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "profileId=prf_12312312&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/payment-links", c.handler)

			m, err := tClient.PaymentLinks.Create(c.args.ctx, c.args.paymentLink, c.args.options)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&PaymentLink{}, m)
			}
		})
	}
}

func (ps *paymentLinksSuite) TestPaymentLinkService_List() {
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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "testmode=true")

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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "profileId=pfl_11211&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/payment-links", c.handler)

			m, err := tClient.PaymentLinks.List(c.args.ctx, c.args.opts)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&PaymentLinksList{}, m)
			}
		})
	}
}

func TestPaymentLinksService(t *testing.T) {
	suite.Run(t, new(paymentLinksSuite))
}
