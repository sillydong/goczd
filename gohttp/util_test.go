package gohttp

import (
	"fmt"
	"strconv"
)

type PersonResponse struct {
	Status  int
	Message string
	Data    string
}

func (resp *PersonResponse) OK() bool {
	return resp.Status == 1
}

func (resp *PersonResponse) Msg() bool {
	return resp.Message
}

type PersonRequest struct {
	UserId   int
	UserName string
}

func (req *PersonRequest) URL() (string, error) {
	return "http://test.test/user/" + strconv.Itoa(req.UserId), nil
}

func GetPerson() {
	resp := &PersonResponse{}
	err := DoGetResponse(&PersonRequest{UserId: 1}, resp)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		if resp.OK() {
			fmt.Printf("%v", resp)
		} else {
			fmt.Printf("%v", resp.Msg())
		}
	}
}

func PostPerson() {
	resp, err := DoPost(&PersonRequest{UserId: 1, UserName: "hello"})
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		if resp.OK() {
			fmt.Printf("%v", resp)
		} else {
			fmt.Printf("%v", resp.Msg())
		}
	}
}
