package leetcode

import (
	"github.com/colachg/go-in-practise/data-structure/list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = linkedlist.LinkedList
type Node = linkedlist.Node

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// node is a pointer to hold the Head node
	node := &Node{Property: 0}
	// here current is a tmp pointer which will be always set by the for loop
	// and this is the key that confused me.
	n1, n2, carry, current := 0, 0, 0, node
	for l1.Head != nil || l2.Head != nil || carry != 0 {
		if l1.Head == nil {
			n1 = 0
		} else {
			n1 = l1.Head.Property
			l1.Head = l1.Head.NextNode
		}

		if l2.Head == nil {
			n2 = 0
		} else {
			n2 = l2.Head.Property
			l2.Head = l2.Head.NextNode
		}

		current.NextNode = &Node{Property: (n1 + n2 + carry) % 10}
		current = current.NextNode
		carry = (n1 + n2 + carry) / 10
	}
	return &ListNode{Head: node.NextNode}
}
