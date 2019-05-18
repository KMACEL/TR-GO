package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	max := 0
	min := 50
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(max-min) + min)
}
