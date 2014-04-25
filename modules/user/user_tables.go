package user

import (
	"errors"
	"github.com/jinzhu/gorm"

	"github.com/iassic/revel-modz/modules/auth"
)

func AddTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(Visitor{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(User{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserBasic{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserDetail{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserAddress{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserPhone{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserProfileElement{}).Error
	if err != nil {
		return err
	}
	return nil
}

func DropTables(db *gorm.DB) error {
	var err error
	err = db.DropTable(Visitor{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(User{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserBasic{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserDetail{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserAddress{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserPhone{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserProfileElement{}).Error
	if err != nil {
		return err
	}
	return nil
}

func FillTables(db *gorm.DB) error {
	for _, u := range fakeFillTables {
		err := db.Save(u).Error
		if err != nil {
			return err
		}

		err = auth.AddUser(db, u.Basic.UserId, "demopass")
		if err != nil {
			return err
		}

	}
	return nil
}
func TestTables(db *gorm.DB) error {
	return errors.New("TODO")
}

var fakeFillTables = []*User{
	&User{

		Basic: UserBasic{
			UserId:   100001,
			UserName: "demo1@domain.com",
			Email:    "demo1@domain.com",
		},
		Detail: UserDetail{
			UserId:    100001,
			Title:     "Mr",
			FirstName: "Demo",
			LastName:  "Test",

			Dob: "04/07/1991",
			Sex: "M",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100001,
				AddressType: "Home", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10001",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},
		Phones: []UserPhone{
			UserPhone{
				UserId:      100001,
				PhoneType:   "Home", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0001",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100002,
			UserName: "demo2@domain.com",
			Email:    "demo2@domain.com",
		},
		Detail: UserDetail{
			UserId:    100002,
			Title:     "Ms",
			FirstName: "Demo",
			LastName:  "Assets",

			Dob: "04/07/1992",
			Sex: "F",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100002,
				AddressType: "Work", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10002",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100002,
				PhoneType:   "Work", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0002",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100003,
			UserName: "demo3@domain.com",
			Email:    "demo3@domain.com",
		},
		Detail: UserDetail{
			UserId:    100003,
			Title:     "Mr",
			FirstName: "Demo",
			LastName:  "Index",

			Dob: "04/07/1993",
			Sex: "M",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100003,
				AddressType: "School", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10003",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100003,
				PhoneType:   "School", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0003",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100004,
			UserName: "demo4@domain.com",
			Email:    "demo4@domain.com",
		},
		Detail: UserDetail{
			UserId:    100004,
			Title:     "Ms",
			FirstName: "Demo",
			LastName:  "Public",

			Dob: "04/07/1994",
			Sex: "F",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100004,
				AddressType: "International", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10004",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100004,
				PhoneType:   "International", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0004",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100005,
			UserName: "demo5@domain.com",
			Email:    "demo5@domain.com",
		},
		Detail: UserDetail{
			UserId:    100005,
			Title:     "Mr",
			FirstName: "Demo",
			LastName:  "Routes",

			Dob: "04/07/1995",
			Sex: "M",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100005,
				AddressType: "Parents", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10005",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100005,
				PhoneType:   "Parents", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0005",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100006,
			UserName: "demo6@domain.com",
			Email:    "demo6@domain.com",
		},
		Detail: UserDetail{
			UserId:    100006,
			Title:     "Mrs",
			FirstName: "Demo",
			LastName:  "Private",

			Dob: "04/07/1996",
			Sex: "F",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100006,
				AddressType: "Family", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10006",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100006,
				PhoneType:   "Home", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0006",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100007,
			UserName: "demo7@domain.com",
			Email:    "demo7@domain.com",
		},
		Detail: UserDetail{
			UserId:    100007,
			Title:     "Dr",
			FirstName: "Demo",
			LastName:  "Readme",

			Dob: "04/07/1997",
			Sex: "M",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100007,
				AddressType: "Home", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10007",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100007,
				PhoneType:   "Home", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0007",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100008,
			UserName: "demo8@domain.com",
			Email:    "demo8@domain.com",
		},
		Detail: UserDetail{
			UserId:    100008,
			Title:     "Dr",
			FirstName: "Demo",
			LastName:  "Grunt",

			Dob: "04/07/1998",
			Sex: "F",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100008,
				AddressType: "Office", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10008",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100008,
				PhoneType:   "Office", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0008",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100009,
			UserName: "demo9@domain.com",
			Email:    "demo9@domain.com",
		},
		Detail: UserDetail{
			UserId:    100009,
			Title:     "President",
			FirstName: "Barack",
			LastName:  "Obama",

			Dob: "04/07/1999",
			Sex: "M",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100009,
				AddressType: "White House", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10009",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100009,
				PhoneType:   "White House", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0009",
			},
		},
	},

	&User{

		Basic: UserBasic{
			UserId:   100010,
			UserName: "demouser010",
			Email:    "demo10@domain.com",
		},
		Detail: UserDetail{
			UserId:    100010,
			Title:     "FirstName Lady",
			FirstName: "Michelle",
			LastName:  "Obama",

			Dob: "04/07/2000",
			Sex: "F",
		},
		Addresses: []UserAddress{
			UserAddress{
				UserId:      100010,
				AddressType: "White House", // Home, Work, etc

				AddressLine1: "4400 Vestal Parkway East",
				AddressLine2: "BU Box 10010",
				City:         "Binghamton",
				State:        "NY",
				Zip:          "13902",
				Country:      "USA",
			},
		},

		Phones: []UserPhone{
			UserPhone{
				UserId:      100010,
				PhoneType:   "White House", // Home, Work, Cell, etc
				PhoneNumber: "607-777-0010",
			},
		},
	},
}
