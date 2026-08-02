package main

import (
	"fmt"
	"strings"
	"time"
)

func TextStreamer(streamText string, stream chan<- string) {
	defer close(stream)
	reader := strings.NewReader(streamText)
	word := make([]byte, 5)
	// fmt.Println(word)
	for {
		n, err := reader.Read(word)
		if n <= 0 {
			return
		}
		if err != nil {
			fmt.Println("There is some error: ", err)
			break
		}

		stream <- fmt.Sprintf("%s",word[:n])
		time.Sleep(200 * time.Millisecond)
	}
}

func TextReceiver(stream <- chan string, done chan struct{}){
	for txt := range stream {
		fmt.Printf(txt)
	}
	done <- struct{}{}
}

func main() {
	streamText := "Our minds are inherently concurrent. At any moment, we juggle memory, intention, perception, and action - parallel processes that somehow compose a unified experience. Writing software that behaves in the same way is deceptively difficult. This chapter explored the essence of concurrent programming in Go by cultivating a mindset of coordination, isolation, and communication. In doing so, we learn to write programs that not only run faster but also operate in harmony with the world around them."

	stream := make(chan string)
	done := make(chan struct{})

	go TextStreamer(streamText, stream)
	go TextReceiver(stream, done)

	<- done
	fmt.Printf("\nStreaming Ended!")
}
