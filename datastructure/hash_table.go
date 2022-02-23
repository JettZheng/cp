package datastructure

func solution(input map[int]int) int {

	var min, minKey int

	min =int(^uint(0) >> 1)

	for k, v := range input {
		if intabs(v,k) > min {
			continue
		}

		if intabs(v,k) == min {
			if k < minKey {
				minKey = k
			}
		}

		if intabs(v,k) < min {
			min = intabs(v,k)
			minKey = k
		}
	}

	return minKey

}

func intabs(a,b int) int {
	if a-b >=0 {
		return a -b 
	}else {
		return b -a
	}
}
