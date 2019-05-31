package logc

import (
	"log"
	"os"
	"time"
)

var Debug = true

// GlobalPrint is
func GlobalPrint(args ...interface{}) {
	if Debug {
		if _, err := os.Stat("./logc"); os.IsNotExist(err) {
			os.MkdirAll("./logc", os.ModePerm)
		}

		f, err := os.OpenFile("logc/globalLogFile_"+GetTimeNamesFormatDaysTYPE2()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
		log.SetOutput(os.Stdout)
	}
}

func GetTimeNamesFormatDaysTYPE2() string {
	return time.Now().Format("2006_01_02")
}
