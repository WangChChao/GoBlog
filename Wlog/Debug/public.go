package Wlog

// Debug 记录Debug级的日志
func (l *logger) Debug(s ...interface{}) {
	l.debug(s...)
}

func (l *logger) Trace(s ...interface{}) {
	l.trace(s...)
}

func (l *logger) Info(s ...interface{}) {
	l.info(s...)
}

func (l *logger) Waring(s ...interface{}) {
	l.waring(s...)
}

func (l *logger) Error(s ...interface{}) {
	l.error(s...)
}

func (l *logger) Fatal(s ...interface{}) {
	l.fatal(s...)
}

func (l *logger) SetWriter(path string, isCreate bool) {
	l.setWriter(path, isCreate)
}

func (l *logger) CloseFile() {
	l.closeFile()
}

func NewLogger(levelStr string, maxFileSize int64) *logger {
	return newLogger(levelStr, maxFileSize)
}
