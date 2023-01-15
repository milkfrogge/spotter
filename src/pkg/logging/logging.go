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
	once.Do(Init) // инициализация лишь один раз (синглтон)
	return logger
}

func Init() {
	// инициализация логгера
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fname := path.Base(frame.File) // путь к файлу, в котором произошла ошибка
			return fmt.Sprintf("%s", frame.Function), fmt.Sprintf("%s: %d", fname, frame.Line)
		},
	})
	logger = l
}
