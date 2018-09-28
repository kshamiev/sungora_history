// Более сложный пример, с использованием пула обработчиков для типовых задач
package gorun

import (
	"fmt"
	"sync"
)

// Task - описание интрефейса работы
type Task interface {
	Execute()
}

// Pool - структура, нам потребуется Мутекс, для гарантий атомарности изменений самого объекта
// Канал входящих задач
// Канал отмены, для завершения работы
// WaitGroup для контроля завершнеия работ
type Pool struct {
	mu    sync.Mutex
	size  int
	tasks chan Task
	kill  chan struct{}
	wg    sync.WaitGroup
}

// Скроем внутреннее усройство за конструктором, пользователь может влиять только на размер пула
func NewPool(size int) *Pool {
	pool := &Pool{
		// Канал задач - буферизированный, чтобы основная программа не блокировалась при постановке задач
		tasks: make(chan Task, 128),
		// Канал kill для убийства "лишних воркеров"
		kill: make(chan struct{}),
	}
	// Вызовем метод resize, чтобы установить соответствующий размер пула
	pool.Resize(size)
	return pool
}

func (p *Pool) Resize(n int) {
	// Захватывам лок, чтобы избежать одновременного изменения состояния
	p.mu.Lock()
	defer p.mu.Unlock()
	for p.size < n {
		p.size++
		p.wg.Add(1)
		go p.worker()
	}
	for p.size > n {
		p.size--
		p.kill <- struct{}{}
	}
}

// Жизненный цикл воркера
func (p *Pool) worker() {
	defer p.wg.Done()
	for {
		select {
		// Если есть задача, то ее нужно обработать
		// Если написать просто через if, то программа заблокируется и будет ждать, пока канал не закроется
		case task, ok := <-p.tasks:
			if ok {
				task.Execute()
			} else {
				return
			}
			// Если пришел сигнал умирать, выходим
		case <-p.kill:
			return
		}
	}
}

func (p *Pool) Close() {
	close(p.tasks)
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) Exec(task Task) {
	p.tasks <- task
}

type ExampleTask string

func (e ExampleTask) Execute() {
	fmt.Println("executing:", string(e))
}

func SampleWork() {
	pool := NewPool(5)

	pool.Exec(ExampleTask("foo"))
	pool.Exec(ExampleTask("bar"))

	pool.Resize(3)

	pool.Resize(6)

	for i := 0; i < 20; i++ {
		pool.Exec(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
	}

	pool.Close()

	pool.Wait()
}
