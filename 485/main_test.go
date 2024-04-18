package main

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{nums: []int{1, 1, 0, 1, 1, 1}},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{nums: []int{1, 0, 1, 1, 0, 1}},
			want: 2,
		},
		{
			name: "Test case 3",
			args: args{nums: []int{1, 1, 0, 1}},
			want: 2,
		},

		// Input: nums = [1,1,0,1,1,1]
		// Output: 3
		// Explanation: The first two digits or the last three digits are consecutive 1s.
		// The maximum number of consecutive 1s is 3.

		// Example 2:
		// Input: nums = [1,0,1,1,0,1]
		// Output: 2

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes(%d) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
