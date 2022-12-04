package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sum int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		rucksack := sc.Text()
		set := map[rune]bool{}
		for i, item := range rucksack {
			if i < len(rucksack)/2 {
				set[item] = true
			} else if set[item] {
				sum += priority(item)
				delete(set, item)
			}
		}
	}
	fmt.Println(sum)
}

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
	}
	return int(item-'A') + 27
}
