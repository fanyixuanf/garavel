/*
@Time : 2022/9/25 18:06
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : welcome
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package web

import (
	"fmt"
	"garavel/global"
	"github.com/gin-gonic/gin"
)

func Welcome (c *gin.Context) {
	c.HTML(200, "view/welcome.html", gin.H{"title": global.G_CONFIG.System.Name, "address": fmt.Sprintf("http://127.0.0.1:%d", global.G_CONFIG.System.Addr)})
}
