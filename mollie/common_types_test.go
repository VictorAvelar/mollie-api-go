package mollie

import (
	"testing"
	"time"
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
			if err := d.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShortDate_MarshalJSON(t *testing.T) {
	t.Run("marshal is successful", func(t *testing.T) {
		n := time.Now()
		d := &ShortDate{}
		d.Time = n
		if _, err := d.MarshalJSON(); err != nil {
			t.Errorf("MarshalJSON() error = %v", err)
		}
	})
}
