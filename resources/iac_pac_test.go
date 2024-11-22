package resources

import (
	"testing"
)

func TestGetScanResult(t *testing.T) {
	type args struct {
		iacPath string
		pacPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				iacPath: "../testdata/valid_iac",
				pacPath: "../testdata/valid_pac",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = GetScanResult(tt.args.iacPath, tt.args.pacPath)
		})
	}
}
