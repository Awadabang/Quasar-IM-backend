package token

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

//different types of error returned by the VerifyToken function
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload(有效负载) contains the payload data of the token
//ID to 防止token被泄露
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

//NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.New() //分配各自ID
	if err != nil {
		return nil, err
	}

	Payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return Payload, nil
}

func (payload *Payload) Valid() error{
	if time.Now().After(payload.ExpiredAt){
		return ErrExpiredToken
	}
	return nil
}