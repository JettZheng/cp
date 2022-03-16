package algorithm

// idea: 1. we iterative all elements in grid, find the value which equals to 1,
func numIslands(grid [][]byte) int {
	var res int

	var height = len(grid)
	var width = len(grid[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '1' {
				spread(&grid, i, j, height, width)
				res++
			}
		}
	}
	return res
}

func spread(grid *[][]byte, i, j, height, width int) {
	if i >= height || j >= width || i < 0 || j < 0 {
		return
	}
	if (*grid)[i][j] == '0' {
		return
	}

	(*grid)[i][j] = '0'

	spread(grid, i-1, j, height, width)
	spread(grid, i, j-1, height, width)
	spread(grid, i+1, j, height, width)
	spread(grid, i, j+1, height, width)
}
