package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSystemError(t *testing.T) {

	TestResultCode := NewResultCode(9321412312, "test %s")
	se := NewAppError(TestResultCode, "params")
	assert.Equal(t, se.Message(), "test params", "errors NewAppError not equal one:%s  other:%s.", se.Message(), "test params")


	se2 := NewAppError(TestResultCode, "test")
	assert.Equal(t, se2.Message(), "test test", "errors NewAppError not equal one:%s  other:%s.", se2.Message(), "test test")

	se3 := NewAppErrorByExistError(TestResultCode, errors.New("errors"), "err")

	assert.Equal(t, se3.GetError().Error(), "errors", "errors NewAppError not equal one:%s  other:%s.", se3.GetError().Error(), "errors")
	assert.Equal(t, se3.Message(), "test err; err:errors", "errors NewAppError not equal one:%s  other:%s.", se3.Message(), "test err; err:errors")
}
