package errc

import (
	"log"
	"os"

	"github.com/KMACEL/TR-Go/08_GO_Web/01_Native/96_Kullanici_Giris_Ornek/logc"
)

var Debug = true

// ErrorCenter is
func ErrorCenter(title string, err error) {
	if Debug {
		if err != nil {
			log.Println("Error - "+title, " : ", err.Error())
			errorFile("Error - "+title, " : ", err.Error())
		}
	}
}

func errorFile(args ...interface{}) {
	if Debug {
		if _, err := os.Stat("./errc"); os.IsNotExist(err) {
			os.MkdirAll("./errc", os.ModePerm)
		}

		f, err := os.OpenFile("errc/errorLogFile_"+logc.GetTimeNamesFormatDaysTYPE2()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error - Error File : Error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
		log.SetOutput(os.Stdout)
	}
}
