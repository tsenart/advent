package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		safe     int
		dampened int
		level    []int
		sc       = bufio.NewScanner(os.Stdin)
	)

	for sc.Scan() {
		fields := strings.Fields(sc.Text())

		if cap(level) < len(fields) {
			level = make([]int, len(fields))
		} else {
			level = level[:len(fields)]
		}

		for i, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				log.Fatal(err)
			}
			level[i] = n
		}

		if !unsafe(level, -1) {
			safe++
			continue
		}

		for i := range level {
			if !unsafe(level, i) {
				dampened++
				break
			}
		}
	}

	fmt.Println(safe, safe+dampened)
}

func unsafe(level []int, skip int) bool {
	var (
		prev      int
		prevDelta int
		first     = true
	)

	for i, n := range level {
		if i == skip {
			continue
		}

		if first {
			first, prev = false, n
			continue
		}

		delta := n - prev
		absDelta := abs(delta)
		if absDelta < 1 || absDelta > 3 || (prevDelta != 0 && sign(prevDelta) != sign(delta)) {
			return true
		}

		prev, prevDelta = n, delta
	}

	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}
