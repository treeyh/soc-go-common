package file

import (
	"path"
	"runtime"
)

// GetCurrentPath 获取调用方当前目录路径
func GetCurrentPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath := path.Dir(filename)
		return abPath
	}
	return ""
}
