package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	began := time.Now()
	sc := bufio.NewScanner(os.Stdin)

	divs := []packet{
		list{list{integer(2)}},
		list{list{integer(6)}},
	}

	idxs := []int{1, 2}

	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}

		_, pkt := parse(line)
		for i, div := range divs {
			if pkt.Compare(div) >= 0 {
				continue
			}
			for j := i; j < len(idxs); j++ {
				idxs[j]++
			}
			break
		}
	}

	fmt.Println(idxs[0]*idxs[1], time.Since(began))
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
