package middleware

import (
	"time"

	"github.com/imyuliz/api-scheduler/frame"
)

func Logger() frame.HandlerFunc {
	return func(c *frame.Context) {
		t := time.Now()
		c.Next()
		c.Infof("[%d] %s in %v", c.HTTPCode, c.Request.RequestURI, time.Since(t))
	}
}
