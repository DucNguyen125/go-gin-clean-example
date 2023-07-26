package middlewares

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RestLogger(context *gin.Context) {
	path := context.Request.URL.Path
	start := time.Now()
	context.Next()
	stop := time.Since(start)
	latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
	statusCode := context.Writer.Status()
	clientIP := context.ClientIP()
	clientUserAgent := context.Request.UserAgent()
	method := context.Request.Method
	logDetail := log.Fields{
		"statusCode": statusCode,
		"latency":    fmt.Sprintf("%dms", latency), // time to process
		"clientIP":   clientIP,
		"method":     method,
		"path":       path,
		"userAgent":  clientUserAgent,
	}
	if query := context.Request.URL.RawQuery; query != "" {
		logDetail["query"] = query
	}
	logger := log.WithFields(logDetail)
	if len(context.Errors) > 0 {
		logger.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
	} else {
		msg := "[GIN]"
		if statusCode >= http.StatusInternalServerError {
			logger.Error(msg)
		} else if statusCode >= http.StatusBadRequest {
			logger.Warn(msg)
		} else {
			logger.Info(msg)
		}
	}
}
