package http

import (
	"context"
	"net/http"

	"github.com/imroc/req"
	"github.com/imyuliz/api-scheduler/utils/json"
)

func Get(c context.Context, url string, v ...interface{}) (*req.Resp, error) {
	return req.Get(url, v)
}

func GetRespCode(resp *req.Resp) string {
	if resp.Response().StatusCode != http.StatusOK {
		return ""
	}
	if resp.Response().Header.Get("Content-type") == "application/json" ||
		resp.Response().Header.Get("Content-type") == "" {

		return json.GetFieldString(resp.String(), "code")
	}
	return ""
}

func mergeHeader(maps ...http.Header) http.Header {
	if len(maps) == 0 {
		return make(http.Header)
	}
	if len(maps) == 1 {
		return maps[0]
	}
	headers := make(http.Header)
	for _, v := range maps {
		if len(v) == 0 {
			continue
		}
		for k, vv := range v {
			headers[k] = vv
		}
	}
	return headers
}
