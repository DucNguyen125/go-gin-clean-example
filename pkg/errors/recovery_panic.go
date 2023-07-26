package errors

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func RecoverPanic(function func()) func() {
	return func() {
		defer func() {
			if err, ok := recover().(error); ok {
				if ok {
					log.Error(errors.Wrapf(err, "unexpected panic"))
				} else {
					log.Error(errors.Errorf("unexpected panic: %+v", recover()))
				}
			}
		}()
		function()
	}
}

func RecoverPanicWithError(function func() error) func() error {
	return func() error {
		defer func() {
			if err, ok := recover().(error); ok {
				if err != nil {
					if ok {
						log.Error(errors.Wrapf(err, "unexpected panic"))
					} else {
						log.Error(errors.Errorf("unexpected panic: %+v", recover()))
					}
				}
			}
		}()
		return function()
	}
}
