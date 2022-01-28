package datastructure

import "reflect"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 206 https://leetcode.com/problems/reverse-linked-list/
func reverseListIteratively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var curr = head
	var next = curr.Next
	var prev *ListNode

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

	var count int
	var pointer = head
	var ppointer *ListNode

	for i := 0; i < k; i++ {
		if pointer != nil {
			ppointer = pointer
			pointer = pointer.Next
			count++
		} else {
			break
		}
	}

	ppointer.Next = nil

	if count == k {
		r := reverseKGroup(pointer, k)
		l := reverseListIteratively(head)

		temp := l
		for temp != nil {
			if temp.Next == nil {
				break
			}
			temp = temp.Next
		}
		temp.Next = r

		return l
	}

	return head
}

// 160 https://leetcode.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB

	// headA and headB length not same,if let them meet, speed is same, make the distance offset.
	for !reflect.DeepEqual(pa, pb) {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}

		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}

	return pa
}

// 21 https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoSortedListsRecursively(l1, l2 *ListNode) *ListNode { //nolint
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		r := mergeTwoSortedListsRecursively(l1.Next, l2)
		l1.Next = r
		return l1
	}

	r := mergeTwoSortedListsRecursively(l1, l2.Next)
	l2.Next = r
	return l2
}

func mergeTwoSortedListsIteratively(l1, l2 *ListNode) *ListNode { //nolint
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var res = &ListNode{}
	var cres = res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cres.Next = l1
			l1 = l1.Next
		} else {
			cres.Next = l2
			l2 = l2.Next
		}
		cres = cres.Next
	}

	if l1 == nil {
		cres.Next = l2
	}
	if l2 == nil {
		cres.Next = l1
	}

	return res.Next
}

// 23 https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeTwoSortedListsRecursively(lists[0],lists[1])
	}

	m := len(lists)/2

	l := mergeKLists(lists[:m])
    r := mergeKLists(lists[m:])

	return mergeTwoSortedListsRecursively(l,r)
}

// 83 https://leetcode.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicatesIteratively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	now := head

	for now != nil && now.Next != nil {
		if now.Val == now.Next.Val {
			now.Next = now.Next.Next
		} else {
			now = now.Next
		}
	}

	return head
}

func deleteDuplicatesRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	now := head
	l := deleteDuplicatesRecursively(now.Next)

	if now.Val != l.Val {
		now.Next = l
		return now
	}

	return l
}

// 19 https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	first := head
	second := head

	for n > 0 {
		second = second.Next
		n--
	}

	if second == nil {
		return head.Next
	}

	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next
	}

	first.Next = first.Next.Next

	return head
}

// 24 https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var curr = head
	var res *ListNode = curr.Next
	var next *ListNode
	for curr != nil && curr.Next != nil {
		next = curr.Next.Next

		curr.Next.Next = curr
		if next == nil {
			curr.Next = nil
		} else {
			if next.Next == nil {
				curr.Next = next
			} else {
				curr.Next = next.Next
			}
		}

		curr = next
	}

	return res
}

func swapPairsRecursively(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    p := head.Next.Next
    
    l := swapPairs(p)
    
    head.Next.Next = nil
    r := reverseListIteratively(head)
    r.Next.Next = l
    
    return r
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var res =&ListNode{0,nil}
	var sres = res

	var curr = head
	var prev *ListNode
	var next *ListNode
	for curr != nil && curr.Next != nil {
		next = curr.Next.Next
		prev = curr.Next

		curr.Next.Next = curr
		curr.Next = nil

		res.Next = prev
		res = res.Next.Next

		curr = next
	}

	res.Next = curr

	return sres.Next
}

// TODO After Stack
// 445 https://leetcode.com/problems/add-two-numbers-ii/description/

// 725 https://leetcode.com/problems/split-linked-list-in-parts/description/

// 328 https://leetcode.com/problems/odd-even-linked-list/description/