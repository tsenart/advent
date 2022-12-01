package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	top3 := []int{0, 0, 0}
	inventory := 0

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			item, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("bad item %q: %v", line, err)
			}

			inventory += item
			continue
		}

		for i := range top3 {
			if inventory <= top3[i] {
				continue
			}

			if i < len(top3)-1 { // shift all lower inventories down
				copy(top3[i+1:], top3[i:])
			}

			top3[i] = inventory
			break
		}

		inventory = 0
	}

	fmt.Printf("solution: %d, top3: %v\n", top3[0]+top3[1]+top3[2], top3)
}
