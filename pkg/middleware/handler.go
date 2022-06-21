package middleware

import (
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

func Handler(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// path := req.Request.URL.Path
	logrus.Println("路由日志")
	// logrus.Println("通过路由PATH,查找流程")
	// logrus.Println("找不到,配置404")
	// logrus.Panicln("找到了,向下执行")
	tid := req.HeaderParameter("trace")
	logrus.Println("找到的header是: " + tid)
	resp.Write([]byte("hello world"))
	logrus.Println("我是可爱")
	// chain.ProcessFilter(req, resp)
}
