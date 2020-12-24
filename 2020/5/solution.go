package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var solutions [2]int
	var ids []int

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		row := bounds{0, 127}
		col := bounds{0, 7}

		for _, dir := range sc.Text() {
			switch dir {
			case 'F':
				row.hi -= half(row)
			case 'B':
				row.lo += half(row)
			case 'R':
				col.lo += half(col)
			case 'L':
				col.hi -= half(col)
			default:
			}
		}

		ids = append(ids, row.lo*8+col.lo)
	}

	sort.Ints(ids)

	solutions[0] = ids[len(ids)-1]

	for i, id := range ids[:len(ids)-1] {
		// Find any gap in the sorted list of seat ids
		if ids[i+1] != id+1 {
			solutions[1] = id + 1
		}
	}

	if solutions[1] == 0 {
		solutions[1] = ids[len(ids)-1] + 1
	}

	fmt.Println(solutions)
}

type bounds struct{ lo, hi int }

func half(b bounds) int {
	return int(math.Ceil(float64(b.hi-b.lo) / 2.0))
}
