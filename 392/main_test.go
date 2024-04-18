package main

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty test case",
			args: args{s: "", t: ""},
			want: true,
		},
		{
			name: "test case 1",
			args: args{s: "abc", t: "ahbgdc"},
			want: true,
		},
		{
			name: "test case 2",
			args: args{s: "axc", t: "ahbgdc"},
			want: false,
		},
		{
			name: "test case 3",
			args: args{s: "xyz", t: "atlassian"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
