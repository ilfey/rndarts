package worker

import (
	"main/pkg/core"

	"github.com/robfig/cron/v3"
)

type Worker struct {
	core *core.Core
	cron *cron.Cron
}

func NewWorker(core *core.Core) *Worker {
	return &Worker{
		core: core,
		cron: cron.New(),
	}
}

// Add task
func (w *Worker) AddTask(tasks ...*Task) ([]cron.EntryID, error) {
	ids := make([]cron.EntryID, len(tasks))

	for _, task := range tasks {
		id, err := w.cron.AddFunc(task.Spec, task.Exec)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Remove task
func (w *Worker) RemoveTask(id cron.EntryID) {
	w.cron.Remove(id)
}

// Start worker
func (w *Worker) Start() {
	w.cron.Start()
}

// Stop worker
func (w *Worker) Stop() {
	w.cron.Stop()
}
