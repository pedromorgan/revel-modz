package auth

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserAuthActivate struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserId      int64
	Token       string
	EmailSentAt time.Time
	ExpiresAt   time.Time
}
