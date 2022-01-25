package datastructure

type ListNode struct {
	Val int
	Next *ListNode
}

// 206 https://leetcode.com/problems/reverse-linked-list/
func reverseListIteratively(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var curr = head
	var prev *ListNode
	var next = head.Next

	for curr != nil {
		next = curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}

func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	r := reverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil

	return r
}
// 25 https://leetcode.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var right = head
	var pprev = head
	var left = pprev

	// calc the length of node enough for k reverse
	var count int
	for i := 0; i < k; i++ {
		if right != nil {
			pprev = right
			right = right.Next
			count++
		} else {
			break
		}
	}

	pprev.Next = nil

	if count == k {
		r := reverseKGroup(right, k)
		l := reverseLinkedList(left)

		var tt = l
		for tt.Next != nil {
			tt = tt.Next
		}
		tt.Next = r
		return l
	}

	return head
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var curr = head
	var next = curr.Next
	var prev  *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	return prev
}

// 160 https://leetcode.com/problems/intersection-of-two-linked-lists/

// 21 https://leetcode.com/problems/merge-two-sorted-lists/

// 23 https://leetcode.com/problems/merge-k-sorted-lists/

// 
