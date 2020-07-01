package file

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestGetCurrentPath(t *testing.T) {

	convey.Convey("log test", t, tests.TestStartUp(func() {
		fmt.Println(GetCurrentPath())
	}, nil))

}
