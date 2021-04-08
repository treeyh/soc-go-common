package copyer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Source struct {
	Name string
}

type Destination struct {
	Name string
}

func TestCopy(t *testing.T) {

	s := &Source{
		Name: "test",
	}
	d := &Destination{}

	ss := &[]*Source{s}
	dd := &[]*Destination{d}
	Copy(ss, dd)

	assert.Equal(t,"test", (*dd)[0].Name )
}
