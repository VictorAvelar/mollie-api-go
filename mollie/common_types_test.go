package mollie

import (
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
)

func TestShortDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"unmarshal fails with invalid date format",
			args{b: []byte("30-12-1991")},
			true,
		},
		{
			"unmarshal is successful",
			args{b: []byte("1991-12-30")},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ShortDate{}
			err := d.UnmarshalJSON(tt.args.b)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestShortDate_MarshalJSON(t *testing.T) {
	t.Run("marshal is successful", func(t *testing.T) {
		n := time.Now()
		d := &ShortDate{}
		d.Time = n
		_, err := d.MarshalJSON()
		assert.Nil(t, err)
	})
}

func TestURLQueryEncode(t *testing.T) {
	cases := []struct {
		name     string
		in       any
		expected string
	}{
		{
			"test amount encode",
			&Amount{Currency: "EUR", Value: "10.00"},
			"amount%5Bcurrency%5D=EUR&amount%5Bvalue%5D=10.00",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			v, _ := query.Values(c.in)
			assert.Equal(t, c.expected, v.Encode())
		})
	}
}
