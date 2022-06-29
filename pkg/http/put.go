package http

import (
	"context"

	"github.com/imroc/req"
)

func Put(c context.Context, url string, v ...interface{}) (*req.Resp, error) {
	return req.Put(url, v)
}
