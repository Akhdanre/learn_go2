package learngogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{}{
			return "New"
		},
	}

	pool.Put("Akhdan")
	pool.Put("Robbani")
	pool.Put("ouken")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(3 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(11 * time.Second)
	fmt.Println("selesai")
}
