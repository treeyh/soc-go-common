package copyer

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

type Source struct {
	Name string
}

type Destination struct {
	Name string
}

func TestCopy(t *testing.T) {

	convey.Convey("Test TestCopy", t, func() {
		s := &Source{
			Name: "test",
		}
		d := &Destination{}

		ss := &[]*Source{s}
		dd := &[]*Destination{d}
		Copy(ss, dd)
		convey.ShouldEqual("test", (*dd)[0].Name)
	})
}
