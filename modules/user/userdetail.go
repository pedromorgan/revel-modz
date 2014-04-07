package user

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserDetail struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId int64

	Title     string
	FirstName string
	Middle    string
	LastName  string
	Suffix    string

	Dob string
	Sex string
}

func addUserDetail(db *gorm.DB, u *UserDetail) error {
	// TODO:  add check for existance?
	err := db.Save(u).Error
	return err
}

func deleteUserDetail(db *gorm.DB, uId int64) error {
	return db.Where(&UserDetail{UserId: uId}).Delete(UserDetail{}).Error
}

func getUserDetailByUserId(db *gorm.DB, userid int64) (*UserDetail, error) {
	var user UserDetail
	err := db.Where(&UserDetail{UserId: userid}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
