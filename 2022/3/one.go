package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sum uint64
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		rucksack := sc.Text()
		var set uint64
		for i, item := range rucksack {
			prio := priority(item)
			mask := uint64(1) << prio
			if i < len(rucksack)/2 {
				set |= mask
			} else if set&mask == mask {
				sum += prio
				set &^= mask // delete
			}
		}
	}
	fmt.Println(sum)
}

func priority(item rune) uint64 {
	if item >= 'a' && item <= 'z' {
		return uint64(item-'a') + 1
	}
	return uint64(item-'A') + 27
}
