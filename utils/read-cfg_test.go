package utils

import "testing"

func TestReadMySQLConfig(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test", args: args{filePath: "../config"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadMySQLConfig(tt.args.filePath); len(got) == 0 {
				t.Errorf("ReadMySQLConfig() = %v", got)
			}
		})
	}
}
