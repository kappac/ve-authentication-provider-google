package logger

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var (
	defaultLogger log.Logger
)

func init() {
	defaultLogger = log.NewLogfmtLogger(
		log.NewSyncWriter(os.Stdout),
	)
	defaultLogger = level.NewFilter(defaultLogger, level.AllowInfo(), level.AllowWarn(), level.AllowError(), level.AllowDebug())
	defaultLogger = log.With(defaultLogger, "ts", log.TimestampFormat(time.Now, "15:04:05.9999"))
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

	l.logger = log.With(l.parentLogger, "entity", l.entityName)

	return l
}

func (l *logger) Info(args ...interface{}) error {
	return level.Info(l.logger).Log(args...)
}

func (l *logger) Infom(message string, args ...interface{}) error {
	mp := l.buildMessagePair(message)
	nArgs := append(mp, args...)
	return l.Info(nArgs...)
}

func (l *logger) Warning(args ...interface{}) error {
	return level.Warn(l.logger).Log(args...)
}

func (l *logger) Warningm(message string, args ...interface{}) error {
	mp := l.buildMessagePair(message)
	nArgs := append(mp, args...)
	return l.Warning(nArgs...)
}

func (l *logger) Error(args ...interface{}) error {
	return level.Error(l.logger).Log(args...)
}

func (l *logger) Errorm(message string, args ...interface{}) error {
	mp := l.buildMessagePair(message)
	nArgs := append(mp, args...)
	return l.Error(nArgs...)
}

func (l *logger) Debug(args ...interface{}) error {
	return level.Debug(l.logger).Log(args...)
}

func (l *logger) Debugm(message string, args ...interface{}) error {
	mp := l.buildMessagePair(message)
	nArgs := append(mp, args...)
	return l.Debug(nArgs...)
}

func (l *logger) getLogger() log.Logger {
	return l.logger
}

func (l *logger) buildMessagePair(message string) []interface{} {
	return []interface{}{"msg", message}
}
