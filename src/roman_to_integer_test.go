package src

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example1", args: args{s: "III"}, want: 3},
		{name: "Example2", args: args{s: "IV"}, want: 4},
		{name: "Example3", args: args{s: "IX"}, want: 9},
		{name: "Example4", args: args{s: "LVIII"}, want: 58},
		{name: "Example5", args: args{s: "MCMXCIV"}, want: 1994},
		{name: "Example6", args: args{s: "MDCCLXXVI"}, want: 1776},
		{name: "Example7", args: args{s: "MLXVI"}, want: 1066},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
