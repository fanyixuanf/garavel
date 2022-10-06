/*
@Time : 2022/10/3 20:15
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : redis
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package initialize

import (
	"garavel/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func NewRedis(config *config.Server, log *zap.Logger) *redis.Client {
	redisCfg := config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		log.Info("redis connect ping response:", zap.String("pong",pong))
		return client
	}
	return client
}