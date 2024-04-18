package main

import "testing"

func Test_twoSum(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{input: []int{}, target: 4},
			want: false,
		},
		{
			name: "small test case",
			args: args{input: []int{1, 2, 3, 4, 5}, target: 7},
			want: true,
		},
		{
			name: "medium test case",
			args: args{input: []int{1, 2, 4, 6, 8, 9, 14, 15}, target: 13},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.input, tt.args.target); got != tt.want {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
