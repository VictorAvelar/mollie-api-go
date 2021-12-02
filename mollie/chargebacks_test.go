package mollie

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type chargebacksSuite struct{ suite.Suite }

func (cbs *chargebacksSuite) SetupSuite() { setEnv() }

func (cbs *chargebacksSuite) TearDownSuite() { unsetEnv() }

func (cbs *chargebacksSuite) TestChargebacksService_Get() {
	type args struct {
		ctx        context.Context
		payment    string
		chargeback string
		options    *ChargebackOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"get chargebacks",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"chb_n9z0tp",
				&ChargebackOptions{
					Include: "details.qrCode",
				},
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cbs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cbs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.GetChargebackResponse))
			},
			noPre,
		},
		{
			"get chargebacks returns an http error from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"chb_n9z0tp",
				&ChargebackOptions{
					Include: "details.qrCode",
				},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			errorHandler,
			noPre,
		},
		{
			"get chargebacks returns an error when creating the request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"chb_n9z0tp",
				&ChargebackOptions{
					Include: "details.qrCode",
				},
			},
			true,
			errBadBaseURL,
			errorHandler,
			func() {
				u, _ := url.Parse(tServer.URL)
				tClient.BaseURL = u
			},
		},
		{
			"get chargebacks returns an error when trying to parse the json response",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"chb_n9z0tp",
				&ChargebackOptions{
					Include: "details.qrCode",
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()
		cbs.T().Run(c.name, func(t *testing.T) {
			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/payments/%s/chargebacks/%s",
					c.args.payment,
					c.args.chargeback,
				),
				c.handler,
			)
			c.pre()

			res, err := tClient.Chargebacks.Get(c.args.ctx, c.args.payment, c.args.chargeback, c.args.options)
			if c.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				require.Nil(t, err)
				assert.Equal(t, res.ID, c.args.chargeback)
			}
		})
	}
}

func (cbs *chargebacksSuite) TestChargebacksService_List() {
	type args struct {
		ctx     context.Context
		options *ListChargebackOptions
	}

	noPre := func() {}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"list chargebacks",
			args{
				context.Background(),
				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cbs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cbs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
			},
			noPre,
		},
		{
			"list chargebacks with options",
			args{
				context.Background(),
				&ListChargebackOptions{
					ProfileID: "pfl_QkEhN94Ba",
				},
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cbs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cbs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
			},
			noPre,
		},
		{
			"list chargebacks return an http error from the remote server",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			errorHandler,
			noPre,
		},
		{
			"list chargebacks return an error when creating the requests",
			args{
				context.Background(),
				nil,
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"list chargebacks return an error when parsing the json response",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()
		cbs.T().Run(c.name, func(t *testing.T) {
			tMux.HandleFunc("/v2/chargebacks", c.handler)

			c.pre()
			res, err := tClient.Chargebacks.List(c.args.ctx, c.args.options)
			if c.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				require.Nil(t, err)
				assert.NotZero(t, res.Count)
			}
		})
	}
}

func (cbs *chargebacksSuite) TestChargebacksService_ListForPayment() {
	type args struct {
		ctx     context.Context
		payment string
		options *ListChargebackOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"list chargebacks attached to a payment",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cbs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cbs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
			},
			noPre,
		},
		{
			"list chargebacks attached to a payment, with options",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				&ListChargebackOptions{
					ProfileID: "pfl_QkEhN94Ba",
				},
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cbs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cbs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListChargebacksResponse))
			},
			noPre,
		},
		{
			"list chargebacks returns an http error from the remote server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			errorHandler,
			noPre,
		},
		{
			"list chargebacks returns an error when building the request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				nil,
			},
			true,
			errBadBaseURL,
			func(rw http.ResponseWriter, r *http.Request) {},
			crashSrv,
		},
		{
			"list chargebacks returns an error when parsing the json response",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()
		cbs.T().Run(c.name, func(t *testing.T) {
			tMux.HandleFunc(
				fmt.Sprintf("/v2/payments/%s/chargebacks", c.args.payment),
				c.handler,
			)

			c.pre()

			res, err := tClient.Chargebacks.ListForPayment(c.args.ctx, c.args.payment, c.args.options)
			if c.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				require.Nil(t, err)
				assert.NotZero(t, res.Count)
				assert.IsType(t, &ChargebackList{}, res)
			}
		})
	}
}

func TestChargebacksService(t *testing.T) {
	suite.Run(t, new(chargebacksSuite))
}
