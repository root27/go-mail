package main

import (
	"log"

	"github.com/root27/go-mail"
)

func main() {

	host := "gmail"

	from := "oguzhaneee@gmail.com"

	to := "oguzhaneee@gmail.com"

	subject := "Variable Test"

	data := struct {
		Name string
		Age  int
	}{
		Name: "Oğuzhan",
		Age:  21,
	}

	err := mail.SendHtmlWithVariables(host, from, "YOUR PASSKEY", to, subject, "./testVariables.html", data)

	if err != nil {
		panic(err)
	}

	log.Printf("Mail sent successfully")

}
