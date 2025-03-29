package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestMiscellaneousService_ApplePaymentSession(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx       context.Context
		appleSess *ApplePaymentSessionRequest
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
			"get apple payment session works as expected.",
			args{
				context.Background(),
				&ApplePaymentSessionRequest{
					Domain: "https://example.com",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListMethodsResponse))
			},
		},
		{
			"get apple payment session, an error is returned from the server",
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
			"get apple payment session, an error occurs when parsing json",
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
			"get apple payment session, invalid url when building request",
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
			tMux.HandleFunc("/v2/wallets/applepay/sessions", c.handler)

			res, m, err := tClient.Wallets.ApplePaymentSession(c.args.ctx, c.args.appleSess)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &ApplePaymentSession{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
