package main

import (
	"./typevalue"
	"fmt"
	"reflect"
)

type itf interface {
	read()
}
type itfImpl struct {
	field string
}

func (p *itfImpl) read() {
	fmt.Println(p.field)
}

func main() {


	var str string = "Hello"
	type People struct {
		name string
		age  int
		city []string
		info map[string]string
		ptr  *string
		itface itf
	}
	p := People{name: "lanyan",
		age:  27,
		city: []string{"Beijing", "GuangDong"},
		info: map[string]string{"key": "value", "db": "redis"},
		ptr:  &str,
		itface: &itfImpl{field: "Filed"},
	}
	typevalue.Display("People", reflect.ValueOf(p))
}
