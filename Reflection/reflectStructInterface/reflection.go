package main

import "reflect"

func main() {
	type myStruct struct {
		field1 int
		field2 string
		field3 float64
	}
	mys := myStruct{2, "hello", 2.4}

	// mysRValue := reflect.ValueOf(mys)
	mysRType := reflect.TypeOf(mys)

	// NumField will be 3, because there are 3 properties of the struct
	for i := 0; i < mysRType.NumField(); i++ {
		//Field(i)  accesses the indexpath of the property we are trying to get to
		// Type of the property
		// fieldRType := mysRType.Field(i)
		// Value of the property
		// fieldRValue := mysRValue.Field(i)
	}

}
