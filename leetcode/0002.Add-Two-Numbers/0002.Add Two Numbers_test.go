package leetcode

import (
	"github.com/colachg/go-in-practise/data-structure/list"
	"reflect"
	"testing"
)

func Test_AddTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	var tests = []struct {
		name string
		args *args
		want *[]int
	}{
		{
			name: "Test case one",
			args: &args{
				l1: []int{2, 6, 7},
				l2: []int{3, 7, 9},
			},
			want: &[]int{5, 3, 7, 1},
		},
		{
			name: "Test case two",
			args: &args{
				l1: []int{2, 2, 2},
				l2: []int{1, 1, 1},
			},
			want: &[]int{3, 3, 3},
		},
		{
			name: "Test case three",
			args: &args{
				l1: []int{1, 4, 6, 7, 9},
				l2: []int{8, 5, 2},
			},
			want: &[]int{9, 9, 8, 7, 9},
		},
		{
			name: "Test case four",
			args: &args{
				l1: []int{9, 9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9},
			},
			want: &[]int{8, 9, 9, 0, 0, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		list1 := linkedlist.ArrayToLinkedList(tt.args.l1)
		list2 := linkedlist.ArrayToLinkedList(tt.args.l2)
		if got := addTwoNumbers(list1, list2).ToArray(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
		}
	}
}
