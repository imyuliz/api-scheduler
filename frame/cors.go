package frame

import "fmt"

func Cors() HandlerFunc {
	return func(c *Context) {
		fmt.Println("hello world")
		// c.Next()
		fmt.Println("hello")
	}
}
