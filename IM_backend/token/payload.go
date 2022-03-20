/*
 * @Author: your name
 * @Date: 2022-03-09 20:57:16
 * @LastEditTime: 2022-03-20 15:57:46
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\token\payload.go
 */
package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
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
	Userid    int64     `json:"userid"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

//NewPayload creates a new token payload with a specific username and duration
func NewPayload(id int64, username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom() //分配各自ID
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Userid:    id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
