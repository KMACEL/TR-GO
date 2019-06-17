package main

import (
	"fmt"

	"github.com/KMACEL/TR-Go/07_GO_Dizayn_Pattern/01_Creational_Yaratilis/01_Singleton/singleton"
)

func main() {
	fmt.Println(singleton.GetInstance().Arttir())
	fmt.Println(singleton.GetInstance().Arttir())
	fmt.Println(singleton.GetInstance().Arttir())

	fmt.Println("--------------------------------------")

	fmt.Println(singleton.SecureGetInstance().Arttir())
	fmt.Println(singleton.SecureGetInstance().Arttir())
	fmt.Println(singleton.SecureGetInstance().Arttir())
}
