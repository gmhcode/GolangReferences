package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errCrewNotFound = errors.New("crew member not found")

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

type findError struct {
	Name, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findSC(name, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Second)

	if v, ok := scMapping[name]; !ok {
		// return -1, fmt.Errorf("Crew member %s not found on server '%s' ", name, server)
		panic("Crew member not found")
		// return -1, findError{name, server, "Crew member not found"}
	} else if ok {
		return v, nil
	}
	return 0, nil
}

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	clearance, err := findSC("Ruko", "Server 1")
// 	fmt.Println("Clearanace level found: ", clearance, " Error code", err)
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// if clearance, err := findSC("RUKU", "server 1"); err != nil {
// 	fmt.Println("Error occured while searching ", err)
// 	//This checks to see if err is of type (findError)
// 	if v, ok := err.(findError); ok {
// 		fmt.Println("Server name ", v.Server)
// 		fmt.Println("Crew Member Name ", v.Name)
// 	}
// } else {
// 	fmt.Println("Crew member has security clearance ", clearance)
// }
// }
//Recover From panics
func main() {
	rand.Seed(time.Now().UnixNano())

	// recover needs to be placed inside a defer function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A Panic recovered", err)
		}
	}()
	if clearance, err := findSC("RUKU", "server 1"); err != nil {
		fmt.Println("Error occured while searching ", err)
		//This checks to see if err is of type (findError)
		if v, ok := err.(findError); ok {
			fmt.Println("Server name ", v.Server)
			fmt.Println("Crew Member Name ", v.Name)
		}
	} else {
		fmt.Println("Crew member has security clearance ", clearance)
	}
}
