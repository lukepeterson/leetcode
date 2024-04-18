package main

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty test case",
			args: args{arr1: []int{}, arr2: []int{}},
			want: []int{},
		},
		{
			name: "small test case",
			args: args{arr1: []int{1, 5, 7}, arr2: []int{2, 4, 9}},
			want: []int{1, 2, 4, 5, 7, 9},
		},
		{
			name: "another test case",
			args: args{arr1: []int{-2, 3, 9}, arr2: []int{0, 4}},
			want: []int{-2, 0, 3, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
