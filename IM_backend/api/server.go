/*
 * @Author: your name
 * @Date: 2022-03-09 00:08:12
 * @LastEditTime: 2022-03-20 00:58:53
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\api\server.go
 */
package api

import (
	"fmt"

	db "github.com/Awadabang/Quasar-IM/db/sqlc"
	"github.com/Awadabang/Quasar-IM/middleware"
	"github.com/Awadabang/Quasar-IM/service"
	"github.com/Awadabang/Quasar-IM/token"
	"github.com/Awadabang/Quasar-IM/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

//NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	//token maker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validPassword)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	//全局设置了跨域
	router.Use(middleware.Cors())
	v1 := router.Group("/")
	{
		v1.POST("login", server.Login)
		v1.POST("verify", server.Verify)
		v1.POST("register", server.Register)
		v1.GET("get_conv", server.Get_conv)

		v1.GET("ws", service.WsHandler)
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
