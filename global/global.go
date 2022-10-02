/*
@Time : 2022/9/25 17:39
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : global
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package global

import (
	"garavel/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ConfigEnv = "G_CONFIG"
	ConfigFile = "env.yaml"
)
var (
	G_DB     *gorm.DB
	G_REDIS  *redis.Client
	G_CONFIG config.Server
	G_VP     *viper.Viper
	G_LOG    *zap.Logger
	G_READDB *gorm.DB
)