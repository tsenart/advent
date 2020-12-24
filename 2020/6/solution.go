package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	scan := true
	any := uint32(0)  // part 1
	all := ^uint32(0) // part 2
	solutions := [2]int{}

	for scan {
		if scan = sc.Scan(); scan && sc.Text() != "" {
			person := uint32(0)

			for _, question := range sc.Text() {
				// Map 'a' to the last bit, 'b' to the one before, etc
				person |= 1 << (int(question) - int(rune('a')))
			}

			any |= person
			all &= person
		} else {
			solutions[0] += bits.OnesCount32(any)
			solutions[1] += bits.OnesCount32(all)
			any, all = 0, ^uint32(0)
		}
	}

	fmt.Println(solutions)
}
