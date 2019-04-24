package main

import "fmt"

type degiskenType int

const (
	degisken1 degiskenType = iota
	degisken2
	degisken3
	degisken4
)

func main() {
	fmt.Println("degisken1", degisken1)
	fmt.Println("degisken2", degisken2)
	fmt.Println("degisken3", degisken3)
	fmt.Println("degisken4", degisken4)

}
