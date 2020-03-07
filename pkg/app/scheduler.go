package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Task interface {
	Name() string                           // информация о задаче
	Action(ctx context.Context) (err error) // выполняемая задача
	WaitFor() time.Duration                 // время до следующего запуска
}

type Scheduler struct {
	pullWork map[string]chan bool // выполняемые в данный момент задачи
	pull     []Task               // пулл всех задач в программе
	lg       logger.Logger        // логер
}

// NewScheduler создание планировщика задач
func NewScheduler(lg logger.Logger) *Scheduler {
	return &Scheduler{
		lg:       lg,
		pullWork: make(map[string]chan bool),
	}
}

// AddStart see Add, Start
func (wf *Scheduler) AddStart(w Task) {
	wf.Add(w)
	wf.Start(w.Name())
}

// Add добавить задачу в Scheduler
func (wf *Scheduler) Add(w Task) {
	wf.pull = append(wf.pull, w)
}

// Start запустить конкретную задачу
func (wf *Scheduler) Start(name string) {
	if _, ok := wf.pullWork[name]; ok {
		return
	}
	for i := range wf.pull {
		if wf.pull[i].Name() == name {
			wf.pullWork[name] = make(chan bool)
			go wf.run(wf.pull[i], wf.pullWork[name])
			return
		}
	}
}

// Stop остановить конкретную задачу
func (wf *Scheduler) Stop(name string) {
	if _, ok := wf.pullWork[name]; !ok {
		return
	}
	wf.pullWork[name] <- true
	<-wf.pullWork[name]
	delete(wf.pullWork, name)
}

// Wait остановить все выполняющиеся задачи
func (wf *Scheduler) Wait() {
	for k := range wf.pullWork {
		wf.pullWork[k] <- true
	}
	for k := range wf.pullWork {
		<-wf.pullWork[k]
		delete(wf.pullWork, k)
	}
}

// GetTasks Получение всех задач
func (wf *Scheduler) GetTasks() map[string]Task {
	res := make(map[string]Task)
	for i := range wf.pull {
		res[wf.pull[i].Name()] = wf.pull[i]
	}
	return res
}

// run менеджер выполенния задачи
func (wf *Scheduler) run(task Task, ch chan bool) {
	for {
		waitFor := task.WaitFor()
		select {
		case <-time.After(waitFor):
			wf.action(task)
		case <-ch:
			ch <- true
			return
		}
	}
}

// action выполнение задачи
func (wf *Scheduler) action(task Task) {
	requestID := uuid.New().String()
	lg := wf.lg.WithField(response.LogUUID, requestID).WithField(response.LogAPI, task.Name())

	ctx := context.Background()
	ctx = context.WithValue(ctx, response.CtxUUID, requestID)
	ctx = logger.WithLogger(ctx, lg)
	ctx = boil.WithDebugWriter(ctx, lg.Writer())

	defer func() {
		if rvr := recover(); rvr != nil {
			lg.Errorf("panic: %+v", rvr)
		}
	}()

	if err := task.Action(ctx); err != nil {
		lg.Errorf("%+v", err)
	}
}
