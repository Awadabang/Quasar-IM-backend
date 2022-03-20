/*
 * @Author: your name
 * @Date: 2022-03-20 15:40:21
 * @LastEditTime: 2022-03-20 16:43:04
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\api\friend_about.go
 */
package api

import (
	"net/http"

	db "github.com/Awadabang/Quasar-IM/db/sqlc"
	"github.com/Awadabang/Quasar-IM/middleware"
	"github.com/Awadabang/Quasar-IM/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) Add_friend(ctx *gin.Context) {
	var req addfriendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.AddFriendParams{
		Owner:    authPayload.Userid,
		FriendID: req.Friendid,
	}

	_, err := server.store.AddFriend(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) Get_friends(ctx *gin.Context) {
	var req getfriendsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.GetOnesFriendsParams{
		Owner:  authPayload.Userid,
		Limit:  req.Page_size,
		Offset: (req.Page_id - 1) * req.Page_size,
	}

	friends, err := server.store.GetOnesFriends(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var rsp []getfriendsResponse

	for _, v := range friends {
		rsp = append(rsp, getfriendsResponse{
			Friend_id: v.FriendID,
			Username:  v.Username,
		})
	}

	ctx.JSON(http.StatusOK, rsp)
}
