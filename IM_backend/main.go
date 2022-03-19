/*
 * @Author: your name
 * @Date: 2022-02-28 22:33:31
 * @LastEditTime: 2022-03-19 22:29:59
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\main.go
 */
package main

import (
	"database/sql"
	"log"

	"github.com/Awadabang/Quasar-IM/api"
	"github.com/Awadabang/Quasar-IM/conf"
	db "github.com/Awadabang/Quasar-IM/db/sqlc"
	"github.com/Awadabang/Quasar-IM/service"
	"github.com/Awadabang/Quasar-IM/util"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//viper
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load of config:", err)
	}

	//Connect to MongoDB
	conf.Init()
	//Connect to Redis
	//conf.Redis_Conn(config.RedisAddr, config.RedisDbName, config.RedisPw)

	//TODO: Mock

	go service.Manager.Start()

	conn, err := sql.Open("mysql", config.MysqlDBSource)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("connot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}
