package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`(\d+)-(\d+)\s+(\w):\s+(\w+)`)
	sc := bufio.NewScanner(os.Stdin)

	var solutions [2]int

	for sc.Scan() {
		fields := re.FindStringSubmatch(sc.Text())

		i, _ := strconv.Atoi(fields[1])
		j, _ := strconv.Atoi(fields[2])
		char := fields[3]
		password := fields[4]

		n := strings.Count(password, char)
		if min, max := i, j; n >= min && n <= max {
			solutions[0]++
		}

		first := string(password[i-1]) == char
		second := string(password[j-1]) == char
		if first != second {
			solutions[1]++
		}
	}

	fmt.Println(solutions)
}
