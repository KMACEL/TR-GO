package main

import "fmt"

func main() {

	//------------------------------------------------------
	// Tanımlama Şekli 1
	//------------------------------------------------------
	var x int
	x = 1
	fmt.Println("Sayısal türde olan 'x' değişkeninin değeri : ", x)

	//------------------------------------------------------
	// Tanımlama Şekli 2
	//------------------------------------------------------
	var y = 2
	fmt.Println("Sayısal türde olan 'y' değişkeninin değeri : ", y)

	//------------------------------------------------------
	// Tanımlama Şekli 3
	//------------------------------------------------------
	z := 3
	fmt.Println("Sayısal türde olan 'z' değişkeninin değeri : ", z)

	//------------------------------------------------------
	// Tanımlama Şekli 4
	//------------------------------------------------------
	var (
		t = 4
		q = "d"
	)
	fmt.Println("Sayısal türde olan 't' değişkeninin değeri : ", t)
	fmt.Println("Sözel türde olan 'q' değişkeninin değeri : ", q)

}
