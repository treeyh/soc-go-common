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

func RandomString2(count int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func RandomString(count int, source []byte) string {
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result = append(result, source[r.Intn(len(source))])
	}
	return string(result)
}

func RandomFileName() string {
	return uuid.NewUuid() + strconv.FormatInt(times.GetNowMillisecond(), 10)
}
