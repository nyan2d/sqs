package sqs

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"unicode"
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

func tname(s string) string {
	if s == "" {
		return "sqs_sessions"
	}
	if !isGreatStr(s) {
		return fmt.Sprintf("sqs_rand%d", len(s))
	}
	if strings.HasPrefix(s, "sqlite_") {
		return fmt.Sprintf("sqs_rand%d", len(s))
	}
	if s[0] == '_' || unicode.IsDigit(rune(s[0])) {
		return fmt.Sprintf("sqs_rand%d", len(s))
	}
	return s
}

func isGreatStr(s string) bool {
	for _, v := range s {
		if !(unicode.IsLetter(v) || unicode.IsDigit(v) || v == '_') {
			return false
		}
	}
	return true
}
