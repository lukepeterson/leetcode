package main

import "testing"

func Test_longestSubArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{nums: []int{3, 2, 1, 3, 1, 1}, k: 5},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{nums: []int{3, 1, 2, 7, 4, 2, 1, 1, 5}, k: 8},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
