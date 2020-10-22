package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker executes code once per time interval, in this case its ever second
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go tickCounter(ticker, done)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(10 * time.Second)
	fmt.Println("Exiting...")
}

func tickCounter(ticker *time.Ticker, done chan bool) {
	i := 0
Loop:
	for {
		select {
		case t := <-ticker.C:
			i++
			fmt.Println("Count ", i, " at ", t)
		case <-done:
			fmt.Println("done signal")
			break Loop
		}
	}
	fmt.Println("Exiting the tick Counter")
}
