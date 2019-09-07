package mollie

import (
	"testing"
)

func TestHalURL_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Unmarshal google.com",
			args{b: []byte("https://google.com")},
			false,
		},
		{
			"Unmarshal parse error",
			args{b: []byte(" http://google.com")},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hl := &HalURL{}
			if err := hl.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
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
			d := &Date{}
			if err := d.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
