package network

import (
	"github.com/treeyh/soc-go-common/core/logger"
	"net"
	"os"
	"sync"
)

var (
	_lock    sync.Mutex
	_localIp = ""
	log      = logger.Logger()
)

func GetIntranetIp() string {

	if _localIp != "" {
		return _localIp
	}

	_lock.Lock()

	defer _lock.Unlock()

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				_localIp = ipNet.IP.String()
				break
			}
		}
	}

	if _localIp == "" {
		_localIp = "127.0.0.1"
	}
	return _localIp
}
