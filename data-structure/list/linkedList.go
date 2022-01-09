package linkedlist

import "fmt"

// Node Class
type Node struct {
	Property int
	NextNode *Node
}

// LinkedList Class
type LinkedList struct {
	Head *Node
}

// ArrayToLinkedList create a linkedlist from input
func ArrayToLinkedList(input []int) *LinkedList {
	if len(input) == 0 {
		return nil
	}
	// here are two pointers
	head := &Node{}
	tmp := head
	for _, num := range input {
		tmp.NextNode = &Node{Property: num}
		tmp = tmp.NextNode
	}
	return &LinkedList{head.NextNode}
}

// ToArray translate a linkedlist to an array
func (linkedList *LinkedList) ToArray() *[]int {
	if linkedList != nil && linkedList.Head != nil {
		var a []int
		for i := linkedList.Head; i != nil; i = i.NextNode {
			a = append(a, i.Property)
		}
		return &a
	}
	return nil
}

// AddToHead add node to the Head of linkedlist
func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.Property = property
	node.NextNode = nil

	if linkedList.Head != nil {
		node.NextNode = linkedList.Head
	}
	linkedList.Head = node
}

//IterateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.Head; node != nil; node = node.NextNode {
		if node == nil {
			fmt.Print("nil")
		} else {
			fmt.Print(node.Property, "->")
		}
	}
}

// LastNode returns the end node of the linked list
func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.Head; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd adds a property to the end of the linked list
func (linkedList LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.Property = property
	node.NextNode = nil

	var lastNode *Node
	lastNode = linkedList.LastNode()

	if lastNode != nil {
		lastNode.NextNode = node
	}
}

//NodeWithValue method returns Node given parameter property
func (linkedList LinkedList) NodeWithValue(property int) *Node {
	var nodeWithValue *Node
	for node := linkedList.Head; node != nil; node = node.NextNode {
		if node.Property == property {
			nodeWithValue = node
			break
		}
	}
	return nodeWithValue
}

//AddAfter method adds a node with nodeProperty after node with property
func (linkedList LinkedList) AddAfter(nodeProperty, property int) {
	node := &Node{}
	node.Property = property
	node.NextNode = nil

	nodeWithValue := linkedList.NodeWithValue(nodeProperty)
	if nodeWithValue != nil {
		node.NextNode = nodeWithValue.NextNode
		nodeWithValue.NextNode = node
	}
}
