package main

import "net/smtp"

const (
	fromAddress       = "r.payami78@g.email"
	fromEmailPassword = "fNQJBeJ23XVju3y8SN"
	smtpServer        = "smtp.ethereal.email"
	smptPort          = "587"
)

// SendEmail will send email to given address
func SendEmail(message string, toAddress []string) (response bool, err error) {

	var auth = smtp.PlainAuth("", fromAddress, fromEmailPassword, smtpServer)
	err = smtp.SendMail(smtpServer+":"+smptPort, auth, fromAddress, toAddress, []byte(message))

	if err == nil {
		return true, nil
	}
	return false, err

}