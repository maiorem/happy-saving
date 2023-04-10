package middlewares

import "gopkg.in/gomail.v2"

func SendEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "example@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "이메일 인증")
	m.SetBody("text/plain", "인증 코드: "+token)

	d := gomail.NewDialer("smtp.gmail.com", 587, "example@example.com", "password")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
