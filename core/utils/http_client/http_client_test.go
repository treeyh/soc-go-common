package http_client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {

	result, statue, err := Get(nil, "http://www.baidu.com", nil)

	fmt.Println(result)
	fmt.Println(statue)
	fmt.Println(err)

	assert.True(t, result != "")
}
