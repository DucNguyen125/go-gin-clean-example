package repository

import (
	"strings"
)

type StorageRepository interface {
	UploadFile(file *strings.Reader, filename string) (path string, err error)
	GenPreSignURL(key string) (string, error)
	GenPreSignURLWithOutTimestamp(key string) (string, error)
}
