package scheduler

import (
	appsv1 "github.com/imyuliz/api-scheduler/scheduler/typed/apps/v1"
	corev1 "github.com/imyuliz/api-scheduler/scheduler/typed/core/v1"
)

type Interface interface {
	CoreV1() corev1.CoreV1Interface
	AppsV1() appsv1.AppsV1Interface
}
