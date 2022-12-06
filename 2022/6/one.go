package main

import (
	"io"
	"math/bits"
	"os"
)

func main() {
	datastream, _ := io.ReadAll(os.Stdin)
	for i := 0; i < len(datastream)-3; i++ {
		marker := uint32(
			1<<(datastream[i]-'a') |
				1<<(datastream[i+1]-'a') |
				1<<(datastream[i+2]-'a') |
				1<<(datastream[i+3]-'a'),
		)
		if bits.OnesCount32(marker) == 4 {
			println(i + 4)
			break
		}
	}
}
