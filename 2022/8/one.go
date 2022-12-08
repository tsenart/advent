package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var grid [][]int8

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		row := make([]int8, len(line))
		for i, c := range line {
			row[i] = int8(c - '0')
		}
		grid = append(grid, row)
	}

	// Running max per traversal direction.
	max := [4][]int8{
		make([]int8, len(grid[0])), // Down
		make([]int8, len(grid)),    // Right
		make([]int8, len(grid[0])), // Up
		make([]int8, len(grid)),    // Left
	}

	// Initialize maximums to -1
	for i := range max {
		for j := range max[i] {
			max[i][j] = -1
		}
	}

	// A tree is counted as visible only once, even if
	// it's visible from multiple starting points.
	visible := make([][]bool, len(grid))
	for i := range visible {
		visible[i] = make([]bool, len(grid[0]))
	}
	total := 0

	visit := func(row, col int, max *int8) {
		if grid[row][col] > *max {
			if *max = grid[row][col]; !visible[row][col] {
				visible[row][col] = true
				total++
			}
		}
	}

	// We do two scans at once: down/right and up/left
	for down := range grid {
		up := len(grid) - 1 - down
		for right := range grid[down] {
			left := len(grid[down]) - 1 - right
			visit(down, right, &max[0][right])
			visit(down, right, &max[1][down])
			visit(up, left, &max[2][left])
			visit(up, left, &max[3][up])
		}
	}

	fmt.Println(total)
}
