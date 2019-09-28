package mollie

import (
	"reflect"
	"testing"
)

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
			if got := NewConfig(tt.args.t, tt.args.auth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
