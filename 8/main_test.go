package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	var wg CustomWaitGroup
	//Канал для отслеживания завершения всех горутин
	done := make(chan struct{})

	// Запуск трех горутин
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(100*i))
		}(i)
	}

	// Запускаем ожидание в отдельной горутине
	go func() {
		wg.Wait()
		close(done)

	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Errorf("Test timed out, Wait() took too long")
	}
}
