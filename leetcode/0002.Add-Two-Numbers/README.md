### Problem  
https://leetcode.com/problems/add-two-numbers/

---
### Tips 

When parsing linkedlist data structure, (for me) it's recommended to init an empty head node.The key operation to link
lists is the two pointers.Let's take this as a example:  

```go
package linkedlist
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
```