package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCapturesService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
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
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Capture{}, capture)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestCapturesService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
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
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
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

		t.Run(c.name, func(t *testing.T) {
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
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &CapturesList{}, list)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
