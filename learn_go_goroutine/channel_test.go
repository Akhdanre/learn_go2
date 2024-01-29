package learngogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)


	// channel <- "akhdan"

	// // var data string
	// // data = <- channel

	// data := <- channel

	// //lansung
	// // fmt.Println(<- channel)

	// fmt.Println(data)


	go func(){
		time.Sleep(2 * time.Second)
		channel <- "Akhdan Robbani"
		fmt.Println("selesai mengirim data ke channel")
	}()
	
	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Akhdan Robbani"
}

func TestChannnelAsParameter(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	
	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}