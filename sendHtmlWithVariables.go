package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func SendHtmlWithVariables(host string, from string, passKey string, to string, subject string, temp string, variables interface{}) error {

	smtpHost := `smtp.` + host + `.com`

	smtpPort := "587"

	auth := smtp.PlainAuth("", from, passKey, smtpHost)

	t, _ := template.ParseFiles(temp)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+subject+"\n%s\n\n", mimeHeaders)))

	t.Execute(&body, variables)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, body.Bytes())

	if err != nil {
		return err
	}

	return nil
}
