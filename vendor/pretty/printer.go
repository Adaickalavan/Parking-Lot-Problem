package pretty

import (
	"fmt"
	"reflect"
)

// Printer pretty prints any array, slice, or string
func Printer(in interface{}) {
	v := reflect.ValueOf(in)
	if v.Len() == 0 {
		return
	}
	for i := 0; i < v.Len()-1; i++ {
		fmt.Printf("%v, ", v.Index(i))
	}
	fmt.Printf("%v\n", v.Index(v.Len()-1))
}
