package mollie

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestOnboardingService_GetOnboardingStatus(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		teardown()
		unsetEnv()
	}()

	tMux.HandleFunc("/"+onboardingTarget, func(rw http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			rw.WriteHeader(http.StatusUnauthorized)
		}

		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(testdata.GetOnboardingStatusResponse))
	})

	res, err := tClient.Onboarding.GetOnboardingStatus(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if res.Status != CompletedOnboardingStatus {
		t.Error("unexpected onboarding status")
	}
}

func TestOnboardingService_SubmitOnboardingData(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		teardown()
		unsetEnv()
	}()

	tMux.HandleFunc("/"+onboardingTarget, func(rw http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			rw.WriteHeader(http.StatusUnauthorized)
		}

		rw.WriteHeader(http.StatusCreated)
	})

	od := OnboardingData{}
	od.Organization.Name = "Testing Org. B.V."

	err := tClient.Onboarding.SubmitOnboardingData(context.TODO(), &od)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOnboardingService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/"+onboardingTarget, errorHandler)

	_, gerr := tClient.Onboarding.GetOnboardingStatus(context.TODO())
	gcerr := tClient.Onboarding.SubmitOnboardingData(context.TODO(), &OnboardingData{})

	tests := []error{gerr, gcerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Fail()
		}
	}
}

func TestOnboardingService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/"+onboardingTarget, errorHandler)

	_, gerr := tClient.Onboarding.GetOnboardingStatus(context.TODO())
	gcerr := tClient.Onboarding.SubmitOnboardingData(context.TODO(), &OnboardingData{})

	tests := []error{gerr, gcerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestOnboardingService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/"+onboardingTarget, encodingHandler)

	_, gerr := tClient.Onboarding.GetOnboardingStatus(context.TODO())

	tests := []error{gerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}
