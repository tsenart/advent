package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	var (
		cols [2][]int
		cnt  = make(map[int]int)
		sc   = bufio.NewScanner(os.Stdin)
	)

	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) != 2 {
			log.Fatalf("unexpected number of fields: %d", len(fields))
		}

		for i, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				log.Fatal(err)
			}

			cols[i] = append(cols[i], n)
			if i == 1 {
				cnt[n]++
			}
		}
	}

	for i := range cols {
		slices.Sort(cols[i])
	}

	var sum, sim int
	for i := range cols[0] {
		sum += abs(cols[0][i] - cols[1][i])
		sim += cols[0][i] * cnt[cols[0][i]]
	}

	log.Println(sum, sim)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
