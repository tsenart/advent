package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	began := time.Now()
	sc := bufio.NewScanner(os.Stdin)

	var (
		heightmap []byte
		width     int
	)

	for sc.Scan() {
		row := sc.Bytes()
		if width == 0 {
			width = len(row)
		}
		heightmap = append(heightmap, row...)
	}

	var (
		start = bytes.IndexByte(heightmap, 'S')
		end   = bytes.IndexByte(heightmap, 'E')
		queue = [][2]int{{start, 0}}
		path  = make([]int, len(heightmap))
	)

	for i := range path {
		path[i] = -1
	}

	heightmap[start] = 'a'
	heightmap[end] = 'z'

	for len(queue) > 0 {
		frontier := queue[0]
		queue = queue[1:]

		i, steps := frontier[0], frontier[1]
		row, col := i/width, i%width

		neighbours := make([]int, 0, 4)
		if row > 0 {
			neighbours = append(neighbours, (row-1)*width+col)
		}

		if row < (len(heightmap)/width)-1 {
			neighbours = append(neighbours, (row+1)*width+col)
		}

		if col > 0 {
			neighbours = append(neighbours, row*width+(col-1))
		}

		if col < width-1 {
			neighbours = append(neighbours, row*width+(col+1))
		}

		for _, n := range neighbours {
			if path[n] == -1 && heightmap[n] <= heightmap[i]+1 {
				if path[n] = i; n == end {
					fmt.Println(steps+1, time.Since(began))
					plot(heightmap, width, start, n, path)
					queue = nil
					break
				}
				queue = append(queue, [2]int{n, steps + 1})
			}
		}
	}
}

func plot(heightmap []byte, width int, start, pos int, path []int) {
	trace := make([]bool, len(heightmap))
	for n := pos; n != start; n = path[n] {
		trace[n] = true
	}

	for i := range heightmap {
		if i%width == 0 {
			fmt.Println()
		}

		if i == pos || trace[i] {
			fmt.Printf("\x1b[31m%c\x1b[0m", heightmap[i])
		} else if path[i] != -1 {
			fmt.Printf("\x1b[30m%c\x1b[0m", heightmap[i])
		} else {
			fmt.Printf("%c", heightmap[i])
		}
	}

	fmt.Println()
}
