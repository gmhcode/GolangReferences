package main

import (
	"encoding/json"
	"fmt"
)

//the json tags are like aliases
type CrewMember struct {
	ID                int      `json:"id,omitempty"`
	Name              string   `json:"name"`
	SecurityClearance int      `json:"clearancelevel"`
	AccessCodes       []string `json:"accessCodes"`
}

// with tags, prints
// {"id":1,"name":"Jaro","clearancelevel":10,"accessCodes":["ADA","LAL"]}
// without tags prints
// {"ID":1,"Name":"Jaro","SecurityClearance":10,"AccessCodes":["ADA","LAL"]}
type ShipInfo struct {
	ShipId    int
	ShipClass string
	Captain   CrewMember
}

// prints
// {"ShipId":1,"ShipClass":"Fighter","Captain":{"id":1,"name":"Jaro","clearancelevel":10,"accessCodes":["ADA","LAL"]}}
func main() {
	//the json tags are like aliases

	cm := CrewMember{1, "Jaro", 10, []string{"ADA", "LAL"}}
	// 	prints
	// {"ID":1,"Name":"Jaro","SecurityClearance":10,"AccessCodes":["ADA","LAL"]}
	si := ShipInfo{1, "Fighter", cm}

	// cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}
	// omitempty ID
	// prints
	// {"name":"Jaro","clearancelevel":10,"accessCodes":["ADA","LAL"]}
	b, err := json.Marshal(&si)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("\n ShipInfo")
	fmt.Println(string(b))
	fmt.Println("\n dictionary")
	m := map[string]int{"item1": 1, "item2": 2}
	// or
	// m := map[int]string{1: "item1", 2: "item2"}

	bm, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(bm))

	printCrewMembers(cm)
	fmt.Println("\n UnMarshal ShipInfo")
	unmarshallShip()
	fmt.Println("\n UnMarshal map")
	unMarshalMap()

	fmt.Println("\n Unmarshal crew members")
	unMarshalArrayCrewMembers()
}

func printCrewMembers(crewMember CrewMember) {
	s := []CrewMember{crewMember, CrewMember{2, "Jim", 5, []string{"TLT", "RAR"}}}

	bSlice, err := json.Marshal(&s)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(bSlice))
}

func unmarshallShip() {
	sbyte := []byte(`{"ShipId":1,"ShipClass":"Fighter","Captain":{"name":"Jaro","clearancelevel":10,"accessCodes":["ADA","LAL"]}}`)
	shipInfo := new(ShipInfo)
	json.Unmarshal(sbyte, shipInfo)
	fmt.Println(shipInfo.ShipId, shipInfo.ShipClass, shipInfo.Captain.Name)
	// Prints
	// 1 Fighter Jaro
}
func unMarshalMap() {

	m := make(map[int]string)
	data := []byte(`{"1": "item1", "2": "item2"}`)
	json.Unmarshal(data, &m)
	fmt.Println(m)

}

func unMarshalArrayCrewMembers() {
	data := []byte(`[{"id":1,"name":"Jaro","clearancelevel":10,"accessCodes":["ADA","LAL"]},{"id":2,"name":"Jim","clearancelevel":5,"accessCodes":["TLT","RAR"]}]`)
	s := []CrewMember{}
	json.Unmarshal(data, &s)
	fmt.Println(s)
}
