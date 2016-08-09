package logger

import (
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/sillydong/goczd/golog"
)

type LoggerMiddleware struct {
}

// Serve serves the middleware
func (l *LoggerMiddleware) Serve(ctx *iris.Context) {
	//all except latency to string
	var date, status, ip, method, path string
	var latency time.Duration
	var startTime, endTime time.Time

	status = strconv.Itoa(ctx.Response.StatusCode())
	ip = ctx.RemoteAddr()
	path = ctx.PathString()
	method = ctx.MethodString()

	startTime = time.Now()
	ctx.Next()
	//no time.Since in order to format it well after
	endTime = time.Now()
	date = endTime.Format("01/02 - 15:04:05")
	latency = endTime.Sub(startTime)

	//finally print the logs
	l.printf("%s %v %4v %s %s %s", date, status, latency, ip, method, path)
}

func (l *LoggerMiddleware) printf(format string, a ...interface{}) {
	golog.Infof(format, a...)
}

// New returns the logger middleware
// receives two parameters, both of them optionals
// first is the logger, which normally you set to the 'iris.Logger'
// if logger is nil then the middlewares makes one with the default configs.
// second is optional configs(logger.Config)
func New() iris.HandlerFunc {
	l := &LoggerMiddleware{}

	return l.Serve
}
