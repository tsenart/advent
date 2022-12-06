package main

import (
	"io"
	"math/bits"
	"os"
)

func main() {
	datastream, _ := io.ReadAll(os.Stdin)
	for i := 0; i < len(datastream)-13; i++ {
		var marker uint32
		for j := 0; j < 14; j++ {
			marker |= 1 << (datastream[i+j] - 'a')
		}
		if bits.OnesCount32(marker) == 14 {
			println(i + 14)
			break
		}
	}
}
