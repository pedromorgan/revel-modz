package maillist

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MaillistMessage struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Draft bool
	Sent  bool

	Author    string
	List      string
	Subject   string `sql:"not null;unique"`
	PlainBody string
	HtmlBody  string
}

func saveDraftMessage(db *gorm.DB, msg *MaillistMessage) error {
	return db.Debug().Save(msg).Error
}

func getDraftMessage(db *gorm.DB, subject string) (*MaillistMessage, error) {
	var mm MaillistMessage
	err := db.Where(&MaillistMessage{Subject: subject}).Find(&mm).Error
	if err != nil {
		return nil, err
	}
	return &mm, nil
}

func deleteDraftMessage(db *gorm.DB, subject string) error {
	return db.Where(&MaillistMessage{Subject: subject}).Delete(MaillistMessage{}).Error
}

func sendMessage(db *gorm.DB, subject string) error {
	msg, err := getDraftMessage(db, subject)
	if err != nil {
		return err
	}

	_ = msg

	return nil
}
