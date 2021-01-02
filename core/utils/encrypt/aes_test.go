package encrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAesDecrypt(t *testing.T) {

	key := "test key"
	encrypt := "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk="

	d, e := AesDecrypt(key, encrypt)
	t.Log(e)
	t.Log(d)

	str := "abcdefg12345"
	str1 := SHA256String(str)
	t.Log(str1)

	assert.Equal(t, str1, "d7115b844d0dabb68b1d7510df9a098cb9185bf4167a48bba3e1c73adf19757b")

	assert.Equal(t, d, "hello world")
}
