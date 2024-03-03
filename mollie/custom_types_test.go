package mollie

import (
	"testing"

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
