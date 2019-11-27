package errors

import (
	"fmt"
	"testing"


)

func TestNewSystemError(t *testing.T) {

	TestResultCode := NewResultCode(9321412312, "test %s")

	se := NewSystemError(TestResultCode, "params")

	fmt.Println(OK)
	fmt.Println(se)
	fmt.Println(se.Message())

	se2 := NewSystemError(TestResultCode)
	fmt.Println(se2)
	fmt.Println(se2.Message())

}
