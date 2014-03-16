package controllers

import (
	"os"
	// "github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
)

func (c User) Account() revel.Result {

	publick := os.Getenv("STRIPE_TEST_PUB")
	privatek := os.Getenv("STRIPE_API_KEY")
    revel.WARN.Println(publick)
    revel.WARN.Println(privatek)

	// u := c.userConnected()

	// // get stuff from DB
	// userbasic := getU

	// // create & file in UserRegister struct
	// userregister := &UserRegister {
	// 	...
	// }

	// c.RenderArgs["userregister"] = userregister

	return c.Render(publick)
}
