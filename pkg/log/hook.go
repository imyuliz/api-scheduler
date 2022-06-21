package log

import (
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	// set file name
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return frame.Function, path.Base(frame.File)
		},
	})
}

type Hook struct {
	TraceID string
}

func NewHook(trace_id string) *Hook {
	return &Hook{TraceID: trace_id}
}

func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["TraceID"] = h.TraceID
	return nil
}
