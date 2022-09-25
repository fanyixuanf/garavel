/*
@Time : 2022/9/25 19:22
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : redis
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
