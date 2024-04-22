package main

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{s: "abcdefg", k: 2},
			want: "bacdfeg",
		},
		{
			name: "Test case 2",
			args: args{s: "abcd", k: 2},
			want: "bacd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr(%s, %d) = %v, want %v", tt.args.s, tt.args.k, got, tt.want)
			}
		})
	}
}
