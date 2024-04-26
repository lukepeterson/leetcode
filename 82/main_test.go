package main

import (
	"reflect"
	"testing"
)

func listFromArray(vals []int) *ListNode {
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

func Test_deleteDuplicates(t *testing.T) {
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
			args: args{head: listFromArray([]int{1, 2, 3, 3, 4, 4, 5})},
			want: listFromArray([]int{1, 2, 5}),
		},
		{
			name: "Test case 2",
			args: args{head: listFromArray([]int{1, 1, 1, 2, 3})},
			want: listFromArray([]int{2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
