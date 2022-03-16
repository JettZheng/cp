package datastructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
        
        res = append([]int{top.Val},res...)
        
        if top.Left != nil {
            stack = append(stack,top.Left)
        }
        
        if top.Right != nil {
            stack = append(stack,top.Right)
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
		pop := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
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

			levelSize --
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

	if l == nil && r != nil {
		return r
	}

	if l != nil && r == nil {
		return l
	}

	return root
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
