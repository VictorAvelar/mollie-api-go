package mollie

import (
	"testing"
	"time"

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
