package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type CrewMember struct {
		ID                int      `json:"id,omitempty"`
		Name              string   `json:"name"`
		SecurityClearance int      `json:"clearancelevel"`
		AccessCodes       []string `json:"accessCodes"`
	}

	cm := CrewMember{1, "Jaro", 10, []string{"ADA", "LAL"}}
	b, err := json.Marshal(&cm)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(b))
}
