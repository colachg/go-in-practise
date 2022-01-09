package linkedlist

import "fmt"

// ListNode Class
type ListNode struct {
	Val  int
	Next *ListNode
}

// ArrayToLinkedList create a linkedlist from input
func ArrayToLinkedList(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}
	// here are two pointers
	head := &ListNode{}
	tmp := head
	for _, num := range input {
		tmp.Next = &ListNode{Val: num}
		tmp = tmp.Next
	}
	return head.Next
}

// ToArray translate a linkedlist to an array
func (listNode *ListNode) ToArray() *[]int {
	if listNode != nil {
		var a []int
		for i := listNode; i != nil; i = i.Next {
			a = append(a, i.Val)
		}
		return &a
	}
	return nil
}

// AddToHead add node to the Head of linkedlist
func (listNode *ListNode) AddToHead(property int) {
	node := &ListNode{}
	node.Val = property
	node.Next = nil

	if listNode != nil {
		node.Next = listNode
	}
	listNode = node
}

//IterateList method iterates over LinkedList
func (listNode *ListNode) IterateList() {
	var node *ListNode
	for node = listNode; node != nil; node = node.Next {
		if node == nil {
			fmt.Print("nil")
		} else {
			fmt.Print(node.Val, "->")
		}
	}
}

// LastNode returns the end node of the linked list
func (listNode *ListNode) LastNode() *ListNode {
	var lastNode *ListNode
	for node := listNode; node != nil; node = node.Next {
		if node.Next == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd adds a property to the end of the linked list
func (listNode *ListNode) AddToEnd(property int) {
	var node = &ListNode{}
	node.Val = property
	node.Next = nil

	var lastNode *ListNode
	lastNode = listNode.LastNode()

	if lastNode != nil {
		lastNode.Next = node
	}
}

//NodeWithValue method returns ListNode given parameter property
func (listNode *ListNode) NodeWithValue(property int) *ListNode {
	var nodeWithValue *ListNode
	for node := listNode; node != nil; node = node.Next {
		if node.Val == property {
			nodeWithValue = node
			break
		}
	}
	return nodeWithValue
}

//AddAfter method adds a node with nodeProperty after node with property
func (listNode *ListNode) AddAfter(nodeProperty, property int) {
	node := &ListNode{}
	node.Val = property
	node.Next = nil

	nodeWithValue := listNode.NodeWithValue(nodeProperty)
	if nodeWithValue != nil {
		node.Next = nodeWithValue.Next
		nodeWithValue.Next = node
	}
}
