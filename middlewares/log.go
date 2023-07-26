package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

var ExcludingApis = map[string]bool{
	"/api/v1/auth/login": true,
}

func getRequestBody(context *gin.Context) string {
	bodyBuffer := &bytes.Buffer{}
	if strings.Contains(context.Request.Header.Get("Content-Type"), "application/json") {
		data, _ := io.ReadAll(context.Request.Body)
		_ = json.Compact(bodyBuffer, data)
		context.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // Write body back
	}
	return strings.ReplaceAll(bodyBuffer.String(), `"`, `'`)
}

func getResponseBody(blw *bodyLogWriter, context *gin.Context) string {
	if strings.Contains(context.Writer.Header().Get("Content-Type"), "application/json") {
		return strings.ReplaceAll(blw.body.String(), `"`, `'`)
	}
	return ""
}

func RestLogger(context *gin.Context) {
	processID := uuid.New().String()
	context.Set("processID", processID)
	path := context.Request.URL.Path
	start := time.Now()
	requestBody := getRequestBody(context)
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
	context.Writer = blw
	context.Next()
	response := getResponseBody(blw, context)
	stop := time.Since(start)
	latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0)) //nolint:gomnd // common
	statusCode := context.Writer.Status()
	clientIP := context.ClientIP()
	clientUserAgent := context.Request.UserAgent()
	method := context.Request.Method
	logDetail := log.Fields{
		"processID":  processID,
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
	if !ExcludingApis[context.Request.URL.Path] {
		logDetail["requestBody"] = requestBody
		logDetail["response"] = response
	}
	logger := log.WithFields(logDetail)
	if len(context.Errors) > 0 {
		logger.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
	} else {
		msg := "[GIN]"
		switch {
		case statusCode >= http.StatusInternalServerError:
			logger.Error(msg)
		case statusCode >= http.StatusBadRequest:
			logger.Warn(msg)
		default:
			logger.Info(msg)
		}
	}
}
