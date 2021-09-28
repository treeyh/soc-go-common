package file

import (
	"fmt"
	"testing"
)

func TestGetCurrentPath(t *testing.T) {

	fmt.Println(GetCurrentPath())

}

func TestGetCurrentAbPath(t *testing.T) {
	fmt.Println("GetTmpDir（当前系统临时目录） = ", GetTmpDir())
	fmt.Println("GetCurrentAbPathByExecutable（仅支持go build） = ", GetCurrentAbPathByExecutable())
	fmt.Println("GetCurrentAbPathByCaller（仅支持go run） = ", GetCurrentAbPathByCaller())
	fmt.Println("GetCurrentAbPath（最终方案-全兼容） = ", GetCurrentAbPath())
}
