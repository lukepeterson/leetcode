package main

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{s: "Let's take LeetCode contest"},
			want: string("s'teL ekat edoCteeL tsetnoc"),
		},
		{
			name: "test case 2",
			args: args{s: "Mr Ding"},
			want: string("rM gniD"),
		},
		{
			name: "test case 3",
			args: args{s: "test case"},
			want: string("tset esac"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
