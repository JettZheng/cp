package advanced

func ConsturctBIT(arr []int) []int{
	n := len(arr)
	BITree := make([]int,n+1)

	for i := 0; i < n; i++ {
		updateBIT(BITree,n,i,arr[i])
	}

	return BITree
}

func updateBIT(bitree []int,n,index,val int){
	index += 1
	for index <= n {
		bitree[index] += val

		index+= lowbit(index)
	}
}

func queryBIT(bitree []int,index int) int{
	 var sum int

	 index+=1
	 for index > 0 {
		 sum +=bitree[index]

		 index = index - lowbit(index)
	 }

	 return sum

}

func lowbit(a int) int {
	return a & -(a)
}

func testbitree() {
	tree := ConsturctBIT([]int{1,2,3,4,5,6,7,8})
	print(queryBIT(tree,7))
}