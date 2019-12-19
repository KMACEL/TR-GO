package main

import (
	"log"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth("", "mertacel@gmail.com", "xyz", "smtp.gmail.com")

	to := []string{"mertacel@gmail.com"}
	msg := []byte("To: mertacel@gmail.com\r\n" +
		"Subject: Bu bir mail başlığıdır\r\n" +
		"\r\n" +
		"huhu ilk mailimi attım\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "mertacel@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
