/*
@Time : 2022/9/25 19:55
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : loadtls
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Next()
	}
}
