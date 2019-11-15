package logger

import (
	"os"

	"github.com/Dimitriy14/golang-restik/src/config"
	"github.com/sirupsen/logrus"
)

var (
	Log Logger
)

// Logger - represents methods for logging
type Logger interface {
	Info(txID string, v ...interface{})
	Infof(txID string, format string, v ...interface{})

	Error(txID string, v ...interface{})
	Errorf(txID string, format string, v ...interface{})

	Warn(txID string, v ...interface{})
	Warnf(txID string, format string, v ...interface{})

	Debug(txID string, v ...interface{})
	Debugf(txID string, format string, v ...interface{})
}

// Load loads logger
func Load() error {
	output := os.Stdout
	if config.Conf.UseLogFile {
		logFile, err := os.Create(config.Conf.LogFile)
		if err != nil {
			return err
		}
		output = logFile
	}

	logLvl, err := logrus.ParseLevel(config.Conf.LogLvl)
	if err != nil {
		return err
	}

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetOutput(output)
	log.SetLevel(logLvl)

	Log = &loggerImpl{log: log}

	return nil
}

type loggerImpl struct {
	log *logrus.Logger
}

func (l *loggerImpl) Info(txID string, v ...interface{}) {
	l.log.WithFields(logrus.Fields{"txID": txID}).Info(v)
}

func (l *loggerImpl) Infof(txID string, format string, v ...interface{}) {
	l.log.WithFields(logrus.Fields{"txID": txID}).Infof(format, v)
}

func (l *loggerImpl) Error(txID string, v ...interface{}) {
	l.log.WithFields(logrus.Fields{"txID": txID}).Error(v)
}

func (l *loggerImpl) Errorf(txID string, format string, v ...interface{}) {
	l.log.WithFields(logrus.Fields{"txID": txID}).Errorf(format, v)
}

func (l *loggerImpl) Warn(txID string, v ...interface{}) {
	l.log.WithFields(logrus.Fields{"txID": txID}).Warning(v)
}

func (l *loggerImpl) Warnf(txID, format string, v ...interface{}) {
	l.log.WithFields(logrus.Fields{"txID": txID}).Warningf(format, v)
}

func (l *loggerImpl) Debug(txID string, v ...interface{}) {
	l.log.WithFields(logrus.Fields{"txID": txID}).Debug(v)
}

func (l *loggerImpl) Debugf(txID, format string, v ...interface{}) {
	l.log.WithFields(logrus.Fields{"txID": txID}).Debugf(format, v)
}
