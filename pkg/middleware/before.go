package middleware

import (
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/imyuliz/api-scheduler/pkg/ctx"
	"github.com/imyuliz/api-scheduler/pkg/log"
	"github.com/imyuliz/api-scheduler/pkg/uuid"
	"github.com/sirupsen/logrus"
)

type fakeConfig int

var TraceKey = "traceid"

func Log(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// get trace header
	tid := req.HeaderParameter(TraceKey)
	if strings.TrimSpace(tid) == "" {
		tid = uuid.NewUUID()
	}
	// set log trace
	logrus.AddHook(log.NewHook(tid))
	// set header
	req.Request.Header.Add(TraceKey, tid)
	resp.Header().Set(TraceKey, tid)

	// 通过请求的 URL 和 Method 找到对应的配置
	// path := req.Request.URL.Path // method := req.Request.Method
	// fmt.Println(path)
	// fmt.Println(method)

	c := ctx.Background()
	c.WithTraceValue(TraceKey, tid)

	logrus.Printf("[Before]HTTP request sent to %s from %s", req.Request.URL.Path, req.Request.RemoteAddr)
	// call registered handler
	chain.ProcessFilter(req, resp)
	logrus.Printf("[After]HTTP request sent to %s from %s", req.Request.URL.Path, req.Request.RemoteAddr)
}
