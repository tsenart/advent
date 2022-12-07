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

	// The puzzle input is already structured in DFS order so
	// we can just compute the answer as we parse it, rather than
	// first construct a tree.

	var stack []int64 // Stack of directory sizes
	var sum int64     // Total sum of dir sizes of at most 100000

	for sc.Scan() {
		switch line := sc.Text(); line[:4] {
		case "$ cd":
			if path := line[5:]; path != ".." {
				stack = append(stack, 0) // cd into new dir
				continue
			}

			// cd out of dir, compute dirSize
			dirSize := stack[len(stack)-1]
			if dirSize <= 100000 {
				sum += dirSize
			}

			if stack = stack[:len(stack)-1]; len(stack) > 0 {
				stack[len(stack)-1] += dirSize
			}
		case "$ ls", "dir ": // Nothing to do, next line
		default: // file
			fs := strings.Fields(line)
			fileSize, _ := strconv.ParseInt(fs[0], 10, 64)
			stack[len(stack)-1] += fileSize
		}
	}

	fmt.Println(sum)
}
