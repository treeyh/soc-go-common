package templates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	Name string
	Age  int64
}

func TestRender(t *testing.T) {
	ts := TestStruct{
		Name: "test",
		Age:  1,
	}

	str := "这是一个测试 {{.}}"
	str2, _ := Render(str, 1)

	assert.Equal(t, "这是一个测试 1", str2, "templates render error.")

	str = "这是第二个模板测试 {{.Name}}{{.Age}}"
	str2, _ = Render(str, ts)

	assert.Equal(t, "这是第二个模板测试 test1", str2, "templates render error.")

	tmap := map[string]string{
		"Name": "1111",
	}
	str = "这是第三个模板测试 {{.Name}}"

	str2, _ = Render(str, tmap)

	assert.Equal(t, "这是第三个模板测试 1111", str2, "templates render error.")

}
