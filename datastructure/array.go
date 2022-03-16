package datastructure

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var j, k int
	var temp = make([]int, m)
	for i := range nums1[:m] {
		temp[i] = nums1[i]
	}

	for i := 0; i < m+n && j < m && k < n; i++ {
		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}

	if j >= m {
		for i := j + 1; i < m+n; i++ {
			nums1[i] = nums2[k]
			k++
		}
		return
	}

	if k >= n {
		for i := k + 1; i < m+n; i++ {
			nums1[i] = temp[j]
			j++
		}
	}

}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	var res [][]int

	// matrix to array

	var d1 []int
	for _, v := range mat {
		d1 = append(d1, v...)
	}

	var temp []int
	for i := range d1 {
		temp = append(temp, d1[i])
		if (i+1)%4 == 0 {
			res = append(res, temp)
			temp = []int{}
		}
	}

	return res
}

// [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	var rows []int
	rows = []int{1}
	res[0] = rows
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		var rows = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				rows[j] = 1
			} else {
				rows[j] = res[i-1][j-1] + res[i-1][j]
			}
		}

		res[i] = rows
	}

	return res
}

// Input: matrix = [[1,3]], target = 3
// Output: true
func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
    if rows == 0 {
        return false
    }
    cols := len(matrix[0])

    var l,r,m int

    r = rows * cols

    for l < r {
        m = l + (r-l)/2

        if matrix[m/cols][m%cols] == target {
            return true
        }

        if matrix[m/cols][m%cols] > target {
            r = m
        }
        if matrix[m/cols][m%cols] < target {
            l = m + 1
        }
    }

    return false
}

func mergeIntervals(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals,func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	res = append(res, intervals[0])
	var j int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[j][1] {
			res = append(res, intervals[i])
			j++
		} else {
			res[j][1] = intervals[i][1]
		}
	}

	return res
}
