package main

import (
	"classFactoryTutorial/appliances"
	"fmt"
)

func main() {

	fmt.Println("Enter preferred appliance type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")
	//Accepts input from the terminal
	var myType int
	fmt.Scan(&myType)

	myAppliance, err := appliances.CreateAppliance(myType)

	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}
}
