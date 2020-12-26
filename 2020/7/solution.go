package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	g := make(graph)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSuffix(sc.Text(), ".")
		fields := strings.SplitN(line, " contain ", 2)
		from := strings.TrimSuffix(fields[0], " bags")

		for _, color := range strings.Split(fields[1], ", ") {
			fields := strings.Fields(color)
			count, _ := strconv.Atoi(fields[0])
			to := strings.Join(fields[1:len(fields)-1], " ")

			g[from] = append(g[from], edge{color: to, kind: contains, count: count})
			g[to] = append(g[to], edge{color: from, kind: contained, count: 1})
		}
	}

	fmt.Println(g.contained("shiny gold"), g.contains("shiny gold")-1)
}

type kind int

const (
	contains  kind = 1
	contained kind = 2
)

type edge struct {
	color string
	kind  kind
	count int
}

type graph map[string][]edge

func (g graph) contained(color string) (count int) {
	var stack []string
	for _, e := range g["shiny gold"] {
		if e.kind == contained {
			stack = append(stack, e.color)
		}
	}

	visited := map[string]bool{}
	for len(stack) > 0 {
		next := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[next] = true

		for _, e := range g[next] {
			if e.kind == contained && !visited[e.color] {
				stack = append(stack, e.color)
			}
		}
	}

	return len(visited)
}

func (g graph) contains(color string) (count int) {
	count = 1
	for _, e := range g[color] {
		if e.kind == contains {
			count += e.count * g.contains(e.color)
		}
	}
	return count
}
