package main

import (
	"fmt"
	"github.com/sfreiberg/gotwilio"
)

func main() {
	accountSid := "*******************************"
	authToken := "********************************"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+***********"
	to := "+8801714771513"
	message := "Welcome to Twilo!, Hello World"
	a, b, c := twilio.SendSMS(from, to, message, "", "")
	fmt.Println(a, b, c)
}
