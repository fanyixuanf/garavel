/*
@Time : 2022/9/25 17:48
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : cors
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Cors) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id\"")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

type Cors struct {}

func NewCors() *Cors {
	return &Cors{}
}
