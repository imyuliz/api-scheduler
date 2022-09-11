package middleware

import (
	"net/http"

	"github.com/imyuliz/api-scheduler/frame"
)

func Cors() frame.HandlerFunc {
	return func(c *frame.Context) {
		c.String(http.StatusOK, "403 NOT FOUND: %s\n", c.Path)
		c.Abort()
	}
}
