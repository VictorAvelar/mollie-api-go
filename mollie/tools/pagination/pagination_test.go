package pagination

import "testing"

func TestExtractFromQueryParam(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name       string
		args       args
		wantLastID string
		wantErr    bool
	}{
		{
			name: "test extracts correct parameter",
			args: args{
				"https://api.mollie.com/v2/payments?from=tr_EkceGSH8Ga&limit=5",
			},
			wantLastID: "tr_EkceGSH8Ga",
			wantErr:    false,
		},
		{
			name: "test wrong url error parameter",
			args: args{
				"h%%s12",
			},
			wantLastID: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLastID, err := ExtractFromQueryParam(tt.args.uri)
			t.Log(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractFromQueryParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLastID != tt.wantLastID {
				t.Errorf("ExtractFromQueryParam() = %v, want %v", gotLastID, tt.wantLastID)
			}
		})
	}
}
