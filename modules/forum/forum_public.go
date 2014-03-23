package forum

import (
	"time"

	"github.com/jinzhu/gorm"
)

// not a database table, for sending over the wire
// Used in the list of topics
type TopicBrief struct {
	TopicId     int64
	TopicTags   []string
	TopicName   string
	NumMessages string

	OpenedAt  time.Time
	OpenedBy  string
	LastMsgAt time.Time
	LastMsgBy string
}

type TopicDetail struct {
	TopicId   int64
	TopicTags []string
	TopicName string

	Messages []MessageWire
}

type MessageWire struct {
	TopicId     int64
	AuthorName  string
	MessageTime time.Time
	MessageBody string
}

func GetTopicList(db *gorm.DB) ([]ForumTopic, error) {
	var topics []ForumTopic

	err := db.Order("updated_at desc").Limit(20).Find(&topics).Error
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func GetAllMessagesByTopicId(db *gorm.DB, id int64) ([]ForumMessage, error) {
	var messages []ForumMessage

	err := db.Where(ForumMessage{TopicId: id}).Order("MessageId").Find(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}
