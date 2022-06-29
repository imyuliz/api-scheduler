package http

import (
	"context"

	"github.com/imroc/req"
)

func Post(c context.Context, url string, v ...interface{}) (*req.Resp, error) {
	return req.Post(url, v)
}
