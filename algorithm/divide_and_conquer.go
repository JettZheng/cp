package algorithm

func maxSubArray(nums []int) int {
	return helper(nums,0,len(nums)-1)
}

func helper(nums []int, l,r int) int {
	if l >= r {
		return nums[r]
	}

	m := l + (r-l)/2

	lmax := helper(nums,l,m-1)
	rmax := helper(nums,m+1,r)

	//left side max
	var lmmax int
	for i := m; i >= l; i-- {
		lmmax = max(nums[m],nums[m]+lmmax)
	}

	//right side max
	var rmmax int
	for i := m+1; i <= r; i++ {
		rmmax = max(nums[m],nums[m]+rmmax)
	}



	return max(max(lmax,rmax),lmmax+rmax)
}
