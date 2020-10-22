package main

import (
	"fmt"
	"time"
)

func main() {

}

func SlowCounter(n int) {
	i := 0

	d := time.Duration(n) * time.Second

	for {
		t := time.NewTimer(d)
		<-t.C
		i++
		fmt.Println(i)
	}
}
