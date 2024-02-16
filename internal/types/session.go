package types

import (
	"time"
)

type Session struct {
	ID        string
	UserID    int
	Token     string
	CreatedAt time.Time
}

func (session *Session) IsEmptySession() bool {
	return session.ID == "" && session.UserID == 0 && session.Token == "" && session.CreatedAt.IsZero()
}
