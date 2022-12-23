package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		began = time.Now()
		grid  = map[point]bool{}
		sc    = bufio.NewScanner(os.Stdin)
		maxY  uint16
	)

	for sc.Scan() {
		prev, path := parse(sc.Bytes())

		for len(path) > 0 {
			var next, src, dst point

			next, path = parse(path)
			src.x, dst.x = minMax(prev.x, next.x)
			src.y, dst.y = minMax(prev.y, next.y)
			_, maxY = minMax(maxY, dst.y)

			for x := src.x; x <= dst.x; x++ {
				for y := src.y; y <= dst.y; y++ {
					grid[point{x, y}] = true
				}
			}

			prev = next
		}
	}

	floor := maxY + 2
	total := 0

	for !grid[point{500, 0}] {
		sand := point{500, 0}

	drop:
		for {
			switch {
			case sand.y+1 == floor:
				break drop
			case !grid[point{sand.x, sand.y + 1}]:
				sand.y++
			case !grid[point{sand.x - 1, sand.y + 1}]:
				sand.x--
				sand.y++
			case !grid[point{sand.x + 1, sand.y + 1}]:
				sand.x++
				sand.y++
			default:
				break drop
			}
		}

		grid[sand] = true
		total++
	}

	fmt.Println(total, time.Since(began))
}

type point struct{ x, y uint16 }

func parse(p []byte) (point, []byte) {
	xp, p, _ := bytes.Cut(p, []byte(","))
	x, _ := strconv.ParseUint(string(xp), 10, 16)
	yp, p, _ := bytes.Cut(p, []byte(" "))
	y, _ := strconv.ParseUint(string(yp), 10, 16)
	_, p, _ = bytes.Cut(p, []byte(" "))
	return point{uint16(x), uint16(y)}, p
}

func minMax(a, b uint16) (uint16, uint16) {
	if a < b {
		return a, b
	}
	return b, a
}
