package log

import (
	"net/http"
	"strings"

	"github.com/imyuliz/api-scheduler/frame/uuid"
	"github.com/sirupsen/logrus"
)

type Log interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	GetTraceID() string
}

type log struct {
	traceID string
	*logrus.Logger
}

func (l *log) GetTraceID() string {
	return l.traceID
}

func NewLogWithRequest(r *http.Request) Log {
	tid := r.Header.Get(TraceKey)
	if strings.TrimSpace(tid) == "" {
		tid = uuid.NewUUID()
	}
	logger := logrus.New()
	logger.AddHook(NewHook(tid))
	logger.SetReportCaller(false)
	logger.SetFormatter(defaultLogFormatter)
	return &log{traceID: tid, Logger: logger}
}

func NewLog() Log {
	tid := uuid.NewUUID()
	logger := logrus.New()
	logger.AddHook(NewHook(tid))
	logger.SetReportCaller(true)
	logger.SetFormatter(defaultLogFormatter)
	return &log{Logger: logger}
}
