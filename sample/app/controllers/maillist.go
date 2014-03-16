package controllers

import (
	"github.com/iassic/revel-modz/modules/maillist"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/sample/app/models"
	"github.com/iassic/revel-modz/sample/app/routes"
)

// Anyone functions
func (c App) Maillist() revel.Result {
	return c.Render()
}

func (c App) MaillistPost(usermaillist *models.UserMaillist, list string) revel.Result {
	usermaillist.Validate(c.Validation)

	if c.Validation.HasErrors() || (list != "weekly" && list != "longer") {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Maillist())
	}

	// check that this email is not in the DB already
	UB, err := user.GetUserBasicByName(c.Txn, usermaillist.Email)
	checkERROR(err)
	if UB != nil {
		c.Validation.Error("Email already taken").Key("usermaillist.Email")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	err = c.addNewMaillistUser(usermaillist.Email, list)
	checkERROR(err)

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = usermaillist.Email + " is now subscribed to the " + list + " mailing list."

	return c.Redirect(routes.App.Result())

}

// Admin functions
func (c Admin) MaillistView() revel.Result {
	return c.Render()
}

func (c Admin) MaillistFilter(list, email string) revel.Result {

	if email != "" {
		maillist_users, err := maillist.GetUserByEmail(c.Txn, email)
		revel.INFO.Println("Got here email", *maillist_users, err)
		if err != nil {
			revel.ERROR.Println(err)
			return c.RenderJson(err)
		}
		return c.RenderJson(maillist_users)

	} else if list == "all" {
		maillist_users, err := maillist.GetAllUsers(c.Txn)
		revel.INFO.Println("Got here ALL", len(maillist_users), err)
		if err != nil {
			revel.ERROR.Println(err)
			return c.RenderJson(err)
		}
		return c.RenderJson(maillist_users)

	} else if list != "" {
		maillist_users, err := maillist.GetUsersByList(c.Txn, list)
		revel.INFO.Println("Got here list:", list, len(maillist_users), err)
		if err != nil {
			revel.ERROR.Println(err)
			return c.RenderJson(err)
		}
		return c.RenderJson(maillist_users)

	} else {
		revel.ERROR.Println("Shouldn't get HERE")
	}
	return nil
}

func (c Admin) MaillistCompose() revel.Result {
	return c.Render()
}

func (c Admin) MaillistComposePost() revel.Result {
	return nil
}

// helper functions

func (c App) addNewMaillistUser(email, list string) error {

	// uuid := get random number (that isn't used already)
	uuid, err := user.GenerateNewUserId(c.Txn)
	if err != nil {
		return err
	}

	err = user.AddUserBasic(TestDB, uuid, email)
	checkERROR(err)

	err = maillist.AddUser(TestDB, uuid, email, list)
	checkERROR(err)

	return nil
}
