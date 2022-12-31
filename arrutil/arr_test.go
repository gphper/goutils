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

func TestUnset(t *testing.T) {

	cases1 := []struct {
		items []int
		item  int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{1, 3, 4}},
		{[]int{1, 2, 3, 4}, 1, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3}},
	}

	for _, v := range cases1 {
		result := Unset[int](v.items, v.item)
		if !reflect.DeepEqual(result, v.want) {
			t.Errorf("items %+v,item %d,want %+v,actual %+v", v.items, v.item, v.want, result)
		}
	}

	cases2 := []struct {
		items []string
		item  string
		want  []string
	}{
		{[]string{"a", "b", "c"}, "a", []string{"b", "c"}},
		{[]string{"a", "b", "c"}, "b", []string{"a", "c"}},
		{[]string{"a", "b", "c"}, "c", []string{"a", "b"}},
	}

	for _, v := range cases2 {
		result := Unset[string](v.items, v.item)
		if !reflect.DeepEqual(result, v.want) {
			t.Errorf("items %+v,item %s,want %+v,actual %+v", v.items, v.item, v.want, result)
		}
	}

}
