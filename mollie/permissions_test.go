package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type permissionsServiceSuite struct{ suite.Suite }

func (os *permissionsServiceSuite) SetupSuite() { setEnv() }

func (os *permissionsServiceSuite) TearDownSuite() { unsetEnv() }

func (os *permissionsServiceSuite) TestPermissionsService_Get() { unsetEnv() }

func (ps *profilesServiceSuite) TestPermissionsService_Get() {
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
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/permissions/%s", c.args.permission), c.handler)

			res, m, err := tClient.Permissions.Get(c.args.ctx, c.args.permission)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Permission{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestPermissionsService_List() {
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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

	for _, c := range cases {
		setup()
		defer teardown()

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/permissions", c.handler)

			res, m, err := tClient.Permissions.List(c.args.ctx)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&PermissionsList{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestPermissionService(t *testing.T) {
	suite.Run(t, new(permissionsServiceSuite))
}
