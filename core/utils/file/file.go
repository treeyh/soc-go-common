package file

import (
	"os"
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

// ExistFile 判断文件是否存在
func ExistFile(filePath string) bool {
	_, err := os.Lstat(filePath)
	return !os.IsNotExist(err)
}

// WriteFile 写文件
func WriteFile(filePath, content string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(content)
}
