package frame

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Context struct {
	//write and request
	Writer  http.ResponseWriter
	Request *http.Request
	//request info
	Method string
	Path   string
	//response
	HttpCode int
	// router chain
	handlers HandlersChain
	index    int
}

// abortIndex represents a typical value used in abort functions.
const abortIndex int8 = math.MaxInt8 >> 1 // 63

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Method:  r.Method,
		Path:    r.URL.Path,
		index:   -1,
	}
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
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
	c.HttpCode = code
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
