package structure

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {

	t.Run("Test function TestNewLinkedList()", func(t *testing.T) {
		tests := []struct {
			name string
			args []int
			want *LinkedList
		}{
			{
				name: "Test normal case",
				args: []int{1, 2, 2},
				want: &LinkedList{Head: &Node{
					property: 1,
					nextNode: &Node{
						property: 2,
						nextNode: &Node{
							property: 2,
							nextNode: nil,
						},
					},
				}},
			},
			{
				name: "Test nil case",
				args: []int{},
				want: nil,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := NewLinkedList(tt.args); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("Test method AddToHead()", func(t *testing.T) {
		tests := []struct {
			name     string
			property int
			want     *LinkedList
		}{
			{name: "Add property 2", property: 2, want: &LinkedList{
				Head: &Node{
					property: 2, nextNode: nil}}},
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
			args *LinkedList
			want *[]int
		}{
			{
				name: "Test case 1 - normal", args: &LinkedList{Head: &Node{
					property: 0,
					nextNode: &Node{
						property: 1,
						nextNode: &Node{
							property: 2,
							nextNode: nil,
						},
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
				args: &LinkedList{},
				want: nil,
			},
			{
				name: "Test case 4 - normal",
				args: &LinkedList{Head: &Node{property: 0}},
				want: &[]int{0},
			},
		}
		for _, tt := range tests {
			if got := tt.args.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		}

	})
}