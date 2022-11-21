package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

const (
	CONFIG_SMTP_HOST     = "smtp.gmail.com"                          // host smtp untuk gmail
	CONFIG_SMTP_PORT     = 587                                       // prot smtp untuk gmail
	CONFIG_SENDER_NAME   = "Muhammad Yusuf <inbox.myusuf@gmail.com>" // nama pengirim
	CONFIG_AUTH_EMAIL    = "inbox.myusuf@gmail.com"                  // email yg kita buatkan app password
	CONFIG_AUTH_PASSWORD = "irkwfpiiyonebuww"                        // dari app password yang telah kita buat
)

func main() {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "muh.iyus@gmail.com", "yusuf@tabeldata.com")                     // email tujuan
	mailer.SetAddressHeader("Cc", "muhammad.yusuf@protechbit.com", "Protechbit Mail Yusuf") // email cc
	mailer.SetHeader("Subject", "Test Golang Email")
	mailer.SetBody("text/html", "Hello, <b>From Golang App Send Email</b>")
	mailer.Attach("./banner.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
