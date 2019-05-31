package main

import (
	"log"
	"net/http"

	"github.com/KMACEL/TR-Go/08_GO_Web/96_Kullanici_Giris_Ornek/databasecenter"
	"github.com/KMACEL/TR-Go/08_GO_Web/96_Kullanici_Giris_Ornek/pages"
	"github.com/gorilla/sessions"
)

func init() {

	pages.Store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 3, // 3 hours
		HttpOnly: true,
	}

}

func main() {
	if databasecenter.StartDatabase() {
		//databasecenter.AddUser("mertacel", "12344321", "admin")

		http.HandleFunc("/", pages.Login)
		http.HandleFunc("/login", pages.Login)
		http.HandleFunc("/logout", pages.Logout)
		http.HandleFunc("/getInfo", pages.GetInfo)
		http.HandleFunc("/otherPage", pages.OtherPage)

		err := http.ListenAndServe(":9000", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}
