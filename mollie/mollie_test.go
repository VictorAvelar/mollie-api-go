package mollie

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v3/mollie/tools/idempotency"
	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	c := http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	conf := NewConfig(true, APITokenEnv)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.client, conf)
			assert.Nil(t, err)
		})
	}
}

func TestNewClientWithEnvVars(t *testing.T) {
	setEnv()
	setup()
	defer unsetEnv()
	defer teardown()

	c := http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.client, tConf)
			require.Nil(t, err)
			assert.NotEmpty(t, got.authentication)
		})
	}
}

func TestClient_NewAPIRequest(t *testing.T) {
	type args struct {
		ctx    context.Context
		method string
		uri    string
		body   interface{}
	}

	type testCtxKey string

	cases := []struct {
		name    string
		args    args
		outBody string
		outURI  string
		wantCtx bool
	}{
		{
			"request with empty context works as expected",
			args{
				ctx:    nil,
				method: http.MethodGet,
				uri:    "test",
				body:   []string{"hello", "world"},
			},
			`["hello","world"]` + "\n",
			"/test",
			false,
		},
		{
			"request with context works and the same as without context",
			args{
				ctx:    context.WithValue(context.Background(), testCtxKey("test-key"), "I will make it to the other side"),
				method: http.MethodGet,
				uri:    "test",
				body:   "some simple string",
			},
			"\"some simple string\"\n",
			"/test",
			true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setEnv()
			setup()
			defer teardown()
			defer unsetEnv()

			req, _ := tClient.NewAPIRequest(c.args.ctx, c.args.method, c.args.uri, c.args.body)

			testHeader(t, req, "Accept", RequestContentType)
			testHeader(t, req, AuthHeader, "Bearer token_X12b31ggg23")

			assert.Equal(t, tServer.URL+c.outURI, req.URL.String())
			body, _ := io.ReadAll(req.Body)
			assert.Equal(t, c.outBody, string(body))
		})
	}
}

func TestClient_NewApiRequest_IdempotencyKeys(t *testing.T) {
	type args struct {
		method string
	}

	tests := []struct {
		name   string
		args   args
		expect string
		dummy  bool
	}{
		{
			"using the std key generator",
			args{http.MethodPost},
			"",
			false,
		},
		{
			"using a nop key generator with the default text",
			args{http.MethodPost},
			"",
			true,
		},
		{
			"using a nop key generator with custom text",
			args{http.MethodPost},
			"testing_mollie_idem_key",
			true,
		},
		{
			"using the std key generator with a non supported method (get)",
			args{http.MethodGet},
			"",
			false,
		},
		{
			"using the std key generator with a non supported method (put)",
			args{http.MethodPut},
			"",
			false,
		},
		{
			"using the std key generator with a non supported method (delete)",
			args{http.MethodDelete},
			"",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv()
			setup()
			defer teardown()
			defer unsetEnv()

			if tt.dummy {
				g := idempotency.NewNopGenerator(tt.expect)
				if tt.expect == "" {
					tt.expect = idempotency.TestKeyExpected
				}
				tClient.SetIdempotencyKeyGenerator(g)
			}

			req, err := tClient.NewAPIRequest(context.Background(), tt.args.method, "/", nil)

			assert.Nil(t, err)

			if tt.args.method != http.MethodPost {
				assert.Empty(t, req.Header.Get(IdempotencyKeyHeader))
			} else {
				assert.NotEmpty(t, req.Header.Get(IdempotencyKeyHeader))
				if tt.dummy {
					testHeader(t, req, IdempotencyKeyHeader, tt.expect)
				} else {
					assert.Len(t, req.Header.Get(IdempotencyKeyHeader), 36)
					assert.Regexp(t, `(?m)^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`, req.Header.Get(IdempotencyKeyHeader))
				}
			}
		})
	}
}

func TestClient_NewAPIRequest_ForceErrors(t *testing.T) {
	type args struct {
		ctx    context.Context
		method string
		uri    string
		body   interface{}
	}

	noPre := func() error {
		return nil
	}

	cases := []struct {
		name string
		args args
		err  error
		pre  func() error
	}{
		{
			"err uri without trailing slash",
			args{
				context.Background(),
				http.MethodGet,
				"test",
				nil,
			},
			errBadBaseURL,
			func() error {
				uri, err := url.Parse("http://localhost")
				if err != nil {
					return err
				}
				tClient = &Client{
					BaseURL: uri,
				}
				return nil
			},
		},
		{
			"err parsing request uri",
			args{
				context.Background(),
				"\\\\\\",
				"test",
				nil,
			},
			fmt.Errorf("net/http: invalid method \"\\\\\\\\\\\\\""),
			noPre,
		},
		{
			"err serializing request body",
			args{
				context.Background(),
				http.MethodPost,
				"test",
				make(chan int),
			},
			fmt.Errorf("encoding_error: json: unsupported type: chan int"),
			noPre,
		},
		{
			"err when parsing requested uri",
			args{
				context.Background(),
				http.MethodPut,
				":",
				nil,
			},
			fmt.Errorf("url_parsing_error: parse \":\": missing protocol scheme"),
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()
		t.Run(c.name, func(t *testing.T) {
			err := c.pre()
			require.Nil(t, err)
			_, err = tClient.NewAPIRequest(c.args.ctx, c.args.method, c.args.uri, c.args.body)
			assert.NotNil(t, err)
			assert.EqualError(t, err, c.err.Error())
		})
	}
}

func TestClient_IsAccessToken(t *testing.T) {
	cases := []struct {
		name  string
		value string
		want  bool
	}{
		{
			"Test dummy_text",
			"dummy_text",
			false,
		},
		{
			"Test pattern is matched",
			"access_testing_token_here",
			true,
		},
		{
			"Test patter is not matched if not in the right position",
			"testing_token_access_here",
			false,
		},
		{
			"Test other tokens are not matched",
			"test_your_token_here",
			false,
		},
		{
			"Test empty is not a matched",
			"",
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			client := &Client{
				authentication: c.value,
			}

			got := client.HasAccessToken()
			assert.Equal(t, c.want, got)
		})
	}
}

func TestClient_NewAPIRequest_OrgTokenOverApiKey(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("org_token")
	req, _ := tClient.NewAPIRequest(context.TODO(), "GET", "test", nil)

	testHeader(t, req, AuthHeader, "Bearer org_token")
}

func TestClient_WithAuthenticationValue_Error(t *testing.T) {
	setup()
	defer teardown()
	err := tClient.WithAuthenticationValue("")

	if err == nil {
		t.Errorf("unexpected error, want %v and got %v", errEmptyAuthKey, err)
	}
}

func TestClient_Do(t *testing.T) {
	type args struct {
		ctx    context.Context
		method string
		uri    string
		body   interface{}
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		want    int
		handler http.HandlerFunc
		pre     func(r *http.Request)
	}{
		{
			"execute successful request",
			args{
				context.Background(),
				http.MethodGet,
				"/test",
				nil,
			},
			false,
			nil,
			http.StatusOK,
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				w.WriteHeader(http.StatusOK)
			},
			func(r *http.Request) {},
		},
		{
			"error when request url is invalid (nil)",
			args{
				context.Background(),
				http.MethodGet,
				"/test",
				nil,
			},
			true,
			fmt.Errorf("http_error: Get \"\": http: nil Request.URL"),
			http.StatusOK,
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				w.Write([]byte("nothing"))
			},
			func(r *http.Request) {
				r.URL = nil
			},
		},
	}

	for _, c := range cases {
		setEnv()
		setup()
		defer teardown()
		defer unsetEnv()
		t.Run(c.name, func(t *testing.T) {
			tMux.HandleFunc(c.args.uri, c.handler)
			req, _ := tClient.NewAPIRequest(c.args.ctx, c.args.method, c.args.uri, c.args.body)
			c.pre(req)
			res, err := tClient.Do(req)
			if c.wantErr {
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, c.want, res.StatusCode)
			}
		})
	}
}

func TestCheckResponse(t *testing.T) {
	res1 := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       ioutil.NopCloser(strings.NewReader("not found ok")),
	}

	res3 := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}

	res2 := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(strings.NewReader("success ok")),
	}

	tests := []struct {
		name string
		code string
		arg  *Response
	}{
		{
			"successful response",
			"",
			&Response{Response: res2},
		},
		{
			"not found response",
			"Not Found",
			&Response{Response: res1},
		},
		{
			"success with empty body",
			"",
			&Response{Response: res3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckResponse(tt.arg); err != nil {
				if !strings.Contains(err.Error(), tt.code) {
					t.Error(err)
				}
			}
		})
	}
}

// <----- Testing helpers ----->

var (
	tMux    *http.ServeMux
	tServer *httptest.Server
	tClient *Client
	tConf   *Config
)

var (
	noPre    = func() {}
	crashSrv = func() {
		u, _ := url.Parse(tServer.URL)
		tClient.BaseURL = u
	}
	setAccessToken = func() {
		tClient.WithAuthenticationValue("access_token_test")
	}
)

// the parameter indicates if you want to prepare your tests against the US sandbox
// just to be used when doing integration testing.
func setup() {
	tMux = http.NewServeMux()
	tServer = httptest.NewServer(tMux)
	tConf = NewAPITestingConfig(true)
	tClient, _ = NewClient(nil, tConf)
	u, _ := url.Parse(tServer.URL + "/")
	tClient.BaseURL = u
}

func teardown() {
	tServer.Close()
}

func setEnv() {
	_ = os.Setenv(APITokenEnv, "token_X12b31ggg23")
	_ = os.Setenv(OrgTokenEnv, "access_ey1923n23123n1k3b123jv12g312h31v32g13")
}

func unsetEnv() {
	_ = os.Unsetenv(APITokenEnv)
	_ = os.Unsetenv(OrgTokenEnv)
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testQuery(t *testing.T, r *http.Request, want string) {
	if r.URL.Query().Encode() != want {
		t.Errorf("Query().Encode() returned unexpected values, want: %q, got %q", want, r.URL.Query().Encode())
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(testdata.InternalServerErrorResponse))
}

func encodingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{hello: [{},]}`))
}

// <----- .Testing helpers ----->
