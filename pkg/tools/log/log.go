package log

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"
)

var blog *logs.BeeLogger

func init() {
	blog = logs.NewLogger(10000)
	blog.SetLogFuncCallDepth(3)
	blog.EnableFuncCallDepth(true)
	blog.SetLogger("console", "")
	blog.SetLogger("multifile", `{"filename":"logs/app.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
}

// Reset will remove all the adapter
func Reset() {
	blog.Reset()
}

// Async set the beelogger with Async mode and hold msglen messages
func Async() *logs.BeeLogger {
	return blog.Async()
}

// SetLevel sets the global log level used by the simple logger.
func SetLevel(l int) {
	blog.SetLevel(l)
}

// SetLogFuncCall set the CallDepth, default is 4
func SetLogFuncCall(b bool) {
	blog.EnableFuncCallDepth(b)
	blog.SetLogFuncCallDepth(4)
}

// SetLogger sets a new logger.
func SetLogger(adapter string, config string) error {
	return blog.SetLogger(adapter, config)
}

// Emergency logs a message at emergency level.
func Emergency(f interface{}, v ...interface{}) {
	blog.Emergency(formatLog(f, v...))
}

// Alert logs a message at alert level.
func Alert(f interface{}, v ...interface{}) {
	blog.Alert(formatLog(f, v...))
}

// Critical logs a message at critical level.
func Critical(f interface{}, v ...interface{}) {
	blog.Critical(formatLog(f, v...))
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	blog.Error(formatLog(f, v...))
}

// Warning logs a message at warning level.
func Warning(f interface{}, v ...interface{}) {
	blog.Warn(formatLog(f, v...))
}

// Warn compatibility alias for Warning()
func Warn(f interface{}, v ...interface{}) {
	blog.Warn(formatLog(f, v...))
}

// Notice logs a message at notice level.
func Notice(f interface{}, v ...interface{}) {
	blog.Notice(formatLog(f, v...))
}

// Informational logs a message at info level.
func Informational(f interface{}, v ...interface{}) {
	blog.Info(formatLog(f, v...))
}

// Info compatibility alias for Warning()
func Info(f interface{}, v ...interface{}) {
	blog.Info(formatLog(f, v...))
}

// Debug logs a message at debug level.
func Debug(f interface{}, v ...interface{}) {
	blog.Debug(formatLog(f, v...))
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Trace(f interface{}, v ...interface{}) {
	blog.Trace(formatLog(f, v...))
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
