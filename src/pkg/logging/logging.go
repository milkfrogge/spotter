package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"sync"
)

var logger *logrus.Logger
var once sync.Once

func GetLogger() *logrus.Logger {
	once.Do(Init)
	return logger
}

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fname := path.Base(frame.File)
			return fmt.Sprintf("%s", frame.Function), fmt.Sprintf("%s: %d", fname, frame.Line)
		},
	})
	logger = l
}
