package inspector_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hlcfan/iop"
	"github.com/hlcfan/iop/inspector"
)

type Person struct {
	ID    int
	Name  string
	Phone string
}

func TestInspectMap(t *testing.T) {
	t.Run("It inspects map", func(t *testing.T) {
		var output bytes.Buffer

		maps := map[string]interface{}{}
		maps["name"] = "alex"
		maps["age"] = 20
		maps["graduated"] = false

		// maps := map[string]Person{
		// 	"alex": {
		// 		ID:    1,
		// 		Name:  "alex",
		// 		Phone: "12345678",
		// 	},
		// }

		vType := reflect.TypeOf(maps)
		vValue := reflect.ValueOf(maps)

		ioP := iop.New()
		ioP.SetOutput(&output)
		mapInspector := inspector.NewMapInspector()
		mapInspector.Inspect(ioP, vType, vValue, 0)

		//TODO: can be flaky, due to map doesn't maintain order
		expected := " map[string]interface {} {\n\tname:\talex,\n\tage:\t20,\n\tgraduated:\tfalse,\n}\n"
		got := output.String()
		if got != expected {
			t.Errorf("Expect: \n%v, but got: \n%v", expected, got)
		}
	})
}
