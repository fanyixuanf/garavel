/*
@Time : 2022/9/25 19:23
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : system
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	NeedInitData  bool   `mapstructure:"need-init-data" json:"needInitData" yaml:"need-init-data"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	Name 		  string `mapstructure:"name" json:"name" yaml:"name"`
	ResourcePath  string `mapstructure:"resource-path" json:"resource-path" yaml:"resource-path"`
}

type Ipinfo struct {
	Token string `mapstructure:"token" json:"token" yaml:"token"`
}

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	MysqlRead  Mysql      `mapstructure:"mysqlread" json:"mysqlread" yaml:"mysqlread"`
	Sqlite     Sqlite     `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Sqlserver  Sqlserver  `mapstructure:"sqlserver" json:"sqlserver" yaml:"sqlserver"`
	Postgresql Postgresql `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	Ipinfo Ipinfo `mapstructure:"ipinfo" json:"ipinfo" yaml:"ipinfo"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path" `
}

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"`
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	From     string `mapstructure:"from" json:"from" yaml:"from"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"`
}
