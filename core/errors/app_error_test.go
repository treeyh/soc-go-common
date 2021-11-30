package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestNewSystemError(t *testing.T) {

	TestResultCode := NewResultCode(9321412312, "test %s")
	se := NewAppError(TestResultCode, "params")
	fmt.Println(fmt.Sprintf("%s", se))
	fmt.Println(fmt.Sprintf("%q", se))

	fmt.Println(reflect.TypeOf(se))
	fmt.Println(reflect.TypeOf(se).String())
	fmt.Println("=============")
	fmt.Println(reflect.TypeOf(TestResultCode))
	fmt.Println(reflect.TypeOf(TestResultCode).String())
	fmt.Println(reflect.TypeOf(TestResultCode).Name())

	est := fmt.Sprintf("%+v", se.GetError())
	fmt.Println(est)

	ess := strings.Split(est, "\n\t")
	fmt.Println(len(ess))
	for _, v := range ess {
		fmt.Println(v)
	}

	assert.Equal(t, se.Message(), "test params", "errors NewAppError not equal one:%s  other:%s.", se.Message(), "test params")

	se2 := NewAppError(TestResultCode, "test")
	assert.Equal(t, se2.Message(), "test test", "errors NewAppError not equal one:%s  other:%s.", se2.Message(), "test test")

	se3 := NewAppErrorByExistError(TestResultCode, errors.New("errors"), "err")

	fmt.Println(se3.Error())
	fmt.Println(se3.GetError().Error())

	assert.Equal(t, se3.Error(), "test err; err:errors: errors")
}
