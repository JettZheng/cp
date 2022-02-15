package algorithm

import "sort"

func search(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l := 0
	r := len(nums)
	var m int

	for l < r {
		m = l + (r-l)/2
		if nums[m] > target {
			r = m
		}
		if nums[m] < target {
			l = m + 1
		}
		if nums[m] == target {
			return m
		}

	}

	return -1
}
