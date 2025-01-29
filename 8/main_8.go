package main

import (
	"sync"
)

// CustomWaitGroup реализует функциональность sync.Waitgroup
type CustomWaitGroup struct {
	counter int
	mu      sync.Mutex
	cond    *sync.Cond
}

func NewCustomWaitgroup() *CustomWaitGroup {
	cwg := &CustomWaitGroup{}
	cwg.cond = sync.NewCond(&cwg.mu)
	return cwg
}

// Увеличиваем счетчик CWG на delta
func (cwg *CustomWaitGroup) Add(delta int) {
	cwg.mu.Lock()
	defer cwg.mu.Unlock()
	cwg.counter += delta

	// Если Add вызывается <0 подаем сигнал ожидателям
	if cwg.counter <= 0 {
		cwg.cond.Broadcast()
	}
}

// Done уменьшает счетчик на 1
func (cwg *CustomWaitGroup) Done() {
	cwg.Add(-1)
}

// Wait блокирует текущую горутину пока счетчтик не станет 0
func (cwg *CustomWaitGroup) Wait() {
	cwg.mu.Lock()
	defer cwg.mu.Unlock()
	for cwg.counter > 0 {
		cwg.cond.Wait()
	}
}
