package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bupaanyone/sioux/server/config"
	"github.com/bupaanyone/sioux/server/utils"
)

const apiPrefix = "api/"

func Init() {
	engine := gin.New()
	engine.Use(middlewareLogger())
	engine.Use(gin.Recovery())

	groupVisitor := engine.Group(config.C.Handler.UrlPrefix)
	groupVisitor.GET(apiPrefix+"ping", ping)

	go func() {
		if err := engine.Run(config.C.Handler.Host); err != nil {
			panic(err)
		}
	}()
}

// @Tags 测试
// @Summary 联通性测试
// @Description 联通性测试接口。
// @Produce plain/text
// @Success 200 {string} string "pong"
// @Router /ping [get]
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func middlewareLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.CtxSetLogIdGin(c)

		start := time.Now().UnixNano()
		c.Next()
		duration := (time.Now().UnixNano() - start) / int64(time.Millisecond)

		utils.Log(c, "[ACCESS] Handled request, status=%d, method=%s, url=%s, client=%s, latency=%dms, size=%d.",
			c.Writer.Status(), c.Request.Method, c.Request.URL.Path, c.ClientIP(), duration, c.Writer.Size())
	}
}

func middlewareCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}
