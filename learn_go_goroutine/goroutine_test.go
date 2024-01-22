package learngogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWordl() {
	fmt.Println("Hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWordl()

	fmt.Println("Ups")
	time.Sleep( 1 * time.Second)
}
