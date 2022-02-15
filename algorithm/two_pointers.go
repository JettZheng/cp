package algorithm

func rotate(nums []int, k int) {

	if len(nums) == 1 {
		return
	}

	k = k % len(nums)

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	var i,j,temp int
	i = 0
	j = len(nums) -1

	for i < j {
		temp = nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
}

func sortedSquares(nums []int) []int {
   var res = make([]int,len(nums))
    var i,j,k int
	i=0
	j=len(nums) -1
	k= len(nums) -1
	for i <= j {
		if nums[i] * nums[i] < nums[j] * nums[j] {
			res[k] = nums[j] * nums[j]
			j--
		}else {
			res[k] = nums[i] * nums[i]
			i++
		}
		k--
	}

	return res
}

func moveZeroes(nums []int)  {
    var j int
	
	for i := range nums {
		if nums[j] == 0 && nums[i] != 0 {
			nums[j],nums[i]= nums[i],nums[j]
		}

		if nums[j] != 0 {
			j++
		}
	}
}

func twoSum(numbers []int, target int) []int {
    var res []int
	var i,j int
	j = len(numbers) -1
	for i <= j {
		if numbers[i] + numbers[j] > target {
			j-=1
		}
		if numbers[i] + numbers[j] < target {
			i+=1
		}
		if numbers[i] + numbers[j] == target{
			return []int {i,j}
		}
	}

	return res
}