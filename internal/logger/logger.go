package logger

import (
	stdlog "log"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kappac/ve-authentication-provider-google/internal/veconfig"
)

const (
	entityKey  = "entity"
	messageKey = "msg"
)

var (
	defaultLogger log.Logger
)

func init() {
	levelOption := level.AllowInfo()

	if veconfig.Global.Debug {
		levelOption = level.AllowDebug()
	}

	defaultLogger = log.NewLogfmtLogger(
		log.NewSyncWriter(os.Stderr),
	)
	defaultLogger = level.NewFilter(defaultLogger, levelOption)
	defaultLogger = log.With(defaultLogger, "ts", log.TimestampFormat(time.Now, "15:04:05.9999"))

	stdlog.SetFlags(0)
	stdlog.SetOutput(log.NewStdlibAdapter(defaultLogger))
}

// Logger describes methods available to log events.
type Logger interface {
	Info(args ...interface{}) error
	Infom(message string, args ...interface{}) error
	Warning(args ...interface{}) error
	Warningm(message string, args ...interface{}) error
	Error(args ...interface{}) error
	Errorm(message string, args ...interface{}) error
	Debug(args ...interface{}) error
	Debugm(message string, args ...interface{}) error
	getLogger() log.Logger
}

type logger struct {
	logger       log.Logger
	parentLogger log.Logger
	entityName   string
}

// New creates logger instance
func New(ous ...OptionUpdater) Logger {
	l := &logger{
		entityName:   "N/A",
		parentLogger: defaultLogger,
	}

	for _, ou := range ous {
		ou(l)
	}

	l.logger = log.With(l.parentLogger, entityKey, l.entityName)

	return l
}

func (l *logger) Info(args ...interface{}) error {
	return level.Info(l.logger).Log(args...)
}

func (l *logger) Infom(message string, args ...interface{}) error {
	return l.Info(
		l.buildArgs(message, args)...,
	)
}

func (l *logger) Warning(args ...interface{}) error {
	return level.Warn(l.logger).Log(args...)
}

func (l *logger) Warningm(message string, args ...interface{}) error {
	return l.Warning(
		l.buildArgs(message, args)...,
	)
}

func (l *logger) Error(args ...interface{}) error {
	return level.Error(l.logger).Log(args...)
}

func (l *logger) Errorm(message string, args ...interface{}) error {
	return l.Error(
		l.buildArgs(message, args)...,
	)
}

func (l *logger) Debug(args ...interface{}) error {
	return level.Debug(l.logger).Log(args...)
}

func (l *logger) Debugm(message string, args ...interface{}) error {
	return l.Debug(
		l.buildArgs(message, args)...,
	)
}

func (l *logger) getLogger() log.Logger {
	return l.logger
}

func (l *logger) buildArgs(message string, args []interface{}) []interface{} {
	mp := []interface{}{messageKey, message}
	return append(mp, args...)
}
