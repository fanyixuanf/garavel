/*
@Time : 2022/10/2 16:53
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : global
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package global

import (
	"garavel/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	ConfigEnv = "G_CONFIG"
	ConfigFile = "env.yaml"
	ResourcePath = "resources/**/*"
)

var (
	G_CONFIG config.Server
	G_VP     *viper.Viper
	G_LOG    *zap.Logger
)