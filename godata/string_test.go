package godata

import (
	"fmt"
	"testing"
)

func Test_RandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomString(12))
	}

}
