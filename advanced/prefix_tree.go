package advanced

type Trie struct {
	HasWord   bool
	Links []*Trie
}

func PConstructor() Trie {
	return Trie{
		Links: make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	if len(word) <= 0 {
		return
	}
	a := int(rune('a'))

	var count int
	curr := this

	for _,v := range word {
		b := int(v) - a
		count++
		if curr.Links[b] == nil {
			t := PConstructor()
			curr.Links[b] = &t
		}
		if count == len(word) {
			curr.Links[b].HasWord = true
		}
		curr = curr.Links[b]
	}
}

func (this *Trie) Search(word string) bool {
	a := int(rune('a'))

	curr := this
	for _, v := range word {
		b := int(v) - a

		if curr.Links[b] == nil {
			return false
		}

		curr = curr.Links[b]
	}

	return curr.HasWord
}

func (this *Trie) StartsWith(prefix string) bool {
	a := int(rune('a'))

	curr := this
	for _, v := range prefix {
		b := int(v) - a

		if curr.Links[b] == nil {
			return false
		}

		curr = curr.Links[b]
	}

	return true
}

func testPrefixTree() {
	tree := PConstructor()
	tree.Insert("apple")
	print(tree.Search("apple"))
	print(tree.Search("app"))
	print(tree.StartsWith("app"))
	tree.Insert("app")
	print(tree.Search("app"))
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
