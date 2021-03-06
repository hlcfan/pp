package inspector

import (
	"fmt"
	"reflect"
)

type StructInspector struct {
}

func NewStructInspector() *StructInspector {
	return &StructInspector{}
}

func (r *StructInspector) Applicable(t reflect.Type, v reflect.Value) bool {
	return v.Kind() == reflect.Struct
}

func (r *StructInspector) Inspect(ioP Printable, t reflect.Type, v reflect.Value, level int) {
	var tabs string
	// TODO: may use buffer
	for i := 0; i < level; i++ {
		tabs += "\t"
	}

	var structType string
	if level == 0 {
		structType = v.Type().String()
	}

	if len(structType) == 0 {
		fmt.Fprintf(ioP.Output(), "%v{\n", structType)
	} else {
		fmt.Fprintf(ioP.Output(), "%v{\n", structType)
	}

	for j := 0; j < v.NumField(); j++ {
		field := v.Field(j)

		typeField := v.Type().Field(j)
		fmt.Fprintf(ioP.Output(), "%s\t%s:\t", tabs, typeField.Name)

		ioP.Inspect(field, level+1)
	}

	var comma string
	if level > 0 {
		comma = ","
	}

	fmt.Fprintf(ioP.Output(), "%s}%s\n", tabs, comma)
}
