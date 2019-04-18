package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Kodları çalıştırıp aşağıdaki linki tarayıcınıza kopyalayıp çalıştırdığınızda
// Çalıştırdığınız terminalde, ilgili paremetreleri göreceksiniz
// http://localhost:9090/test/test2/?username=mert&password=1234

func cookies(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: "Mert", Expires: expiration}
	http.SetCookie(w, &cookie)

	http.SetCookie(w, &(http.Cookie{Name: "wifename", Value: "Kubra", Expires: expiration}))
	http.SetCookie(w, &(http.Cookie{Name: "childname", Value: "Erdem", Expires: expiration}))

	getCookie, _ := r.Cookie("username")
	fmt.Fprint(w, getCookie)

	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie)
	}
}

func main() {
	http.HandleFunc("/", cookies)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
