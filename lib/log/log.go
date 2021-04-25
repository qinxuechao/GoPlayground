package log

import (
	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: false, FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05.000"})
}
