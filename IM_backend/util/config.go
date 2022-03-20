/*
 * @Author: your name
 * @Date: 2022-02-28 23:07:47
 * @LastEditTime: 2022-03-21 00:06:41
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\util\config.go
 */
package util

import (
	"time"

	"github.com/spf13/viper"
)

//Config stores all configuration of the application
//The values are read by viper from a config file or environment varibles
type Config struct {
	AppMode             string        `mapstructure:"APP_MODE"`
	HttpPort            string        `mapstructure:"HTTP_PORT"`
	RedisAddr           string        `mapstructure:"REDIS_ADDR"`
	RedisDbName         string        `mapstructure:"REDIS_DBNAME"`
	RedisPw             string        `mapstructure:"REDIS_PW"`
	MongoDBSource       string        `mapstructure:"MONGO_DB_SOURCE"`
	MongoDBName         string        `mapstructure:"MONGO_DB_NAME"`
	MysqlDBSource       string        `mapstructure:"MYSQL_DB_SOURCE"`
	DBTESTSOURCE        string        `mapstructure:"DB_TEST_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	DefaultAvatar       string        `mapstructure:"DEFAULT_AVATAR"`
}

//LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
