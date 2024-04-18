package main

import (
	"reflect"
	"testing"
)

func Test_prefixSum(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{input: []int{5, 2, 1, 6, 3, 8}},
			want: []int{5, 7, 8, 14, 17, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixSum(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixSum(%d) = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}
