package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func InitLogger() error {
	l, err := zap.NewDevelopment()

	if err != nil {
		_ = fmt.Errorf("Cant create zap logger. Reason %v", err)
		return err
	}

	sugar = l.Sugar()

	return nil
}

func ZapLog() *zap.SugaredLogger {
	return sugar
}
