package copyer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Source struct {
	Name string `json:"name"`
}

type Destination struct {
	Name string `json:"name"`
}

func TestCopy(t *testing.T) {

	s := &Source{
		Name: "test",
	}
	d := &Destination{}

	ss := &[]*Source{s}
	dd := &[]*Destination{d}
	Copy(context.Background(), ss, dd)

	assert.Equal(t, "test", (*dd)[0].Name)

	ls := make([]Source, 0)
	ls = append(ls, Source{
		Name: "test",
	}, Source{
		Name: "test2",
	})

	ds := make([]Destination, len(ls))

	err := CopyList(context.Background(), ls, &ds)
	assert.NoError(t, err)
	assert.True(t, len(ds) > 0)

	t.Log(ds)

}
