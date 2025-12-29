package modao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	logg "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"github.com/rrzu/modao/common"
)

type DebugKey string

var debugKey DebugKey

func RegisterDebugKey(key DebugKey) {
	debugKey = key
}

type OptModaoLogger struct {
	OnDebug bool
}

type IModaoLogger interface {
	logg.Interface
	SetOption(opt OptModaoLogger) logg.Interface
}

// NewClickhouseModaoLogger 在 clickhouse.go 文件中添加以下代码
func NewClickhouseModaoLogger(logger *logrus.Logger, config logg.Config) IModaoLogger {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = logg.Green + "%s\n" + logg.Reset + logg.Green + "[info] " + logg.Reset
		warnStr = logg.BlueBold + "%s\n" + logg.Reset + logg.Magenta + "[warn] " + logg.Reset
		errStr = logg.Magenta + "%s\n" + logg.Reset + logg.Red + "[error] " + logg.Reset
		traceStr = logg.Green + "%s\n" + logg.Reset + logg.Yellow + "[%.3fms] " + logg.BlueBold + "[rows:%v]" + logg.Reset + " %s"
		traceWarnStr = logg.Green + "%s " + logg.Yellow + "%s\n" + logg.Reset + logg.RedBold + "[%.3fms] " + logg.Yellow + "[rows:%v]" + logg.Magenta + " %s" + logg.Reset
		traceErrStr = logg.RedBold + "%s " + logg.MagentaBold + "%s\n" + logg.Reset + logg.Yellow + "[%.3fms] " + logg.BlueBold + "[rows:%v]" + logg.Reset + " %s"
	}

	return &DefaultModaoLogger{
		logger:       logger,
		config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

// DefaultModaoLogger 默认 modao 日志
type DefaultModaoLogger struct {
	logger *logrus.Logger // 使用 “系统日志对象” 作为写日志对象
	config logg.Config    // 日志配置

	opt OptModaoLogger // 特定配置

	infoStr, warnStr, errStr            string // 日志格式
	traceStr, traceErrStr, traceWarnStr string // trace 日志格式
}

// SetOption 设置特定配置
func (l *DefaultModaoLogger) SetOption(opt OptModaoLogger) logg.Interface {
	l.opt = opt
	return l
}

// LogMode log mode
func (l *DefaultModaoLogger) LogMode(level logg.LogLevel) logg.Interface {
	newLogger := *l
	newLogger.config.LogLevel = level
	return &newLogger
}

// Info print info
func (l *DefaultModaoLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logg.Info {
		l.logger.Infof(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l *DefaultModaoLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logg.Warn {
		l.logger.Warnf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l *DefaultModaoLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logg.Error {
		l.logger.Errorf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l *DefaultModaoLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.config.LogLevel <= logg.Silent {
		return
	}

	var onDebug = l.opt.OnDebug

	// sql 和 影响行数
	sql, rows := fc()
	rowsStr := common.TernaryAny(rows == -1, "-", cast.ToString(rows))

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.config.LogLevel >= logg.Error && (!errors.Is(err, logg.ErrRecordNotFound) || !l.config.IgnoreRecordNotFoundError):
		if onDebug {
			l.logger.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rowsStr, sql)
		} else {
			l.logger.Tracef(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rowsStr, sql)
		}
	case elapsed > l.config.SlowThreshold && l.config.SlowThreshold != 0 && l.config.LogLevel >= logg.Warn:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.config.SlowThreshold)
		if onDebug {
			l.logger.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rowsStr, sql)
		} else {
			l.logger.Tracef(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rowsStr, sql)
		}
	case l.config.LogLevel == logg.Info:
		if onDebug {
			l.logger.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rowsStr, sql)
		} else {
			l.logger.Tracef(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rowsStr, sql)
		}
	}
}
