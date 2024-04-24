package main

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{ransomNote: "foo", magazine: "leetcode"},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{ransomNote: "todel", magazine: "leetcode"},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{ransomNote: "a", magazine: "b"},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{ransomNote: "aa", magazine: "ab"},
			want: false,
		},
		{
			name: "Test case 5",
			args: args{ransomNote: "aa", magazine: "aab"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
