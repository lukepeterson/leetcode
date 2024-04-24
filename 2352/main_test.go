package main

import "testing"

func Test_equalPairs(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1 (3x3)",
			args: args{grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}},
			want: 1,
		},
		{
			name: "Test case 2 (4x4)",
			args: args{grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}},
			want: 3,
		},
		{
			name: "Test case 3 (2x2)",
			args: args{grid: [][]int{{11, 1}, {1, 11}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalPairs(tt.args.grid); got != tt.want {
				t.Errorf("equalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
