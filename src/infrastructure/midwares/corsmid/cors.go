package corsmid

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
	   method := c.Request.Method
	   c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	   c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	   c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, authoritaion")
	   c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	   if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
 
	   // 处理请求
	   c.Next()
	}
 }