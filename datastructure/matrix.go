package datastructure

func spiralOrder(matrix [][]int) []int {
	var res []int

	var m = len(matrix)
	var n = len(matrix[0])

	var start, end, lstart, lend int
	end = n - 1
	lend = m - 1

	for start <= end && lstart <= lend {
		for i := start; i <= end; i++ {
			res = append(res, matrix[lstart][i])
		}
		lstart++
		if lstart > lend || start > end {
			break
		}

		for i := lstart; i <= lend; i++ {
			res = append(res, matrix[i][end])
		}
		end--
		if lstart > lend || start > end {
			break
		}

		for i := end; i >= start; i-- {
			res = append(res, matrix[lend][i])
		}
		lend--
		if lstart > lend || start > end {
			break
		}

		for i := lend; i >= lstart; i-- {
			res = append(res, matrix[i][start])
		}
		start++
	}

	return res
}
