package main

import (
	"fmt"
	"log"
	"net/http"
)

// Yönelndirme, URL'de verilen yola göre sayfa değişimleri yapar.

// http://localhost:9090
func anaSayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba Burası Projemizin Anasayfasıdır...")
}

// http://localhost:9090/merhaba
func merhabaDunya(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba Dünya!")
}

// http://localhost:9090/acel
func acel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba, ACEL Sayfasına hoş geldiniz!")
}

func main() {
	http.HandleFunc("/", anaSayfa)
	http.HandleFunc("/merhaba", merhabaDunya)
	http.HandleFunc("/acel", acel)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
