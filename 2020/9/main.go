package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var cipher []uint64
	for sc.Scan() {
		n, _ := strconv.ParseUint(sc.Text(), 10, 64)
		cipher = append(cipher, n)
	}

	var invalid uint64
	for i := 25; i < len(cipher); i++ {
		a, b := factors(cipher[i], cipher[i-25:i])
		if a == 0 || b == 0 {
			invalid = cipher[i]
			break
		}
	}

	var weakness uint64
crack:
	for i := 0; i < len(cipher); i++ {
		sum := cipher[i]
		min, max := sum, sum
		for j := i + 1; j < len(cipher) && sum <= invalid; j++ {
			if cipher[j] < min {
				min = cipher[j]
			}

			if cipher[j] > max {
				max = cipher[j]
			}

			if sum += cipher[j]; sum == invalid {
				weakness = min + max
				break crack
			}
		}
	}

	fmt.Println(invalid, weakness)
}

func factors(n uint64, preamble []uint64) (a, b uint64) {
	for _, a := range preamble {
		for _, b := range preamble {
			if a+b == n {
				return a, b
			}
		}
	}
	return 0, 0
}
