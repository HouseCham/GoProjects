package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
)

var myEmail string = "ramsesramirezvallejo@gmail.com"
var myPassword string = "***"
var sendGridToken string = "***"

func sendMailSimple(subject string, htmlBody string, to []string) {
	auth := smtp.PlainAuth(
		"Chamito",
		myEmail,
		myPassword,
		"smtp.gmail.com",
	)

	Headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := fmt.Sprintf("Subject: %s\n%s\n\n%s", subject, Headers, htmlBody)

	err := smtp.SendMail("smtp.gmail.com:587", auth, myEmail, to, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
}

func sendMailSimpleHTML(subject string, templatePath string, to []string) {

	//* Get html template
	var htmlBody bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(&htmlBody, struct {Name string}{Name: "Chamito"})

	auth := smtp.PlainAuth(
		"",
		myEmail,
		myPassword,
		"smtp.gmail.com",
	) 

	Headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := fmt.Sprintf("Subject: %s\n%s\n\n%s", subject, Headers, htmlBody.String())

	err = smtp.SendMail("smtp.gmail.com:587", auth, myEmail, to, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
}

func sendGoMail(templatePath string, subject string, to string) {
	//* Get html template
	var htmlBody bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(&htmlBody, struct {Name string}{Name: "Chamito"})

	m := gomail.NewMessage()
	m.SetHeader("From", myEmail)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", "azteca271999@gmail.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody.String())
	//m.Attach("./cat.webp")

	d := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func sendSendGrid(templatePath string) {
	from := mail.NewEmail("Ramsés", myEmail)
	subject := "This is an email sent using sendgrid-go"
	to := mail.NewEmail("Yaya", "chamses1999@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"

	//* Get html template
	var htmlBody bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(&htmlBody, struct {Name string}{Name: "Chamito"})

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlBody.String())
	client := sendgrid.NewSendClient(sendGridToken)

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	//sendMailSimpleHTML("Test email", "./email_test.html", []string{"chamses1999@gmail.com"})
	//sendGoMail("./email_test.html", "Email sent using gomail package", "chamses1999@gmail.com")
	sendSendGrid("./email_test.html")
}