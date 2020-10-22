package main

import (
	"fmt"
	"reflect"
)

func main() {
	main2()
}

//will change the value of something using reflection
func main2() {
	var x1 float32 = 5.7
	v := reflect.ValueOf(&x1)

	// v.SetFloat(2.2)
	fmt.Println("v settable?", v.CanSet()) //*float32 ==> x1
	vpElem := v.Elem()                     // x1
	fmt.Println("vpElem settable", vpElem.CanSet())

	fmt.Println("vpElem", vpElem)
	vpElem.SetFloat(2.2)
	fmt.Println("vpElem", vpElem)
	fmt.Println("xl", x1)

}

func main1() {
	var x1 float32 = 5.7
	inspectIfTypeFloat(x1)
}

//inspects the type and value of x1
func inspectIfTypeFloat(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Println("type: ", v.Type()) //same as reflect.TypeOf(x1)
	fmt.Println("Is type float64?", v.Kind() == reflect.Float64)
	fmt.Println("Float Value: ", v.Float())

	// This converts the original value back "5.7"
	convertBack(v)
}

func convertBack(v reflect.Value) {
	interfaceValue := v.Interface()
	switch t := interfaceValue.(type) {
	case float32:
		fmt.Println("Original float32 value", t)
	}
}
