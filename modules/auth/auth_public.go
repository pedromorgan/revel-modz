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
	//doing this in signup.go
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

//returns success, user id, and error
func CheckUserActivationToken(db *gorm.DB, token string, now time.Time) (bool, int64, error) {
	var act UserAuthActivate
	err := db.Where(&UserAuthActivate{Token: token}).First(&act).Error
	if err == gorm.RecordNotFound {
		// fail
		return false, 0, errors.New("Activation failed")
	}
	if err != nil {
		return false, 0, err
	}

	if now.After(act.ExpiresAt) {
		// fail
		return false, 0, errors.New("Activation timed out")
	}

	// success
	// remove activate token
	err = db.Delete(&act).Error
	if err != nil {
		return false, 0, err
	}

	// get user auth data
	var u UserAuth
	err = db.Where(&UserAuth{UserId: act.UserId}).First(&u).Error
	if err != nil {
		return false, 0, err
	}
	// set activated to true
	u.Activated = true
	// update the user auth data
	err = db.Save(&u).Error
	if err != nil {
		return false, 0, err
	}

	return true, act.UserId, nil
}
