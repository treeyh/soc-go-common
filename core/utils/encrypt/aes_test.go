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

	str := "cr@ejyi.com014e1bdefc7e42f3894511f48133d278"
	str1 := SHA256String(str)
	t.Log(str1)

	assert.Equal(t, str1, "9f6e9086119fc86b31d14128055625030c1d887d5c09cea62592382c3bedb590")

	assert.Equal(t, d, "hello world")
}
