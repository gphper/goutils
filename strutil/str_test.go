package strutil

import (
	"testing"
)

func TestJoinStrs(t *testing.T) {

	cases := []struct {
		args   []interface{}
		result string
	}{
		{[]interface{}{"Hello", "World"}, "HelloWorld"},
	}

	for _, v := range cases {
		if JoinStrs(v.args...) != v.result {
			t.Fatalf("JoinStrs functcdion failed, args: %v, result：%v", v.args, v.result)
		}
	}

}

func TestInterfaceToString(t *testing.T) {
	cases := []struct {
		args   interface{}
		result string
	}{
		{20, "20"},
		{10.12, "10.12"},
		{float32(12.1234), "12.1234"},
		{uint(32), "32"},
		{int8(32), "32"},
		{uint8(32), "32"},
		{int16(32), "32"},
		{uint16(32), "32"},
		{int32(32), "32"},
		{uint32(32), "32"},
		{int64(32), "32"},
		{uint64(32), "32"},
		{"hello", "hello"},
		{true, "True"},
		{[]byte{72, 101, 108, 108, 111, 87, 111, 114, 108, 100}, "HelloWorld"},
		{
			struct {
				Name string
			}{"gphper"},
			`{"Name":"gphper"}`,
		},
	}

	for _, v := range cases {
		res := InterfaceToStr(v.args)
		if res != v.result {
			t.Fatalf("InterfaceToString functcdion failed, args: %v, expect：%v,result：%v ", v.args, v.result, res)
		}
	}
}
