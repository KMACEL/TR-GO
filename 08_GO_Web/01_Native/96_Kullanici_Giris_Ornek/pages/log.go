package pages

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KMACEL/TR-Go/08_GO_Web/01_Native/96_Kullanici_Giris_Ornek/errc"
	"github.com/KMACEL/TR-Go/08_GO_Web/01_Native/96_Kullanici_Giris_Ornek/logc"
)

func logPages(page string, r *http.Request, sessionName string, sessionUsername interface{}, deviceID interface{}) {
	var req requestJSON
	req.RequestPage = page
	req.Request = fmt.Sprint(r)
	req.SessionName = sessionName
	req.SessionUsername = fmt.Sprint(sessionUsername)
	req.DeviceID = fmt.Sprint(deviceID)

	getReq, err := json.Marshal(req)
	errc.ErrorCenter("logPages "+page+"json.Marsha ", err)

	logc.GlobalPrint(string(getReq))
}
