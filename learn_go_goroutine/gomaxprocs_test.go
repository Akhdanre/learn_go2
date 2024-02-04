package learngogoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)


func TestGetGomaxprocs(t *testing.T){
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++{
		group.Add(1)
		go func ()  {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu :", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread :",totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine :",totalGoroutine)
}
func TestChangeGomaxprocs(t *testing.T){
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++{
		group.Add(1)
		go func ()  {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu :", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread :",totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine :",totalGoroutine)
}