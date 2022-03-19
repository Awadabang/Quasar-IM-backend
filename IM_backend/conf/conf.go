/*
 * @Author: Awadabang
 * @Date: 2022-02-28 22:33:31
 * @LastEditTime: 2022-03-20 01:06:48
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\conf\conf.go
 */
package conf

import (
	"context"
	"log"
	"strconv"

	"github.com/Awadabang/Quasar-IM/util"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//初始化
var (
	MongoDBClient *mongo.Client
	RedisClient   *redis.Client
	MongoDBSource string
	MnongoDBName  string
	RedisAddr     string
	RedisDBName   string
	RedisPw       string
)

func Init() {
	//viper
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load of config:", err)
	}
	MongoDBSource = config.MongoDBSource
	MnongoDBName = config.MongoDBName
	RedisAddr = config.RedisAddr
	RedisDBName = config.RedisDbName
	RedisPw = config.RedisPw

	//MongoDB
	MongoDB_Conn(MongoDBSource)
	//Redis
	Redis_Conn(RedisAddr, RedisDBName, RedisPw)
}

func MongoDB_Conn(connString string) {
	clientOptions := options.Client().ApplyURI(connString)
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	logging.Info("MongoDB successfully connect")
}

func Redis_Conn(RedisAddr string, RedisDbName string, RedisPw string) {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64) //string to uint64
	client := redis.NewClient(&redis.Options{       //登录Redis
		Addr:     RedisAddr,
		Password: RedisPw, // 无密码，注释掉就好了
		DB:       int(db),
	})
	_, err := client.Ping().Result() //验证是否ping通
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
