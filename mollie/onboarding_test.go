package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type onboardingServiceSuite struct{ suite.Suite }

func (os *onboardingServiceSuite) SetupSuite() { setEnv() }

func (os *onboardingServiceSuite) TearDownSuite() { unsetEnv() }

func (os *onboardingServiceSuite) TestOnboardingService_GetOnboardingStatus() {
	cases := []struct {
		name    string
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get onboarding status works as expected.",
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOnboardingStatusResponse))
			},
		},
		{
			"get onboarding status, an error is returned from the server",
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get onboarding status, an error occurs when parsing json",
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get onboarding status, invalid url when building request",
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/onboarding/me", c.handler)

			res, m, err := tClient.Onboarding.GetOnboardingStatus(context.Background())
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Onboarding{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (os *onboardingServiceSuite) TestOnboardingService_SubmitOnboardingData() {
	cases := []struct {
		name    string
		data    *OnboardingData
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get onboarding status works as expected.",
			&OnboardingData{},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "POST")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOnboardingStatusResponse))
			},
		},
		{
			"get onboarding status, an error is returned from the server",
			&OnboardingData{},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get onboarding status, invalid url when building request",
			&OnboardingData{},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/onboarding/me", c.handler)

			res, err := tClient.Onboarding.SubmitOnboardingData(context.Background(), c.data)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}
func TestOnboardingService(t *testing.T) {
	suite.Run(t, new(onboardingServiceSuite))
}
