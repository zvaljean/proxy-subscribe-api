package data

import (
	"testing"
)

var filepath = "/home/valjean/workspace/own/code/golang/proxy-subscribe/configs/user.csv"

func TestParseCsv(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
	}{
		{
			name:     "parse csv file",
			filepath: filepath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s-%s", tt.name, tt.filepath)
			csv := ParseCsvForList(tt.filepath, 1)

			for key, value := range *csv {
				t.Logf("the string value %d: -> %s\n", key, value)
			}
		})
	}
}

func TestParseCsvForMap(t *testing.T) {
	type args struct {
		filepath string
		key      int
		value    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test parse csv for map",
			args: args{
				filepath: filepath,
				key:      1,
				value:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := ParseCsvForMap(tt.args.filepath, tt.args.key, tt.args.value)
			for key, val := range *data {
				t.Logf("key: %s, value: %s", key, val)
			}
		})
	}
}
