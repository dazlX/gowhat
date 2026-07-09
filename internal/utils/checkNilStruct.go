package utils

import "reflect"

func CheckNilStruct(obj any) bool {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsZero() {
			return false
		}
	}
	return true
}