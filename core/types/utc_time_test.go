package types

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/utils/copyer"
	"github.com/treeyh/soc-go-common/core/utils/json"
)

type OrderUtc struct {
	OrderId    string  `json:"orderId"`
	CreateTime UtcTime `json:"createTime"`
	Int64      Int64   `json:"int64"`
}

type OrderUtc1 struct {
	OrderId    string  `json:"orderId"`
	CreateTime UtcTime `json:"createTime"`
	Int64      Int64   `json:"int64"`
}

func TestUtcTime_MarshalJSON(t *testing.T) {

	order := OrderUtc{
		OrderId:    "10001",
		CreateTime: UtcTime(time.Now()),
		Int64:      Int64(64),
	}

	orderBytes, err := json.ToJson(order)
	assert.NoError(t, err)

	t.Log("orderBytes = " + orderBytes)

	order1 := &OrderUtc{}
	json.FromJson(orderBytes, order1)
	t.Log(order1.CreateTime)
	assert.Equal(t, order1.Int64.ToInt64(), int64(64))

	t.Log(order1.Int64)

	order2 := &OrderUtc1{}
	copyer.Copy(context.Background(), &order, order2)
	//
	t.Log(order2.CreateTime)
	assert.Equal(t, order2.Int64.ToInt64(), int64(64))
	t.Log(order2.Int64)

}

func TestUtcTime_MarshalJSONZone(t *testing.T) {

	order := Order3{
		OrderId:    "10001",
		CreateTime: time.Now(),
		Int64:      int64(1234567890),
	}

	orderBytes, err := json.ToJson(order)
	assert.NoError(t, err)

	t.Log("orderBytes = " + orderBytes)

	order.CreateTime = InByOffset(order.CreateTime, 3*3600)
	orderBytes2, err := json.ToJson(order)
	assert.NoError(t, err)
	t.Log("orderBytes2 = " + orderBytes2)

	order1 := &OrderUtc{}
	_ = json.FromJson(orderBytes, order1)
	t.Log("== order1 ==")
	t.Log(order1.CreateTime)
	t.Log(order1.CreateTime.ToTime().Zone())

	order2 := &OrderUtc{}
	json.FromJson(orderBytes2, order2)
	t.Log("== order2 ==")
	t.Log(order2.CreateTime)
	t.Log(order2.CreateTime.ToTime().Zone())
	orderBytes3 := json.ToJsonIgnoreError(order2)
	t.Log("orderBytes3 = " + orderBytes3)

	order3 := &OrderUtc{}
	json.FromJson(orderBytes, order3)
	t.Log("== order3 ==")
	t.Log(order3.CreateTime)

	order4 := &OrderUtc{}
	json.FromJson(orderBytes2, order4)
	t.Log("== order4 ==")
	t.Log(order4.CreateTime)
	typeof := reflect.TypeOf(order4.CreateTime)
	t.Log(typeof)
	for i := 0; i < typeof.NumMethod(); i++ {
		fmt.Printf("method is %s, type is %s, kind is %s.\n", typeof.Method(i).Name, typeof.Method(i).Type, typeof.Method(i).Type.Kind())
	}

	order5 := &Order3{}
	json.FromJson(orderBytes, order5)
	t.Log("== order5 ==")
	t.Log(order5.CreateTime)
	t.Log(order5.CreateTime.Zone())

	order6 := &Order3{}
	json.FromJson(orderBytes2, order6)
	t.Log("== order6 ==")
	t.Log(order6.CreateTime)
	t.Log(order6.CreateTime.Zone())
	t.Log("order6 = " + json.ToJsonIgnoreError(order6))

	order7 := &Order3{}
	json.FromJson(orderBytes3, order7)
	t.Log("== order7 ==")
	t.Log(order7.CreateTime)
	t.Log(order7.CreateTime.Zone())
	t.Log("order7 = " + json.ToJsonIgnoreError(order7))
}

func TestUtcTime(t *testing.T) {

	now, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", "2019-06-06T11:11:11+08:00", time.Local)
	t.Log(err)
	t.Log(now)
	ti := UtcTime(now)
	t.Log(ti)
	t.Log(time.Time(ti))

	time1, err := time.Parse(consts.AppSystemTimeFormat8, "2019-06-06T11:11:11+0700")
	t.Log(time1)

}

// func TestTimeZone(t *testing.T) {
// 	now, _ := time.Parse(consts.AppSystemTimeFormat8, "2020-01-01T10:12:00+0500")
// 	// now3, _ := time.Parse(consts.AppSystemTimeFormat8, "2020-01-01T09:12:00+0400")
// 	t.Log(now)
// 	t.Log(now.UnixNano())
// 	t.Log(now.Zone())
//
// 	nLoc := time.FixedZone("+0545", 5.75*3600)
// 	nTime := now.In(nLoc)
//
// 	t.Log(nTime)
// 	t.Log(nTime.UnixNano())
// 	t.Log(nTime.Zone())
//
// 	timeOffset := int(5.75 * 3600)
// 	tzTime := now.In(globalTimeZoneMap[timeOffset])
//
// 	t.Log(tzTime)
// 	t.Log(tzTime.UnixNano())
// 	t.Log(tzTime.Zone())
//
// 	tzTime2 := InByOffset(now, timeOffset)
//
// 	t.Log(tzTime2)
// 	t.Log(tzTime2.UnixNano())
// 	t.Log(tzTime2.Zone())
// }
