package mollie

import (
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
)

func TestContextValues_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    ContextValues
		wantErr bool
	}{
		// Add test cases here
		{
			name: "Correct decoding JSON",
			data: []byte(`{"type1": "value1", "type2": "value2"}`),
			want: ContextValues{
				"type1": "value1",
				"type2": "value2",
			},
			wantErr: false,
		},
		{
			name:    "Handle empty JSON",
			data:    []byte(`{}`),
			want:    ContextValues{},
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			data:    []byte(`{"type1": "value1", "type2": "value2"`),
			want:    make(ContextValues),
			wantErr: true,
		},
		{
			name:    "Incorrect type in json returns an empty map",
			data:    []byte(`{"type1":["value1", "value2"]}`),
			want:    make(ContextValues),
			wantErr: false,
		},
		{
			name: "Test correct case described on issue #251",
			data: []byte(`{"context": {
      "paymentId": "tr_xxxxxxxx"
  }}`),
			want:    make(ContextValues),
			wantErr: false,
		},
		{
			name:    "Test failing case described on issue #251",
			data:    []byte(`{"context": []}`),
			want:    make(ContextValues),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cv ContextValues
			err := cv.UnmarshalJSON(tt.data)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, cv)
			}
		})
	}
}

func TestAmount_URLEncodingSimple(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want string
	}{
		{
			name: "Test URL encoding simple.",
			a: Amount{
				Value:    "10.00",
				Currency: "EUR",
			},
			want: "currency=EUR&value=10.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := query.Values(tt.a)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, v.Encode())
		})
	}
}

func TestAmount_URLEncodingNested(t *testing.T) {
	tests := []struct {
		name string
		a    struct {
			Amount Amount `url:"amount"`
		}
		want string
	}{
		{
			name: "Test URL encoding nested amount in struct.",
			a: struct {
				Amount Amount `url:"amount"`
			}{
				Amount{
					Value:    "10.00",
					Currency: "EUR",
				},
			},
			want: "amount%5Bcurrency%5D=EUR&amount%5Bvalue%5D=10.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := query.Values(tt.a)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, v.Encode())
		})
	}
}
