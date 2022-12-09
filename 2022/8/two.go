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

	var max int
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			north := row - 1
			for north > 0 && grid[north][col] < grid[row][col] {
				north--
			}

			south := row + 1
			for south < len(grid)-1 && grid[south][col] < grid[row][col] {
				south++
			}

			west := col - 1
			for west > 0 && grid[row][west] < grid[row][col] {
				west--
			}

			east := col + 1
			for east < len(grid[row])-1 && grid[row][east] < grid[row][col] {
				east++
			}

			score := (row - north) * (south - row) * (col - west) * (east - col)
			if score > max {
				max = score
			}
		}
	}

	fmt.Println(max)
}
