package controllers

import (
	"os"

	"code.google.com/p/go.net/websocket"
	"github.com/iassic/revel-modz/modules/ws_comm"
	"github.com/revel/revel"
	"github.com/revel/revel/mail"
)

func (c User) Comm(ws *websocket.Conn) revel.Result {
	user := c.userConnected()

	comm := ws_comm.New()
	comm.AddHandler("echo", echoHandler)

	if user.UserId == 200001 {
		comm.AddHandler("email", emailHandler)
	}

	comm.Serve(ws)

	revel.INFO.Println("closing WS connection")
	return nil
}

func echoHandler(msg string, outbound chan string) {
	revel.INFO.Println("Echo: ", msg)
	outbound <- "ack " + msg
}

func emailHandler(msg string, outbound chan string) {
	revel.ERROR.Printf("%q\n", msg)
	if msg == "send test" {
		revel.WARN.Println("sending test message")
		err := sendTestMessage()
		if err != nil {
			outbound <- "error sending message"
			return
		}
		outbound <- "ack email sent"
		return
	}
	outbound <- "error unknown email command"
}

func sendTestMessage() error {
	mail_server := os.Getenv("MAIL_SERVER")
	mail_sender := os.Getenv("MAIL_SENDER")
	mail_passwd := os.Getenv("MAIL_PASSWD")

	message := mail.NewTextAndHtmlMessage(
		[]string{"tony@iassic.com"},
		"Test Message",
		"Test Text Body",
		"Test <b>Html</b> <i>Body</i><br>",
	)
	// message.Cc = []string{"admin@domain.com"}
	// message.Bcc = []string{"secret@domain.com"}
	sender := mail.Sender{
		From:    mail_sender,
		ReplyTo: mail_sender,
	}

	mailer := mail.Mailer{
		Server:   mail_server,
		Port:     587,
		UserName: mail_sender,
		Password: mail_passwd,
		// Host: "iassic.com",
		// Auth: smtp.Auth,
		Sender: &sender,
	}

	revel.ERROR.Printf("%+v\n", mailer)
	err := mailer.SendMessage(message)
	revel.ERROR.Printf("%+v\n", err)
	return err
}
