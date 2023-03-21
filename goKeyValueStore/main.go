package main

import (
	"fmt"
)

type Test[k comparable, v any] struct{
	data map[k]v
}

func makeTest[k comparable, v any]() *Test[k,v]{
	return &Test[k,v]{
		data: make(map[k]v),
	}
}

func (t *Test[k, v])put(key k, value v){
	t.data[key] = value
	fmt.Printf("Key: %v\nValue: %v", key, value)
}

func main(){
	hello := makeTest[string, string]()
	hello.put("mayhul", "jindal")

	// pointers
}
