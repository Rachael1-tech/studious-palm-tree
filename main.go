package main

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASSWORD")
	to := os.Getenv("SMTP_TO")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hi it's Rae Rae!!")
	m.SetBody("text/html", "<h1>Hello Rae Rae!</h1><p>This is a test email.</p>")

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}
	fmt.Println("Email sent!") 
}