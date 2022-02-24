package advanced

// [1,3,4,5,6,7,8]
// []

type SegmentTreeNode struct {
	Start int
	End   int
	Sum   int
	Left  *SegmentTreeNode
	Right *SegmentTreeNode
}

func buildTree(start, end int, arr []int) *SegmentTreeNode {
	if start == end {
		return &SegmentTreeNode{
			Start: start,
			End:   end,
			Sum:   arr[start],
			Left:  &SegmentTreeNode{},
			Right: &SegmentTreeNode{},
		}
	}

	m := start + (end-start)/2

	left := buildTree(start, m, arr)
	right := buildTree(m+1, end, arr)

	return &SegmentTreeNode{
		Start: start,
		End:   end,
		Sum:   left.Sum + right.Sum,
		Left:  left,
		Right: right,
	}
}


func updateTree(root *SegmentTreeNode,index,val int) {
	if root == nil {
		return
	}
	if root.Start == root.End && root.End == index {
		root.Sum = val
		return
	}

	mid  := root.Start + (root.End - root.Start) / 2
	
	if index <= mid {
		updateTree(root.Left,index,val)
	}else {
		updateTree(root.Right,index,val)
	}

	if root.Left != nil {
		root.Sum += root.Left.Sum
	}

	if root.Right != nil {
		root.Sum += root.Right.Sum
	}
	
}

func queryRange(root *SegmentTreeNode,l,r int) int {
	if root.Start == l && root.End == r {
		return root.Sum
	}

	m := root.Start + (root.End - root.Start) / 2

	if l > m {
		return queryRange(root.Right,l,r)
	}

	if r <= m {
		return queryRange(root.Left,l,r)
	}

	return queryRange(root.Left,l,m)+ queryRange(root.Right,m+1,r)
}

func testSegementTree() {
	root := buildTree(0,5,[]int{1,2,3,4,5,6})
	updateTree(root,1,100)
	print(queryRange(root,0,2))
}
