package maillist

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(MaillistUser{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(MaillistMessage{}).Error
	if err != nil {
		return err
	}
	return nil

}

func DropTables(db *gorm.DB) error {
	var err error
	err = db.DropTable(MaillistUser{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(MaillistMessage{}).Error
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
