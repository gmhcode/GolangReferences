package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myStruct struct {
		Field1 int     `alias:"f1" desc:"field number 1"`
		Field2 string  `alias:"f2" desc:"field number 2"`
		Field3 float64 `alias:"f3" desc:"field number 3"`
	}
	mys := myStruct{2, "hello", 2.4}
	InspectStructType(&mys)

}

func InspectStructType(mys interface{}) {
	mysRValue := reflect.ValueOf(mys)

	//if its not a pointer, return
	if mysRValue.Kind() != reflect.Ptr {
		println("is ptr")
		return
	}

	fmt.Println(mysRValue)
	// prints
	// &{2 hello 2.4}
	mysRValue = mysRValue.Elem()
	fmt.Println(mysRValue)
	//prints
	// {2 hello 2.4}
	// if not a struct, return
	if mysRValue.Kind() != reflect.Struct {
		println("is struct")
		return
	}
	mysRValue.Field(0).SetInt(15)

	mysRType := mysRValue.Type()
	// NumField will be 3, because there are 3 properties of the struct
	for i := 0; i < mysRType.NumField(); i++ {
		//Field(i)  accesses the indexpath of the property we are trying to get to
		// Type of the property
		fieldRType := mysRType.Field(i) //datatype: StructField
		// Value of the property
		fieldRValue := mysRValue.Field(i)
		fmt.Printf("Field Name: '%s', field type: '%s' , field value: '%v' \n", fieldRType.Name, fieldRType.Type, fieldRValue.Interface())

		fmt.Println("Struct tags, alias: ", fieldRType.Tag.Get("alias"), "description: ", fieldRType.Tag.Get("desc"))
	}
	// prints
	// Field Name: 'Field1', field type: 'int' , field value: '2'
	// Field Name: 'Field2', field type: 'string' , field value: 'hello'
	// Field Name: 'Field3', field type: 'float64' , field value: '2.4'
}
