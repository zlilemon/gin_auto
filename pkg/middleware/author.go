package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
