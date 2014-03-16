package forum

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// database tables
type ForumTopic struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	TopicId   int64
	TopicName string
	TopicTags []string
}

type ForumMessage struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// keys
	TopicId   int64
	MessageId int64

	// message details
	AuthorId    int64 // UserId
	MessageBody string
}

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
