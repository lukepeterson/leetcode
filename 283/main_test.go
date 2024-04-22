package main

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{nums: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "test case 2",
			args: args{nums: []int{0}},
			want: []int{0},
		},
		{
			name: "test case 3",
			args: args{nums: []int{1, 0, 0}},
			want: []int{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes(%d) got %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
