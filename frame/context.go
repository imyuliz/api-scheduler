package frame

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/imyuliz/api-scheduler/frame/log"
	"github.com/imyuliz/api-scheduler/frame/pool"
	"github.com/imyuliz/api-scheduler/frame/trace"
	"github.com/imyuliz/api-scheduler/scheduler"
)

type Context struct {
	Writer      http.ResponseWriter // http writer
	Request     *http.Request       // http request
	Method      string              // request method
	Path        string              // request path
	HTTPCode    int                 //response code
	handlers    HandlersChain       // router chain
	index       int8                // chain index
	DataStorage pool.DataStorage    // request data storage
	Trace       trace.Trace         // trace id record
	log.Log                         // log config
	Clientset   scheduler.Interface // resoure clientset

}

// abortIndex represents a typical value used in abort functions.
const abortIndex int8 = math.MaxInt8 >> 1 // 63

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	l := log.NewLogWithRequest(r)
	// TODO: trace key config
	// TODO: clientset init
	traceKey := log.TraceKey
	return &Context{
		Writer:  w,
		Request: r,
		Method:  r.Method,
		Path:    r.URL.Path,
		index:   -1,
		Log:     l,
		Trace:   trace.NewTrace(traceKey, l.GetTraceID()),
	}
}

func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) String(code int, message string, v ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Write([]byte(fmt.Sprintf(message, v...)))
}

func (c *Context) Write(data []byte) {
	c.Writer.Write(data)
}

func (c *Context) Status(code int) {
	c.HTTPCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	data, err := json.Marshal(obj)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
	c.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Write([]byte(html))
}
