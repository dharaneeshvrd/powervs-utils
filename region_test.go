package utils

import (
	"testing"
)

func TestGetRegion(t *testing.T) {
	type args struct {
		zone string
	}
	tests := []struct {
		name       string
		args       args
		wantRegion string
		wantErr    bool
	}{
		{
			"London",
			args{"lon06"},
			"lon1",
			false,
		},
		{
			"Dallas",
			args{"us-south"},
			"us-south",
			false,
		},
		{
			"Dallas 12",
			args{"dal12"},
			"dal",
			false,
		},
		{
			"Sao Paulo 01",
			args{"sao01"},
			"sao",
			false,
		},
		{
			"Washington DC",
			args{"us-east"},
			"us-east",
			false,
		},
		{
			"Toronto",
			args{"tor01"},
			"tor",
			false,
		},
		{
			"Frankfurt",
			args{"eu-de-1"},
			"eu-de",
			false,
		},
		{
			"Sydney",
			args{"syd01"},
			"syd",
			false,
		},
		{
			"India",
			args{"blr01"},
			"",
			true,
		},
		{
			"Tokyo",
			args{"tok04"},
			"tok",
			false,
		},
		{
			"Montreal",
			args{"mon01"},
			"mon",
			false,
		},
		{
			"Osaka",
			args{"osa21"},
			"osa",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRegion, err := GetRegion(tt.args.zone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRegion != tt.wantRegion {
				t.Errorf("GetRegion() gotRegion = %v, want %v", gotRegion, tt.wantRegion)
			}
		})
	}
}
