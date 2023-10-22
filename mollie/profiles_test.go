package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/suite"
)

type profilesServiceSuite struct{ suite.Suite }

func (ps *profilesServiceSuite) SetupSuite() { setEnv() }

func (ps *profilesServiceSuite) TearDownSuite() { unsetEnv() }

func (ps *profilesServiceSuite) TestProfilesService_Get() {
	type args struct {
		ctx     context.Context
		profile string
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
			"get profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"get profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get profile, an error occurs when parsing json",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s", c.args.profile), c.handler)

			res, m, err := tClient.Profiles.Get(c.args.ctx, c.args.profile)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Profile{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_GetCurrent() {
	cases := []struct {
		name    string
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get current profile works as expected.",
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"get current profile, an error is returned from the server",
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get current profile, an error occurs when parsing json",
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get current profile, invalid url when building request",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s", "me"), c.handler)

			res, m, err := tClient.Profiles.Current(context.Background())
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Profile{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_List() {
	type args struct {
		ctx     context.Context
		options *ProfileListOptions
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
			"list profiles works as expected.",
			args{
				context.Background(),
				&ProfileListOptions{},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfilesListResponse))
			},
		},
		{
			"list profiles with options, works as expected.",
			args{
				context.Background(),
				&ProfileListOptions{
					Limit: 100,
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "limit=100&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfilesListResponse))
			},
		},
		{
			"list profiles, an error is returned from the server",
			args{
				context.Background(),
				&ProfileListOptions{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"list profiles, an error occurs when parsing json",
			args{
				context.Background(),
				&ProfileListOptions{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list profiles, invalid url when building request",
			args{
				context.Background(),
				&ProfileListOptions{},
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
			tMux.HandleFunc("/v2/profiles", c.handler)

			res, m, err := tClient.Profiles.List(c.args.ctx, c.args.options)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&ProfileList{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_Create() {
	type args struct {
		ctx     context.Context
		profile *Profile
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
			"create profile works as expected.",
			args{
				context.Background(),
				&Profile{
					Name: "testing name",
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"create profile, an error is returned from the server",
			args{
				context.Background(),
				&Profile{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"create profile, an error occurs when parsing json",
			args{
				context.Background(),
				&Profile{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create profile, invalid url when building request",
			args{
				context.Background(),
				&Profile{},
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
			tMux.HandleFunc("/v2/profiles", c.handler)

			res, m, err := tClient.Profiles.Create(c.args.ctx, c.args.profile)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Profile{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_Update() {
	type args struct {
		ctx       context.Context
		profileID string
		profile   *Profile
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
			"update profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				&Profile{
					Name: "testing name",
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "PATCH")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"update profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				&Profile{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"update profile, an error occurs when parsing json",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				&Profile{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				&Profile{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s", c.args.profileID), c.handler)

			res, m, err := tClient.Profiles.Update(c.args.ctx, c.args.profileID, c.args.profile)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Profile{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_Delete() {
	type args struct {
		ctx     context.Context
		profile string
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
			"delete profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "DELETE")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"delete profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"delete profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s", c.args.profile), c.handler)

			res, err := tClient.Profiles.Delete(c.args.ctx, c.args.profile)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_EnablePaymentMethod() {
	type args struct {
		ctx     context.Context
		profile string
		method  PaymentMethod
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
			"enable payment method for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.EnablePaymentMethodResponse))
			},
		},
		{
			"enable payment method for profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"enable payment method for profile, an error occurs when parsing json",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"enable payment method for profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s/methods/%s", c.args.profile, c.args.method), c.handler)

			res, m, err := tClient.Profiles.EnablePaymentMethod(c.args.ctx, c.args.profile, c.args.method)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&PaymentMethodDetails{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_DisablePaymentMethod() {
	type args struct {
		ctx     context.Context
		profile string
		method  PaymentMethod
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
			"disable payment method for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "DELETE")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
			},
		},
		{
			"disable payment method for profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"disable payment method for profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s/methods/%s", c.args.profile, c.args.method), c.handler)

			res, err := tClient.Profiles.DisablePaymentMethod(c.args.ctx, c.args.profile, c.args.method)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_EnableGiftCardIssuer() {
	type args struct {
		ctx     context.Context
		profile string
		issuer  GiftCardIssuer
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
			"enable gifcard issuer for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.EnablePaymentMethodResponse))
			},
		},
		{
			"enable gifcard issuer for profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"enable gifcard issuer for profile, an error occurs when parsing json",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"enable gifcard issuer for profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s/methods/giftcard/issuers/%s", c.args.profile, c.args.issuer), c.handler)

			res, m, err := tClient.Profiles.EnableGiftCardIssuer(c.args.ctx, c.args.profile, c.args.issuer)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&GiftCardEnabled{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_DisableGiftCardIssuer() {
	type args struct {
		ctx     context.Context
		profile string
		issuer  GiftCardIssuer
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
			"disable giftcard issuer for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "DELETE")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
			},
		},
		{
			"disable giftcard issuer for profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"disable giftcard issuer for profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s/methods/giftcard/issuers/%s", c.args.profile, c.args.issuer), c.handler)

			res, err := tClient.Profiles.DisableGiftCardIssuer(c.args.ctx, c.args.profile, c.args.issuer)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_EnableGiftCardIssuerForCurrent() {
	type args struct {
		ctx    context.Context
		issuer GiftCardIssuer
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
			"enable gifcard issuer for profile works as expected.",
			args{
				context.Background(),
				Good4fun,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.EnableGiftCardIssuerResponse))
			},
		},
		{
			"enable gifcard issuer for profile, an error is returned from the server",
			args{
				context.Background(),
				Good4fun,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"enable gifcard issuer for profile, an error occurs when parsing json",
			args{
				context.Background(),
				Good4fun,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"enable gifcard issuer for profile, invalid url when building request",
			args{
				context.Background(),
				Good4fun,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/me/methods/giftcard/issuers/%s", c.args.issuer), c.handler)

			res, m, err := tClient.Profiles.EnableGiftCardIssuerForCurrent(c.args.ctx, c.args.issuer)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&GiftCardEnabled{}, m)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ps *profilesServiceSuite) TestProfilesService_DisableGiftCardIssuerForCurrent() {
	type args struct {
		ctx    context.Context
		issuer GiftCardIssuer
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
			"disable giftcard issuer for profile works as expected.",
			args{
				context.Background(),
				Good4fun,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(ps.T(), r, "DELETE")
				testQuery(ps.T(), r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
			},
		},
		{
			"disable giftcard issuer for profile, an error is returned from the server",
			args{
				context.Background(),
				Good4fun,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"disable giftcard issuer for profile, invalid url when building request",
			args{
				context.Background(),
				Good4fun,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/me/methods/giftcard/issuers/%s", c.args.issuer), c.handler)

			res, err := tClient.Profiles.DisableGiftCardIssuerForCurrent(c.args.ctx, c.args.issuer)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService(t *testing.T) {
	suite.Run(t, new(profilesServiceSuite))
}
