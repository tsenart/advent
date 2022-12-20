package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	began := time.Now()
	sc := bufio.NewScanner(os.Stdin)

	var packets []packet
	for sc.Scan() {
		if line := sc.Bytes(); len(line) != 0 {
			_, pkt := parse(line)
			packets = append(packets, pkt)
		}
	}

	dividers := []packet{
		list{list{integer(2)}},
		list{list{integer(6)}},
	}

	packets = append(packets, dividers...)

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].Compare(packets[j]) < 0
	})

	product := 1
	for _, p := range dividers {
		i, _ := sort.Find(len(packets), func(i int) int {
			return p.Compare(packets[i])
		})
		product *= i + 1
	}

	fmt.Println(product, time.Since(began))
}

type packet interface {
	Compare(packet) int
}

func parse(p []byte) ([]byte, packet) {
	if len(p) == 0 {
		return p, nil
	}

	if p[0] != '[' {
		i := bytes.IndexAny(p, ",]")
		if i <= 0 {
			return p, nil
		}

		n, _ := strconv.Atoi(string(p[:i]))
		return p[i:], integer(n)
	}

	var l list
	for p[0] != ']' {
		var pkt packet
		if p, pkt = parse(p[1:]); pkt != nil {
			l = append(l, pkt)
		}
	}

	return p[1:], l
}

type list []packet

func (l list) Compare(other packet) int {
	o, ok := other.(list)
	if !ok {
		o = list{other}
	}

	for i, p := range l {
		if i == len(o) {
			break
		} else if c := p.Compare(o[i]); c != 0 {
			return c
		}
	}

	return len(l) - len(o)
}

type integer int

func (i integer) Compare(other packet) int {
	if o, ok := other.(integer); ok {
		return int(i - o)
	}
	return list{i}.Compare(other)
}
