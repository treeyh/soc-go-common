package random

import (
	"fmt"
	"testing"
)

func TestGetTime0(t *testing.T) {
	fmt.Println(RandomString(6))

	str := "0123456789"
	bytes := []byte(str)
	fmt.Println(RandomString2(6, bytes))

}
