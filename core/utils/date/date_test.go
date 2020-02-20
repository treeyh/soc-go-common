package date

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/types"
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {

	t1 := types.Time(time.Now())

	str := FormatTime(t1)
	fmt.Println(str)
	fmt.Println(ParseTime(str))

}

func TestFormat(t *testing.T) {

	ti := "2019-07-07 12:13:14"
	dt, _ := Parse(ti)

	convey.Convey("Test Format ", t, func() {
		convey.Convey(" test format date str true ", func() {
			convey.So(Format(dt), convey.ShouldEqual, ti)
		})

		convey.Convey(" test format date str false ", func() {
			convey.So(Format(dt), convey.ShouldNotEqual, "2019-07-06 12:13:14")
		})
	})

}

func TestGetDateInt(t *testing.T) {
	fmt.Println(GetDateInt(types.NowTime()))

	fmt.Println(GetDateTimeLong(types.NowTime()))

	fmt.Println(GetDateIntByTime(time.Now()))

	fmt.Println(GetDateTimeLongByTime(time.Now()))
}
