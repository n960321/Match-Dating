package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)


func timeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		log.Info().Int("StatusCode", statusCode).Str("ReqUri", reqUri).Str("Method", reqMethod).Dur("latencyTime", latencyTime).Msg("Request End")
	}
}
