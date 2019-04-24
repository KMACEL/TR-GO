package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//------------------------------------------------------
	// Kullanıcıdan Veri Alma
	//------------------------------------------------------
	var number int

	fmt.Print("Lütfen sayı giriniz : ")
	fmt.Scan(&number)

	fmt.Println("Girilen Sayı : ", number)

	//------------------------------------------------------
	// Boşluk bırakrısanız yukarıdaki kullanım işinize yaramayacaktır.
	// Eğer boşluklarla bir string tipinde veri alacaksanız aşağıdaki şekilde kullanınız
	//------------------------------------------------------
	reader := bufio.NewReader(os.Stdin) //create new reader, assuming bufio imported
	var storageString string

	fmt.Print("Lütfen İstediğiniz kadar giriniz : ")
	storageString, _ = reader.ReadString('\n')

	fmt.Println(storageString)

}
