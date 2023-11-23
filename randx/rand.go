// Package randx.md
// Date: 2022/9/7 00:57
// Author: Amu
// Description:
package randx

import (
	crand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"time"
)

const (
	Numeral      = "0123456789"
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandInt generate random ini between min and max, maybe in, not be max
func RandInt(min, max int) int {
	if min == max {
		return min
	}

	if max < min {
		max, min = min, max
	}

	return rand.Intn(max-min) + min
}

// RandBytes generate random byte slice
func RandBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	b := make([]byte, length)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return nil
	}
	return b
}

// RandString 生成随机字符串
func RandString(length int) string {
	return random(Letters, length)
}

// RandUpper generate a random upper case string
func RandUpper(length int) string {
	return random(UpperLetters, length)
}

// RandLower generate a random lower case string
func RandLower(length int) string {
	return random(LowerLetters, length)
}

// RandNumeral generate a random numeral string of specified length
func RandNumeral(length int) string {
	return random(Numeral, length)
}

// RandNumeralOrLetter generate a random numeral or letter string
func RandNumeralOrLetter(length int) string {
	return random(Numeral+Letters, length)
}

// random generate a random string based on given string range.
func random(s string, length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = s[rand.Int63()%int64(len(s))]
	}
	return string(b)
}

// UUID4 generate a random UUID of version 4 according to RFC 4122.
func UUID4() (string, error) {
	uuid := make([]byte, 16)

	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
