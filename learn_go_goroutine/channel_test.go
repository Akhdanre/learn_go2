package learngogoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// channel <- "akhdan"

	// // var data string
	// // data = <- channel

	// data := <- channel

	// //lansung
	// // fmt.Println(<- channel)

	// fmt.Println(data)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Akhdan Robbani"
		fmt.Println("selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Akhdan Robbani"
}

func TestChannnelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "test only in"
	//data := <- channel // will error cause channel only in
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInAndOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)

}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Akhdan"
		channel <- "Robbani"
		channel <- "Oukenze"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("selesai")

}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulanga ke " + strconv.Itoa(i)
			time.Sleep(2 * time.Second)
		}
		defer close(channel)

	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	go GiveMeResponse(channel1)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 3 {
			break
		}
	}
}
