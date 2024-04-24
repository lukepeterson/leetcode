package main

import (
	"reflect"
	"testing"
)

func Test_findWinners(t *testing.T) {
	type args struct {
		matches [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test case 1",
			args: args{matches: [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}},
			want: [][]int{{1, 2, 10}, {4, 5, 7, 8}},
		},
		{
			name: "Test case 2",
			args: args{matches: [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}},
			want: [][]int{{1, 2, 5, 6}, {}},
		},
		{
			name: "Test case 3",
			args: args{matches: [][]int{{1, 100000}}},
			want: [][]int{{1}, {100000}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWinners(tt.args.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWinners(%d) = %v, want %v", tt.args.matches, got, tt.want)
			}
		})
	}
}
