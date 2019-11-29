package templates

import (
	"github.com/smartystreets/goconvey/convey"
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

	convey.Convey("TestNewUuid", t, func() {
		str := "这是一个测试 {{.}}"
		str2, _ := Render(str, 1)

		convey.ShouldEqual("这是一个测试 1", str2)

		str = "这是第二个模板测试 {{.Name}}{{.Age}}"
		str2, _ = Render(str, ts)

		convey.ShouldEqual("这是第二个模板测试 test1", str2)

		tmap := map[string]string{
			"Name": "1111",
		}
		str = "这是第三个模板测试 {{.Name}}"

		str2, _ = Render(str, tmap)

		convey.ShouldEqual("这是第三个模板测试 1111", str2)
	})

}
