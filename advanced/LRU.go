package advanced

type LRUCache struct {
	Head    *Node
	Tail    *Node
	Cap     int
	NodeMap map[int]*Node
}

type Node struct {
	Key int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	head := new(Node)
	head.Next = new(Node)
	head.Next.Prev = head

	return LRUCache{
		Head:    head,
		Tail:    head.Next,
		Cap:     capacity,
		NodeMap: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node,ok := this.NodeMap[key];ok {
		this.remove(node)
		this.add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node,ok := this.NodeMap[key];ok {
		this.remove(node)
		lnode := new(Node)
		lnode.Key = key
		lnode.Val = value
		this.add(lnode)
	}else {
		if len(this.NodeMap) < this.Cap {
			node := new(Node)
			node.Key = key
			node.Val = value	
			this.add(node)
			this.NodeMap[key] = node
		} else {
			node := new(Node)
			node.Key = key
			node.Val = value	
			this.add(node)
			this.NodeMap[key] = node
	
			this.remove(this.Tail.Prev)
		}
	}
}

func (this *LRUCache) add(node *Node) {
	temp := this.Head.Next
	temp.Prev = node
	this.Head.Next = node

	node.Prev = this.Head
	node.Next = temp

	this.NodeMap[node.Key] = node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	delete(this.NodeMap,node.Key)
}

func testLRU() {
	lru:=Constructor(2)
	lru.Get(2)
	lru.Put(2,6)
	lru.Get(1)
	lru.Put(1,5)
	lru.Put(1,2)
	lru.Get(1)
	lru.Get(2)
}