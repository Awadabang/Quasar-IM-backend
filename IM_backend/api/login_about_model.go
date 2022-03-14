package api

import (
	"time"

	db "github.com/Awadabang/Quasar-IM/db/sqlc"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register
type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
