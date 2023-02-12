package mollie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_ToggleTesting(t *testing.T) {
	c := NewAPITestingConfig(false)

	assert.True(t, c.testing)
	c.ToggleTesting()
	assert.False(t, c.testing)
}

func TestConfig_ToggleIdempotency(t *testing.T) {
	c := NewAPITestingConfig(false)

	assert.False(t, c.reqIdempotency)
	assert.True(t, c.ToggleIdempotency())
	assert.True(t, c.reqIdempotency)
}

func TestConfig_SwitchAuthStrategy(t *testing.T) {
	c := NewAPITestingConfig(false)

	assert.Equal(t, APITokenEnv, c.auth)
	c.SwitchAuthStrategy(OrgTokenEnv)
	assert.Equal(t, OrgTokenEnv, c.auth)
}

func TestNewConfig(t *testing.T) {
	type args struct {
		t    bool
		auth string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			"config set to testing and with API token",
			args{
				t:    true,
				auth: APITokenEnv,
			},
			&Config{
				testing: true,
				auth:    "MOLLIE_API_TOKEN",
			},
		},
		{
			"config set to testing and with ORG token",
			args{
				t:    true,
				auth: OrgTokenEnv,
			},
			&Config{
				testing: true,
				auth:    "MOLLIE_ORG_TOKEN",
			},
		},
		{
			"config set to production and with ORG token",
			args{
				t:    true,
				auth: OrgTokenEnv,
			},
			&Config{
				testing: true,
				auth:    "MOLLIE_ORG_TOKEN",
			},
		},
		{
			"config set to production and with API token",
			args{
				t:    true,
				auth: APITokenEnv,
			},
			&Config{
				testing: true,
				auth:    "MOLLIE_API_TOKEN",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewConfig(tt.args.t, tt.args.auth)
			assert.Equal(t, got, tt.want)
		})
	}
}

// ExampleNewConfig demonstrates how to initialize a Config
// struct with the specified values for testing and token source.
func ExampleNewConfig() {
	conf := NewConfig(true, APITokenEnv)
	fmt.Printf(
		"testing config, testing: %v, req_idempotency: %v, token source: %s.",
		conf.testing,
		conf.reqIdempotency,
		conf.auth,
	)
	// Output: testing config, testing: true, req_idempotency: false, token source: MOLLIE_API_TOKEN.
}

// ExampleNewAPITestingConfig demonstrates how to initialize a Config
// struct with testing mode enabled, token source from the default API
// token env variable (MOLLIE_API_TOKEN) and request idempotency feature
// enabled.
func ExampleNewAPITestingConfig() {
	conf := NewAPITestingConfig(true)
	fmt.Printf(
		"testing api config, testing: %v, req_idempotency: %v, token source: %s.",
		conf.testing,
		conf.reqIdempotency,
		conf.auth,
	)
	// Output: testing api config, testing: true, req_idempotency: true, token source: MOLLIE_API_TOKEN.
}

// ExampleNewOrgTestingConfig demonstrates how to initialize a Config
// struct with testing mode enabled, token source from the default Org
// token env variable (MOLLIE_ORG_TOKEN) and request idempotency feature
// enabled.
func ExampleNewOrgTestingConfig() {
	conf := NewOrgTestingConfig(true)
	fmt.Printf(
		"testing org config, testing: %v, req_idempotency: %v, token source: %s.",
		conf.testing,
		conf.reqIdempotency,
		conf.auth,
	)
	// Output: testing org config, testing: true, req_idempotency: true, token source: MOLLIE_ORG_TOKEN.
}

// ExampleNewAPIConfig demonstrates how to initialize a Config
// struct with testing mode disabled, token source from the default API
// token env variable (MOLLIE_API_TOKEN) and request idempotency feature
// enabled.
func ExampleNewAPIConfig() {
	conf := NewAPIConfig(true)
	fmt.Printf(
		"testing api config, testing: %v, req_idempotency: %v, token source: %s.",
		conf.testing,
		conf.reqIdempotency,
		conf.auth,
	)
	// Output: testing api config, testing: false, req_idempotency: true, token source: MOLLIE_API_TOKEN.
}

// ExampleNewOrgConfig demonstrates how to initialize a Config  struct
// with testing mode disabled, token source from the default Org token
// env variable (MOLLIE_ORG_TOKEN) and request idempotency feature enabled.
func ExampleNewOrgConfig() {
	conf := NewOrgConfig(true)
	fmt.Printf(
		"testing org config, testing: %v, req_idempotency: %v, token source: %s.",
		conf.testing,
		conf.reqIdempotency,
		conf.auth,
	)
	// Output: testing org config, testing: false, req_idempotency: true, token source: MOLLIE_ORG_TOKEN.
}
