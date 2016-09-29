package recovery

import (
	"runtime/debug"

	"github.com/kataras/iris"
	"github.com/sillydong/goczd/golog"
)

type Recovery struct {
}

func (r *Recovery) Serve(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			golog.Criticalf("Recovery from panic\n%s\n%s\n", err, debug.Stack())
			//ctx.Panic just sends  http status 500 by default, but you can change it by: iris.OnPanic(func( c *iris.Context){})
			ctx.Panic()
		}
	}()
	ctx.Next()
}

// New restores the server on internal server errors (panics)
// receives an optional logger, the default is the Logger with an os.Stderr as its output
// returns the middleware
func New() iris.HandlerFunc {
	/*r := recovery{os.Stderr}
	if out != nil && len(out) == 1 {
		r.out = out[0]
	}*/
	r := &Recovery{}
	return r.Serve
}
