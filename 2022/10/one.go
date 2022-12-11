package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var q []int
	sum, cycle, x := 0, 0, 1

	for sc.Scan() || len(q) > 0 {
		cycle++

		if len(q) > 0 {
			n := q[0]
			q = q[1:]
			x += n
		}

		if (cycle+20)%40 == 0 && cycle <= 220 {
			strength := x * cycle
			sum += strength
		}

		switch line := sc.Text(); {
		case line == "":
			continue
		case line[:4] == "noop":
			q = append(q, 0)
		case line[:4] == "addx":
			x, _ := strconv.Atoi(line[5:])
			q = append(q, 0, x)
		}
	}

	fmt.Println(sum)
}
