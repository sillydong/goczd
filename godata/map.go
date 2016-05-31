package godata

import (
	"fmt"
	"reflect"
)

//map转slice
func Map2Slice(params map[string]string, sep string) []string {
	sparam := make([]string, len(params))
	for key, val := range params {
		sparam = append(sparam, key+sep+val)
	}
	return sparam
}

//struct转map[string]interface{}
func Struct2MapInterface(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

//struct转map[string]string
func Struct2MapString(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = fmt.Sprintf("%v", v.Field(i).Interface())
	}
	return data
}
