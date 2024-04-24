package main

import "testing"

func Test_largestUniqueNumber(t *testing.T) {
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
			args: args{[]int{5, 7, 3, 9, 4, 9, 8, 3, 1}},
			want: 8,
		},
		{
			name: "Test case 2",
			args: args{[]int{9, 9, 8, 8}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestUniqueNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestUniqueNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
