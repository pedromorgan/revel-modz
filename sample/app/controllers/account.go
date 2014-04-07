package controllers

import (
	// "github.com/iassic/revel-modz/modules/user"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/iassic/revel-modz/sample/app/models"
	"github.com/revel/revel"
)

func (c User) Account() revel.Result {

	u := c.userConnected()
	UA, err := user.GetUserAddressesById(c.Txn, u.UserId)
	checkERROR(err)
	UD, err := user.GetUserDetailById(c.Txn, u.UserId)
	checkERROR(err)
	UP, err := user.GetUserPhonesById(c.Txn, u.UserId)
	checkERROR(err)

	ur := &models.UserRegister{
		Username: u.Email,
		Email:    u.Email,
	}

	if UD != nil {
		ur.Fname = UD.FirstName
		ur.MidInit = UD.Middle
		ur.Lname = UD.LastName
		ur.Dob = UD.Dob
		ur.Sex = UD.Sex
	}

	if UA != nil && len(UA) > 0 {
		ur.Address1 = UA[0].AddressLine1
		ur.Address2 = UA[0].AddressLine2
		ur.City = UA[0].City
		ur.State = UA[0].State
		ur.Zipcode = UA[0].Zip
		ur.Country = UA[0].Country
	}
	if UP != nil && len(UP) > 0 {
		ur.PhoneNumber = UP[0].PhoneNumber
	}
	c.RenderArgs["ur"] = ur

	return c.Render()
}
