package sqs

import (
	"crypto/rand"
	"math/big"
)

func generateToken(size int) string {
	runes := []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	r := make([]rune, size)
	for i := 0; i < size; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(runes))))
		r[i] = runes[num.Int64()]
	}
	return string(r)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func intToBool(i int) bool {
	return i != 0
}
