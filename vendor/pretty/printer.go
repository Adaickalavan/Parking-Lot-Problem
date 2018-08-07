package pretty

import (
	"fmt"
	"reflect"
)

// Printer pretty prints any array, slice, or string
func Printer(in interface{}) error {
	v := reflect.ValueOf(in)
	if (v.Kind() != reflect.Slice) &&
		(v.Kind() != reflect.Array) &&
		(v.Kind() != reflect.String) {
		return fmt.Errorf("Incompatible input type: %v", v.Kind())
	}
	if v.Len() == 0 {
		return nil
	}
	for i := 0; i < v.Len()-1; i++ {
		fmt.Printf("%v, ", v.Index(i))
	}
	fmt.Printf("%v\n", v.Index(v.Len()-1))
	return nil
}
