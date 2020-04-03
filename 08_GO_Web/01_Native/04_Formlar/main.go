package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func anaSayfa(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Form Verisi : ", r.Form)
	fmt.Println("URL Yolu : ", r.URL.Path)
	fmt.Println("Şema : ", r.URL.Scheme)
	fmt.Println("Belirli Parametre : ", r.Form["username"])

	fmt.Println("-----------------------------------")
	fmt.Println("Bütün Parametreler")
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("-----------------------------------")

	fmt.Fprintf(w, "Merhaba Dünya!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sorgu Method : ", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		fmt.Println("İlk açıldığında")
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println("Butona Tıklanınca")

		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", anaSayfa)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
