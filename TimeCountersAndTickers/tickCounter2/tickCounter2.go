package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker executes code once per time interval, in this case its ever second
	ticker := time.NewTicker(1 * time.Second)
	go tickCounter(ticker)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	time.Sleep(10 * time.Second)
	fmt.Println("Exiting...")
}

func tickCounter(ticker *time.Ticker) {
	i := 0
	for t := range ticker.C {
		i++
		fmt.Println("Count ", i, " at ", t)
	}
}
