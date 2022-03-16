package golang

// question: why temp := append(nums[:i],nums[i+1:]...) will change nums

func testSlice() {
	var a = []int{1,2,3,4}
	var b = append(a[:2],5)
	
	var d = a[1:2:3]
	print(d)
	// append 在底层数组容量不足时扩容，足够时在slice指向的切片追加
	var c = make([]int,1200,1200)

	// slice[a:b:c]
	// a:b or 1:3 -> gives length
	// a:c or 1:5 -> gives capacity
	var  n  =1199

	for n >=0 {
		c[n] = 1
		n--
	}

	c = append(c, 3)


	print(len(a))
	print(cap(a))
	print(a)
	print(b)
}
