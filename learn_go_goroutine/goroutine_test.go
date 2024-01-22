package learngogoroutine

import (
	"fmt"
	"testing"
	"time"
)
func DisplayNumber(number int){
	fmt.Println("Display", number)
}

func RunHelloWordl() {
	fmt.Println("Hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWordl()

	fmt.Println("Ups")
	time.Sleep( 1 * time.Second)
}

func TestManyGoroutine(t *testing.T){
	for i:= 0 ; i < 100000; i++{
		 go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
