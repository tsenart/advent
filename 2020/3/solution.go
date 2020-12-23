package main

import (
	"bufio"
	"fmt"
	"os"
)

type slope struct{ x, y int }

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var grid [][]rune
	for sc.Scan() {
		var row []rune
		for _, r := range sc.Text() {
			row = append(row, r)
		}
		grid = append(grid, row)
	}

	solution := trees(grid, slope{3, 1})
	fmt.Printf("Part one: %d\n", solution)

	for _, slope := range []slope{
		{1, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		solution *= trees(grid, slope)
	}

	fmt.Printf("Part two: %d\n", solution)
}

func trees(grid [][]rune, slope slope) (count int) {
	x, y := 0, 0
	for y < len(grid) {
		if grid[y][x%len(grid[y])] == '#' {
			count++
		}

		x += slope.x
		y += slope.y
	}

	return count
}
