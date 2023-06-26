package middleware

import (
	"fmt"
	"gin_auto/pkg/comm"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func M1(c *gin.Context) {
	fmt.Println("come to M1")
}

func GetIp(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}

	return reqIP
}

func AuthCheck(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqIP := GetIp(c)
		if doCheck {
			if reqIP == "127.0.0.1" {
				// ip 白名单校验通过
				fmt.Println("AuthCheck Success, reqIP=:", reqIP)
				c.Next()
			} else {
				fmt.Println("AuthCheck failed, reqIP=:", reqIP)
				c.Abort()
			}

		} else {
			//不做校验
			c.Next()
		}
	}
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.Query("token")
		if token == "" {
			code = 201
		} else {
			claims, err := comm.ParseToken(token)
			if err != nil {
				code = 202
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 203
			}
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  200,
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
