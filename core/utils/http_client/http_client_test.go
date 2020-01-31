package http_client

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestGet(t *testing.T) {
	convey.Convey("http get test", t, tests.TestStartUp(func() {
		result, statue, err := Get(nil, "http://www.baidu.com", nil)

		fmt.Println(result)
		fmt.Println(statue)
		fmt.Println(err)
		convey.ShouldBeTrue(result != "")
	}, nil))
}
