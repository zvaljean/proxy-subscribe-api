package functions

import "testing"

func Test_parseFlag(t *testing.T) {
	type args struct {
		conf string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case-1",
			args: args{
				conf: "2410091",
			},
		},
		{
			name: "case-2",
			args: args{
				conf: "241009",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := parseFlag(tt.args.conf)
			t.Logf("%s: flag, %s,", tt.name, flag)
		})
	}
}
