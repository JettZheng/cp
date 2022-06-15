package algorithm

/*
 1.[l,r)
 2.binary search should use sorted array
 3.binary search should use two pointers to check if l < r
*/
func binarySearchIteratively(nums []int, target int) int {
	l := 0
	r := len(nums)-1
	var m int

	for l <= r {
		m = l + (r-l)/2
		if nums[m] > target {
			r = m - 1
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

func binarySearchLower(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var m int
	var i  = -1
	for l <= r {
		m = l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		}
		if nums[m] < target {
			l = m + 1
		}
		if nums[m] == target {
			i = m
			r = m - 1
		}
	}

	return i
}

func binarySearchUpper(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var m int
	var i  = -1
	for l <= r {
		m = l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		}
		if nums[m] < target {
			l = m + 1
		}
		if nums[m] == target  {
			i = m
			l = m + 1
		}
	}

	return i
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

func minArray(numbers []int) int {  
    l := 0
    r := len(numbers) -1

    for l < r {
        m := l + (r-l) /2

        if numbers[m] < numbers[r] {
            r = m
        } else if numbers[m] > numbers[r] {
            l = m + 1
        } else {
            r -- 
        }
    }

    return numbers[l]
}

// 4566712 6
// 7123456
func hsearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		if arr[l] == target {
			return l
		}

		m := l + (r - l)/2

		if arr[m] == target {
			r = m
			continue
		}

		if arr[m] > arr[0] {
			if target >= arr[0] && target < arr[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if arr[m] < arr[0] {
			if target > arr[m] && target <= arr[len(arr)-1] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			l++
		}
	}

	return -1
}
