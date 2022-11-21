package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const (
	CONFIG_SMTP_HOST     = "smtp.gmail.com"                          // host smtp untuk gmail
	CONFIG_SMTP_PORT     = 587                                       // prot smtp untuk gmail
	CONFIG_SENDER_NAME   = "Muhammad Yusuf <inbox.myusuf@gmail.com>" // nama pengirim
	CONFIG_AUTH_EMAIL    = "inbox.myusuf@gmail.com"                  // email yg kita buatkan app password
	CONFIG_AUTH_PASSWORD = "irkwfpiiyonebuww"                        // dari app password yang telah kita buat
)

func main() {
	to := []string{"muh.iyus@gmail.com", "yusuf@tabeldata.com"} // email tujuan
	cc := []string{"muhammad.yusuf@protechbit.com"}             // email cc
	subject := "Test Golang Email"
	message := "Hello, From Golang App Send Email"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
