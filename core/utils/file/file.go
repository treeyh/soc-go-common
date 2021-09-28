package file

import (
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	log = logger.Logger()
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
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// ReadSmallFile 读取小文件，一次性读取
func ReadSmallFile(filePath string) (*string, errors.AppError) {
	tmpContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.FileReadFail, err)
	}
	content := string(tmpContent)
	return &content, nil
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

// GetFileModTime 获取文件修改时间
func GetFileModTime(filePath string) (time.Time, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return time.Unix(0, 0), err
	}
	return fileInfo.ModTime(), nil
}

// IsDir 是否是目录
func IsDir(filePath string) (bool, error) {
	fileInfo, err := os.Stat("test.log")
	if err != nil {
		return false, err
	}

	//是否是目录
	return fileInfo.IsDir(), nil
}

// GetDirSon 返回目录下子文件/目录列表
func GetDirSon(filePath string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(filePath)
}

// GetDirWalk 递归获取路径子文件，目录列表
func GetDirWalk(filePath string) (map[string]os.FileInfo, error) {
	files := make(map[string]os.FileInfo)

	err := filepath.Walk(filePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files[path] = info
			return nil
		})

	return files, err
}

// GetCurrentAbPath 最终方案-全兼容
func GetCurrentAbPath() string {
	dir := GetCurrentAbPathByExecutable()
	if strings.Contains(dir, GetTmpDir()) {
		return GetCurrentAbPathByCaller()
	}
	return dir
}

// GetTmpDir 获取系统临时目录，兼容go run
func GetTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// GetCurrentAbPathByExecutable 获取当前执行文件绝对路径
func GetCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// GetCurrentAbPathByCaller 获取当前执行文件绝对路径（go run）
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
