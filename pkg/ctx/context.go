package ctx

import (
	"context"

	"github.com/imyuliz/api-scheduler/pkg/pool"
)

type Context interface {
	context.Context
	WithTraceValue(key, value string)
	TraceTrace(key string) string
	pool.DataStorage
}

func Background() Context {
	return &defaultContext{
		Context:     context.Background(),
		DataStorage: pool.NewDataStorage(),
	}
}

type defaultContext struct {
	context.Context
	pool.DataStorage
}

func (dc *defaultContext) WithTraceValue(key, value string) {
	dc.Context = context.WithValue(dc.Context, key, value)
}

func (dc *defaultContext) TraceTrace(key string) string {
	return dc.Context.Value(key)
}
