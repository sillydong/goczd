package godata
import "reflect"

//判断对象是否在数组中
func InArray(array interface{}, v interface{}) (exists bool) {
	exists=false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice,reflect.Array:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(v, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}
	return
}

//判断对象在数组中的位置
func ArrayIndex(array interface{}, v interface{}) (index int) {
	index=-1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(v, s.Index(i).Interface()) == true {
				index = i
				return
			}
		}
	}
	return
}

//判断是否数组
func IsArray(array interface{})bool{
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice,reflect.Array:
		return true
	}
	return false
}
