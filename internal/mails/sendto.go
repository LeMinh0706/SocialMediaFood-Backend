package mails

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/logger"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"go.uber.org/zap"
)

func BuildMessage(mail Mail) string {
	msg := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ","))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

func SendMailResetPassword(to []string, link string, config util.Config) error {
	content := Mail{
		From:    EmailAddress{Address: config.EmailAdmin, Name: "Foodioo Support"},
		To:      to,
		Subject: "Foodioo reset password",
		Body:    fmt.Sprintf("You can reset your password in Foodio system by this link: %s", link),
	}
	message := BuildMessage(content)

	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost)
	err := smtp.SendMail(config.SMTPHost+config.SMTPPort, auth, config.EmailAdmin, to, []byte(message))
	if err != nil {
		logger.Logger.Error("Send email fail: ", zap.Error(err))
	}
	return nil
}
