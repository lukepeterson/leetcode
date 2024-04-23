package main

import "testing"

func Test_largestAltitude(t *testing.T) {
	type args struct {
		gain []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{gain: []int{-5, 1, 5, 0, -7}},
			want: 1,
		},
		{
			name: "Test case 2",
			args: args{gain: []int{-4, -3, -2, -1, 4, 3, 2}},
			want: 0,
		},
		{
			name: "Test case 3",
			args: args{gain: []int{-59, -27, -7, -24, 69, 80, -22, -53}},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestAltitude(tt.args.gain); got != tt.want {
				t.Errorf("largestAltitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
