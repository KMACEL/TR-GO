package pages

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// GetInfo is
func GetInfo(w http.ResponseWriter, r *http.Request) {
	log.SetOutput(os.Stdout)

	var username string
	conditionsMap := map[string]interface{}{}

	r.ParseForm()

	session, err := Store.Get(r, "authenticated-user-session")

	if err != nil {
		log.Println("Unable to retrieve session data!", err)
	}

	log.Println("Session name : ", session.Name())
	log.Println("Username : ", session.Values["username"])
	log.Println("Status : ", session.Values["SendMessage"])

	if r.Method == "GET" {
		conditionsMap["SendMessage"] = session.Values["SendMessage"]
		session.Values["SendMessage"] = -1
		session.Save(r, w)
	} else {
		conditionsMap["SendMessage"] = -1
	}

	logPages("GetInfo", r, session.Name(), session.Values["username"], nil)

	conditionsMap["Username"] = session.Values["username"]
	username = fmt.Sprint(session.Values["username"])

	if r.Form["deviceId"] != nil {

		deviceID := r.Form["deviceId"][0]

		if len(deviceID) == 0 {
			conditionsMap["DeviceIDNullError"] = true
		} else {
			conditionsMap["DeviceIDNullError"] = false
		}

		getDeviceID = deviceID

		logPages("GetInfo", r, "", username, deviceID)

		if len(getDeviceID) <= 0 {
			conditionsMap["DeviceIDError"] = true
		} else {
			conditionsMap["DeviceIDError"] = false
		}

		if !conditionsMap["DeviceIDError"].(bool) && !conditionsMap["DeviceIDNullError"].(bool) {
			http.Redirect(w, r, "/otherPage", 301)
		}
	} else {
		conditionsMap["DeviceIDError"] = true
	}

	if r.Method == "GET" {
		conditionsMap["DeviceIDError"] = false
		conditionsMap["DeviceIDNullError"] = false
	}

	if err := template.Must(template.ParseFiles("templates/getInfo.html")).Execute(w, conditionsMap); err != nil {
		log.Println(err)
	}
}
