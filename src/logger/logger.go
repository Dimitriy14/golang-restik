package logger

import (
	"fmt"
	"os"

	"github.com/Dimitriy14/golang-restik/src/config"
	"github.com/google/logger"
)

var (
	Log Logger
	//Postgres logger
	PL PostgresLogger
)

// Logger - represents methods for logging
type Logger interface {
	Info(txID string, v ...interface{})
	Infof(txID string, format string, v ...interface{})

	Error(txID string, v ...interface{})
	Errorf(txID string, format string, v ...interface{})

	Warn(txID string, v ...interface{})
	Warnf(txID string, format string, v ...interface{})
}

type PostgresLogger interface {
	Print(v ...interface{})
}

// Load loads logger
func Load() error {
	lf, err := os.Create(config.Conf.LogFile)
	if err != nil {
		return err
	}

	l := logger.Init("Restik", true, true, lf)

	Log = &loggerImpl{log: l}
	PL = &postgresLoggerImpl{log: l}
	return nil
}

type loggerImpl struct {
	log *logger.Logger
}

func (l *loggerImpl) Info(txID string, v ...interface{}) {
	l.log.Info(txID, v)
}

func (l *loggerImpl) Infof(txID string, format string, v ...interface{}) {
	l.log.Infof(fmt.Sprintf("%s\t%s", txID, format), v)
}

func (l *loggerImpl) Error(txID string, v ...interface{}) {
	l.log.Error(txID, v)
}

func (l *loggerImpl) Errorf(txID string, format string, v ...interface{}) {
	l.log.Errorf(fmt.Sprintf("%s\t%s", txID, format), v)
}

func (l *loggerImpl) Warn(txID string, v ...interface{}) {
	l.log.Warning(txID, v)
}

func (l *loggerImpl) Warnf(txID, format string, v ...interface{}) {
	l.log.Warningf(fmt.Sprintf("%s\t%s", txID, format), v)
}

type postgresLoggerImpl struct {
	log *logger.Logger
}

func (l *postgresLoggerImpl) Print(v ...interface{}) {
	l.log.Info(v)
}
