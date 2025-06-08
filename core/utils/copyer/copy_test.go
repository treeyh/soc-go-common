package copyer

import (
	"context"
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

type Source struct {
	Name string `json:"name"`
}

type Destination struct {
	Name string `json:"name"`
}

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Age  int32  `json:"age"`
}

// func (user *User) DoubleAge() int32 {
// 	return 2 * user.Age
// }

type Employee struct {
	Name      string `json:"name"`
	Age       int32  `json:"age"`
	DoubleAge int32  `json:"doubleAge"`
	SuperRole string `json:"superRole"`
}

// func (employee *Employee) Role(role string) {
// 	employee.SuperRole = "Super " + role
// }

func TestCopy(t *testing.T) {

	user := User{Name: "Jinzhu", Age: 18, Role: "Admin"}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)

	sss := &Source{Name: "test"}
	ddd := &Destination{}

	err2 := Copy(context.Background(), sss, ddd)
	fmt.Println(err2)
	fmt.Printf("%#v\n", ddd)
	assert.Equal(t, "test", ddd.Name)

	// ss := &[]*Source{s}
	// dd := &[]*Destination{d}
	// Copy(context.Background(), s, d)

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
