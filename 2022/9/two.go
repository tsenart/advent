package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const x, y = 0, 1

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var (
		knots [10][2]int
		set   = make(map[[2]int]bool)
	)

	for sc.Scan() {
		line := sc.Text()
		step, _ := strconv.Atoi(line[2:])

		var move [2]int
		switch line[:1] {
		case "U":
			move[y] = step
		case "D":
			move[y] = -step
		case "R":
			move[x] = step
		case "L":
			move[x] = -step
		}

		for {
			var d int
			for i := x; i <= y; i++ {
				if move[i] != 0 {
					d = delta(move[i])
					move[i] += -d
					knots[0][i] += d
					break
				}
			}

			if d == 0 {
				break
			}

			for i := 0; i < len(knots)-1; i++ {
				dx := knots[i][x] - knots[i+1][x]
				dy := knots[i][y] - knots[i+1][y]

				if abs(dx) > 1 || abs(dy) > 1 {
					knots[i+1][x] += delta(dx)
					knots[i+1][y] += delta(dy)
				}
			}

			if !set[knots[len(knots)-1]] {
				set[knots[len(knots)-1]] = true
			}
		}
	}

	fmt.Println(len(set))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func delta(x int) int {
	switch {
	case x == 0:
		return 0
	case x < 0:
		return -1
	default:
		return 1
	}
}
