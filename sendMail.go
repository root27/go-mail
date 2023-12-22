package mail

import "net/smtp"

func Send(host string, from string, passKey string, to string, subject string, body string) error {

	smtpHost := `smtp.` + host + `.com`

	smtpPort := "587"

	message := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n",
	)

	auth := smtp.PlainAuth("", from, passKey, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)

	if err != nil {
		return err
	}

	return nil

}
