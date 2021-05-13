package http_client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {

	result, statue, err := Get(nil, "http://www.baidu.com", nil, nil)

	fmt.Println(result)
	fmt.Println(statue)
	fmt.Println(err)

	assert.Equal(t, statue, 200)
	assert.NoError(t, err)

	assert.True(t, result != "")
}
