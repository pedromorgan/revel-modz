package maillist

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type MaillistUser struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId    int64 `sql:"not null"`
	Email     string
	List      string `sql:"not null"`
	Activated bool
}

func checkUserExistsById(db *gorm.DB, uId int64) (bool, error) {
	var mu MaillistUser
	err := db.Where(&MaillistUser{UserId: uId}).Find(&mu).Error
	if err == gorm.RecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if uId == mu.UserId {
		return true, nil
	}
	return false, nil
}

func checkUserExistsByEmail(db *gorm.DB, email string) (bool, error) {
	var mu MaillistUser
	err := db.Where(&MaillistUser{Email: email}).Find(&mu).Error
	if err == gorm.RecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if email == mu.Email {
		return true, nil
	}
	return false, nil
}

func addUser(db *gorm.DB, mu *MaillistUser) error {
	idFound, err := checkUserExistsById(db, mu.UserId)
	if err != nil {
		return err
	}
	emailFound, err := checkUserExistsByEmail(db, mu.Email)
	if err != nil {
		return err
	}
	if !idFound && !emailFound {
		return db.Save(mu).Error
	}
	return errors.New("User already exists")
}

func deleteUserById(db *gorm.DB, uId int64) error {
	return db.Where(&MaillistUser{UserId: uId}).Delete(MaillistUser{}).Error
}

func getUserById(db *gorm.DB, uId int64) (*MaillistUser, error) {
	var mu MaillistUser
	err := db.Where(&MaillistUser{UserId: uId}).Find(&mu).Error
	if err != nil {
		return nil, err
	}
	return &mu, nil
}

func getUserByEmail(db *gorm.DB, email string) (*MaillistUser, error) {
	var mu MaillistUser
	err := db.Where(&MaillistUser{Email: email}).Find(&mu).Error
	if err != nil {
		return nil, err
	}
	return &mu, nil
}

func getAllUsers(db *gorm.DB) ([]MaillistUser, error) {
	var mus []MaillistUser
	err := db.Find(&mus).Error
	if err != nil {
		return nil, err
	}
	return mus, nil
}

func getUsersByList(db *gorm.DB, list string) ([]MaillistUser, error) {
	var mus []MaillistUser
	err := db.Where(&MaillistUser{List: list}).Find(&mus).Error
	if err != nil {
		return nil, err
	}
	return mus, nil
}
