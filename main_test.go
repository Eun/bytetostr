package main

import (
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"dec and spaces", args{"72 101 108 108 111 32 87 111 114 108 100"}, "Hello World", false},
		{"hex and spaces", args{"0x48 0x65 0x6C 0x6C 0x6F 0x20 0x57 0x6F 0x72 0x6C 0x64"}, "Hello World", false},
		{"hex, spaces and commas", args{"0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64"}, "Hello World", false},
		{"one giant hex", args{"0x480x650x6C0x6C0x6F0x200x570x6F0x720x6C0x64"}, "Hello World", false},
		{"unicode", args{"0x2713"}, "✓", false},
		{"unicode", args{"0x00002713"}, "✓", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convert(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
