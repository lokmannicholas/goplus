package sliceplus

import (
	"fmt"
	"reflect"
)

func Remove(v interface{}, pos int) interface{} {
	slice, success := IsSlice(v, reflect.Slice)
	if !success {
		return nil
	}
	if pos >= slice.Len() || pos < 0 {
		return nil
	}
	var newSlice = make([]interface{}, 0, slice.Len()-1)
	for i := 0; i < slice.Len(); i++ {
		if i != pos {
			newSlice = append(newSlice, slice.Index(i).Interface())
		}
	}
	return newSlice
}

func Delete(v interface{}, v2 interface{}) interface{} {
	slice, success := IsSlice(v, reflect.Slice)
	if !success {
		return nil
	}
	if !Contains(v, v2) {
		return nil
	}
	item, success := IsSlice(v2, reflect.Slice)
	if success {
		var temp = v
		for j := 0; j < item.Len(); j++ {
			temp = Delete(temp, item.Index(j).Interface())
			if temp == nil {
				return nil
			}
		}
		return temp
	}
	var index = -1
	for i := 0; i < slice.Len(); i++ {
		if matchValue(slice.Index(i).Interface(), v2) {
			index = i
			break
		}
	}
	return Remove(v, index)
}

func Find(v interface{}, v2 interface{}) int {
	slice, success := IsSlice(v, reflect.Slice)
	if !success {
		if v == v2 {
			return 0
		}
	}

	item, success := IsSlice(v2, reflect.Slice)
	if success {
		for j := 0; j < item.Len(); j++ {
			if Find(v, item.Index(j).Interface()) > -1 {
				return j
			}
		}
	}
	for i := 0; i < slice.Len(); i++ {
		if matchValue(slice.Index(i).Interface(), v2) {
			return i
		}
	}

	return -1
}
func Contains(v interface{}, v2 interface{}) bool {
	slice, success := IsSlice(v, reflect.Slice)
	if !success {
		if v == v2 {
			return true
		}
	}

	item, success := IsSlice(v2, reflect.Slice)
	if success {
		for j := 0; j < item.Len(); j++ {
			if !Contains(v, item.Index(j).Interface()) {
				return false
			}
		}
		return true
	}
	for i := 0; i < slice.Len(); i++ {
		if matchValue(slice.Index(i).Interface(), v2) {
			return true
		}
	}

	return false
}

func matchValue(v interface{}, v2 interface{}) bool {
	return fmt.Sprintf("%v", v) == fmt.Sprintf("%v", v2)
}

func IsSlice(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
