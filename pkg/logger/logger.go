package logger

import (
	"github.com/seed95/product-service/pkg/logger/keyval"
)

type Logger interface {
	Debug(message string, keyAndValues ...keyval.Pair)
	Info(message string, keyAndValues ...keyval.Pair)
	Warn(message string, keyAndValues ...keyval.Pair)
	Error(message string, keyAndValues ...keyval.Pair)
	Panic(message string, keyAndValues ...keyval.Pair)
}

func LogReqRes(logger Logger, message string, err error, commonKeyVal ...keyval.Pair) {
	if err != nil {
		commonKeyVal = append(commonKeyVal, keyval.Error(err))
		logger.Error(message, commonKeyVal...)
	} else {
		logger.Info(message, commonKeyVal...)
	}
}
