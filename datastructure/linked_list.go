package datastructure

type ListNode struct {
	Val  int
	Next *ListNode
}

// L206 https://leetcode.com/problems/reverse-linked-list/
func reverseListIteratively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, next *ListNode
	var curr = head

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

// L25 https://leetcode.com/problems/reverse-nodes-in-k-group/
func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	var curr = head

	var count int

	for count < k && curr != nil {
		prev = curr
		curr = curr.Next
		count++
	}

	if count < k {
		return head
	}

	r := reverseKGroup(curr, k)

	prev.Next = nil
	l := reverseListIteratively(head)

	var res = l

	for l.Next != nil {
		l = l.Next
	}

	l.Next = r

	return res
}

func reverseKGroup(head *ListNode, k int) *ListNode { // 1234 21 43
	if head == nil || head.Next == nil {
		return head
	}
	var count = k

	var moveNode = head
	var prev *ListNode

	for count > 0 && moveNode != nil {
		prev = moveNode
		moveNode = moveNode.Next
		count--
	}

	prev.Next = nil

	l := reverseListIteratively(head)

	var res = l
	for l != nil && l.Next != nil {
		l = l.Next
	}

	l.Next = reverseKGroup(moveNode, k)

	return res
}

// L160 https://leetcode.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB

	// headA and headB length not same,if let them meet, speed is same, make the distance offset.
	for pa != pb {
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

// L21 https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoSortedListsRecursively(l1, l2 *ListNode) *ListNode { //nolint
	if l1 == nil {
		return l1
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

// L23 https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeTwoSortedListsRecursively(lists[0], lists[1])
	}

	m := len(lists) / 2

	l := mergeKLists(lists[:m])
	r := mergeKLists(lists[m:])

	return mergeTwoSortedListsRecursively(l, r)
}

// L83 https://leetcode.com/problems/remove-duplicates-from-sorted-list/
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

// L19 https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	first := head
	second := head

	for n > 0 && second != nil {
		second = second.Next
		n--
	}

	if second == nil {
		return first.Next
	}

	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next

	return head
}

// 1234 21 43
// 2134
//
// L24 https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return head
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

	var res = &ListNode{0, nil}
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
// L445 https://leetcode.com/problems/add-two-numbers-ii/description/

// L725 https://leetcode.com/problems/split-linked-list-in-parts/description/
func splitListToParts(root *ListNode, k int) []*ListNode {
	var l int
	var p = root
	for p != nil {
		p = p.Next
		l++
	}

	mod := l % k
	groups := l / k

	res := make([]*ListNode, k)
	var curr = root
	for i := 0; curr != nil && i < k; i++ {
		var curSize int
		if mod > 0 {
			curSize = groups + 1
			mod--
		} else {
			curSize = groups
		}

		var chead = curr
		for j := 0; j < curSize-1; j++ {
			curr = curr.Next
		}

		next := curr.Next
		curr.Next = nil

		res[i] = chead
		curr = next
	}

	return res
}

// L328 https://leetcode.com/problems/odd-even-linked-list/description/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var odd = head
	var even = odd.Next

	var fodd = odd
	var feven = even

	for odd != nil && even != nil && odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = odd.Next.Next

		odd = odd.Next
		even = even.Next
	}
	odd.Next = feven

	return fodd
}

// L234 https://leetcode.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var right *ListNode
	if fast == nil { //奇数
		right = reverseListIteratively(slow)
		slow.Next = nil
	}

	if fast != nil { //偶数
		right = reverseListIteratively(slow.Next)
		slow.Next = nil
	}

	return isEqual(head, right)
}

func isEqual(l1, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil && l2 == nil {
		return true
	}

	return false
}

// 12345
// 14325
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	rhead := head
	var rheadPrev *ListNode

	for left > 1 {
		rheadPrev = rhead
		rhead = rhead.Next
		left--
		right--
	}

	curr := rhead
	next := rhead.Next
	var prev *ListNode

	for right > 0 {
		next = curr.Next

		curr.Next = prev

		prev = curr
		curr = next
		right--
	}

	if rheadPrev != nil {
		rheadPrev.Next = prev
	} else {
		head = prev
	}

	rhead.Next = next

	return head
}

//12345 2
//s f
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	phead := head

	slow := phead
	fast := phead

	var meetpoint *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			meetpoint = slow
			break
		}
	}

	if meetpoint != nil {
		for head != nil {
			if head == meetpoint {
				return head
			}
			head = head.Next
			meetpoint = meetpoint.Next
		}
	}

	return nil
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	var slow = head
	var fast = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	a := reverseListIteratively(slow)

	curr := head
	for a != nil && curr != nil && a.Next != nil && curr.Next != nil {
		temp := curr.Next
		curr.Next = a

		temp2 := a.Next
		a.Next = temp

		curr = temp
		a = temp2
	}
}

func reorderList3(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	var slow = head
	var fast = head
	var prevSlow *ListNode

	for fast != nil && fast.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prevSlow.Next = nil
	r := reverseListIteratively(slow)

	mergeList(head, r)

	return
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList2(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseListIteratively(l2)
	mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
