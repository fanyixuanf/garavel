/*
@Time : 2022/10/2 16:38
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : garavel.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import (
	"flag"
	"fmt"
	"garavel/bootstrap"
	myconfig "garavel/config"
	"garavel/global"
	"garavel/utils"
	"github.com/fsnotify/fsnotify"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var (
	config *myconfig.Server
	confPath string
	err    error
	level  zapcore.Level
	writer zapcore.WriteSyncer
	log *zap.Logger
	rootPath = utils.RootPath()
)

func main() {
	app, err :=bootstrap.InitializeApp(config, log)
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}

}

func init() {
	initConfig()
	initLog()
}

func initConfig() {
	flag.StringVar(&confPath, "c", "", "choose config file.")
	flag.Parse()
	if confPath == "" {
		// 优先级: 命令行 > 环境变量 > 默认值
		if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" {
			confPath =  filepath.Join(rootPath, "../", global.ConfigFile)
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", confPath)
		} else {
			confPath = configEnv
			fmt.Printf("您正在使用G_CONFIG环境变量,config的路径为%v\n", confPath)
		}
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", confPath)
	}
	v := viper.New()
	v.SetConfigFile(confPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
		if err := v.Unmarshal(&global.G_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}
	if err := v.Unmarshal(&global.G_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.G_VP = v
}

func initLog() {
	if ok, _ := utils.PathExists(config.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.Zap.Director)
		_ = os.Mkdir(config.Zap.Director, os.ModePerm)
	}

	switch config.Zap.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	writer, err = getWriteSyncer() // 日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		log = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		log = zap.New(getEncoderCore())
	}
	if config.Zap.ShowLine {
		log.WithOptions(zap.AddCaller())
	}
	global.G_LOG = log
}

// getWriteSyncer zap logger中加入file-rotatelogs
func getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		config.Zap.Director+string(os.PathSeparator)+"%Y-%m-%d.log",
		zaprotatelogs.WithLinkName(config.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (configs zapcore.EncoderConfig) {
	configs = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch  {
	case config.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		configs.EncodeLevel = zapcore.LowercaseLevelEncoder
	case config.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		configs.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case config.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		configs.EncodeLevel = zapcore.CapitalLevelEncoder
	case config.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		configs.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return configs
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if config.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(config.Zap.Prefix + "2006-01-02 - 15:04:05.000"))
}