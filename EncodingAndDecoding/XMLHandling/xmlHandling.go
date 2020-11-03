package main

import (
	"encoding/xml"
	"fmt"
)

// attr makes it print <Captain name="Jaro">
type CrewMember struct {
	ID                int      `xml:"id,omitempty"`
	Name              string   `xml:"name,attr"`
	SecurityClearance int      `xml:"clearancelevel"`
	AccessCodes       []string `xml:"accessCodes>code"`
	// prints
	//  <accessCodes>
	// 		<code>ADA</code>
}

type ShipInfo struct {
	// tag will now be <ship> instead of <ShipInfo>
	XMLName   xml.Name   `xml:"ship"`
	ShipId    int        `xml:"ShipDetails>ShipId"`
	ShipClass string     `xml:"ShipDetails>ShipClass"`
	Captain   CrewMember `xml:"captain"`
}

func main() {

	cm := CrewMember{ID: 1, Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}
	si := ShipInfo{ShipId: 1, ShipClass: "Fighter", Captain: cm}

	b, err := xml.MarshalIndent(si, " ", "	")

	if err != nil {
		fmt.Println("ERR:", err)
		return
	}
	fmt.Println(string(b))
	// Marshal prints
	// <ShipInfo><ShipId>1</ShipId><ShipClass>Fighter</ShipClass><Captain><name>Jaro</name><clearancelevel>10</clearancelevel><accessCodes>ADA</accessCodes><accessCodes>LAL</accessCodes></Captain></ShipInfo>

	//MarshallIndent prints
	// 	 <ShipInfo>
	//         <ShipId>1</ShipId>
	//         <ShipClass>Fighter</ShipClass>
	//         <Captain>
	//                 <name>Jaro</name>
	//                 <clearancelevel>10</clearancelevel>
	//                 <accessCodes>ADA</accessCodes>
	//                 <accessCodes>LAL</accessCodes>
	//         </Captain>
	//  </ShipInfo>

}
