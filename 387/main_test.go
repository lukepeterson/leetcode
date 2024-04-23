package main

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{s: "leetcode"},
			want: 0,
		},
		{
			name: "Test case 2",
			args: args{s: "loveleetcode"},
			want: 2,
		},
		{
			name: "Test case 3",
			args: args{s: "aabb"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar(%s) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
