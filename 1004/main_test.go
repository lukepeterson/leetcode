package main

import "testing"

func Test_longestOnes(t *testing.T) {
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
			args: args{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2},
			want: 6,
		},
		{
			name: "test case 2",
			args: args{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes(%d) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
