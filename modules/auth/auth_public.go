package auth

import (
	"errors"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jinzhu/gorm"
)

func AddUser(db *gorm.DB, uId int64, password string) error {
	err := addUser(db, uId, password)
	if err != nil {
		return err
	}

	// TODO: send activation email and add activation record
	err = addActivationRecord(db, uId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *gorm.DB, uId int64) error {
	return db.Where(&UserAuth{UserId: uId}).Delete(UserAuth{}).Error
}

// returns true,nil on successful authentication; false,error otherwise
func Authenticate(db *gorm.DB, uId int64, password string) (bool, error) {
	var ua UserAuth
	err := db.Where(&UserAuth{UserId: uId}).Find(&ua).Error
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(ua.HashedPassword, []byte(password))
	if err != nil {
		return false, errors.New("Password Fail")
	}

	return true, nil
}

func UpdatePassword(db *gorm.DB, uId int64, new_password string) error {
	hPass, _ := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	ua := &UserAuth{
		UserId:         uId,
		HashedPassword: hPass,
	}
	found, err := checkUserExistsById(db, uId)
	if err != nil {
		return err
	}

	if found {
		return db.Save(ua).Error
	} else {
		return errors.New("user doesn't exists")
	}
}

func AddUserActivationToken(db *gorm.DB, uId int64, token string, sentAt, expires time.Time) error {
	act := UserAuthActivate{
		UserId:      uId,
		Token:       token,
		EmailSentAt: sentAt,
		ExpiresAt:   expires,
	}

	err := db.Save(&act).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckUserActivationToken(db *gorm.DB, uId int64, token string) error {
	return errors.New("Activate function not implemented")
}
