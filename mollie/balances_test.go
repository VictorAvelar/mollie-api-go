package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type balancesServiceSuite struct{ suite.Suite }

func (bs *balancesServiceSuite) SetupSuite() { setEnv() }

func (bs *balancesServiceSuite) TearDownSuite() { unsetEnv() }

func (bs *balancesServiceSuite) TestBalancesService_Get() {
	type args struct {
		ctx     context.Context
		balance string
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
			"get balance by id works as expected",
			args{
				context.Background(),
				"test_balance",
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(bs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(bs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.GetBalanceResponse))
			},
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		bs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/balances/%s",
					c.args.balance,
				),
				c.handler,
			)

			res, capture, err := tClient.Balances.Get(c.args.ctx, c.args.balance)
			if c.wantErr {
				bs.NotNil(err)
				bs.EqualError(err, c.err.Error())
			} else {
				bs.Nil(err)
				bs.IsType(&Balance{}, capture)
				bs.Same(c.args.ctx, res.Request.Context())
				bs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestBalancesService(t *testing.T) {
	suite.Run(t, new(balancesServiceSuite))
}
