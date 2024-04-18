package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty test case",
			args: args{""},
			want: true,
		},
		{
			name: "single character test case",
			args: args{"a"},
			want: true,
		},
		{
			name: "small test case",
			args: args{"aa"},
			want: true,
		},
		{
			name: "medium test case",
			args: args{"radar"},
			want: true,
		},
		{
			name: "medium test case 2",
			args: args{"atlassian"},
			want: false,
		},
		{
			name: "large test case",
			args: args{"a man, a plan, a canal, panama"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
