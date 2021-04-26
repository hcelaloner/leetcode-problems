package src

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example1", args: args{x: 121}, want: true},
		{name: "Example2", args: args{x: 122}, want: false},
		{name: "Example3", args: args{x: -121}, want: false},
		{name: "Example4", args: args{x: -122}, want: false},
		{name: "Example5", args: args{x: 10}, want: false},
		{name: "Example6", args: args{x: 11}, want: true},
		{name: "Example7", args: args{x: -101}, want: false},
		{name: "Example8", args: args{x: 313}, want: true},
		{name: "Example9", args: args{x: 1000021}, want: false},
		{name: "Example10", args: args{x: 10101}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
