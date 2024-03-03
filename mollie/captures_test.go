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
		options *CaptureOptions
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
				nil,
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
			"get captures works expands query params correctly",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"cpt_4qqhO89gsT",
				&CaptureOptions{
					Embed: []EmbedValue{EmbedPayment},
				},
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "embed=payments")
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
				nil,
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
				nil,
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

			res, capture, err := tClient.Captures.Get(c.args.ctx, c.args.payment, c.args.capture, c.args.options)
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

func TestCapturesService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment string
		capture CreateCapture
		options *CaptureOptions
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
			"create captures works as expected",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				CreateCapture{
					Amount: &Amount{
						Value:    "20.00",
						Currency: "EUR",
					},
					Description: "Order #12345",
				},
				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.CreateCaptureResponse))
			},
			noPre,
		},
		{
			"create captures works as with access token expected",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				CreateCapture{
					Amount: &Amount{
						Value:    "20.00",
						Currency: "EUR",
					},
					Description: "Order #12345",
				},
				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.CreateCaptureWithAccessTokenResponse))
			},
			setAccessToken,
		},
		{
			"create captures returns an http error from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				CreateCapture{},
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"create captures returns an error when creating the request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				CreateCapture{},
				nil,
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"create captures returns an error when trying to parse the json response",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				CreateCapture{},
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

		t.Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/payments/%s/captures",
					c.args.payment,
				),
				c.handler,
			)

			res, capture, err := tClient.Captures.Create(c.args.ctx, c.args.payment, c.args.capture)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
				if tClient.HasAccessToken() {
					assert.True(t, capture.Testmode)
				}
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
		options *CaptureOptions
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
				&CaptureOptions{
					Embed: []EmbedValue{EmbedPayment},
				},
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "embed=payments")

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
				&CaptureOptions{},
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
				&CaptureOptions{},
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
				&CaptureOptions{},
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

			res, list, err := tClient.Captures.List(c.args.ctx, c.args.payment, c.args.options)
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
