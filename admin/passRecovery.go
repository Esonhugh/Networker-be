package main

import (
	"Network-be/Server/db"
	"Network-be/config"
	"Network-be/data/PO"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func printHelp() {
	println("[*] Usage: " + os.Args[0] + " {username whom you want recover}")
}

func GenPassCode(width int) string {
	numeric := "0987654321"
	lowerCase := "qwertyuiopasdfghjklzxcvbnm"
	upperCase := "QWERTYUIOPASDFGHJKLZXCVBNM"
	whole := []byte(numeric + lowerCase + upperCase)
	r := len(whole)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%c", whole[rand.Intn(r)])
	}
	return sb.String()
}

const tmpl = `<center>
<h1>You Successfully Changed the password into the ConfigSever</h1>
<p>Your new password is: %s</p>
<p>If you successfully logged in, don't forget change your password</p>
</center>`

func SendPassByEmail(to, body string) error {
	e := email.NewEmail()
	e.From = config.GlobalConfig.GetString("email.user")
	e.To = []string{to}
	// e.Bcc = []string{}
	e.Subject = "[DN42ConfigServer] Password Changes"
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

func main() {
	argc := len(os.Args)
	if argc == 2 {
		println("Warning: You are reset the user:", os.Args[1], "'s password in db.")
		config.Init()
		db.InitDB()
		var user PO.Auth
		db.DBService.MainDB.Where("username = ?", os.Args[1]).Find(&user)
		password := GenPassCode(16)
		println("user: ", os.Args[1])
		println("pass: ", password)
		user.SetPassword(password)
		db.DBService.MainDB.Save(&user)
		println("DB Saved")
		println("Sending mail start...")
		body := fmt.Sprintf(tmpl, password)

		if err := SendPassByEmail(user.Email, body); err == nil {
			println("Email sent")
		} else {
			println("Email die,", err)
		}
	} else {
		printHelp()
	}
}
