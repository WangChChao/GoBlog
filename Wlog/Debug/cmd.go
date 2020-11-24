package Wlog

// isWritLog 判断当前是否需要输出日志
func (l *logger) isWritLog(lv LogLevel) bool {
	if l.level <= lv {
		return true
	}
	return false
}

// debug debug级的日志
func (l *logger) debug(s ...interface{}) {
	if l.isWritLog(DEBUG) {
		writLog("debug", l, s...)
	}
}

// trace Trace级的日志
func (l *logger) trace(s ...interface{}) {
	if l.isWritLog(TRACE) {
		writLog("trace", l, s...)
	}
}

// info Info级的日志
func (l *logger) info(s ...interface{}) {
	if l.isWritLog(INFO) {
		writLog("info", l, s...)
	}
}

// waring Waring级的日志
func (l *logger) waring(s ...interface{}) {
	if l.isWritLog(WARNING) {
		writLog("warning", l, s...)
	}
}

// error Error级的日志
func (l *logger) error(s ...interface{}) {
	if l.isWritLog(ERROR) {
		writLog("error", l, s...)
	}
}

// fatal Fatal级的日志
func (l *logger) fatal(s ...interface{}) {
	if l.isWritLog(FATAL) {
		writLog("fatal", l, s...)
	}
}
