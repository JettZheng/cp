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

func subsets(nums []int) [][]int {
    var res [][]int
    var backtrack func (i int,curr []int)
    
    backtrack = func(i int,curr []int){
        if i > len(nums){
            return
        }
        
        if i == len(nums) {
            res = append(res,curr)
            return
        }
        
        curr = append(curr,nums[i])


        
        backtrack(i+1,curr)
        
        curr = curr[:len(curr)-1]
        

        backtrack(i+1,curr)
    }
    
    backtrack(0,[]int{})
    
    return res
}


