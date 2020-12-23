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
	var ns []int64

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		ns = append(ns, n)
	}

	var iterations int

search:
	for i, n := range ns {
		for j, n2 := range ns {
			for k, n3 := range ns {
				iterations++
				if i != j && i != k && n+n2+n3 == 2020 {
					fmt.Printf("Found %d * %d * %d = %d in %d iterations\n", n, n2, n3, n*n2*n3, iterations)
					break search
				}
			}
		}
	}
}
