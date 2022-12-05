package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var count int
	for sc.Scan() {
		var pair [2][2]int
		for i, s := range strings.Split(sc.Text(), ",") {
			ps := strings.SplitN(s, "-", 2)
			pair[i][0], _ = strconv.Atoi(ps[0])
			pair[i][1], _ = strconv.Atoi(ps[1])
		}

		switch { // sort
		case pair[0][0] > pair[1][0]:
			fallthrough
		case pair[0][0] == pair[1][0] && pair[0][1] < pair[1][1]:
			pair[0], pair[1] = pair[1], pair[0]
		}

		if pair[1][0] <= pair[0][1] {
			count++
		}
	}

	fmt.Println(count)
}
