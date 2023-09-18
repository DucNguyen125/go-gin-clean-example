package logger

import (
	"context"
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func LogHandler(ctx context.Context, err error) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		processID := ctx.Value("processID")
		log.WithFields(log.Fields{
			"file":      fmt.Sprintf("%s:%d", file, line),
			"processID": processID,
		}).Error(err.Error())
	}
}
