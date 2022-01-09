package leetcode

import "github.com/colachg/go-in-practise/data-structure/list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = linkedlist.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// node is a pointer to hold the Head node
	node := &ListNode{Val: 0}
	// here current is a tmp pointer which will be always set by the for loop
	// and this is the key that confused me.
	n1, n2, carry, current := 0, 0, 0, node
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return node.Next
}
