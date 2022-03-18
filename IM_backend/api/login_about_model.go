package api

import (
	"time"

	db "github.com/Awadabang/Quasar-IM/db/sqlc"
	"github.com/Awadabang/Quasar-IM/token"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register
type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,password"`
}

// Login
type userResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

type loginResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Verify
func hidePayload(user *token.Payload) hidepayloadResponse {
	return hidepayloadResponse{
		Username: user.Username,
	}
}

type hidepayloadResponse struct {
	Username string `json:"username"`
}

type verifyResponse struct {
	Access_token string              `json:"access_token"`
	User         hidepayloadResponse `json:"user"`
}
