package main

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{nums: []int{2, 3, 1, 2, 4, 3}, target: 7},
			want: 2,
		},
		{
			name: "Test case 2",
			args: args{nums: []int{1, 4, 4}, target: 4},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, target: 11},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen(%d) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
