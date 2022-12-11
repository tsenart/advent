package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	began := time.Now()
	sc := bufio.NewScanner(os.Stdin)

	var (
		head, tail [2]int // [x, y]
		set        = make(map[[2]int]bool)
		count      = 0
	)

	const x, y = 0, 1

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

		for move[x] != 0 || move[y] != 0 {
			switch {
			case move[x] < 0:
				move[x]++
				head[x]--
			case move[x] > 0:
				move[x]--
				head[x]++
			case move[y] < 0:
				move[y]++
				head[y]--
			case move[y] > 0:
				move[y]--
				head[y]++
			}

			dx, dy := head[x]-tail[x], head[y]-tail[y]
			switch {
			case dx > 1:
				tail[x]++
				tail[y] += dy
			case dx < -1:
				tail[x]--
				tail[y] += dy
			case dy > 1:
				tail[y]++
				tail[x] += dx
			case dy < -1:
				tail[y]--
				tail[x] += dx
			}

			if !set[tail] {
				set[tail] = true
				count++
			}
		}
	}

	fmt.Println(count, time.Since(began))
}
