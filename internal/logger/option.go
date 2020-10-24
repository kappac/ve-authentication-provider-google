package logger

// OptionUpdater updates Error properties.
type OptionUpdater func(*logger)

// WithEntity sets entity name to be displayed in a log string
func WithEntity(e string) OptionUpdater {
	return func(l *logger) {
		l.entityName = e
	}
}

// WithLogger sets parent logger to be used by new instance
func WithLogger(pl Logger) OptionUpdater {
	return func(l *logger) {
		l.parentLogger = pl.getLogger()
	}
}
