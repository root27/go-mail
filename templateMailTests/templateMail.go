package main

import (
	"fmt"

	mail "github.com/root27/go-mail"
)

func main() {

	host := "gmail"

	from := "oguzhaneee@gmail.com"

	to := "oguzhaneee@gmail.com"

	subject := "Html test"

	err := mail.SendWithHtml(host, from, "YOUR PASSKEY", to, subject, "./test.html")

	if err != nil {
		panic(err)
	}

	fmt.Println("Mail sent successfully")

}
