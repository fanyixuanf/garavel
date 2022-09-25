/*
@Time : 2022/9/25 19:14
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : config
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package boostrap

import (
	"flag"
	"fmt"
	"garavel/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var config string

func init() {
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
		if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" {
			config = global.ConfigFile
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", global.ConfigFile)
		} else {
			config = configEnv
			fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
		}
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.G_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.G_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.G_VP = v
}
