package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	began := time.Now()
	sc := bufio.NewScanner(os.Stdin)

	type monkey struct {
		items       []int
		op          string
		arg         int
		factor      int
		throw       map[bool]int
		inspections int
	}

	var monkeys []*monkey
	var m *monkey

	for sc.Scan() || m != nil {
		line := strings.TrimSpace(sc.Text())
		switch {
		case line == "":
			monkeys = append(monkeys, m)
			m = nil
		case strings.HasPrefix(line, "Monkey"):
			m = &monkey{throw: map[bool]int{}}
		case strings.HasPrefix(line, "Starting items"):
			items := strings.Split(strings.TrimPrefix(line, "Starting items: "), ", ")
			for _, item := range items {
				n, _ := strconv.Atoi(item)
				m.items = append(m.items, n)
			}
		case strings.HasPrefix(line, "Operation"):
			op := strings.Fields(strings.TrimPrefix(line, "Operation: new = old "))
			if m.op = op[0]; op[1] != "old" {
				m.arg, _ = strconv.Atoi(op[1])
			}

		case strings.HasPrefix(line, "Test"):
			m.factor, _ = strconv.Atoi(strings.TrimPrefix(line, "Test: divisible by "))

		case strings.HasPrefix(line, "If"):
			fs := strings.Fields(line)
			b, _ := strconv.ParseBool(strings.TrimSuffix(fs[1], ":"))
			to, _ := strconv.Atoi(fs[len(fs)-1])
			m.throw[b] = to
		}
	}

	factor := 1
	for _, m := range monkeys {
		factor *= m.factor
	}

	for round := 1; round <= 10000; round++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				m.inspections++

				arg := m.arg
				if arg == 0 {
					arg = item
				}

				switch m.op {
				case "*":
					item *= arg
				case "+":
					item += arg
				}

				item %= factor
				to := m.throw[item%m.factor == 0]
				monkeys[to].items = append(monkeys[to].items, item)
			}

			m.items = m.items[:0]
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	fmt.Println(monkeys[0].inspections*monkeys[1].inspections, time.Since(began))
}
