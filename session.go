package sqs

import (
	"time"
)

type Session struct {
	UserID    int
	CreatedAt time.Time
	ValidTo   time.Time
	IsInvalid bool
}

func NewSession(userID int, validTo time.Time) *Session {
	return &Session{
		UserID:    userID,
		CreatedAt: time.Now(),
		ValidTo:   validTo,
		IsInvalid: false,
	}
}

func (s *Session) Clone() *Session {
	return &Session{
		UserID:    s.UserID,
		CreatedAt: s.CreatedAt,
		ValidTo:   s.ValidTo,
		IsInvalid: s.IsInvalid,
	}
}
