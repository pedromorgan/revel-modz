package stripe

import (
	"github.com/drone/go.stripe"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"os"
	"time"
)

type Customer struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId   int64  `sql:"not null;unique"`
	UserName string `sql:"not null;unique"`

	StripeId    string
	StripeEmail string
	StripeDesc  string
}

type Subscription struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId   int64 `sql:"not null;unique"`
	PlanName string
}

func AddCustomerByCard(db *gorm.DB, uId int64, uName, sEmail, sDesc string, card *stripe.CardParams) error {
	_key := os.Getenv("STRIPE_API_KEY")
	stripe.SetKey(_key)

	// create a Stripe Customer
	params := stripe.CustomerParams{
		Email: sEmail,
		Desc:  sDesc,
		Card:  card,
	}
	// save the Stripe Customer
	sCustomer, err := stripe.Customers.Create(&params)
	if err != nil {
		return err
	}

	// create a îassic Customer
	iCustomer := &Customer{
		UserId:      uId,
		UserName:    uName,
		StripeId:    sCustomer.Id, // b/c of this, need to save îassic Customer second
		StripeEmail: sEmail,
		StripeDesc:  sDesc,
	}
	// save the îassic Customer
	err = db.Save(iCustomer).Error
	if err != nil {
		return err
	}

	return nil
}

func AddCustomerByToken(db *gorm.DB, uId int64, uName, sEmail, sDesc, token string) error {
	_key := os.Getenv("STRIPE_API_KEY")
	stripe.SetKey(_key)

	// create a Stripe Customer
	params := stripe.CustomerParams{
		Email: sEmail,
		Desc:  sDesc,
		Token: token,
	}
	// save the Stripe Customer
	sCustomer, err := stripe.Customers.Create(&params)
	if err != nil {
		return err
	}

	// create a îassic Customer
	iCustomer := &Customer{
		UserId:      uId,
		UserName:    uName,
		StripeId:    sCustomer.Id, // b/c of this, need to save îassic Customer second
		StripeEmail: sEmail,
		StripeDesc:  sDesc,
	}
	// save the îassic Customer
	err = db.Save(iCustomer).Error
	if err != nil {
		return err
	}

	return nil
}

func GetCustomer(db *gorm.DB, uId int64) (*Customer, error) {
	var c Customer
	err := db.Where(&Customer{UserId: uId}).First(&c).Error
	if err == gorm.RecordNotFound {
		return nil, err
	}
	if err != nil {
		revel.TRACE.Println(err)
		return nil, err
	}
	return &c, nil
}

func AddSubscription(db *gorm.DB, uId int64, plan string) error {
	_key := os.Getenv("STRIPE_API_KEY")
	stripe.SetKey(_key)

	iCustomer, err := GetCustomer(db, uId)
	if err != nil {
		return err
	}

	params := stripe.SubscriptionParams{
		Plan:    plan,
		Prorate: true,
	}

	iSubsription := &Subscription{
		UserId:   uId,
		PlanName: plan,
	}

	err = db.Debug().Save(iSubsription).Error
	if err != nil {
		return err
	}

	_, err = stripe.Subscriptions.Update(iCustomer.StripeId, &params)
	if err != nil {
		// TODO, should we do a rollback on the DB here???
		return err
	}

	return nil
}

func GetSubscription(db *gorm.DB, uId int64) (*Subscription, error) {
	var s Subscription
	err := db.Where(&Subscription{UserId: uId}).First(&s).Error
	if err == gorm.RecordNotFound {
		return nil, err
	}
	if err != nil {
		revel.TRACE.Println(err)
		return nil, err
	}
	return &s, nil
}

func CancelSubscription(db *gorm.DB, uId int64) error {

	_key := os.Getenv("STRIPE_API_KEY")
	stripe.SetKey(_key)

	// find îassic customer
	iCustomer, err := GetCustomer(db, uId)
	if err != nil {
		return err
	}
	// find îassic subscription
	iSubscription, err := GetSubscription(db, uId)
	if err != nil {
		return err
	}

	// delete îassic subscription
	err = db.Delete(&iSubscription).Error
	if err != nil {
		return err
	}

	// delete stripe subscription
	_, err = stripe.Subscriptions.Cancel(iCustomer.StripeId)
	if err != nil {
		return err
	}

	return nil
}
