package controllers

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"

	htmpl "html/template"
	ttmpl "text/template"

	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/maillist"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
	"github.com/revel/revel/mail"

	"github.com/iassic/revel-modz/sample/app/models"
	"github.com/iassic/revel-modz/sample/app/routes"
)

func (c App) Signup() revel.Result {
	return c.Render()
}

func (c App) SignupPost(usersignup *models.UserSignup) revel.Result {
	usersignup.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	// check that this email is not in the DB already
	UB, err := user.GetUserBasicByName(c.Txn, usersignup.Email)
	checkERROR(err)

	if UB != nil {
		c.Validation.Error("Email already taken").Key("usersignup.Email")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	// uuid := get random number (that isn't used already)
	uuid, err := user.GenerateNewUserId(c.Txn)
	checkERROR(err)

	// update visitor info in DB with UserId
	c.updateVisitorWithUserIdPanic()

	// add user to tables
	// TODO do something more with the errors
	err = user.AddUserBasic(TestDB, uuid, usersignup.Email)
	checkERROR(err)

	err = auth.AddUser(TestDB, uuid, usersignup.Password)
	checkERROR(err)

	//generate token
	token := generateSecret()
	revel.INFO.Println(token)

	now := time.Now()
	expires := now.Add(time.Hour * 72)

	err = auth.AddUserActivationToken(TestDB, uuid, token, now, expires)
	checkERROR(err)

	sendActivationEmail(usersignup.Email, token)

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = "you should be receiving an email at " +
		usersignup.Email + " to confirm and activate your account."

	return c.Redirect(routes.App.Result())

}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) RegisterPost(userregister *models.UserRegister) revel.Result {
	userregister.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Register())
	}

	// check that this email is not in the DB already
	UB, err := user.GetUserBasicByName(c.Txn, userregister.Email)
	checkERROR(err)

	if UB != nil {
		c.Validation.Error("Email already taken").Key("userregister.Email")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	// uuid := get random number (that isn't used already)
	uuid, err := user.GenerateNewUserId(c.Txn)
	checkERROR(err)

	// update visitor info in DB with UserId
	c.updateVisitorWithUserIdPanic()

	// add user to tables
	// TODO do something more with the errors
	err = user.AddUserBasic(TestDB, uuid, userregister.Email)
	checkERROR(err)

	err = auth.AddUser(TestDB, uuid, userregister.Password)
	checkERROR(err)

	// TODO  which mailing lists did they check off?
	err = maillist.AddUser(TestDB, uuid, userregister.Email, "weekly")
	checkERROR(err)

	// TODO add address / phone DB insert
	// ...
	addy := &user.UserAddress{
		UserId:       uuid,
		AddressType:  "default",
		AddressLine1: userregister.Address1,
		AddressLine2: userregister.Address2,
		City:         userregister.City,
		State:        userregister.State,
		Zip:          userregister.Zipcode,
		Country:      userregister.Country,
	}
	err = user.AddUserAddress(TestDB, addy)
	checkERROR(err)

	err = user.AddUserPhone(TestDB, uuid, "default", userregister.PhoneNumber)
	checkERROR(err)

	//generate token
	token := generateSecret()
	revel.INFO.Println(token)

	now := time.Now()
	expires := now.Add(time.Hour * 72)

	err = auth.AddUserActivationToken(TestDB, uuid, token, now, expires)
	checkERROR(err)

	sendActivationEmail(userregister.Email, token)

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = "you should be receiving an email at " +
		userregister.Email + " to confirm and activate your account."

	return c.Redirect(routes.App.Result())
}

func (c App) ActivatePost(token string) revel.Result {

	// print token
	revel.WARN.Println("ActToken =", token)

	now := time.Now()

	// check token value and expiration
	success, uuid, err := auth.CheckUserActivationToken(TestDB, token, now)

	if err != nil || !success {

		c.Flash.Out["heading"] = "Activation FAILED :["
		c.Flash.Out["message"] = fmt.Sprint(err)

		return c.Redirect(routes.App.Result())
	} else {
		c.Flash.Out["heading"] = "Thanks for Activating!"
		c.Flash.Out["message"] = "You should be logged into your account."

		UB, err := user.GetUserBasicById(c.Txn, uuid)
		checkERROR(err)

		if UB == nil {
			revel.ERROR.Println("UB nil in activation post")
			panic("UB nil in activation post")
		}

		c.Session["user"] = UB.UserName
		c.RenderArgs["user_basic"] = UB

		return c.Redirect(routes.App.Result())
	}
	return nil
}

const alphaNumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func generateSecret() string {
	chars := make([]byte, 64)
	for i := 0; i < 64; i++ {
		chars[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(chars)
}

func sendActivationEmail(email, token string) error {
	revel.INFO.Println(email)
	revel.INFO.Println(token)

	params := activationEmailParams{"http://localhost:9000", token}

	text_tmpl := ttmpl.Must(ttmpl.New("text_body").Parse(email_text_body))
	text_buf := new(bytes.Buffer)
	text_wtr := bufio.NewWriter(text_buf)
	err := text_tmpl.Execute(text_wtr, params)
	if err != nil {
		revel.ERROR.Println("executing text_tmpl:", err)
	}
	text_wtr.Flush()

	html_tmpl := htmpl.Must(htmpl.New("html_body").Parse(email_html_body))
	html_buf := new(bytes.Buffer)
	html_wtr := bufio.NewWriter(html_buf)
	err = html_tmpl.Execute(html_wtr, params)
	if err != nil {
		revel.ERROR.Println("executing html_tmpl:", err)
	}
	html_wtr.Flush()

	mail_server := os.Getenv("MAIL_SERVER")
	mail_sender := os.Getenv("MAIL_SENDER")
	mail_passwd := os.Getenv("MAIL_PASSWD")

	message := mail.NewTextAndHtmlMessage(
		[]string{email},
		"Hello from iassic",
		text_buf.String(),
		html_buf.String(),
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

	err = mailer.SendMessage(message)
	revel.ERROR.Printf("%+v\n", err)
	return err
}

type activationEmailParams struct {
	WebLocation string
	Token       string
}

var email_text_body = `
Thanks for joining iassic.com!

To activate your account copy the following link into your browser:

{{.WebLocation}}/activate/{{.Token}}
`
var email_html_body = `
Thanks for joining iassic.com!<br><br>

To activate your account click the following link:
<br>
<a href="{{.WebLocation}}/activate/{{.Token}}">Activate Account</a>
`
