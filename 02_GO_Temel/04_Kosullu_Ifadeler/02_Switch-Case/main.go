package main

import "fmt"

func main() {

	//------------------------------------------------------
	// SWITCH - CASE - DEFAULT Kullanımı
	//------------------------------------------------------
	a := 10

	switch a {
	case 5:
		fmt.Println("A değeri 5")
	case 9:
		fmt.Println("A değeri 9")
	case 10:
		fmt.Println("A değeri 10")
	default:
		fmt.Println("Yukarıdakilerin hiç biri")
	}

	//------------------------------------------------------
	// SWITCH - CASE - DEFAULT Kullanımı 2
	//------------------------------------------------------
	b := 20

	switch {
	case b < 5:
		fmt.Println("B değeri 5")
	case b < 10:
		fmt.Println("B değeri 10")
	case b <= 20:
		fmt.Println("B değeri 20")
	default:
		fmt.Println("Yukarıdakilerin hiç biri")
	}

	//------------------------------------------------------
	// SWITCH - CASE - DEFAULT Kullanımı 3
	//------------------------------------------------------

	c := 20
	var x interface{} // İleride Anlatılacak
	x = c

	switch x.(type) {
	case int:
		fmt.Println("Değişken Integer Türünde")
	case string:
		fmt.Println("Değişken String Türünde")
	case float32:
		fmt.Println("Değişken Floar32 Türünde")
	default:
		fmt.Println("Yukarıdakilerin hiç biri")
	}

	//------------------------------------------------------
	// SWITCH - CASE - DEFAULT - fallthrough Kullanımı
	//------------------------------------------------------

	d := 20

	switch {
	case d < 5:
		fmt.Println("D değeri 5")
	case d > 10:
		fmt.Println("D değeri 10")
		fallthrough
	case d <= 20:
		fmt.Println("D değeri 20")
	default:
		fmt.Println("Yukarıdakilerin hiç biri")
	}

}
