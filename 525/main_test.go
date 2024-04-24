package main

import "testing"

func Test_findMaxLength(t *testing.T) {
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
			args: args{[]int{0, 1, 1, 0, 1, 1, 0, 1}},
			want: 4,
		},
		{
			name: "Test case 2",
			args: args{[]int{0, 1}},
			want: 2,
		},
		{
			name: "Test case 3",
			args: args{[]int{0, 1, 0}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
