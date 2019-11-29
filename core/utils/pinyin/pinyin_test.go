package pinyin

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConvertCode(t *testing.T) {

	convey.Convey("Test Format ", t, func() {
		str := "这是One IssueOfFirst任务"
		convey.ShouldEqual(ConvertCode(str), "ZSOIOFRW")

		str = "One这是OfIssue First1123456"
		convey.ShouldEqual(ConvertCode(str), "OZSOIF1123456")

		str = "OneOneTwoTwoThirdThirdFourFourFive"
		convey.ShouldEqual(ConvertCode(str), "OOTTTTFFF")
		convey.ShouldEqual(ConvertCode(str), "OOTTTTFF")
	})

}
