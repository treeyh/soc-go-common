package errors

import (
	"errors"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewSystemError(t *testing.T) {

	convey.Convey("Test TestNewSystemError", t, func() {
		TestResultCode := NewResultCode(9321412312, "test %s")
		se := NewSystemError(TestResultCode, "params")
		//fmt.Println("se:" + se.Message())
		convey.ShouldEqual(se.Message(), "test params")

		se2 := NewSystemError(TestResultCode, "test")
		//fmt.Println("se2:" + se2.Message())
		//fmt.Println("se1:" + se.Message())
		convey.ShouldEqual(se2.Message(), "test test")

		se3 := NewSystemErrorExistError(TestResultCode, errors.New("errors"), "err")
		//fmt.Println("se2:" + se2.Message())
		//fmt.Println("se1:" + se.Message())
		//fmt.Println("se3:" + se3.Message())
		//fmt.Println("se3 error:" + se3.GetError().Error())

		convey.ShouldEqual(se3.GetError().Error(), "errors")
		convey.ShouldEqual(se3.Message(), "test err")
	})
}
