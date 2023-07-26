package logger

import (
	"base-gin-golang/config"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func Init(cfg *config.Environment) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableQuote:    true,
		DisableColors:   true,
	})
	debugLogFile, err := os.OpenFile("logs/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Error(err)
	}
	errorLogFile, err := os.OpenFile("logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Error(err)
	}
	if cfg.ExportLog {
		log.SetOutput(io.Discard)
		log.SetFormatter(&log.JSONFormatter{})
		log.AddHook(&writer.Hook{
			Writer: errorLogFile,
			LogLevels: []log.Level{
				log.PanicLevel,
				log.FatalLevel,
				log.ErrorLevel,
			},
		})
		log.AddHook(&writer.Hook{
			Writer: debugLogFile,
			LogLevels: []log.Level{
				log.PanicLevel,
				log.FatalLevel,
				log.ErrorLevel,
				log.WarnLevel,
				log.InfoLevel,
				log.DebugLevel,
			},
		})
	}
}
