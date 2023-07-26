package string

import (
	"fmt"
	"strconv"
)

type Service interface {
	ConvertStringToInt(value string) (int, error)
	Sscanf(str string, format string, v ...interface{}) (int, error)
}

type stringService struct{}

func NewStringService() Service {
	return &stringService{}
}

func (s *stringService) ConvertStringToInt(value string) (int, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *stringService) Sscanf(str string, format string, a ...interface{}) (int, error) {
	n, err := fmt.Sscanf(str, format, a...)
	return n, err
}
