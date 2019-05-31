package pages

import (
	"html/template"
	"log"
	"net/http"

	"github.com/KMACEL/TR-Go/08_GO_Web/96_Kullanici_Giris_Ornek/databasecenter"
	"github.com/KMACEL/TR-Go/08_GO_Web/96_Kullanici_Giris_Ornek/errc"
)

// Login is
func Login(w http.ResponseWriter, r *http.Request) {
	conditionsMap := map[string]interface{}{}

	session, err := Store.Get(r, "authenticated-user-session")
	errc.ErrorCenter("Login -> Store.Get ", err)

	if session != nil {
		conditionsMap["Username"] = session.Values["username"]
	}

	if r.FormValue("Password") != "" && r.FormValue("Username") != "" {
		username := r.FormValue("Username")
		password := r.FormValue("Password")

		if !databasecenter.CheckAuthentication(username, password) {
			log.Println("Either username or password is wrong")
			conditionsMap["LoginError"] = true
		} else {
			log.Println("Logged in :", username)
			conditionsMap["Username"] = username
			conditionsMap["LoginError"] = false

			// create a new session and redirect to dashboard
			session, errStoreNew := Store.New(r, "authenticated-user-session")
			errc.ErrorCenter("Login -> Store.New ", errStoreNew)

			session.Values["username"] = username
			err := session.Save(r, w)
			if err != nil {
				errc.ErrorCenter("Login ->  session.Save ", err)
				log.Println(err)
			}

			logPages("Login", r, session.Name(), session.Values["username"], nil)

			http.Redirect(w, r, "/getInfo", http.StatusFound)
		}

	}

	if err := template.Must(template.ParseFiles("gtpl/login.gtpl")).Execute(w, conditionsMap); err != nil {
		log.Println(err)
	}
}

// Logout is
func Logout(w http.ResponseWriter, r *http.Request) {
	//read from session
	session, _ := Store.Get(r, "authenticated-user-session")

	// remove the username
	session.Values["username"] = ""
	session.Values["SendMessage"] = -1

	err := session.Save(r, w)

	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
