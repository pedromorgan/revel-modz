package forum

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTables(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(ForumTopic{}).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(ForumMessage{}).Error
	if err != nil {
		return err
	}
	return err
}

func DropTables(db *gorm.DB) error {
	var err error
	err = db.DropTable(ForumTopic{}).Error
	if err != nil {
		return err
	}
	err = db.DropTable(ForumMessage{}).Error
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
