package strs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	str := "abcdefg"

	assert.Equal(t, Str2Bytes(str), []byte(str), "strs Str2Bytes error.")
	assert.Equal(t, Bytes2Str(Str2Bytes(str)), str, "strs Bytes2Str error.")
}
