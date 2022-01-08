package structure

import "fmt"

// Node Class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList Class
type LinkedList struct {
	Head *Node
}

// NewLinkedList create a linkedlist from input
func NewLinkedList(input []int) *LinkedList {
	if len(input) == 0 {
		return nil
	}
	// here are two pointers
	head := &Node{}
	tmp := head
	for _, num := range input {
		tmp.nextNode = &Node{property: num}
		tmp = tmp.nextNode
	}
	return &LinkedList{head.nextNode}
}

// ToArray translate a linkedlist to an array
func (linkedList *LinkedList) ToArray() *[]int {
	if linkedList != nil && linkedList.Head != nil {
		var a []int
		for i := linkedList.Head; i != nil; i = i.nextNode {
			a = append(a, i.property)
		}
		return &a
	}
	return nil
}

// AddToHead add node to the Head of linkedlist
func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.Head != nil {
		node.nextNode = linkedList.Head
	}
	linkedList.Head = node
}

//IterateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.Head; node != nil; node = node.nextNode {
		if node == nil {
			fmt.Print("nil")
		} else {
			fmt.Print(node.property, "->")
		}
	}
}

// LastNode returns the end node of the linked list
func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.Head; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd adds a property to the end of the linked list
func (linkedList LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node
	lastNode = linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
	}
}

//NodeWithValue method returns Node given parameter property
func (linkedList LinkedList) NodeWithValue(property int) *Node {
	var nodeWithValue *Node
	for node := linkedList.Head; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWithValue = node
			break
		}
	}
	return nodeWithValue
}

//AddAfter method adds a node with nodeProperty after node with property
func (linkedList LinkedList) AddAfter(nodeProperty, property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil

	nodeWithValue := linkedList.NodeWithValue(nodeProperty)
	if nodeWithValue != nil {
		node.nextNode = nodeWithValue.nextNode
		nodeWithValue.nextNode = node
	}
}
