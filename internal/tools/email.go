package tools

import (
	"auth/internal/config"
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(email string, subject, body string, cfg *config.ConfigSMTP) {
	to := []string{email}

	message := []byte(subject + "\r\n\r\n" + body)

	auth := smtp.PlainAuth("", cfg.SmtpFrom, cfg.SmtpPass, cfg.SmtpHost)

	err := smtp.SendMail(cfg.SmtpHost+":"+cfg.SmtpPass, auth, cfg.SmtpFrom, to, message)
	if err != nil {
		log.Println("Ошибка отправки письма")
	}

	sendFake(email, subject, body)
}

func sendFake(email string, subject, body string) {
	fmt.Println("=== Mock Email Sent ===")
	fmt.Println("To:", email)
	fmt.Println("Subject:", subject)
	fmt.Println("Body:", body)
	fmt.Println("=======================")
}
