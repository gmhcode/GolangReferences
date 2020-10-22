package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker executes code once per time interval, in this case its ever second
	ticker := time.NewTicker(1 * time.Second)

	done := tickCounter(ticker)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(10 * time.Second)
	fmt.Println("Exiting...")
}

func tickCounter(ticker *time.Ticker) chan bool {
	done := make(chan bool)

	go func() {
		i := 0
	Loop:
		for {
			//This is basically a switch statement for listeners
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
	}()
	return done
}
