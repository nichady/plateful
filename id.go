package main

import (
	"math/rand"
	"time"
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	runes  = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
)

func RandomID() string {
	b := make([]rune, 5)
	for i := range b {
		b[i] = runes[random.Intn(len(runes))]
	}
	return string(b)
}
