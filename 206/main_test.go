package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Empty list",
			args: args{head: nil},
			want: nil,
		},
		{
			name: "One element",
			args: args{head: &ListNode{Val: 1}},
			want: &ListNode{Val: 1},
		},
		{
			name: "Two elements",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
		},
		{
			name: "Three elements",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
			want: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
