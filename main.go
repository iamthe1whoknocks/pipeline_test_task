package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

const (
	dataLength = 10 // int slice length in incoming data
)

func main() {
	out, err := generateIntSlice(dataLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)

}

// generate random int slice
func generateIntSlice(length int) ([]int, error) {
	out := make([]int, 0, length)

	for i := 0; i < length; i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return nil, fmt.Errorf("generateIntSlice : %w", err)
		}
		out = append(out, int(val.Int64()))
	}
	return out, nil
}
