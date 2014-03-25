package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"os"
	"github.com/drone/go.stripe"
)

func (c User) StripePost(stripeToken string) revel.Result {
	var _key string
	fmt.Println(stripeToken)
	fmt.Printf("%+v\n",c.Params)
    
	_key = os.Getenv("STRIPE_API_KEY")

	stripe.SetKey(_key)
	params:= stripe.ChargeParams{
		Desc: "Fee",
		Amount: 2000,
		Currency: stripe.USD,
		Token: stripeToken,
	}
	charge,err:=stripe.Charges.Create(&params)
	checkERROR(err)
	if err != nil {
	c.Flash.Out["heading"] = "Error charging card"
	c.Flash.Out["message"] = fmt.Sprint(err)

	} else {
	c.Flash.Out["heading"] = "Successfully Charged"
	c.Flash.Out["message"] = "Thank you for paying"
	}

	fmt.Printf("%+v\n", charge)
  // for now, just print out the arguments
	return c.Redirect("/u/result")
}