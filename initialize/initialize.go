/*
@Time : 2022/10/2 21:06
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : initialize
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package initialize

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(Routers, NewDB, GormWriteHandler,NewRedis)

type NewDb struct {
	wdb *gorm.DB
	redis *redis.Client
}

func NewDB (log *zap.Logger, wdb * gorm.DB, redis *redis.Client) (*NewDb, func(), error) {
	cleanup := func() {
		log.Info("closing the database")
	}
	return &NewDb{
		wdb: wdb,
		redis: redis,
	}, cleanup, nil
}
