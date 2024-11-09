package utils

import (
	"math/rand"
)

func GenerateRandomID() int {
	return rand.Intn(10000000) + 1
}
