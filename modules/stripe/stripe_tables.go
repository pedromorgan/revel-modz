package stripe

import (
	"fmt"
	"os"

	"github.com/drone/go.stripe"
	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(Customer{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(Subscription{}).Error
	if err != nil {
		return err
	}
	return nil
}
func DropTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(Customer{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(Subscription{}).Error
	if err != nil {
		return err
	}
	return nil
}

func FillTables(db *gorm.DB) error {
	_key := os.Getenv("STRIPE_API_KEY")
	stripe.SetKey(_key)

	for i, c := range stripeCustomers {
		uId := int64(100000 + i)
		uName := fmt.Sprintf("demo%d@domain.com", i)
		err := AddCustomerByCard(db, uId, uName, c.Email, c.Desc, c.Card)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func TestTables(db *gorm.DB) error {
	_key := os.Getenv("STRIPE_API_KEY")
	stripe.SetKey(_key)

	for i, c := range stripeCustomers {
		uId := int64(100000 + i)
		err := AddSubscription(db, uId, c.Plan)
		if err != nil {
			fmt.Println(err)
		}

	}
	return nil
}

// The way to create a test customer
// we should probably use the user/emails from the test data in module/user
//
// params := stripe.CustomerParams{
//
// }

// var stripeTokens = []stripe.TokenParams{
// 	//Currency: "usd",
// 	Card: &stripe.CardParams{
// 		Name:     "George Costanza",
// 		Number:   "4242424242424242",
// 		ExpYear:  2014,
// 		ExpMonth: 12,
// 	},
// }

var stripeCustomers = []stripe.CustomerParams{
	stripe.CustomerParams{
		Email: "george.costanza@mail.com",
		Desc:  "short, bald",
		Plan:  "Standard",
		Card: &stripe.CardParams{
			Name:     "George Costanza",
			Number:   "4242424242424242",
			ExpYear:  2014,
			ExpMonth: 12,
		},
	},
	stripe.CustomerParams{
		Email: "george.costanza@mail.com",
		Desc:  "short, bald",
		Plan:  "Standard",
		Card: &stripe.CardParams{
			Name:     "George Costanza",
			Number:   "4000000000000101",
			ExpYear:  2014,
			ExpMonth: 12,
		},
	},
	stripe.CustomerParams{
		Email: "george.costanza@mail.com",
		Desc:  "short, bald",
		Plan:  "Standard",
		Card: &stripe.CardParams{
			Name:     "George Costanza",
			Number:   "4000000000000341",
			ExpYear:  2014,
			ExpMonth: 12,
		},
	},
	stripe.CustomerParams{
		Email: "george.costanza@mail.com",
		Desc:  "short, bald",
		Plan:  "Standard",
		Card: &stripe.CardParams{
			Name:     "George Costanza",
			Number:   "4000000000000002",
			ExpYear:  2014,
			ExpMonth: 12,
		},
	},
	stripe.CustomerParams{
		Email: "george.costanza@mail.com",
		Desc:  "short, bald",
		Plan:  "Standard",
		Card: &stripe.CardParams{
			Name:     "George Costanza",
			Number:   "4000000000000127",
			ExpYear:  2014,
			ExpMonth: 12,
		},
	},
	stripe.CustomerParams{
		Email: "george.costanza@mail.com",
		Desc:  "short, bald",
		Plan:  "Standard",
		Card: &stripe.CardParams{
			Name:     "George Costanza",
			Number:   "4000000000000069",
			ExpYear:  2014,
			ExpMonth: 12,
		},
	},
	stripe.CustomerParams{
		Email: "george.costanza@mail.com",
		Desc:  "short, bald",
		Plan:  "Standard",
		Card: &stripe.CardParams{
			Name:     "George Costanza",
			Number:   "4000000000000119",
			ExpYear:  2014,
			ExpMonth: 12,
		},
	},
}
