package main

import "fmt"

func main() {
	sl := new(level)
	sl.raiseShieldLevel(4)
	sl.raiseShieldLevel(5)
	fmt.Println(*sl)
}

type level int

func (lv *level) raiseShieldLevel(i int) {
	if *lv == 0 {
		*lv = 1
	}
	*lv = (*lv) * level(i)
}
