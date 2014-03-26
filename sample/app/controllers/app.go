package controllers

import (
	"fmt"
	"strconv"

	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
)

type App struct {
	DbController
}

func (c App) RenderArgsFill() revel.Result {
	u := c.userConnected()
	if u != nil {
		c.RenderArgs["user_basic"] = u

		// look up role in RBAC module
		isAdmin := u.UserName == "admin@domain.com"
		if isAdmin {
			// set up things for an admin role
			c.Session["admin"] = "true"
		}
		return nil
	}
	v := c.visitorConnected()
	if v != nil {
		c.RenderArgs["visitor"] = v
		c.Session["v"] = fmt.Sprint(v.VisitorId)
		return nil
	}

	return nil
}

func (c App) userConnected() *user.UserBasic {
	if c.RenderArgs["user_basic"] != nil {
		return c.RenderArgs["user_basic"].(*user.UserBasic)
	}
	if username, ok := c.Session["user"]; ok {
		u, err := user.GetUserBasicByName(c.Txn, username)
		checkERROR(err)
		if u == nil {
			revel.ERROR.Println("user field in Session[] not found in DB")
			return nil
		}

		// update user info in DB, what info?

		if c.RenderArgs["visitor"] != nil {
			// update visitor info in DB with UserId
			c.updateVisitorWithUserIdPanic()
			// remove visitor fields in RenderArgs and Session?
		}

		return u
	}
	return nil
}

func (c App) visitorConnected() *user.Visitor {
	if c.RenderArgs["visitor"] != nil {
		return c.RenderArgs["visitor"].(*user.Visitor)
	}
	if visitor_id, ok := c.Session["v"]; ok {
		v_id, err := strconv.ParseInt(visitor_id, 10, 64)
		checkERROR(err)
		v, err := user.GetVisitorByVisitorId(c.Txn, v_id)
		checkERROR(err)
		if v == nil {
			revel.ERROR.Println("visitor field in Session[] not found in DB")
			delete(c.Session, "v")
			goto new_visitor_label
		}

		c.updateVisitor(v)
		return v
	}

	// if we get here, we have a new visitor or they have deleted their cookies
new_visitor_label:

	revel.WARN.Println("New visitor")
	v, err := c.addNewVisitor()
	checkERROR(err)

	return v
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Result() revel.Result {
	return c.Render()
}


// Defined in signup.go
// func (c App) Signup() revel.Result
// func (c App) SignupPost(usersignup *models.UserSignup) revel.Result
// func (c App) Register() revel.Result
// func (c App) RegisterPost(userregister *models.UserRegister) revel.Result
// func (c App) addNewUser(email, password string) (*user.UserBasic, error)

// Defined in maillist.go
// func (c App) Maillist() revel.Result
// func (c App) MaillistPost(usermaillist *models.UserMaillist) revel.Result
// func (c App) addNewMaillistUser(email, list string) (*maillist.MaillistUser, error)

// Defined in auth.go
// func (c App) Login() revel.Result
// func (c App) LoginPost(userlogin *models.UserLogin) revel.Result
// func (c App) Logout() revel.Result
