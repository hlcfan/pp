package iop

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
)

var std = New()

type IOPrinter struct {
	mutex sync.Mutex
	Out   io.Writer
}

func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

// func Puts(variable interface{}) {
// 	fmt.Fprintf("===Variable: %#v\n", variable)
// }

func Inspect(variable interface{}) {
	std.Inspect(variable)
}

func New() *IOPrinter {
	return &IOPrinter{
		Out: os.Stdout,
	}
}

func (p *IOPrinter) SetOutput(out io.Writer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Out = out
}

func (p *IOPrinter) Inspect(variable interface{}) {
	t := reflect.TypeOf(variable)
	v := reflect.ValueOf(variable)
	// fmt.Printf("===Kind: %#v\n", t.Kind() == reflect.Slice)
	switch t.Kind() {
	case reflect.Slice:
		p.inspectSlice(t, v)
	case reflect.Int:
		// p.inspectSlice(t, v)
	}
}

func (p *IOPrinter) inspectSlice(t reflect.Type, v reflect.Value) {
	fmt.Println("===================")
	// fmt.Println("===Ele type: ", t)
	// fmt.Println("===Ele type: ", t.Elem())
	fmt.Fprintf(p.Out, "%s {\n", t)
	for i := 0; i < v.Len(); i++ {
		ele := v.Index(i)
		fmt.Fprintln(p.Out, "\t\t{")
		for j := 0; j < ele.NumField(); j++ {
			valueField := ele.Field(j)
			typeField := ele.Type().Field(j)
			fmt.Fprintf(p.Out, "\t\t\t%s:\t%v,\n", typeField.Name, valueField.Interface())
		}
		fmt.Fprintln(p.Out, "\t\t},")
		// fmt.Printf("===Index: %#v\n", val)
		// fmt.Printf("\t%#v\n", ele)
	}

	fmt.Fprintln(p.Out, "}")

	// fmt.Printf("%#v\n", v)
	fmt.Println("===================")
}