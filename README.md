# Go Email Package

Simple send email package for Go.

## Installation

```bash

go get github.com/root27/go-email

```

## Functions

- Send Email with text

This function only sends text emails. You can find usage and details of the function below.


```go

func SendEmail(host string, port int, from string, password string, to []string, subject string, body string) error

```

### Example

```go

package main

import (
    "fmt"
    "github.com/root27/go-email"
)
func main(){
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

```

- Send Email with html template

This function sends html emails. You can find usage and details of the function below.

```go

func SendWithHtml(host string, from string, passKey string, to string, subject string, htmlTemplate string) error

```

### Example

```go

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

```

- Send Email with html template and variables

This function sends html emails with variables. If you have any variables in your html template, you can use this function. You can find usage and details of the function below.

```go

func SendHtmlWithVariables(host string, from string, passKey string, to string, subject string, htmlTemplate string, variables interface{}) error

```

### Example

```go

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

``````








