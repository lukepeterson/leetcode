package main

import (
	"testing"
)

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 0",
			args: args{s: "ibibhix", k: 3},
			want: 2,
		},
		{
			name: "Test case 1",
			args: args{s: "abciiidef", k: 3},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{s: "aeiou", k: 2},
			want: 2,
		},
		{
			name: "Test case 3",
			args: args{s: "leetcode", k: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels(%s, %d) = %v, want %v", tt.args.s, tt.args.k, got, tt.want)
			}
		})
	}
}
