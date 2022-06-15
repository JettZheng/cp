package datastructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeColor struct {
	Node  *TreeNode
	Color int //0 white 1 grey
}

func preorderRecursive(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderRecursive(root.Left)...)
	res = append(res, preorderRecursive(root.Right)...)
	return res
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack = []*TreeNode{root}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append([]int{top.Val}, res...)

		if top.Left != nil {
			stack = append(stack, top.Left)
		}

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}

	return res
}

func preorderTraversalI(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack = []*TreeNode{root}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, pop.Val)

		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}
	}

	return res
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	var top = root
	for len(stack) > 0 || top != nil {
		for top != nil {
			stack = append(stack, top)
			top = top.Left
		}

		top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, top.Val)

		top = top.Right
	}

	return res
}

func traversalByColorMark(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var colorRoot = &TreeNodeColor{}
	colorRoot.Node = root
	var stack = []*TreeNodeColor{colorRoot}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Node == nil {
			continue
		}

		if top.Color == 1 {
			res = append(res, top.Node.Val)
		} else {
			var rightNode = &TreeNodeColor{}
			rightNode.Node = top.Node.Right
			stack = append(stack, rightNode)

			var leftNode = &TreeNodeColor{}
			leftNode.Node = top.Node.Left
			stack = append(stack, leftNode)

			top.Color = 1
			stack = append(stack, top)
		}
	}

	return res
}

func levelOrderTraversal(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		var temp []int
		for levelSize > 0 {
			top := queue[0]
			queue = queue[1:]

			temp = append(temp, top.Val)

			if top.Left != nil {
				queue = append(queue, top.Left)
			}

			if top.Right != nil {
				queue = append(queue, top.Right)
			}

			levelSize--
		}

		res = append(res, temp)
	}

	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1

	if l > r {
		return l
	}

	return r
}

var res int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var lm, rm, rootEdge int
	if root.Left != nil {
		lm = maxDepth(root.Left) - 1
		rootEdge += 1
	}

	if root.Right != nil {
		rm = maxDepth(root.Right) - 1
		rootEdge += 1
	}

	res = max(res, lm+rm+rootEdge)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l == nil && r == nil {
		return nil
	}

	if r != nil {
		return r
	}

	return l
}

func rightSideView(root *TreeNode) []int {
	var res []int

	var queue []*TreeNode
	queue = append(queue, root)

	var levelCount int
	for len(queue) > 0 {
		levelCount = len(queue)

		var head *TreeNode

		for levelCount > 0 {
			head = queue[0]
			queue = queue[1:]

			if levelCount == 1 {
				res = append(res, head.Val)
			}
			levelCount--

			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
	}
	return res
}

// TODO
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 0
}

func maxThree(a, b, c int) int {
	var temp int
	if a > b {
		temp = a
	} else {
		temp = b
	}

	if temp > c {
		return temp
	}
	return c
}

type Part struct {
	Node   *TreeNode
	Number int
}

func widthOfBinaryTree(root *TreeNode) int {
	maxWidth := 0
	if root == nil {
		return maxWidth
	}
	queue := []*Part{{root, 1}}
	for len(queue) > 0 {
		levelSize := len(queue)
		maxWidth = Max(maxWidth, queue[levelSize-1].Number-queue[0].Number+1)

		for levelSize > 0 {
			top := queue[0]
			queue = queue[1:]

			if top.Node.Left != nil {
				queue = append(queue, &Part{top.Node.Left, 2 * top.Number})
			}

			if top.Node.Right != nil {
				queue = append(queue, &Part{top.Node.Right, 2*top.Number + 1})
			}
			levelSize--
		}

	}
	return maxWidth
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}
	var level = 1

	for len(queue) > 0 {
		levelSize := len(queue)

		var isReverse bool
		if level%2 == 0 {
			isReverse = true
		}
		var temp []int
		for levelSize > 0 {
			top := queue[0]
			queue = queue[1:]

			if isReverse {
				temp = append([]int{top.Val}, temp...)
			} else {
				temp = append(temp, top.Val)
			}

			if top.Left != nil {
				queue = append(queue, top.Left)
			}

			if top.Right != nil {
				queue = append(queue, top.Right)
			}

			levelSize--
		}

		level++
		res = append(res, temp)
	}

	return res
}


var MMax int
func diameterOfBinaryTree2(root *TreeNode) int {
       if root == nil {
        return 0
    }
    dfs(root)

    return MMax
}

func dfs(root *TreeNode) int {

    if root.Left == nil && root.Right == nil {
        return 0
    }

    var l,r int
    if root.Left != nil {
        l = dfs(root.Left) + 1
    }
    if root.Right != nil {
        r = dfs(root.Right) + 1
    }

    MMax = max(MMax,l+r)

    return max(l,r)
}