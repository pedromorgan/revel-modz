package forum

import (
	"time"
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
