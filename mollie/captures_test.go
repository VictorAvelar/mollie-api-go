package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type capturesServiceSuite struct{ suite.Suite }

func (cs *capturesServiceSuite) SetupSuite() { setEnv() }

func (cs *capturesServiceSuite) TearDownSuite() { unsetEnv() }

func (cs *capturesServiceSuite) TestCapturesService_Get() {
	type args struct {
		ctx     context.Context
		payment string
		capture string
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
			"get captures works as expected",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.GetCaptureResponse))
			},
			noPre,
		},
		{
			"get captures returns an http error from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			errorHandler,
			noPre,
		},
		{
			"get captures returns an error when creating the request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"get captures returns an error when trying to parse the json response",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
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

		cs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/payments/%s/captures/%s",
					c.args.payment,
					c.args.capture,
				),
				c.handler,
			)

			res, capture, err := tClient.Captures.Get(c.args.ctx, c.args.payment, c.args.capture)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
				cs.IsType(&Capture{}, capture)
				cs.Same(c.args.ctx, res.Request.Context())
				cs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (cs *capturesServiceSuite) TestCapturesService_List() {
	type args struct {
		ctx     context.Context
		payment string
		capture string
	}

	type key string

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"list captures works as expected",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListCapturesResponse))
			},
			noPre,
		},
		{
			"list captures returns an http error from the server",
			args{
				context.WithValue(context.Background(), key("test"), "test-value"),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			errorHandler,
			noPre,
		},
		{
			"list captures returns an error when creating the request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"list captures returns an error when trying to parse the json response",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
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

		cs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/payments/%s/captures",
					c.args.payment,
				),
				c.handler,
			)

			res, list, err := tClient.Captures.List(c.args.ctx, c.args.payment)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
				cs.IsType(&CapturesList{}, list)
				cs.Same(c.args.ctx, res.Request.Context())
				cs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestCapturesService(t *testing.T) {
	suite.Run(t, new(capturesServiceSuite))
}
