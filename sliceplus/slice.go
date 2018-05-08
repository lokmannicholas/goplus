package sliceplus

import (
	"reflect"
)

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
		if slice.Index(i).Interface() == v2 {
			return true
		}
	}

	return false
}

func IsSlice(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
