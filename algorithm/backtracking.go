package algorithm

// 
func permute(nums []int) [][]int {
    var res [][]int
    if len(nums) == 1 {
        res = append(res,nums)
        return res
    }

    for i := range nums {
        root := nums[i]
		var temp []int
        temp = append(temp,nums[:i]...)
		temp = append(temp,nums[i+1:]...)
        resLevel := permute(temp)

        for j := range resLevel {
            resLevel[j] = append(resLevel[j],root)
            res = append(res,resLevel[j])
        }
    }

    return res

}


