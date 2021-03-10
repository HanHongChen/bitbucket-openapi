package logger

import (
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var NudrDataRepositoryLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	NudrDataRepositoryLog = log.WithFields(logrus.Fields{"component": "OAPI", "category": "NudrDataRepo"})
}

func GetLogger() *logrus.Logger {
	return log
}

func SetLogLevel(level logrus.Level) {
	NudrDataRepositoryLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	NudrDataRepositoryLog.Infoln("set report call :", bool)
	log.SetReportCaller(bool)
}
