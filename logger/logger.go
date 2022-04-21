package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
)

func New() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf("%s()>", funcName), fmt.Sprintf(" %s:%d", filename, f.Line)
		},
	}
	return logger
}
