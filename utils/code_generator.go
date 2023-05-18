package utils

import (
	"fmt"
	"math/rand"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenCodes() {

	fmt.Println(randSeq(5))

}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
