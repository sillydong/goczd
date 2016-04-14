package godata

import "testing"

func Test_InArray(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	str := "a"
	if InArray(arr, str) {
		t.Logf("%v [%s]", arr, str)
	} else {
		t.Errorf("not in array")
	}
}

func Test_IsArray(t *testing.T) {
	//arr := make([]string,0)
	arr := []string{"a", "b", "c"}
	//arr := 123
	if !IsArray(arr) {
		t.Errorf("%v is array", arr)
	}
}
