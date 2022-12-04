package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	total := 0
	count := 0
	group := ^uint64(0)

	for sc.Scan() {
		var rucksack uint64
		for _, item := range sc.Text() {
			rucksack |= 1 << priority(item)
		}

		group &= rucksack

		if count++; count == 3 {
			total += bits.TrailingZeros64(group)
			group = ^uint64(0)
			count = 0
		}
	}

	fmt.Println(total)
}

func priority(item rune) (prio int) {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
	}
	return int(item-'A') + 27
}
