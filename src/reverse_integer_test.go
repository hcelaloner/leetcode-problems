package src

import (
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example1", args: args{x: 123}, want: 321},
		{name: "Example2", args: args{x: -123}, want: -321},
		{name: "Example3", args: args{x: 120}, want: 21},
		{name: "Example4", args: args{x: 0}, want: 0},
		{name: "Example5", args: args{x: math.MinInt32}, want: 0},
		{name: "Example6", args: args{x: math.MinInt32 + 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v for arg= %v", got, tt.want, tt.args.x)
			}
		})
	}
}
