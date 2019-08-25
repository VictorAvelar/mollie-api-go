package mollie

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
)

// ---------- testing utilities ----------
var (
	testMux    *http.ServeMux
	testClient *APIClient
	testServer *httptest.Server
)

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	testClient, _ = NewClient(nil, nil, testServer.URL)
}

func teardown() {
	testServer.Close()
}

func setEnv() {
	_ = os.Setenv(APITokenEnv, "token_X12b31ggg23")
	_ = os.Setenv(OrgTokenEnv, "ey1923n23123n1k3b123jv12g312h31v32g13")
}

func unsetEnv() {
	_ = os.Unsetenv(APITokenEnv)
	_ = os.Unsetenv(OrgTokenEnv)
}

// ---------- .testing utilities ----------

func TestNewClient(t *testing.T) {

	type args struct {
		ctx context.Context
		cl  httpClient
		uri string
	}

	u, err := url.Parse("http://localhost")
	if err != nil {
		t.Fail()
	}
	want := &APIClient{
		Ctx:               context.Background(),
		HTTPClient:        http.DefaultClient,
		BaseURL:           u,
		APIKey:            "",
		OrganizationToken: "",
	}

	tests := []struct {
		name      string
		check     *APIClient
		arguments args
		wantErr   bool
		err       error
	}{
		{
			name:  "test client is build properly",
			check: want,
			arguments: args{
				ctx: nil,
				cl:  nil,
				uri: "http://localhost",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:  "test url parser breaks",
			check: want,
			arguments: args{
				ctx: nil,
				cl:  nil,
				uri: " http:/:/localhost",
			},
			wantErr: true,
			err:     errors.New("parse  http:/:/localhost: first path segment in URL cannot contain colon"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.arguments.ctx, tt.arguments.cl, tt.arguments.uri)
			if tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Fatalf("failed while initializing the client: got: %v, want: %v", err, tt.err)
				}
			} else {
				if !reflect.DeepEqual(got, tt.check) {
					t.Fatalf("failed while initializing the client: got: %v, want: %v", got, tt.check)
				}
			}
		})
	}
}

func TestNewClient_WithEnvVars(t *testing.T) {
	setEnv()
	defer unsetEnv()
	u, err := url.Parse("http://localhost:3000")
	if err != nil {
		t.Fail()
	}
	want := &APIClient{
		Ctx:               context.Background(),
		HTTPClient:        http.DefaultClient,
		BaseURL:           u,
		APIKey:            "token_X12b31ggg23",
		OrganizationToken: "ey1923n23123n1k3b123jv12g312h31v32g13",
	}
	got, _ := NewClient(nil, nil, "http://localhost:3000")

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("client initialization error, got: %v, want %v", got, want)
	}
}

func TestWithAPIToken(t *testing.T) {
	setup()
	defer teardown()
	tests := []struct {
		name       string
		check      string
		wantErr    bool
		err        error
		afterTest  func()
		beforeTest func()
	}{
		{
			name:    "test an empty string returns an error",
			check:   "",
			wantErr: true,
			err:     errEmptyAPIKey,
		},
		{
			name:    "test a valid test token is set",
			check:   "test_$DJAO@A##MKao23u#N",
			wantErr: false,
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testClient.WithAPIKey(tt.check)
			if tt.wantErr {
				if !reflect.DeepEqual(got, tt.err) {
					t.Fatalf("failed while attaching API key, got: %v, want: %v", got, tt.err)
				}
			} else {
				if testClient.APIKey != tt.check {
					t.Fail()
				}
			}
		})
	}
}

func TestWithOrganizationToken(t *testing.T) {
	setup()
	defer teardown()
	tests := []struct {
		name       string
		check      string
		wantErr    bool
		err        error
		afterTest  func()
		beforeTest func()
	}{
		{
			name:    "test an empty string returns an error",
			check:   "",
			wantErr: true,
			err:     errEmptyAPIKey,
		},
		{
			name:    "test a valid test token is set",
			check:   "styduasjkdlmqwhdbw",
			wantErr: false,
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testClient.WithOrganizationToken(tt.check)
			if tt.wantErr {
				if !reflect.DeepEqual(got, tt.err) {
					t.Fatalf("failed while attaching API key, got: %v, want: %v", got, tt.err)
				}
			} else {
				if testClient.OrganizationToken != tt.check {
					t.Fail()
				}
			}
		})
	}
}

func TestNewAPIRequest_WellformedRequest(t *testing.T) {
	setup()
	defer teardown()

	inURI, outURI := "/testing", testServer.URL+"/v2/testing"
	inMethod, outMethod := "", http.MethodGet

	got, err := testClient.NewAPIRequest(inMethod, inURI, nil, false)
	if err != nil {
		t.Fatalf("failed while building the API Request: %v", err)
	}

	if got.URL.String() != outURI {
		t.Fatalf("mailformed request uri: want: %v, got: %v", outURI, got.URL)
	}

	if got.Method != outMethod {
		t.Fatalf("method assignment failed: got: %v, want: %v", got.Method, outMethod)
	}
}

func TestNewAPIRequest_WithAPIKeyAuthHeader(t *testing.T) {
	setup()
	defer teardown()

	var testKey = "test_demoAPIKey1234"

	testClient.WithAPIKey(testKey)

	req, err := testClient.NewAPIRequest("", "/testing", nil, true)
	if err != nil {
		t.Fatalf("failed while building the API Request: %v", err)
	}
	got := req.Header.Get(AuthHeader)
	want := "Bearer " + testKey

	if got != want {
		t.Fatalf("API Key header is not set: got %v, want %v", got, want)
	}
}

func TestNewAPIRequest_WithOrgTokenAuthHeader(t *testing.T) {
	setup()
	defer teardown()

	var apiKey = "test_demoAPIKey1234"
	var orgToken = "token_demoOrgToken"

	// Accoring to the token hierarchy Org Token should be used.
	testClient.WithOrganizationToken(orgToken)
	testClient.WithAPIKey(apiKey)

	req, err := testClient.NewAPIRequest("", "/testing", nil, true)
	if err != nil {
		t.Fatalf("failed while building the API Request: %v", err)
	}
	got := req.Header.Get(AuthHeader)
	want := "Bearer " + orgToken

	if got != want {
		t.Fatalf("API Key header is not set: got %v, want %v", got, want)
	}
}

func TestNewAPIRequest_AuthError(t *testing.T) {
	setup()
	defer teardown()

	_, err := testClient.NewAPIRequest("", "/testing", nil, true)

	if err != errEmptyAPIKey {
		t.Fatalf("mismatching error message: got %v, want %v", err, errEmptyAPIKey)
	}
}
