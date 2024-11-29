package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func passWordGenerator(length int, special, letters, nus bool) (password string) {
	pw := make([]byte, length)
	var types []int
	if special {
		types = append(types, 0)
	}
	if letters {
		types = append(types, 1)
	}
	if nus {
		types = append(types, 2)
	}
	for i := range pw {
		curr := types[rand.Intn(len(types))]
		if curr == 0 {
			pw[i] = specialBytes[rand.Intn(len(specialBytes))]
		} else if curr == 1 {
			pw[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else if curr == 2 {
			pw[i] = numBytes[rand.Intn(len(numBytes))]
		}
	}
	return string(pw)
}

func main() {
	fmt.Println(passWordGenerator(10, true, true, false))
}
