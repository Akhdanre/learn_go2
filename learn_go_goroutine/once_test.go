package learngogoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func CounterValue() {
	counter++
}

func TestSyncOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)

			once.Do(CounterValue)

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter ", counter)
}
