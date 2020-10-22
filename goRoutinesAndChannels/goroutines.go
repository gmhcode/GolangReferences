package main

import (
	"fmt"
	"time"
)

func main() {
	// go waitAndSay("World")
	// fmt.Println("Hello")
	// time.Sleep(3 * time.Second)

	c := make(chan bool)
	go waitAndSayChan(c, "World")
	fmt.Println("Hello")

	//we send a signal to c in order to allow waitAndSay to continue
	c <- true

	//we wait to receive another signal on c before we exit
	<-c

}

func waitAndSay(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func waitAndSayChan(c chan bool, s string) {
	//the function pauses here until c becomes true
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}

func bufferedChannels() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"
	ch <- "Hello again"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
