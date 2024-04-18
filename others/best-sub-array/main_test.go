package main

import "testing"

func Test_findBestSubarray(t *testing.T) {
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
			name: "Test case 1",
			args: args{nums: []int{3, -1, 4, 12, -8, 5, 6}, k: 4},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestSubarray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findBestSubarray(%d) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
