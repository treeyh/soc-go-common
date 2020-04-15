package random

import (
	"github.com/treeyh/soc-go-common/core/utils/times"
	"github.com/treeyh/soc-go-common/core/utils/uuid"
	"math/rand"
	"strconv"
	"time"
)

func Token() string {
	return uuid.NewUuid()
}

func RandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func RandomFileName() string {
	return uuid.NewUuid() + strconv.FormatInt(times.GetNowMillisecond(), 10)
}
