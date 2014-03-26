package controllers

import (
	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/maillist"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"

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

	err = auth.AddUser(TestDB, UB.UserId, usersignup.Password)
	checkERROR(err)

	/**********

		generate token

		generate email

		send email  (with token as part of link)

		href="http://localhost:9000/activate?token=" + token

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

	err = auth.AddUser(TestDB, UB.UserId, userregister.Password)
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

	return c.Redirect(routes.App.Results())
}
