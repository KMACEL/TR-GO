package main

import "fmt"

func main() {
	const sabitDegisken int = 10

	fmt.Println("Sabit Değişken : ", sabitDegisken)

	// sabitDegisken = 22 tekrar atama yapılmaz, hata

	const (
		sabit1 = 12
		sabit2 = "merhaba"
		sabit3 = true
		sabit4 = 12.5
	)

	fmt.Println("Sabit Değişken : ", sabit1)
	fmt.Println("Sabit Değişken : ", sabit2)
	fmt.Println("Sabit Değişken : ", sabit3)
	fmt.Println("Sabit Değişken : ", sabit4)

}
