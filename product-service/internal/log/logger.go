package log

import "github.com/hailsayan/sophocles/pkg/logger"

var Logger logger.Logger

func SetLogger(logger logger.Logger) {
	Logger = logger
}
