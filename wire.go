//go:build wireinject
// +build wireinject

/*
@Time : 2022/10/1 21:19
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : write.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

func initializeEvent() Event {
	wire.Build(Event, Greeter, Message)

}