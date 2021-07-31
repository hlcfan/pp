package pp_test

import (
	"bytes"
	"database/sql"
	"testing"
	"time"

	"github.com/hlcfan/pp"
)

type person struct {
	ID        int
	Name      string
	Phone     string
	Graduated bool
	CreatedAt sql.NullTime
	Addresses map[int]address
	vehicles  []vehicle
}

type vehicle struct {
	plate string
}

type address struct {
	PostalCode int
}

func TestPuts(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		var output bytes.Buffer
		pp.SetOutput(&output)

		loc, _ := time.LoadLocation("Asia/Singapore")
		people := []person{
			{
				ID:        1,
				Name:      "alex",
				Phone:     "12345678",
				Graduated: true,
				CreatedAt: sql.NullTime{
					Valid: true,
					Time: time.Date(
						2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
				},
				Addresses: map[int]address{
					1: {PostalCode: 123},
				},
				vehicles: []vehicle{
					{
						plate: "CA-1234",
					},
				},
			},
			{
				ID:        2,
				Name:      "bob",
				Phone:     "87654321",
				Graduated: false,
				CreatedAt: sql.NullTime{
					Valid: true,
					Time: time.Date(
						2021, 06, 5, 20, 34, 58, 651387237, loc),
				},
				Addresses: map[int]address{
					2: {PostalCode: 876},
				},
			},
		}

		pp.Puts(people)

		expected := "[]pp_test.person {\n    {\n        ID:        1,\n        Name:      alex,\n        Phone:     12345678,\n        Graduated: true,\n        CreatedAt: {\n                   Time:  2009-11-17 20:34:58.651387237 +0000 UTC,\n                   Valid: true,\n        },\n        Addresses: map[int]pp_test.address {\n                   1:  {\n                       PostalCode: 123,\n                   },\n        },\n        vehicles: []pp_test.vehicle {\n                  {\n                      plate: CA-1234,\n                  },\n        },\n    },\n    {\n        ID:        2,\n        Name:      bob,\n        Phone:     87654321,\n        Graduated: false,\n        CreatedAt: {\n                   Time:  2021-06-05 20:34:58.651387237 +0800 +08,\n                   Valid: true,\n        },\n        Addresses: map[int]pp_test.address {\n                   2:  {\n                       PostalCode: 876,\n                   },\n        },\n        vehicles: []pp_test.vehicle {\n        },\n    },\n}\n"
		got := output.String()
		// fmt.Printf("=Got: %#v\n", got)
		// fmt.Printf("=Expected: %#v\n", expected)
		if got != expected {
			t.Errorf("Expect:\n%s, but got:\n%s", expected, got)
		}
	})

	t.Run("map", func(t *testing.T) {
		var output bytes.Buffer
		pp.SetOutput(&output)

		m := map[string]string{"foo": "bar", "hello": "world"}
		pp.Puts(m)
		got := output.String()

		pass := false
		expected := []string{
			"map[string]string {\n    foo:   bar,\n    hello: world,\n}\n",
			"map[string]string {\n    hello: world,\n    foo:   bar,\n}\n",
		}

		for _, exp := range expected {
			if got == exp {
				pass = true
				break
			}
		}

		if !pass {
			t.Errorf("Expect one of: %v, but got: %s", expected, got)
		}
	})

	t.Run("simple slice", func(t *testing.T) {
		var output bytes.Buffer
		pp.SetOutput(&output)

		people := []person{
			{
				ID: 1,
				CreatedAt: sql.NullTime{
					Valid: true,
					Time: time.Date(
						2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
				},
			},
		}

		pp.Puts(people)

		expected := "[]pp_test.person {\n    {\n        ID:        1,\n        Name:      ,\n        Phone:     ,\n        Graduated: false,\n        CreatedAt: {\n                   Time:  2009-11-17 20:34:58.651387237 +0000 UTC,\n                   Valid: true,\n        },\n        Addresses: map[int]pp_test.address {\n        },\n        vehicles: []pp_test.vehicle {\n        },\n    },\n}\n"
		got := output.String()
		// fmt.Printf("=Got: %#v\n", got)
		// fmt.Printf("=Expected: %#v\n", expected)
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})

	t.Run("multiple arguments", func(t *testing.T) {
		var output bytes.Buffer
		pp.SetOutput(&output)

		m := []string{"foo", "bar", "hello", "world"}
		pp.Puts("Slice: ", m)
		got := output.String()
		expected := "Slice: \n[]string {\n    foo,\n    bar,\n    hello,\n    world,\n}\n"
		// fmt.Printf("=Got: %#v\n", got)
		// fmt.Printf("=Exp: %#v\n", expected)
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})

	t.Run("it prints with label", func(t *testing.T) {
		var output bytes.Buffer
		pp.SetOutput(&output)

		m := []string{"foo", "bar"}
		pp.Label("Info").Puts(m)
		got := output.String()
		expected := "Info: []string {\n    foo,\n    bar,\n}\n"
		if got != expected {
			t.Errorf("Expect: %s, but got: %s", expected, got)
		}
	})
}
