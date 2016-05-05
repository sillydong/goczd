package gorpc

import (
	"github.com/hprose/hprose-go"
	"github.com/sillydong/goczd/golog"
	"net/http"
	"reflect"
)

//response
type response struct {
	Status bool
	Data   interface{}
	Error  string
}

func (r *response) run() map[string]interface{} {
	return map[string]interface{}{
		"status": r.Status,
		"data":   r.Data,
		"error":  r.Error,
	}
}

func NewResponse(status bool, data interface{}, err string) map[string]interface{} {
	return (&response{Status: status, Data: data, Error: err}).run()
}

//event
type event struct{}

func (e *event) OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context) {
}

func (e *event) OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context) {
}
func (e *event) OnSendError(err error, context hprose.Context) {
	golog.Errorf("%+v", err)
	golog.Errorf("%+v", context)
}

//init rpc
func InitRpc(listentcp, listenhttp, listensocket string, debug bool, methods *hprose.Methods) {
	if listensocket != "" {
		golog.Info("listening on unix:" + listensocket)
		service := hprose.NewUnixServer("unix:" + listensocket)
		service.DebugEnabled = debug
		service.ServiceEvent = &event{}
		service.Methods = methods
		service.Start()
	} else if listentcp != "" {
		golog.Info("listening on tcp://" + listentcp)
		service := hprose.NewTcpServer("tcp://" + listentcp)
		service.DebugEnabled = debug
		service.ServiceEvent = &event{}
		service.Methods = methods
		service.Start()
	} else if listenhttp != "" {
		golog.Info("listening on http://" + listenhttp)
		service := hprose.NewHttpService()
		service.DebugEnabled = debug
		service.ServiceEvent = &event{}
		service.Methods = methods
		http.ListenAndServe(listenhttp, service)
	} else {
		panic("missing listen configuration")
	}
}
