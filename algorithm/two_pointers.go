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

// 2371 9
/*
2 0
3 1
7 2
1 3
*/
func twoSum(nums []int, target int) []int {
	var dict = make(map[int]int,len(nums))

	var res []int
	for i := range nums {
		if index,ok := dict[target - nums[i]];ok {
			res = append(res, []int{nums[index],nums[i]}...)
			break
		} else {
			dict[nums[i]] = i
		}
	}

	return res
}


func threeSum(nums []int) [][]int {
    var targets = make([]int,len(nums))

	for i := range nums {
		targets[i] = 0-nums[i]
	}

	var res [][]int

	for i := range nums {
		arr := append(nums[:i],nums[i+1:]...)
		temp := twoSum(arr,targets[i])
		if len(temp) > 0 {
            temp = append(temp,nums[i])
			res = append(res, temp)
		}
	}

	return res
}

func testappend(){

	var a = []int{-2,-4,4,3,1}
	var b = append(a[:1],a[2:]...)
	print(b)
}