package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("/Users/tomas/Code/advent/2022/14/example.in")
	sc := bufio.NewScanner(f)

	type point struct{ x, y int }

	var (
		paths [][]point
		min   = point{x: math.MaxInt, y: math.MaxInt}
		max   point
	)

	for sc.Scan() {
		var path []point
		for _, p := range strings.Split(sc.Text(), " -> ") {
			xy := strings.SplitN(p, ",", 2)
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])

			path = append(path, point{x, y})

			if x > max.x {
				max.x = x
			} else if x < min.x {
				min.x = x
			}

			if y > max.y {
				max.y = y
			} else if y < min.y {
				min.y = y
			}
		}
		paths = append(paths, path)
	}

	width := (max.x - min.x) + 1
	height := max.y + 1
	normalize := func(p point) point {
		if min.x > 0 {
			p.x -= min.x
		}
		return p
	}

	fmt.Printf("width: %d, height: %d, min: %+v, max: %+v", width, height, min, max)
	grid := make([]byte, width*height)

	for _, path := range paths {
		for len(path) > 1 {
			src, dst := normalize(path[0]), normalize(path[1])
			dx := dst.x - src.x
			dy := dst.y - src.y

		line:
			for {
				grid[src.y*width+src.x] = '#'
				switch {
				case dx != 0:
					s := sign(dx)
					src.x += s
					dx -= s
				case dy != 0:
					s := sign(dy)
					src.y += s
					dy -= s
				default:
					break line
				}
			}

			path = path[1:]
		}
	}

	debug(grid, width)
}

func sign(n int) int {
	if n > 0 {
		return 1
	}
	return -1
}

func debug(grid []byte, width int) {
	for i, cell := range grid {
		if i%width == 0 {
			fmt.Printf("\n%d ", i/width)
		}

		if cell == 0 {
			fmt.Printf("%s", ".")
		} else {
			fmt.Printf("%c", cell)
		}
	}
	fmt.Println()
}
