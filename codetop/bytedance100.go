package codetop

/*
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
"tmmzuxt"
*/
func lengthOfLongestSubstring(s string) int {
	var checkmap = make(map[byte]int, 0)

	var max, l int

	for i := range s {
		if v, ok := checkmap[s[i]]; ok && v >= l {
			l = v + 1
		} else {
			max = getMax(max, i-l+1)
		}

		checkmap[s[i]] = i
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
reverse linked list
1234 4321
*/

type Node struct {
	Val  int
	Next *Node
}

func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, next *Node
	curr := head

	for curr != nil {
		next = curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}

/*
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

type LRUCache struct {
	Cap  int
	Dict map[int]*DNode
	Head *DNode
	Tail *DNode
}

type DNode struct {
	Key  int
	Val  int
	Prev *DNode
	Next *DNode
}

func Constructor(capacity int) LRUCache {
	var res = LRUCache{
		Cap:  capacity,
		Dict: make(map[int]*DNode, 0),
		Head: new(DNode),
		Tail: new(DNode),
	}

	res.Head.Next = res.Tail
	res.Tail.Prev = res.Head

	return res
}

func (this *LRUCache) remove(Node *DNode) {
	if v, ok := this.Dict[Node.Key]; ok {
		if v.Next != nil {
			v.Next.Prev = v.Prev
		}

		if v.Prev != nil {
			v.Prev.Next = v.Next
		}
		delete(this.Dict, Node.Key)
	}

	return
}

func (this *LRUCache) add(key int, val int) {
	var temp = &DNode{
		Key: key,
		Val: val,
	}

	temp.Next = this.Head.Next
	this.Head.Next = temp

	temp.Prev = this.Head
	temp.Next.Prev = temp

	this.Dict[key] = temp
	return
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Dict[key]; ok {
		// Get reset the recent location
		this.remove(v)
		this.add(key, v.Val)
		return v.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.Dict[key]; ok {
		this.remove(v)
		this.add(key, value)
	} else {
		if len(this.Dict) >= this.Cap {
			this.remove(this.Tail.Prev)
		}

		this.add(key, value)
	}
}

func getLargestElement(nums []int, k int) {

}