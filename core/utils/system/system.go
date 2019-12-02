package system

import (
	"github.com/treeyh/soc-go-common/core/consts"
	"runtime"
)

// IsLinux 判断当前系统是否为linux
func IsLinux() bool {
	return runtime.GOOS == consts.GOOSLinux
}

// IsMac 判断当前系统是否为mac
func IsMac() bool {
	return runtime.GOOS == consts.GOOSMac
}

// IsWindows 判断当前系统是否为windows
func IsWindows() bool {
	return runtime.GOOS == consts.GOOSWindows
}
