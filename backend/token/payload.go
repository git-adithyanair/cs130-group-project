package token

import (
	"time"
)

type Payload struct {
	UserID   int64     `json:"user_id"`
	Email    string    `json:"email"`
	IssuedAt time.Time `json:"issued_at"`
}

func NewPayload(userID int64, email string) (*Payload, error) {
	payload := &Payload{
		UserID:   userID,
		Email:    email,
		IssuedAt: time.Now(),
	}
	return payload, nil
}
