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
	"context"
	"garavel/bootstrap"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)



func main() {

	rootCmd := &cobra.Command{
		Use:   "app",
		Run: func(cmd *cobra.Command, args []string) {
			app, cleanup, err := bootstrap.InitializeApp(config, log)
			if err != nil {
				panic(err)
			}
			defer cleanup()

			// 启动应用
			log.Info("start app ")
			if err := app.Run(); err != nil {
				panic(err)
			}

			// 优雅关闭应用
			quit := make(chan os.Signal)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit

			log.Info("shutdown app")

			// 设置 5 秒的超时时间
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// 关闭应用
			if err := app.Stop(ctx); err != nil {
				panic(err)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
