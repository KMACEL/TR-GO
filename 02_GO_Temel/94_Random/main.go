package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30
	fmt.Println(rand.Intn(max-min) + min)
}
