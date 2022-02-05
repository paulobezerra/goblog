package utils

import (
	"fmt"
	"reflect"
)

func GetMethods(data interface{}) {
	dataType := reflect.TypeOf(data)
	fmt.Println(dataType.NumMethod())
	for i := 0; i < dataType.NumMethod(); i++ {
		method := dataType.Method(i)
		fmt.Println(method.Name)
	}
}
