package token

import (
	"time"
)

type Payload struct {
	UserID   string    `json:"user_id"`
	Email    string    `json:"email"`
	IssuedAt time.Time `json:"issued_at"`
}

func NewPayload(userID string, email string) (*Payload, error) {
	payload := &Payload{
		UserID:   userID,
		Email:    email,
		IssuedAt: time.Now(),
	}
	return payload, nil
}
