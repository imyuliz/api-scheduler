package middleware

import (
	"context"
	"errors"
	"fmt"
	"sync"

	appsv1 "github.com/imyuliz/api-scheduler/api/apps/v1"
	"github.com/imyuliz/api-scheduler/frame"
	v1 "github.com/imyuliz/api-scheduler/pkg/apis/meta/v1"
)

func Flow() frame.HandlerFunc {
	return func(c *frame.Context) {
		namesapce := c.Request.Header.Get("namespace")
		url := c.Request.URL.Path
		options := v1.GetOptions{}
		flow, err := c.Clientset.AppsV1().Flows(namesapce).Get(context.Background(), url, options)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		c.Info("flow config is %+v", flow)

		// TODO: 开始处理逻辑

	}
}

func flowStory(c *frame.Context, flow *appsv1.Flow) error {
	if flow.Spec.Paused {
		return errors.New("URL has been paused")
	}
	// handlers 是一个二维数组,里面并发执行，外面串联执行
	handlers := flow.Spec.Template.Spec.Handlers
	if len(handlers) <= 0 {
		return errors.New("handlers is empty")
	}
	for i := range handlers {
		if len(handlers[i]) <= 0 {
			continue
		}

		plot(c, handlers[i])

	}

	return nil
}

func plot(c *frame.Context, reqs appsv1.Requests) {
	l := len(reqs)
	if l <= 0 {
		return
	}
	if l == 1 {
		// TODO: 发起请求
		return
	}
	// 并发发起请求
	wg := sync.WaitGroup{}
	wg.Add(l)
	for _, v := range reqs {
		go func(r appsv1.Request) {
			// TODO: 发起请求
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func action(c frame.Context, r appsv1.Request) error {
	// TODO: 判断是否符合条件
	// r.Conditions
	if r.HTTPGet != nil {
		// http get request
		if len(r.HTTPGet.Address) <= 0 {
			return errors.New("address is empty")
		}
		if len(r.HTTPGet.HTTPHeaders) > 0 {
			for _, v := range r.HTTPGet.HTTPHeaders {
				if len(v.Value) > 0 {
				}
				val := c.DataStorage.Get(v.Value)
				if val == nil && v.Required {
					return fmt.Errorf("header %s is required, but value is empty", v.Name)
				}
				val.V

			}
		}
		return
	}

	if r.HTTPPost != nil {
		// http post request
		return
	}

	if r.HTTPPut != nil {
		// http put request
		return
	}

	if r.HTTPPatch != nil {
		// http patch request
		return
	}

	if r.HTTPDelete != nil {
		// http delete request
		return
	}

}
