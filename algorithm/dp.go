package algorithm

import "math"



func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res = math.MinInt32

	var dpmax = make([]int,len(nums))
	var dpmin = make([]int,len(nums))

	dpmax[0] = nums[0]
	dpmin[0] = nums[0]

	res = max(dpmax[0],res)
	for i := 1; i <= len(nums)-1; i ++ {
		dpmax[i] = max(nums[i],max(nums[i] * dpmax[i-1],nums[i] * dpmin[i-1]))
		dpmin[i] = minTwo(nums[i],minTwo(nums[i] * dpmin[i-1],nums[i] * dpmax[i-1]))
		res = max(dpmax[i],res)
	}

	return res
}

func jett(nums []int,l, r int) int {
	if l >= r {
		return nums[r]
	}

	m  := l + (r-l)/2

	lmax := jett(nums,l,m-1)
	rmax := jett(nums,m+1,r)


	var lm,rm,sum int

	for i := m-1; i >= 0; i-- {
		lm = max(lm,nums[i]+sum)
		sum +=nums[i]
	}


	return max(max(lmax,rmax),lm+rm+nums[m])
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

// jump game.
// [2,3,1,1,4]
// res[4] = true
// dp[3] = true
// dp[2] = 

func canJump2(nums []int) bool {
	lastPos := len(nums) - 1

	for i := 0; i >=0; i-- {
		if nums[lastPos] + i >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}

func canJump(nums []int) bool {
	n := len(nums)
	res := make([]bool,n)
	
	for i:= n-1;i>=0;i-- {
		if i >= n - 1 {
			res[i] = true
		}
		if nums[i] + i >= n -1 {
			res[i] = true
		}
		for j := 1; j <= nums[i]; j++ {
			if i + j >= n -1 {
				res[i] = true
				break
			}
			if res[i+j]{
				res[i] = true
				break
			}
		}
	}

	return res[0]
}

func maxProfit(prices []int) int {
    _,p :=dp(prices,0)
    return p
}

func dp(prices []int,i int) (lpoint int,profit int) {
    if i == len(prices) - 1 {
        return i,0
    }
    
    if i == len(prices) - 2 {
        if prices[i+1] > prices[i] {
            return i, prices[i+1] - prices[i]
        } else {
            return i, 0
        }
    }
    
    l,p := dp(prices,i+1)
    
    if prices[i] >= prices[l] {
        return l,p
    }
    
    return i,prices[i]-prices[l] + p
}

/* "abcbcad"
	state(j,j) = true
	state(j-1,j+1) = state(j,j) && s[j-1] + s[j+1]
*/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var state = make([][]bool,len(s))

	for i := 0; i < len(s); i++ {
		state[i] = make([]bool,len(s))
	}

	var start,end,max int

	for j := 0; j < len(s); j++ {
		state[j][j] = true//?

		for i := 0; i < j; i++ {
			if j - i == 1 {
				state[i][j] = s[i] == s[j]
			}else {
				state[i][j] = state[i+1][j-1] && s[i] == s[j]
			}

			if j - i + 1 > max {
				max = j - i + 1
				start = i
				end = j
			}
		}
	}

	return s[start:end]
}

/* 
300.https://leetcode.com/problems/longest-increasing-subsequence/
idea:

1.使用动态规划
状态：用一维数组记录每个以该元素结尾的最大长度是多少 dp[i]
状态转移： dp[i] = nums[i]> nums[j] 则说明形成有序，可以在该元素上追加一个， 0->i-1 遍历迭代记录每个位置是否形成有序，找出最大值。

遍历整个dp数组，找出最大值即可
*/

func lengthOfLIS(nums []int) int {
    var dp = make([]int, len(nums))
    
    for i := range dp {
        dp[i] = 1
    }
    
    var res = 1
    for i := 1; i < len(nums); i ++ {
        var dmax = 1
        for j := 0; j < i; j ++ {
            if nums[i] > nums[j] {
                dmax = max(dmax,dp[j] + 1)
            }
        }
        
        dp[i] = dmax
        res = max(res,dmax)
    }
    
    return res
}

//  dp[i] = dp[i-1] + dp[i-2]
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    var a,b,c int
    a = 1
    b = 2
    
    for i:=2;i<n;i++ {
        c = a+b
        a = b
        b = c
    }
    
    return c
}

/*
	状态：保存每一个点的最大经过路径
	状态转移：如果是该列的第一个，只能计算上方的最大值，往下走 一个选择
	如果是该列其他个，则需要计算 上方走下来，还是左边走过来的 两个值比较大小
*/
func maxPath(nums [][]int) int {
	var dp = make([][]int,len(nums))

	a := len(nums[0])
	for i := range dp {
		dp[i] = make([]int, a)
	}

	dp[0][0] = nums[0][0]

	var res int
	var count = 2
	for i := 1; i < len(nums); i++ {
		for j := 0; j < count; j ++ {
			if j - 1 >= 0 {
				dp[i][j] = max(dp[i-1][j],dp[i][j-1])+nums[i][j]
			} else {
				dp[i][j] = dp[i-1][j]+nums[i][j]
			}

			res = max(res,dp[i][j])
		}
		count++

	}

	return res
}


func maximalSquare(matrix [][]byte) int {
    var res int
    m := len(matrix)
    n := len(matrix[0])

    var dp = make([][]int,m)
    for i := range dp {
        dp[i] = make([]int,n)
    }
    for i := m-1; i >= 0; i -- {
        for j := n-1; j >= 0;j-- {
            if i + 1 == m || j + 1 == n {
                if matrix[i][j] == '1' {
                    dp[i][j] = 1
                }
                continue
            }
        }
    }

    for i := m-2; i >= 0; i -- {
        for j := n-2; j >= 0;j-- {
            if matrix[i][j] == '1' {
                dp[i][j] = min(dp[i+1][j],dp[i+1][j+1],dp[i][j+1]) + 1
            }

            res = max(res,dp[i][j])
        }
    }
    return res
}

func min(a,b,c int) int {
    var temp int
    if a > b {
       temp = b
    } else {
        temp = a
    }
    if temp > c {
        return c
    }
    return temp
}

func minTwo(a,b int) int {
    if a > b {
       return b
    }

    return a
}

// [5,6,7,1,2,3,4]
// [2,3,4,5,6,7,1]
func findMin(nums []int) int {
    low, high := 0, len(nums) - 1
    for low < high {
        pivot := low + (high - low) / 2
        if nums[pivot] < nums[high] {
            high = pivot
        } else {
            low = pivot + 1
        }
    }
    return nums[low]
}