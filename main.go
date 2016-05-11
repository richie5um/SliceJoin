package main

import (
	"fmt"
	"reflect"
)

func test1() {
	v1 := [1]int{2}
	v2 := []int{1}

	result, err := joinx(v1, v2)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)
	fmt.Printf("Output: %v\n", s)
}

func test2() {
	v1 := []string{"a"}
	v2 := [1]string{"b"}

	result, err := joinx(v1, v2)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)
	fmt.Printf("Output: %v\n", s)
}

func main() {
	test1()
	test2()
}

func joinx(i1, i2 interface{}) (interface{}, error) {
	i1Type := reflect.TypeOf(i1).Elem()
	i2Type := reflect.TypeOf(i2).Elem()

	if i1Type != i2Type {
		return nil, fmt.Errorf("Different Element Types: %v %v", i1Type, i2Type)
	}

	i1Value := reflect.ValueOf(i1)
	if i1Value.Kind() == reflect.Array {
		i1s := reflect.MakeSlice(reflect.SliceOf(i1Type), 0, i1Value.Len())
		for i := 0; i < i1Value.Len(); i++ {
			i1s = reflect.Append(i1s, i1Value.Index(i))
		}

		i1Value = i1s
	}

	i2Value := reflect.ValueOf(i2)
	if i2Value.Kind() == reflect.Array {
		i2s := reflect.MakeSlice(reflect.SliceOf(i2Type), 0, i2Value.Len())
		for i := 0; i < i2Value.Len(); i++ {
			i2s = reflect.Append(i2s, i2Value.Index(i))
		}

		i2Value = i2s
	}

	if i1Value.Kind() == reflect.Slice && i2Value.Kind() == reflect.Slice {
		return reflect.AppendSlice(i1Value, i2Value).Interface(), nil
	}

	return nil, fmt.Errorf("Not implemented")
}
