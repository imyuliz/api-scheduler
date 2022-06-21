package http

import "github.com/imroc/req"

type Header req.Header

type QueryParam req.QueryParam

type BodyJson struct {
	v string
}
