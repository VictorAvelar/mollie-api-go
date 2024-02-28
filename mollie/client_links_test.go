package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCreateClientLink(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
		cd  *ClientDetails
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
			"create new client link",
			args{
				context.Background(),
				&ClientDetails{},
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.CreateClientLinkResponse))
			},
			noPre,
		},
		{
			"create client link, an error is returned from the server",
			args{
				context.Background(),
				&ClientDetails{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"create client link, an error occurs when parsing json",
			args{
				context.Background(),
				&ClientDetails{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
		{
			"create client link, invalid url when building request",
			args{
				context.Background(),
				&ClientDetails{},
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()
		t.Run(c.name, func(t *testing.T) {
			tMux.HandleFunc(
				"/v2/client-links",
				c.handler,
			)
			c.pre()

			res, cb, err := tClient.ClientLinks.CreateClientLink(c.args.ctx, c.args.cd)
			if c.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &ClientLink{}, cb)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestCreateFinalizeClientLink(t *testing.T) {
	type args struct {
		ctx        context.Context
		clientLink string
		options    *ClientLinkFinalizeOptions
	}
	tests := []struct {
		name              string
		args              args
		wantClientLinkURI string
	}{
		{
			"constructs client link finalize step correctly.",
			args{
				context.Background(),
				"https://my.mollie.com/dashboard/client-link/finalize/csr_vZCnNQsV2UtfXxYifWKWH",
				&ClientLinkFinalizeOptions{
					ClientID: "app_j9Pakf56Ajta6Y65AkdTtAv",
					State:    "decafbad",
					Scope:    "onboarding.read+organization.read+payments.write+payments.read+profiles.write",
				},
			},
			"https://my.mollie.com/dashboard/client-link/finalize/csr_vZCnNQsV2UtfXxYifWKWH?clientID=app_j9Pakf56Ajta6Y65AkdTtAv&scope=onboarding.read%2Borganization.read%2Bpayments.write%2Bpayments.read%2Bprofiles.write&state=decafbad",
		},
		{
			"constructs client link finalize with complex values",
			args{
				context.Background(),
				"https://my.mollie.com/dashboard/client-link/finalize/csr_vZCnNQsV2UtfXxYifWKWH",
				&ClientLinkFinalizeOptions{
					ClientID: "",
					State:    "\ns\\s\\s\\s\n",
					Scope:    "",
				},
			},
			"https://my.mollie.com/dashboard/client-link/finalize/csr_vZCnNQsV2UtfXxYifWKWH?state=%0As%5Cs%5Cs%5Cs%0A",
		},
		{
			"constructs client link finalize with complex values",
			args{
				context.Background(),
				"https://my.mollie.com/dashboard/client-link/finalize/csr_vZCnNQsV2UtfXxYifWKWH",
				&ClientLinkFinalizeOptions{},
			},
			"https://my.mollie.com/dashboard/client-link/finalize/csr_vZCnNQsV2UtfXxYifWKWH?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClientLinkURI := tClient.ClientLinks.CreateFinalizeClientLink(tt.args.ctx, tt.args.clientLink, tt.args.options)
			if gotClientLinkURI != tt.wantClientLinkURI {
				t.Errorf("ClientLinksService.CreateFinalizeClientLink() = %v, want %v", gotClientLinkURI, tt.wantClientLinkURI)
			}
		})
	}
}
