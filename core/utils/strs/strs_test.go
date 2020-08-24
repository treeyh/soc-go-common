package strs

import (
	"fmt"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	str := "abc阳阳1天天23"
	fmt.Println(Str2Bytes(str))
	fmt.Println([]byte(str))

	bts := Str2Bytes(str)

	fmt.Println(Bytes2Str(bts))
	fmt.Println(string(bts))

}
