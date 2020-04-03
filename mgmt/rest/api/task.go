package api

import (
	"github.com/micruzz82/snap/core"
	"github.com/micruzz82/snap/core/serror"
	"github.com/micruzz82/snap/pkg/schedule"
	"github.com/micruzz82/snap/scheduler/wmap"
)

type Tasks interface {
	CreateTask(schedule.Schedule, *wmap.WorkflowMap, bool, ...core.TaskOption) (core.Task, core.TaskErrors)
	GetTasks() map[string]core.Task
	GetTask(string) (core.Task, error)
	StartTask(string) []serror.SnapError
	StopTask(string) []serror.SnapError
	RemoveTask(string) error
	WatchTask(string, core.TaskWatcherHandler) (core.TaskWatcherCloser, error)
	EnableTask(string) (core.Task, error)
}
