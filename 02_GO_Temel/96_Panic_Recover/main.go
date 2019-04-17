package main

import "fmt"

type hataType int

const (
	esitHata hataType = iota
	buyukHata
)

func main() {
	a := 2
	fmt.Println(1)

	defer func() {
		fmt.Println(2)
		if x := recover(); x != nil {
			fmt.Println(3)
			if x == esitHata {
				fmt.Println(4)
				a = a + 1
				fmt.Println("A değiştirildi...", x)
			}

			if x == buyukHata {
				fmt.Println(5)
				a = a - 1
				fmt.Println("A değiştirildi...", x)
			}
		}
	}()

	fmt.Println(6)

	if a > 2 {
		fmt.Println(7)
		panic(esitHata)
	}

	if a > 1 {
		fmt.Println(8)
		panic(buyukHata)
	}

	fmt.Println(9)
}
