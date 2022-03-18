package api

import (
	"database/sql"
	"net/http"

	db "github.com/Awadabang/Quasar-IM/db/sqlc"
	"github.com/Awadabang/Quasar-IM/util"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

// API: Register Login Verify

func (server *Server) Register(ctx *gin.Context) {
	// binding the request of register
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// make the password to hash
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// create the user in DB
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
	}
	// insert the user into DB
	_, err = server.store.CreateUser(ctx, arg)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}

//测试access_token的合法性
func (server *Server) Verify(c *gin.Context) {
	var auth string = c.GetHeader("Authorization")
	user, err := server.tokenMaker.VerifyToken(auth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	rsp := verifyResponse{
		Access_token: auth,
		User:         hidePayload(user),
	}

	c.JSON(http.StatusOK, rsp)
}
