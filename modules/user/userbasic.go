package user

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserBasic struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId   int64  `sql:"not null;unique"`
	UserName string `sql:"not null;unique"`
	Email 	string
}

func addUserBasic(db *gorm.DB, u *UserBasic) error {
	// TODO:  add check for existance?
	err := db.Save(u).Error
	return err
}

func deleteUserBasic(db *gorm.DB, uId int64) error {
	return db.Where(&UserBasic{UserId: uId}).Delete(UserBasic{}).Error
}

func getUserBasicById(db *gorm.DB, id int64) (*UserBasic, error) {
	var user UserBasic
	err := db.Where(&UserBasic{Id: id}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getUserBasicByUserId(db *gorm.DB, userid int64) (*UserBasic, error) {
	var user UserBasic
	err := db.Where(&UserBasic{UserId: userid}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getUserBasicByName(db *gorm.DB, name string) (*UserBasic, error) {
	var user UserBasic
	err := db.Where(&UserBasic{UserName: name}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
