package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger = zap.NewNop()

func Initialize(level string) error {
	l, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	Log = l
	return nil
}
