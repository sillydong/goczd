package goxorm
import (
	"github.com/go-xorm/core"
	"github.com/donnie4w/go-logger/logger"
	"fmt"
)

type XormLogger struct {
	level core.LogLevel
}

func NewXormLogger(logdir, filename string, l core.LogLevel) *XormLogger {
	logger.SetRollingDaily(logdir, filename)
	return &XormLogger{
		level:l,
	}
}

func (s *XormLogger) Err(v ...interface{}) (err error) {
	if s.level > core.LOG_OFF && s.level <= core.LOG_ERR {
		logger.Error(v...)
	}
	return
}

func (s *XormLogger) Errf(format string, v ...interface{}) (err error) {
	if s.level > core.LOG_OFF && s.level <= core.LOG_ERR {
		logger.Error(fmt.Sprintf(format, v))
	}
	return
}

func (s *XormLogger) Debug(v ...interface{}) (err error) {
	if s.level > core.LOG_OFF && s.level <= core.LOG_DEBUG {
		logger.Debug(v...)
	}
	return
}

func (s *XormLogger) Debugf(format string, v ...interface{}) (err error) {
	if s.level > core.LOG_OFF && s.level >= core.LOG_DEBUG {
		logger.Debug(fmt.Sprintf(format, v))
	}
	return
}

func (s *XormLogger) Info(v ...interface{}) (err error) {
	if s.level > core.LOG_OFF && s.level >= core.LOG_INFO {
		logger.Info(v...)
	}
	return
}

func (s *XormLogger) Infof(format string, v ...interface{}) (err error) {
	if s.level > core.LOG_OFF && s.level >= core.LOG_INFO {
		logger.Info(fmt.Sprintf(format, v))
	}
	return
}

func (s *XormLogger) Warning(v ...interface{}) (err error) {
	if s.level > core.LOG_OFF && s.level >= core.LOG_WARNING {
		logger.Warn(v...)
	}
	return
}

func (s *XormLogger) Warningf(format string, v ...interface{}) (err error) {
	if s.level > core.LOG_OFF && s.level >= core.LOG_WARNING {
		logger.Warn(fmt.Sprintf(format, v))
	}
	return
}

func (s *XormLogger) Level() core.LogLevel {
	return s.level
}

func (s *XormLogger) SetLevel(l core.LogLevel) (err error) {
	s.level = l
	return
}
