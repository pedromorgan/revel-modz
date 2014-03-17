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
	AuthorId    string // UserId
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

func GetAllTopics(db *gorm.DB) ([]ForumTopic, error) {
	return fakeTopics, nil
}

func GetAllMessagesByTopicId(db *gorm.DB, id int64) ([]ForumMessage, error) {
	msgs := make([]ForumMessage, 0)
	for _,m:=range fakeMessages{
		if m.TopicId == id {
			msgs = append(msgs, m)
		}
	}
}

var fakeTopics = []ForumTopic{
	ForumTopic{1,"Topic1", []string{"tagA", "tagB"}	},
	ForumTopic{2,"Topic2", []string{"tagC", "tagD"}	},
	ForumTopic{3,"Topic3", []string{"tagE", "tagF"}	},
	ForumTopic{4,"Topic4", []string{"tagG", "tagH"}	},
	ForumTopic{5,"Topic5", []string{"tagI", "tagJ"}	},
	ForumTopic{6,"Topic6", []string{"tagK", "tagL"}	},
}

var fakeMessages = []ForumMessage{}
	ForumMessage{1,1,"T1 Message1", "user1", "t1 msg1 body"},
	ForumMessage{1,2,"T1 Message2", "user2", "t1 msg2 body"},
	ForumMessage{1,3,"T1 Message3", "user3", "t1 msg3 body"},
	
	ForumMessage{2,1,"T2 Message1", "user2", "t2 msg1 body"},
	ForumMessage{2,2,"T2 Message2", "user3", "t2 msg2 body"},
	ForumMessage{2,3,"T2 Message3", "user4", "t2 msg3 body"},

	ForumMessage{3,1,"T3 Message1", "user3", "t3 msg1 body"},
	ForumMessage{3,2,"T3 Message2", "user4", "t3 msg2 body"},
	ForumMessage{3,3,"T3 Message3", "user5", "t3 msg3 body"},

	ForumMessage{4,1,"T4 Message1", "user4", "t4 msg1 body"},
	ForumMessage{4,2,"T4 Message2", "user5", "t4 msg2 body"},
	ForumMessage{4,3,"T4 Message3", "user6", "t4 msg3 body"},

	ForumMessage{5,1,"T5 Message1", "user5", "t5 msg1 body"},
	ForumMessage{5,2,"T5 Message2", "user6", "t5 msg2 body"},
	ForumMessage{5,3,"T5 Message3", "user7", "t5 msg3 body"},

	ForumMessage{6,1,"T6 Message1", "user6", "t6 msg1 body"},
	ForumMessage{6,2,"T6 Message2", "user7", "t6 msg2 body"},
	ForumMessage{6,3,"T6 Message3", "user8", "t6 msg3 body"},
}
