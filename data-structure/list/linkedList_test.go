package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {

	t.Run("Test function ArrayToLinkedList()", func(t *testing.T) {
		tests := []struct {
			name string
			args []int
			want *ListNode
		}{
			{
				name: "Test normal case",
				args: []int{1, 2, 2},
				want: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				}},
			{
				name: "Test nil case",
				args: []int{},
				want: nil,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ArrayToLinkedList(tt.args); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ArrayToLinkedList() = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("Test method AddToHead()", func(t *testing.T) {
		tests := []struct {
			name     string
			property int
			want     *ListNode
		}{
			{
				name:     "Add property 2",
				property: 2,
				want: &ListNode{
					Val:  2,
					Next: nil,
				}},
		}
		for _, tt := range tests {
			tt.want.AddToHead(tt.property)
			got := tt.want
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		}

	})
	t.Run("Test method ToArray()", func(t *testing.T) {
		tests := []struct {
			name string
			args *ListNode
			want *[]int
		}{
			{
				name: "Test case 1 - normal", args: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					}}, want: &[]int{0, 1, 2},
			},
			{
				name: "Test case 2 - nil",
				args: nil,
				want: nil,
			},
			{
				name: "Test case 3 - nil",
				args: &ListNode{},
				want: &[]int{0},
			},
			{
				name: "Test case 4 - normal",
				args: &ListNode{Val: 0},
				want: &[]int{0},
			},
		}
		for _, tt := range tests {
			if got := tt.args.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s, got: %v, want: %v", tt.name, got, tt.want)
			}
		}

	})
}
