package pp

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"

	"github.com/hlcfan/pp/inspector"
)

var std = New()

type PPrinter struct {
	mutex      sync.Mutex
	Out        io.Writer
	inspectors []Inspectable
	maxDepth   int
}

type Inspectable interface {
	Applicable(reflect.Type, reflect.Value) bool
	Inspect(inspector.Printable, reflect.Type, reflect.Value, int)
}

func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

// func Puts(variable interface{}) {
// 	fmt.Fprintf("===Variable: %#v\n", variable)
// }

func Inspect(variable interface{}) {
	v := reflect.ValueOf(variable)
	std.Inspect(v, 0)
}

func New() *PPrinter {
	return &PPrinter{
		Out: os.Stdout,
		// maxDepth: 2,
		inspectors: []Inspectable{
			inspector.NewSliceInspector(),
			inspector.NewMapInspector(),
			inspector.NewIntegerInspector(),
			inspector.NewTimeInspector(),
			inspector.NewStructInspector(),
			inspector.NewStringInspector(),
			inspector.NewBoolInspector(),
			inspector.NewInterfaceInspector(),
			// inspector.NewFallbackInspector(),
		},
	}
}

func (p *PPrinter) SetOutput(out io.Writer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Out = out
}

func (p *PPrinter) Inspect(variable reflect.Value, level int) {
	if p.maxDepth > 0 && level > p.maxDepth {
		return
	}

	var inspector Inspectable

	// v := reflect.ValueOf(variable)
	v := variable
	t := reflect.TypeOf(v)

	// fmt.Printf("===To inspect: %#v\n", variable)
	fmt.Printf("===To inspect kind: %s\n", v.Kind())
	for _, i := range p.inspectors {
		if i.Applicable(t, v) {
			fmt.Println("===Found")
			inspector = i
			break
		}
	}

	// fmt.Printf("===Inspectable: %#v\n", inspector)
	if inspector == nil {
		return
	}

	// Should be pass in PPrinter to cater case for nested object, so that
	// `Inspect` can be called nestedly.
	inspector.Inspect(p, t, v, level)
	// // fmt.Printf("===Kind: %#v\n", t.Kind() == reflect.Slice)
	// switch t.Kind() {
	// case reflect.Slice:
	// 	p.inspectSlice(t, v)
	// case reflect.Int:
	// 	// p.inspectSlice(t, v)
	// }
}

func (p *PPrinter) Output() io.Writer {
	return p.Out
}

func (p *PPrinter) inspectSlice(t reflect.Type, v reflect.Value) {
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