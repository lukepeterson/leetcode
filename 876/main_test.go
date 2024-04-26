package main

import (
	"reflect"
	"testing"
)

func newList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test case 1",
			args: args{head: newList([]int{1, 2, 3, 4, 5})},
			want: newList([]int{3, 4, 5}),
		},
		{
			name: "Test case 2",
			args: args{head: newList([]int{1, 2, 3, 4, 5, 6})},
			want: newList([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
