package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Kodları çalıştırıp aşağıdaki linki tarayıcınıza kopyalayıp çalıştırdığınızda
// Çalıştırdığınız terminalde, ilgili paremetreleri göreceksiniz
// http://localhost:9090/test/test2/?username=mert&password=1234

func merhabaDunya(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", merhabaDunya)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
