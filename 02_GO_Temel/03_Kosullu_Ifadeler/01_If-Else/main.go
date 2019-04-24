package main

import "fmt"

func main() {

	//------------------------------------------------------
	// IF Kullanımı
	//------------------------------------------------------
	a := 5

	if a < 10 {
		fmt.Println("A 10 dan küçük : ", a)
	}

	if a > 10 {
		fmt.Println("A 10 dan büyük : ", a)
	}

	if a == 10 {
		fmt.Println("A 10 a eşit : ", a)
	}

	if a <= 10 {
		fmt.Println("A 10 a küçük eşit : ", a)
	}

	if a >= 10 {
		fmt.Println("A 10 a büyük eşik : ", a)
	}

	if a != 10 {
		fmt.Println("A 10 a eşit değil : ", a)
	}

	//------------------------------------------------------
	// IF - ELSE Kullanımı
	//------------------------------------------------------

	b := 15

	if b < 10 {
		fmt.Println("B 10 dan küçük : ", b)

	} else {
		fmt.Println("B 10 dan küçük değil : ", b)

	}

	//------------------------------------------------------
	// IF - ELSE IF - ELSE Kullanımı
	//------------------------------------------------------
	c := 20
	if c < 10 {
		fmt.Println("C 10 dan küçük : ", c)

	} else if c > 10 {
		fmt.Println("C 10 dan büyük : ", c)

	} else if c == 10 {
		fmt.Println("C 10'a eşit : ", c)

	} else {
		fmt.Println("Yukarıdaki koşulların hiç biri : ", c)

	}

}
