package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPermissionsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		permission PermissionGrant
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
			"get permission works as expected.",
			args{
				context.Background(),
				PaymentsRead,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPermissionsResponse))
			},
		},
		{
			"get permission, an error is returned from the server",
			args{
				context.Background(),
				PaymentsWrite,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get permission, an error occurs when parsing json",
			args{
				context.Background(),
				PaymentsWrite,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get permission, invalid url when building request",
			args{
				context.Background(),
				PaymentsWrite,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/permissions/%s", c.args.permission), c.handler)

			res, m, err := tClient.Permissions.Get(c.args.ctx, c.args.permission)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Permission{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestPermissionsService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
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
			"get permission works as expected.",
			args{
				context.Background(),
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPermissionsResponse))
			},
		},
		{
			"get permission, an error is returned from the server",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get permission, an error occurs when parsing json",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get permission, invalid url when building request",
			args{
				context.Background(),
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	setEnv()
	defer unsetEnv()
	for _, c := range cases {
		setup()
		defer teardown()
		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/permissions", c.handler)

			res, m, err := tClient.Permissions.List(c.args.ctx)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PermissionsList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
