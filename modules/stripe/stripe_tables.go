package stripe

import (
	"github.com/drone/go.stripe"
	"github.com/jinzhu/gorm"
)

// The way to create a test customer
// we should probably use the user/emails from the test data in module/user
//
// params := stripe.CustomerParams{
//     Email:  "george.costanza@mail.com",
//     Desc:   "short, bald",
//     Card:   &stripe.CardParams {
//         Name     : "George Costanza",
//         Number   : "4242424242424242",
//         ExpYear  : 2014,
//         ExpMonth : 12,
//     },
// }
