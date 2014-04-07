package controllers

import (
	gorm "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/modules/analytics"
	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/maillist"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/iassic/revel-modz/modules/user-files"
)

var (
	TestDB *gorm.DB

	fill = false
)

type DbController struct {
	*revel.Controller
	Txn *gorm.DB
}

func (c *DbController) Begin() revel.Result {
	txn := TestDB.Begin()
	err := txn.Error
	checkPANIC(err)
	c.Txn = txn
	return nil
}

func (c *DbController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	err := c.Txn.Commit().Error
	checkPANIC(err)

	c.Txn = nil
	return nil
}

func (c *DbController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	err := c.Txn.Rollback().Error
	checkPANIC(err)

	c.Txn = nil
	return nil
}

func InitDB() {

	var driver, spec string
	var found bool
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.ERROR.Fatal("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.ERROR.Fatal("No db.spec found.")
	}

	// Open a connection.
	ndb, err := gorm.Open(driver, spec)
	checkPANIC(err)

	ndb.SetLogger(gorm.Logger{revel.INFO})
	// ndb.LogMode(true)

	TestDB = &ndb

	revel.INFO.Println("Connection made to DB")
}

func SetupTables() {
	revel.INFO.Println("Setting up Prod DB")
	addTables()
}

func SetupDevDB() {
	revel.INFO.Println("Setting up Dev DB")
	dropTables()
	addTables()

	user.FillTables(TestDB)

	// testUserDB()
}

func dropTables() {
	revel.INFO.Println("Dropping tables")
	// analytics.DropTables(TestDB)
	auth.DropTables(TestDB)
	user.DropTables(TestDB)
	maillist.DropTables(TestDB)
	// userfiles.DropTables(TestDB)
}

func addTables() {
	revel.INFO.Println("AutoMigrate tables")
	analytics.AddTables(TestDB)
	auth.AddTables(TestDB)
	user.AddTables(TestDB)
	maillist.AddTables(TestDB)
	userfiles.AddTables(TestDB)
}

type DbFillUser struct {
	UserId   int64
	UserName string
	Password string
	Maillist string
}

var dev_users = []*DbFillUser{
	&DbFillUser{UserId: 100001, UserName: "demo1@domain.com", Password: "demopass", Maillist: "weekly"},
	&DbFillUser{UserId: 100002, UserName: "demo2@domain.com", Password: "demopass", Maillist: "weekly"},
	&DbFillUser{UserId: 100003, UserName: "demo3@domain.com", Password: "demopass", Maillist: "longer"},
	&DbFillUser{UserId: 100004, UserName: "demo4@domain.com", Password: "demopass", Maillist: "longer"},
	&DbFillUser{UserId: 200001, UserName: "admin@domain.com", Password: "adminpass", Maillist: ""},
}

func fillUserTables() {
	var err error
	for _, devuser := range dev_users {
		err = user.AddUserBasic(TestDB, devuser.UserId, devuser.UserName)
		checkERROR(err)

		err = auth.AddUser(TestDB, devuser.UserId, devuser.Password)
		checkERROR(err)
	}
	revel.INFO.Println("Filled User DBs")
}
func fillMailTables() {
	var err error
	for _, um := range dev_users {
		if um.Maillist == "" {
			continue
		}
		err = maillist.AddUser(TestDB, um.UserId, um.UserName, um.Maillist)
		checkERROR(err)
	}
	revel.INFO.Println("Filled maillist DBs")
}

func testUserDB() {
	for _, up := range dev_users {

		u, err := user.GetUserBasicById(TestDB, up.UserId)
		checkERROR(err)
		if u == nil {
			revel.ERROR.Println("Failed to look up user by id:", up.UserId)
		}

		u, err = user.GetUserBasicByName(TestDB, up.UserName)
		checkERROR(err)
		if u == nil {
			revel.ERROR.Println("Failed to look up user by name:", up.UserName)
		}

		passed, err := auth.Authenticate(TestDB, up.UserId, up.Password)
		checkERROR(err)
		if !passed {
			revel.ERROR.Printf("Failed to authenticate user: %+v\n", *up)
		}
	}
}
