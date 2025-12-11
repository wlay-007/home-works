package email

import (
	"fmt"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

func SendEmail(hash string, username string, password string, toUser string) {
	verificationURL := fmt.Sprintf("http://localhost:8081/verify/%s", hash)
	fmt.Println(verificationURL)
	e := &email.Email{
		To:      []string{toUser},
		From:    fmt.Sprintf("Verification <%s>", username),
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML: []byte(fmt.Sprintf(`
<h1>Подтверждение почты</h1>
<p>Привет! Перейди по ссылке, чтобы подтвердить свою почту:</p>
<p><a href="%[1]s">%[1]s</a></p>
`, verificationURL)),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", username, password, "smtp.gmail.com"))
	if err != nil {
		fmt.Println(err)
	}
}
