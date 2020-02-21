package app

import (
	"context"
	"sync"
	"time"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/pkg/logger"
)

type Task interface {
	// информация о задаче
	Name() string
	// выполняемая задача
	Action(ctx context.Context) (err error)
	// время до следующего запуска
	WaitFor() time.Duration
}

type Scheduler struct {
	pullWork map[string]bool
	pull     []Task
	lg       logger.Logger
	wg       sync.WaitGroup // для контроля завершния работы
	kill     chan string    // канал для убийства обработчиков
}

// NewScheduler создание шины воркеров
func NewScheduler(lg logger.Logger) *Scheduler {
	return &Scheduler{
		lg:       lg,
		kill:     make(chan string, 100),
		pullWork: make(map[string]bool),
	}
}

// добавить задачу в Scheduler
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
			wf.pullWork[name] = true

			go wf.run(wf.pull[i])

			return
		}
	}
}

// Stop остановить конкретную задачу
func (wf *Scheduler) Stop(name string) {
	if _, ok := wf.pullWork[name]; !ok {
		return
	}

	delete(wf.pullWork, name)
	wf.kill <- name
}

// Wait остановить все выполняющиеся задачи
func (wf *Scheduler) Wait() {
	wf.kill <- ""
	wf.wg.Wait()
	close(wf.kill)
}

// GetTask Получение всех задач
func (wf *Scheduler) GetTask() map[string]Task {
	res := make(map[string]Task)
	for i := range wf.pull {
		res[wf.pull[i].Name()] = wf.pull[i]
	}

	return res
}

// run менеджер выполенния задачи
func (wf *Scheduler) run(task Task) {
	wf.wg.Add(1)
	defer wf.wg.Done()

	for {
		waitFor := task.WaitFor()
		select {
		case <-time.After(waitFor):
			wf.action(task)
		case name := <-wf.kill:
			switch name {
			case "": // завершаем все воркеры
				wf.kill <- ""
				return
			case task.Name(): // завершаем текущий воркер
				return
			default: // ложное срабатывание, перенаправляем для заврешения целевого воркера
				for i := range wf.pull {
					if wf.pull[i].Name() == name {
						wf.kill <- name
					}
				}
			}
		}
	}
}

// action выполнение задачи
func (wf *Scheduler) action(task Task) {
	lg := wf.lg.WithField("task", task.Name())

	ctx := context.Background()
	ctx = boil.WithDebugWriter(ctx, lg.Writer())
	ctx = logger.WithLogger(ctx, lg)

	defer func() {
		if rvr := recover(); rvr != nil {
			lg.Errorf("%+v", rvr)
		}
	}()

	if err := task.Action(ctx); err != nil {
		lg.Errorf("%+v", err)
	}
}
