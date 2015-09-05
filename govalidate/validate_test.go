package govalidate

import "testing"

func Test_IsMobile(t *testing.T) {
	if IsMobile("13151285088") {
		t.Log("1正确")
	}
	if !IsMobile("1315128508") {
		t.Log("2正确")
	}
	if IsMobile("17151285088") {
		t.Log("3正确")
	}
}

func Test_IsDate(t *testing.T) {

}

func Test_IsEmail(t *testing.T) {

}

func Test_IsUrl(t *testing.T) {

}

func Test_IsIp(t *testing.T) {

}
