package controllers

import (
	"math/rand"
	"os"

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

	sendActivationEmail(usersignup.Email, token)
	/**********
		generate email

		send email  (with token as part of link)

		href="http://localhost:9000/activate?token=" + token
		<a href="http://localhost:9000/activate/TOKEN">Activate Account</a>

	 	Do DB stuff (later)

		**********/

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

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = "you should be receiving an email at " +
		userregister.Email + " to confirm and activate your account."

	return c.Redirect(routes.App.Result())
}

func (c App) ActivatePost(token string) revel.Result {

	revel.WARN.Println("ActToken =", token)

	c.Flash.Out["heading"] = "Thanks for Activating!"
	c.Flash.Out["message"] = "More of a message will be here eventually"

	return c.Redirect(routes.App.Result())
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

	mail_server := os.Getenv("MAIL_SERVER")
	mail_sender := os.Getenv("MAIL_SENDER")
	mail_passwd := os.Getenv("MAIL_PASSWD")

	message := mail.NewTextAndHtmlMessage(
		[]string{email},
		"Hello from iassic",
		email_text_body,
		email_html_body,
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

var email_text_body = `
Thanks for joining iassic.com!

To activate your account copy the following link into your browser:

{{web_location}}/activate?id={{.token}}
`
var email_html_body = `
Thanks for joining iassic.com!<br><br>

To activate your account click the following link:
<br>
<a href="{{web_location}}/activate?id={{.token}}">Activate Account</a>
`
