package main

import (
	"fmt"
	"time"
)

func main() {
	go ticketCounter(1)
	time.Sleep(5 * time.Second)
}

func ticketCounter(n int) {
	ticker := time.NewTicker(time.Duration(n) * time.Second)
	i := 0
	//runs this for loop every second of n
	for t := range ticker.C {
		i++
		fmt.Println("Count ", i, " at ", t)
	}
}
