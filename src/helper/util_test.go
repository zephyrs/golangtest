package helper

import "testing"

func TestTrace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{"bbb"}, want: "bbb"},
		{"2", args{"aaa"}, "aaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trace(tt.args.s); got != tt.want {
				t.Errorf("Trace() = %v, want %v", got, tt.want)
			}
		})
	}
}
