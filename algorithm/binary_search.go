package algorithm

/*
 1.[l,r)
 2.binary search should use sorted array
 3.binary search should use two pointers to check if l < r
*/
func binarySearchIteratively(nums []int, target int) int {
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

func search(nums []int, target int) int {
	l := 0
	r := len(nums)-1

    return binarySearchRecursively(nums,target,l,r)
}

func binarySearchRecursively(nums []int, target int,l,r int) int {
	m := l + (r-l)/2

	if l > r {
		return -1
	}

	if nums[m] == target {
		return m
	}

	if nums[m] > target {
		return binarySearchRecursively(nums,target,l,m-1)
	}

	if nums[m] < target {
		return binarySearchRecursively(nums,target,m+1,r)
	}

	return -1
}
