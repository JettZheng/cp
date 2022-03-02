package algorithm

func findKthLargest(nums []int, k int) int {
	pivotValue := nums[len(nums)-1]

	var p int

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= pivotValue {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}

	nums[p], nums[len(nums)-1] = nums[len(nums)-1], nums[p]

	if p == k-1 {
		return nums[p]
	} else if p > k-1 {
		return findKthLargest(nums[:p], k)
	} else {
		return findKthLargest(nums[p+1:], k-p-1)
	}
}
