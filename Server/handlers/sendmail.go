package handlers

import (
	"Network-be/config"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const tmpl = `<center>
<h1>You Successfully register into the ConfigSever</h1>
<p>Your verify ticket is: %s</p>
<p>Please click the link below to verify your account</p>
<a href="//%s/verify?ticket=%s"> Verify </a>
</center>`

func SendVerifyByEmail(to, body string) error {
	e := email.NewEmail()
	e.From = config.GlobalConfig.GetString("email.user")
	e.To = []string{to}
	// e.Bcc = []string{}
	e.Subject = "[DN42ConfigServer] Verify Ticket"
	// e.Text = []byte("Here is the text version of the message.\nThis is line 2 of the text.")
	e.HTML = []byte(body)
	return e.Send(
		// Smtp Server
		fmt.Sprintf("%s:%s",
			config.GlobalConfig.GetString("email.host"),
			config.GlobalConfig.GetString("email.port")),
		// Auth
		smtp.PlainAuth("",
			config.GlobalConfig.GetString("email.user"),
			config.GlobalConfig.GetString("email.auth"),
			config.GlobalConfig.GetString("email.host")))
}

func CreateContent(verify string) string {
	return fmt.Sprintf(tmpl, verify,
		config.GlobalConfig.GetString("server.host"), verify)
}
