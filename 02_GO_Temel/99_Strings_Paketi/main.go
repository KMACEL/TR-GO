package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	value := "kuş, kedi ve balık iyi arkadaştırlar."

	r1 := strings.NewReplacer("kuş", "martış",
		"kedi", "tekir",
		"balık", "karabaş")

	// Replace all pairs.
	result := r1.Replace(value)
	fmt.Println(result)

	//#################################################################################
	//Örnek2
	//#################################################################################

	orjinal := "kedi ve köpek"
	//  Replacer (karşılaştırma).
	r := strings.NewReplacer("kedi", "araba",
		"ve", "veya",
		"köpek", "traktör")

	t0 := time.Now()

	// Version 1:  Replacer.
	for i := 0; i < 1000000; i++ {
		result := r.Replace(orjinal)
		if result != "araba veya traktör" {
			fmt.Println(0)
		}
	}

	t1 := time.Now()

	// Version 2: Replace .
	for i := 0; i < 1000000; i++ {
		temp := strings.Replace(orjinal, "kedi", "araba", -1)
		temp = strings.Replace(temp, "ve", "veya", -1)
		result := strings.Replace(temp, "köpek", "traktör", -1)
		if result != "araba veya traktör" {
			fmt.Println(0)
		}
	}

	t2 := time.Now()

	// Sonuç : Replacer daha hızlıdır.
	fmt.Println(t1.Sub(t0))
	fmt.Println(t2.Sub(t1))
}
