package uuid

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGetUuid(t *testing.T) {

	fmt.Println(NewUuid())
}

func TestNewUuid(t *testing.T) {
	convey.Convey("TestNewUuid", t, func() {
		convey.Convey(" test new uuid ", func() {
			convey.So(NewUuid(), convey.ShouldNotBeEmpty)
		})
	})
}
