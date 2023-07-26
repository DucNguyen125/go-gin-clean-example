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

var ExcludingApisForLog = map[string]bool{
	"/api/v1/auth/login": true,
}

func getRequestBody(ctx *gin.Context) string {
	bodyBuffer := &bytes.Buffer{}
	if strings.Contains(ctx.Request.Header.Get("Content-Type"), "application/json") {
		data, _ := io.ReadAll(ctx.Request.Body)
		_ = json.Compact(bodyBuffer, data)
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // Write body back
	}
	return strings.ReplaceAll(bodyBuffer.String(), `"`, `'`)
}

func getResponseBody(blw *bodyLogWriter, ctx *gin.Context) string {
	if strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/json") {
		return strings.ReplaceAll(blw.body.String(), `"`, `'`)
	}
	return ""
}

func (m *middleware) RestLogger(ctx *gin.Context) {
	processID := uuid.New().String()
	ctx.Set("processID", processID)
	path := ctx.Request.URL.Path
	start := time.Now()
	requestBody := getRequestBody(ctx)
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw
	ctx.Next()
	response := getResponseBody(blw, ctx)
	stop := time.Since(start)
	latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0)) //nolint:gomnd // common
	statusCode := ctx.Writer.Status()
	clientIP := ctx.ClientIP()
	clientUserAgent := ctx.Request.UserAgent()
	method := ctx.Request.Method
	logDetail := log.Fields{
		"processID":  processID,
		"statusCode": statusCode,
		"latency":    fmt.Sprintf("%dms", latency), // time to process
		"clientIP":   clientIP,
		"method":     method,
		"path":       path,
		"userAgent":  clientUserAgent,
	}
	if query := ctx.Request.URL.RawQuery; query != "" {
		logDetail["query"] = query
	}
	if !ExcludingApisForLog[ctx.Request.URL.Path] {
		logDetail["requestBody"] = requestBody
		logDetail["response"] = response
	}
	logger := log.WithFields(logDetail)
	if len(ctx.Errors) > 0 {
		logger.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
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
