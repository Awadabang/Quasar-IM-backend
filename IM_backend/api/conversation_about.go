/*
 * @Author: your name
 * @Date: 2022-02-28 22:33:31
 * @LastEditTime: 2022-03-20 15:17:24
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\api\conversation_about.go
 */
package api

import (
	"net/http"

	"github.com/Awadabang/Quasar-IM/middleware"
	"github.com/Awadabang/Quasar-IM/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) Get_conv(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	ctx.JSON(http.StatusOK, authPayload.Username)
}
