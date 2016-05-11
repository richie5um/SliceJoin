package main

import (
	"fmt"
	"reflect"
	"strings"
)

func test1() {
	v1 := [1]int{2}
	v2 := []int{1}

	result, err := join(v1, v2)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)
	fmt.Printf("Output: %v\n", s)
}

func test2() {
	v1 := []string{"a"}
	v2 := [1]string{"b"}

	result, err := join(v1, v2)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)
	fmt.Printf("Output: %v\n", s)
}

func test3() {
	v1 := "a,b,c"
	v2 := "c,d,e"

	result, err := join(v1, v2)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)
	fmt.Printf("Output: %v\n", s)
}

func test4() {
	v1 := []string{"a"}
	v2 := "b,c,d"

	result, err := join(v1, v2)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)
	fmt.Printf("Output: %v\n", s)
}

func test5() {
	v := "b,c,d"

	result, err := join(v, nil)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)
	fmt.Printf("Output: %v\n", s)
}

func test6() {
	v := "b,c,d"

	result, err := join(nil, v)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)
	fmt.Printf("Output: %v\n", s)
}

func test7() {
	result, err := join(nil, nil)
	if nil != err {
		panic(err)
	}

	fmt.Printf("Output: %v\n", result)
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
}

func join(i1, i2 interface{}) (interface{}, error) {
	if nil == i1 {
		i1 = ""
	}
	if nil == i2 {
		i2 = ""
	}

	i1Value := reflect.ValueOf(i1)
	i2Value := reflect.ValueOf(i2)

	splitComma := func(c rune) bool {
		return ',' == c
	}

	if i1Value.Kind() == reflect.String {
		i1Value = reflect.ValueOf(strings.FieldsFunc(i1.(string), splitComma))
		i1 = i1Value.Interface()
	}

	if i2Value.Kind() == reflect.String {
		i2Value = reflect.ValueOf(strings.FieldsFunc(i2.(string), splitComma))
		i2 = i2Value.Interface()
	}

	i1Type := reflect.TypeOf(i1).Elem()
	i2Type := reflect.TypeOf(i2).Elem()

	if i1Type != i2Type {
		return nil, fmt.Errorf("Different Element Types: %v %v", i1Type, i2Type)
	}

	if i1Value.Kind() == reflect.Array {
		i1s := reflect.MakeSlice(reflect.SliceOf(i1Type), 0, i1Value.Len())
		for i := 0; i < i1Value.Len(); i++ {
			i1s = reflect.Append(i1s, i1Value.Index(i))
		}
		i1Value = i1s
	}

	if i2Value.Kind() == reflect.Array {
		i2s := reflect.MakeSlice(reflect.SliceOf(i2Type), 0, i2Value.Len())
		for i := 0; i < i2Value.Len(); i++ {
			i2s = reflect.Append(i2s, i2Value.Index(i))
		}
		i2Value = i2s
	}

	if i1Value.Kind() == reflect.Slice && i2Value.Kind() == reflect.Slice {
		if 0 < i1Value.Len() && 0 < i2Value.Len() {
			return reflect.AppendSlice(i1Value, i2Value).Interface(), nil
		} else if 0 < i1Value.Len() {
			return i1Value.Interface(), nil
		} else if 0 < i2Value.Len() {
			return i2Value.Interface(), nil
		} else {
			return nil, nil
		}
	}

	return nil, fmt.Errorf("Not implemented")
}
