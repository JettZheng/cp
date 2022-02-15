package datastructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144 https://leetcode.com/problems/binary-tree-preorder-traversal/
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

func preorderTraversal(root *TreeNode) []int {
	var res []int

	dfs(root, &res)

	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)

	return
}

func preorderTraversalI(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    
    var stack = []*TreeNode{root}
    
    for len(stack) != 0 {
        l := len(stack) - 1
        pop := stack[l]
        stack = stack[:l]
        res = append(res,pop.Val)
        
        if pop.Right != nil {
            stack = append(stack,pop.Right)
        }
        if pop.Left != nil {
            stack = append(stack,pop.Left)
        }
    }
    
    return res
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return  0
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

    var lm,rm,rootEdge int
    if root.Left != nil {
       lm = maxDepth(root.Left) - 1
        rootEdge+=1
    }
    
    if root.Right != nil {
         rm = maxDepth(root.Right) -1
        rootEdge+=1
    }
    
    res = max(res,lm+rm+rootEdge)
    
    return  res
}

func max(a,b int)int {
    if a > b {
        return a
    }
    
    return b
}