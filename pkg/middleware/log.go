package frame

import (
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		c.Infof("[%d] %s in %v", c.HTTPCode, c.Request.RequestURI, time.Since(t))
	}
}
