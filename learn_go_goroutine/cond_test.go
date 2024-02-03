package learngogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var Locker = sync.Mutex{}
var cond = sync.NewCond(&Locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()

	fmt.Println("done", value)

	cond.L.Unlock()
}

func TestCondition(t *testing.T) {

	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// langusng mengeksekusi semua tanpa menunggu
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()

}
