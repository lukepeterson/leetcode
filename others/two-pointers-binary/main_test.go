package main

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{input: "1101100111"},
			want: 5,
		},
		{
			name: "test case 2",
			args: args{input: "0101011"},
			want: 4,
		},
		{
			name: "test case 3",
			args: args{input: "01111011"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.input); got != tt.want {
				t.Errorf("longestSubstring(%s) = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}
