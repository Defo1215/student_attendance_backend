package cors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		origin := c.Request.Header.Get("Origin")
//		if origin != "" {
//			//接收客户端发送的origin （重要！）
//			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
//			//服务器支持的所有跨域请求的方法
//			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
//			//允许跨域设置可以返回其他子段，可以自定义字段
//			c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
//			// 允许浏览器（客户端）可以解析的头部 （重要）
//			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
//			//设置缓存时间
//			c.Header("Access-Control-Max-Age", "172800")
//			//允许客户端传递校验信息比如 cookie (重要)
//			c.Header("Access-Control-Allow-Credentials", "true")
//			c.Set("Content-Type", "application/json")
//		}
//
//		//允许类型校验
//		if method == "OPTIONS" {
//			c.JSON(http.StatusOK, "ok!")
//		}
//
//		defer func() {
//			if err := recover(); err != nil {
//				fmt.Printf("Panic info is: %v\n", err)
//			}
//		}()
//		c.Next()
//	}
//}

func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,token,info")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
