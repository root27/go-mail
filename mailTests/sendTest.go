package main

import (
	"fmt"

	mail "github.com/root27/go-mail"
)

func main() {

	host := "gmail"

	from := "oguzhaneee@gmail.com"

	to := "oguzhaneee@gmail.com"

	subject := "test"

	message := "Hello World"

	err := mail.Send(host, from, "YOUR PASSKEY", to, subject, message)

	if err != nil {
		panic(err)
	}

	fmt.Println("Mail sent successfully")

}
