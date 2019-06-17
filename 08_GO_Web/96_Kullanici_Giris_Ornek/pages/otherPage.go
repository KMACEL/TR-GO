package pages

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/KMACEL/TR-Go/08_GO_Web/96_Kullanici_Giris_Ornek/errc"
)

// OtherPage is
func OtherPage(w http.ResponseWriter, r *http.Request) {
	log.SetOutput(os.Stdout)
	conditionsMap := map[string]interface{}{}

	session, err := Store.Get(r, "authenticated-user-session")
	if err != nil {
		errc.ErrorCenter("Diger Sayfa -> Store.Get ", err)
		log.Println("Unable to retrieve session data!", err)
	}

	conditionsMap["Username"] = session.Values["username"]

	r.ParseForm()

	if r.Form["deviceId"] != nil {
		conditionsMap["DeviceID"] = r.Form["deviceId"][0]
	} else {
		conditionsMap["DeviceID"] = getDeviceID
	}

	logPages("otherPage", r, session.Name(), session.Values["username"], getDeviceID)

	if r.Form["deviceId"] != nil {
		getDeviceID = r.Form["deviceId"][0]
	}
	if len(getDeviceID) == 0 {
		//fmt.Fprintf(w, "Device ID not null!")
		//return
	}

	if r.Method == "POST" {
		logPages("otherPage", r, session.Name(), session.Values["username"], getDeviceID)
		session.Values["SendMessage"] = 1
		//session.Values["SendMessage"] = 0
		session.Save(r, w)
		http.Redirect(w, r, "/getInfo", 301)

	}

	if err := template.Must(template.ParseFiles("templates/otherPage.html")).Execute(w, conditionsMap); err != nil {
		errc.ErrorCenter("Other Page -> template.Must ", err)
		log.Println(err)
	}

}
