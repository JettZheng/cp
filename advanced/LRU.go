package advanced

type LRUCache struct {
    Cap int
    Dict map[int]*Node
    Head *Node
    Tail *Node
}

type Node struct {
    Key int
    Val int
    Prev *Node
    Next *Node
}

func Constructor(capacity int) LRUCache {
    head := new(Node)
    tail := new(Node)
    head.Next = tail
    tail.Prev = head
	
    return LRUCache {
        Cap: capacity,
        Dict: make(map[int]*Node),
        Head:head,
        Tail:tail,
    }
}


func (this *LRUCache) Get(key int) int {
    if node,ok := this.Dict[key];ok {
        this.remove(node)
        this.add(node.Key,node.Val)
        return node.Val
    }
    
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if node,ok := this.Dict[key]; ok {
        this.remove(node)
        this.add(key,value)
    }else {
        if len(this.Dict) ==this.Cap {
            this.remove(this.Tail.Prev)
            this.add(key,value)
        }else if len(this.Dict) < this.Cap {
            this.add(key,value)
        }
    }
}

func (this *LRUCache) remove(node *Node){
    delete(this.Dict,node.Key)
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
}

func (this *LRUCache) add(key,val int){
    var temp =&Node{}
    temp.Key = key
    temp.Val = val
    temp.Next = this.Head.Next
    this.Head.Next.Prev = temp
    
    temp.Prev = this.Head
    this.Head.Next = temp
    
    this.Dict[key] = temp
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