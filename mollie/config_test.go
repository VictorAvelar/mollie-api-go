package mollie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		t    bool
		ikg  bool
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
			got := NewConfig(tt.args.t, tt.args.ikg, tt.args.auth)
			assert.Equal(t, got, tt.want)
		})
	}
}

func ExampleNewConfig() {
	fmt.Println(NewConfig(true, true, APITokenEnv))
	// Output: &{true MOLLIE_API_TOKEN true}
}
