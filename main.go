package main

import (
	"fmt"
	"reflect"
)

func main() {
	v1 := []string{"a"}
	v2 := [1]string{"b"}

	fmt.Printf("Type v1: %v\n", GetElementType(v1).String())
	fmt.Printf("Type v2: %v\n", GetElementType(v2).String())

	result, err := joinx(v1, v2)
	if nil != err {
		panic(err)
	}

	s := reflect.ValueOf(result)

	fmt.Printf("Output: %v\n", s)
}

func GetElementType(input interface{}) reflect.Type {
	return reflect.TypeOf(input).Elem()
}

func joinx(i1, i2 interface{}) (interface{}, error) {
	i1Type := GetElementType(i1)
	i2Type := GetElementType(i2)

	if i1Type != i2Type {
		return nil, fmt.Errorf("Different Element Types: %v %v", i1Type, i2Type)
	}

	i1Value := reflect.ValueOf(i1)
	if i1Value.Kind() == reflect.Array {
		i1s := make([]interface{}, i1Value.Len())
		for i := 0; i < i1Value.Len(); i++ {
			i1s[i] = i1Value.Index(i).Interface()
		}

		i1Value = reflect.ValueOf(i1s)
		fmt.Printf("i1s Converted: %v\n", i1s)
	}

	i2Value := reflect.ValueOf(i2)
	if i2Value.Kind() == reflect.Array {
		//i2s := make([]interface{}, i2Value.Len())
		i2s := reflect.MakeSlice(reflect.SliceOf(i2Type), i2Value.Len(), i2Value.Len())
		for i := 0; i < i2Value.Len(); i++ {
			i2s = reflect.Append(i2s, i2Value.Index(i))
		}

		i2Value = i2s
		fmt.Printf("i2s: %v\n", i2s)
	}

	fmt.Printf("i1Value: %v\n", i1Value.Kind())
	fmt.Printf("i2Value: %v\n", i2Value.Kind())

	if i1Value.Kind() == reflect.Slice && i2Value.Kind() == reflect.Slice {
		return reflect.AppendSlice(i1Value, i2Value).Interface(), nil
	}

	return nil, fmt.Errorf("Not implemented")
}

// func toSlice(input inteface{}) []interface{} {
//     s := reflect.ValueOf(input)
//     if s.Kind() != reflect.Array {
//         panic("toSlice() given a non-array type")
//     }

//     inputValue := reflect.ValueOf(s)

// 	r := reflect.MakeSlice(l1v.Type(), 0, 0)

// 	for i := 0; i < l1v.Len(); i++ {
// 		r = reflect.Append(r, l1v.Index(i))
// 	}

//     ret := make([]interface{}, s.Len())

//     for i:=0; i<s.Len(); i++ {
//         ret[i] = s.Index(i).Interface()
//     }

//     return ret
// }

// func joinx(l1, l2 interface{}) (interface{}, error) {
// 	if l1 == nil || l2 == nil {
// 		return make([]interface{}, 0), nil
// 	}

// 	l1v := reflect.ValueOf(l1)
// 	l2v := reflect.ValueOf(l2)

// 	if l1v.String() == "array" {
// 		l1 = l1[:]
// 	}

// 	if l2v.String() == "array" {
// 		l2 = l2[:]
// 	}

// 	if l1v.Kind() == l2v.Kind() && l1v.String() == "slice" {
// 		return reflect.AppendSlice(l1, l2), nil
// 	}

// 	fmt.Printf("OutputLv1: %v\n", l1v.Kind())
// 	fmt.Printf("OutputLv2: %v\n", l2v.Kind())

// 	if l1v.Kind() != l2v.Kind() {
// 		return nil, fmt.Errorf("Incorrect Types: %v %v", l1v, l2v)
// 	}

// 	r := reflect.MakeSlice(l1v.Type(), 0, 0)

// 	for i := 0; i < l1v.Len(); i++ {
// 		r = reflect.Append(r, l1v.Index(i))
// 	}

// 	for i := 0; i < l2v.Len(); i++ {
// 		r = reflect.Append(r, l2v.Index(i))
// 	}

// 	return r, nil
// }
