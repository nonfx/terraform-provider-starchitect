package utils

import (
	"reflect"
	"testing"
)

func Test_getTaxonsByIAC(t *testing.T) {
	type args struct {
		iacPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				iacPath: "../testdata/valid_iac",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTaxonsByIAC(tt.args.iacPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTaxonsByIAC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTaxonsByIAC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultPAC(t *testing.T) {
	type args struct {
		iacPath    string
		pacVersion string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				iacPath: "../testdata/valid_iac",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultPAC(tt.args.iacPath, tt.args.pacVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDefaultPAC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDefaultPAC() = %v, want %v", got, tt.want)
			}
		})
	}
}
