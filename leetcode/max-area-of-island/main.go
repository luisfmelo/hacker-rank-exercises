package main

import "fmt"

func dfs(i, j int, grid [][]int) int {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return 0
	} else if grid[i][j] == 1 {
		grid[i][j] = 2
		return 1 + dfs(i-1, j, grid) + dfs(i+1, j, grid) + dfs(i, j-1, grid) + dfs(i, j+1, grid)
	}
	grid[i][j] = 2
	return 0
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := range grid {
		for j := range grid[0] {
			area := dfs(i, j, grid)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func main() {
	matrix := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(matrix) == 6)
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}) == 0)
}
