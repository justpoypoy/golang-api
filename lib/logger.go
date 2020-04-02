package lib

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	Helper "github.com/justpoypoy/api/helper" // Lib custom
)

func LoggerJSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var JSONLog LogJSON
		start := time.Now()
		c.Next()
		latency := Helper.GetDurationInMillseconds(start)
		JSONLog = LogJSON{
			ClientIP:   Helper.GetClientIP(c),
			TimeStamp:  start.Format("2006-01-02 15:04:05"),
			Method:     c.Request.Method,
			Path:       c.Request.RequestURI,
			StatusCode: c.Writer.Status(),
			Latency:    latency,
		}
		log.Println(JSONLog)
	}
}
