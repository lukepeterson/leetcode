package main

import (
	"reflect"
	"testing"
)

func Test_answerQueries(t *testing.T) {
	type args struct {
		input   []int
		queries [][]int
		limit   int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			// For example, given nums = [1, 6, 3, 2, 7, 2], queries = [[0, 3], [2, 5], [2, 4]], and limit = 13,
			// the answer is [true, false, true]. For each query, the subarray sums are [12, 14, 12].
			name: "test case 1",
			args: args{
				input:   []int{1, 6, 3, 2, 7, 2},
				queries: [][]int{{0, 3}, {2, 5}, {2, 4}},
				limit:   13,
			},
			want: []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerQueries(tt.args.input, tt.args.queries, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nanswerQueries(%d, %d, %d) = %v, want %v", tt.args.input, tt.args.queries, tt.args.limit, got, tt.want)
			}
		})
	}
}
