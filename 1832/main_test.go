package main

import "testing"

func Test_checkIfPangram(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{sentence: "thequickbrownfoxjumpsoverthelazydog"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{sentence: "leetcode"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPangram(tt.args.sentence); got != tt.want {
				t.Errorf("checkIfPangram() = %v, want %v", got, tt.want)
			}
		})
	}
}
