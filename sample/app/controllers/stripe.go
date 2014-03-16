package controllers

import (
	"fmt"
	"github.com/drone/go.stripe"
	"github.com/revel/revel"
	"os"
)

func (c User) StripePost(stripeToken, plan string) revel.Result {
	var _key string
	var price int64
	price = 0
	fmt.Println(stripeToken)
	fmt.Printf("%+v\n", c.Params)

	_key = os.Getenv("STRIPE_API_KEY")
	stripe.SetKey(_key)

	if plan == "Standard" {
		price = 2000
	} else if plan == "Semester" {
		price = 1000
	} else if plan == "Corporate" {
		price = 10000
	} else {
		c.Flash.Out["heading"] = "Error Plan not found"
		c.Flash.Out["message"] = "Please select one of the avaliable plans"
	}
	params := stripe.ChargeParams{
		Desc:     plan,
		Amount:   price,
		Currency: stripe.USD,
		Token:    stripeToken,
	}
	charge, err := stripe.Charges.Create(&params)
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
