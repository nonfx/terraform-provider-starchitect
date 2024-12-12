package resources

import (
	"testing"
)

func TestGetScanResult(t *testing.T) {
	tests := []struct {
		name       string
		iacPath    string
		pacPath    string
		pacVersion string
		logPath    string
		wantErr    bool
	}{
		{
			name:       "Valid IAC and PAC paths",
			iacPath:    "../testdata/valid_iac",
			pacPath:    "../testdata/valid_pac",
			pacVersion: "",
			logPath:    "test_logs",
			wantErr:    false,
		},
		{
			name:       "Invalid IAC path",
			iacPath:    "../testdata/invalid_path",
			pacPath:    "../testdata/valid_pac",
			pacVersion: "",
			logPath:    "test_logs",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, score := GetScanResult(tt.iacPath, tt.pacPath, tt.pacVersion, tt.logPath)
			if (result == "" || score == "") != tt.wantErr {
				t.Errorf("GetScanResult() error = %v, wantErr %v", result, tt.wantErr)
				return
			}
		})
	}
}
