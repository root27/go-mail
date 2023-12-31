# Go Email Package

Simple send email package for Go.

## Installation

```bash

go get github.com/root27/go-mail

```

## Functions

- Send Email with text

This function only sends text emails. You can find usage and details of the function below.


```go

func Send(host string, from string, passKey string, to string, subject string, body string) error

```

### Example

```go

package main

import (
    "fmt"
    "github.com/root27/go-mail"
)
func main(){
	host := "HOST NAME" // e.g gmail, yandex

	from := "MAIL ADDRESS FROM WHICH YOU WILL SEND"

	to := "MAIL ADDRESS TO WHICH YOU WILL SEND"

	//PASSKEY IS THE KEY OF THE MAIL ADDRESS FROM WHICH YOU WILL SEND

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

	host := "HOST NAME" // e.g gmail, yandex

	from := "MAIL ADDRESS FROM WHICH YOU WILL SEND"

	to := "MAIL ADDRESS TO WHICH YOU WILL SEND"

	//PASSKEY IS THE KEY OF THE MAIL ADDRESS FROM WHICH YOU WILL SEND

	subject := "Html test"

	htmlPath := "Path of the html file"

	err := mail.SendWithHtml(host, from, "YOUR PASSKEY", to, subject, htmlPath)

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

	host := "HOST NAME" // e.g gmail, yandex

	from := "MAIL ADDRESS FROM WHICH YOU WILL SEND"

	to := "MAIL ADDRESS TO WHICH YOU WILL SEND"

	//PASSKEY IS THE KEY OF THE MAIL ADDRESS FROM WHICH YOU WILL SEND

	subject := "Variable Test"


	htmlPath := "Path of the html file"

	data := struct {
		Name string
		Age  int
	}{
		Name: "Oğuzhan",
		Age:  21,
	}

	err := mail.SendHtmlWithVariables(host, from, "YOUR PASSKEY", to, subject, htmlPath, data)

	if err != nil {
		panic(err)
	}

	log.Printf("Mail sent successfully")

}

```

LICENSE

[MIT](./LICENSE)








