package http

import (
	"github.com/imroc/req"
	"github.com/imyuliz/api-scheduler/pkg/ctx"
)

func Put(c ctx.Context, url string, v ...interface{}) (*req.Resp, error) {
	return req.Put(url, v)
}
