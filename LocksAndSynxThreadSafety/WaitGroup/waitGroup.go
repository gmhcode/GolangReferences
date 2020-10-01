package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("work done for ", i)
		}(i)
	}
	//Completed the waitGroup when the Add() counter reaches zero. decrement wait groups by calling done on them
	wg.Wait()
}
