package inspector

import (
	"fmt"
	"reflect"
)

type MapInspector struct {
}

func NewMapInspector() *MapInspector {
	return &MapInspector{}
}

func (r *MapInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.Map
}

func (r *MapInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	var tabs, indentation string
	// TODO: may use buffer
	for i := 0; i < level; i++ {
		tabs += "\t"
	}

	if level > 0 {
		indentation = " "
	}

	fmt.Fprintf(ioP.Output(), "%s%s {\n", indentation, v.Type())
	for _, key := range v.MapKeys() {
		v := v.MapIndex(key)
		// fmt.Printf("===Ele: %s\n", v.Type())
		// fmt.Fprintf(ioP.Output(), "%s\t%s:\t%v,\n", tabs, e, v)
		fmt.Fprintf(ioP.Output(), "%s\t%v: ", tabs, key.Interface())
		ioP.Inspect(v, level+1)
	}

	fmt.Fprintf(ioP.Output(), "%s}\n", tabs)
}
