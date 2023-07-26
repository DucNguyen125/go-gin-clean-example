package config

import "errors"

type Constants struct {
	DebugMode bool `env:"DEBUG_MODE,default=false"`
	Port      int  `env:"PORT,default=3000"`
	ExportLog bool `env:"EXPORT_LOG,default=true"`
}

var (
	ErrInvalidEnv = errors.New("invalid env")
)
