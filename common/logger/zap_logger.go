package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

// initialize zap logger
func InitZapLog() error {
	l, err := zap.NewDevelopment()

	if err != nil {
		_ = fmt.Errorf("Cant create zap logger. Reason %v", err)
		return err
	}

	sugar = l.Sugar()

	return nil
}

// Through this func. we use the logger over the app
func ZapLog() *zap.SugaredLogger {
	return sugar
}
