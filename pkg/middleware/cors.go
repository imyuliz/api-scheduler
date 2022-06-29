package frame

import "net/http"

func Cors() HandlerFunc {
	return func(c *Context) {
		c.String(http.StatusOK, "403 NOT FOUND: %s\n", c.Path)
		c.Abort()
	}
}
