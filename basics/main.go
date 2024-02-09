package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChannel()
}

func bufferedChannel() {
	const BufferedChannelLength = 2

	timeFormatLayout := "15:04:05"

	ch := make(chan string, BufferedChannelLength)

	go func(ch chan string) {
		for i := 0; i < BufferedChannelLength+3; i++ {
			fmt.Println("sending a message")
			time.Sleep(time.Second * 1)
			ch <- fmt.Sprintf("message number %d, date: %s", i+1, time.Now().Format(timeFormatLayout))
			fmt.Println("message sent")
		}

		close(ch)
	}(ch)

	for msg := range ch {
		fmt.Printf("in date: %s, the message is: %s\n", time.Now().Format(timeFormatLayout), msg)
		time.Sleep(time.Second * 1)
	}
}

func simpleChannelUsing() {
	ch := make(chan string)

	go func(ch chan string) {
		fmt.Println("taking a nap...")

		time.Sleep(time.Second * 2)

		ch <- "hi there"
	}(ch)

	fmt.Println("waiting for a message...")

	fmt.Printf("the message is: %s\n", <-ch)
}
