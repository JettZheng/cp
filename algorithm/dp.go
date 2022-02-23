package algorithm


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