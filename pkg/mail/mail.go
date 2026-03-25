package mail

import (
	"fmt"
	"net/smtp"
)

type MailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

type MailService struct {
	config MailConfig
}

func NewMailService(config MailConfig) *MailService {
	return &MailService{config: config}
}

func (s *MailService) SendEmail(to string, subject string, body string) error {
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	auth := smtp.PlainAuth("", s.config.User, s.config.Password, s.config.Host)

	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-version: 1.0;\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\";\r\n"+
		"\r\n"+
		"%s\r\n", s.config.From, to, subject, body)

	return smtp.SendMail(addr, auth, s.config.From, []string{to}, []byte(msg))
}
