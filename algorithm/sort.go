package algorithm

// quickSort algorithm is random pick a pivot,then according pivot to do partition.
func quickSort(nums *[]int,l,r int) {
	if l < r  {
		pivot := partition(*nums,l,r)

	quickSort(nums,l,pivot-1)
	quickSort(nums,pivot+1,r)
	}
}

func partition(nums []int,l,r int) int {
	pivotValue := nums[r]

	var p = l
	for i := l; i < r; i++ {
		if nums[i] < pivotValue{
			nums[i],nums[p] = nums[p],nums[i]
			p++
		}
	}

	nums[p],nums[r] = nums[r],nums[p]

	return p
}