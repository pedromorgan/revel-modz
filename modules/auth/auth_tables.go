package auth

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(UserAuth{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(UserAuthActivate{}).Error
	if err != nil {
		return err
	}
	return nil
}

func DropTables(db *gorm.DB) error {
	var err error
	err = db.DropTable(UserAuth{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(UserAuthActivate{}).Error
	if err != nil {
		return err
	}
	return nil
}

func FillTables(db *gorm.DB) error {
	return errors.New("TODO")
}
func TestTables(db *gorm.DB) error {
	return errors.New("TODO")
}
