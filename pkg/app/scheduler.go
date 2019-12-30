package app

import (
	"context"
	"sync"
	"time"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/pkg/logger"
)

type SchedulerTask interface {
	// информация о задаче
	Info() (name string, t time.Duration)
	// выполняемая задача
	Action(ctx context.Context) (err error)
}

type Scheduler struct {
	pull []SchedulerTask
	lg   logger.Logger
	wg   sync.WaitGroup // для контроля завершния работы
	kill chan bool      // канал для убийства обработчиков
}

func NewScheduler(lg logger.Logger) *Scheduler {
	return &Scheduler{
		lg:   lg,
		kill: make(chan bool, 1),
	}
}

// добавить задачу в Scheduler
func (wf *Scheduler) Add(w SchedulerTask) {
	wf.pull = append(wf.pull, w)
}

// запустить все добавленные задачи
func (wf *Scheduler) Run() {
	for i := range wf.pull {
		go wf.run(wf.pull[i])
	}
}

// остановить все выполняющиеся задачи
func (wf *Scheduler) Wait() {
	wf.kill <- true
	wf.wg.Wait()
	close(wf.kill)
}

func (wf *Scheduler) run(task SchedulerTask) {
	wf.wg.Add(1)
	defer wf.wg.Done()

	_, t := task.Info()
	ticker := time.NewTicker(t)

	for {
		select {
		case <-ticker.C:
			wf.action(task)
		case <-wf.kill:
			wf.kill <- true
			return
		}
	}
}

func (wf *Scheduler) action(task SchedulerTask) {
	workerName, _ := task.Info()

	lg := wf.lg.WithField("task", workerName)

	ctx := context.Background()
	ctx = boil.WithDebugWriter(ctx, lg.Writer())
	ctx = logger.WithLogger(ctx, lg)

	defer func() {
		if rvr := recover(); rvr != nil {
			wf.lg.Errorf("task: %s %+v", workerName, rvr)
		}
	}()

	if err := task.Action(ctx); err != nil {
		wf.lg.Errorf("task: %s %+v", workerName, err)
	}
}
