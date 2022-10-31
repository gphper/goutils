package arrutil

import (
	"reflect"
	"testing"
)

func TestArrColumn(t *testing.T) {
	type Person struct {
		Username string
		age      int
	}

	persons := [3]Person{
		{"gphper", 18},
		{"gopher", 10},
		{"gper", 20},
	}

	pslice := persons[0:2]

	tmp, err := ArrColumn(persons, "Username")
	if err != nil || reflect.DeepEqual(tmp, []string{"gphper", "gopher", "gper"}) {
		t.Fail()
	}

	_, err = ArrColumn(persons, "age")
	if err == nil {
		t.Fail()
	}

	_, err = ArrColumn(10, "age")
	if err == nil {
		t.Fail()
	}

	tmp, err = ArrColumn(pslice, "Username")
	if err != nil || reflect.DeepEqual(tmp, []string{"gphper", "gopher"}) {
		t.Fail()
	}
}
