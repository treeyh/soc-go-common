package types

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/utils/copyer"
	"testing"
	"time"

	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/utils/json"
)

type Order struct {
	OrderId    string `json:"orderId"`
	CreateTime Time   `json:"createTime"`
	Int64      Int64  `json:"int64"`
}

type Order1 struct {
	OrderId    string `json:"orderId"`
	CreateTime Time   `json:"createTime"`
	Int64      Int64  `json:"int64"`
}

type Order3 struct {
	OrderId    string    `json:"orderId"`
	CreateTime time.Time `json:"createTime"`
	Int64      int64     `json:"int64"`
}

func TestUnixTime(t *testing.T) {
	str := time.Now().Format(consts.AppSystemTimeFormat8)
	t.Log(str)
}

func TestUnixTime_MarshalJSON(t *testing.T) {

	order := Order{
		OrderId:    "10001",
		CreateTime: Time(time.Now()),
		Int64:      Int64(64),
	}

	orderBytes, err := json.ToJson(order)
	assert.NoError(t, err)

	t.Log(orderBytes)

	order1 := &Order{}
	json.FromJson(orderBytes, order1)
	t.Log(order1.CreateTime)
	assert.Equal(t, order1.Int64.ToInt64(), int64(64))

	t.Log(order1.Int64)

	order2 := &Order1{}
	copyer.Copy(context.Background(), &order, order2)
	//
	t.Log(order2.CreateTime)
	assert.Equal(t, order2.Int64.ToInt64(), int64(64))
	t.Log(order2.Int64)

}

func TestTime(t *testing.T) {

	now, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", "2019-06-06T11:11:11+08:00", time.Local)
	t.Log(err)
	t.Log(now)
	ti := Time(now)
	t.Log(ti)
	t.Log(time.Time(ti))

	time1, err := time.Parse(consts.AppSystemTimeFormat8, "2019-06-06T11:11:11+0700")
	t.Log(time1)

}

func TestTimeZone(t *testing.T) {
	now, _ := time.Parse(consts.AppSystemTimeFormat8, "2020-01-01T10:12:00+0500")
	//now3, _ := time.Parse(consts.AppSystemTimeFormat8, "2020-01-01T09:12:00+0400")
	t.Log(now)
	t.Log(now.UnixNano())
	t.Log(now.Zone())

	nLoc := time.FixedZone("+0545", 5.75*3600)
	nTime := now.In(nLoc)

	t.Log(nTime)
	t.Log(nTime.UnixNano())
	t.Log(nTime.Zone())

	timeOffset := int(5.75 * 3600)
	tzTime := now.In(globalTimeZoneMap[timeOffset])

	t.Log(tzTime)
	t.Log(tzTime.UnixNano())
	t.Log(tzTime.Zone())

	tzTime2 := InByOffset(now, timeOffset)

	t.Log(tzTime2)
	t.Log(tzTime2.UnixNano())
	t.Log(tzTime2.Zone())
}

func TestTime2(t *testing.T) {

	tttt := &Order{
		OrderId:    "123123",
		CreateTime: Time(time.Now()),
		Int64:      Int64(64),
	}

	t.Log(tttt.CreateTime)
	t.Log(time.Time(tttt.CreateTime))
	t.Log(tttt.Int64)
	t.Log(int64(tttt.Int64))

	j := `{"OrderId":"123123","CreateTime":"2020-02-09 00:38:47","Int64":"64"}`

	tt := &Order{}
	json.FromJson(j, tt)
	t.Log(tt.CreateTime)
	t.Log(tt.Int64)

	assert.Equal(t, tt.Int64.ToInt64(), int64(64))

	t.Log(json.ToJsonIgnoreError(tt))

	ttt := &Order3{
		OrderId:    tt.OrderId,
		CreateTime: time.Time(tt.CreateTime),
		Int64:      64,
	}
	str := json.ToJsonIgnoreError(ttt)
	t.Log(str)
	tt2 := &Order{}
	json.FromJson(str, tt2)
	t.Log(tt2.Int64)
	str2 := json.ToJsonIgnoreError(tt2)
	t.Log(str2)

	tt3 := &Order3{}
	json.FromJson(j, tt3)
	t.Log(tt3.CreateTime)
	t.Log(tt3.Int64)

}

func TestTime0(t *testing.T) {
	fmt.Println(time.Now().Format(consts.AppTimeFormat))
	fmt.Println(",,,", Time0().String())
	fmt.Println(time.Now())
	//var s time.Time
	////var o Time = time.Now()
	//fmt.Println(copyer.Copy(context.Background(), Time(time.Now()), s))
	//fmt.Println(s)
	//fmt.Println(time.Now().UTC())
	//a, _ := time.Parse(consts.AppTimeFormat, time.Now().Format(consts.AppTimeFormat))
	//fmt.Println(a, a.Add(-time.Hour*8))
}

type TestStruct struct {
	TimeField Time
}

func TestTime_IsNull(t *testing.T) {
	jsonStr := "{\"TimeField\":\"1970-01-01 00:00:01\"}"

	st := &TestStruct{}
	json.FromJson(jsonStr, st)

	t.Log(st.TimeField)
	t.Log(st.TimeField.IsNotNull())
}

func TestInt64(t *testing.T) {

	j := `{"OrderId":"123123","Int64": 64,"CreateTime":"2020-02-09 00:38:47"}`
	tt3 := &Order3{}
	json.FromJson(j, tt3)
	t.Log(tt3.CreateTime)
	t.Log(tt3.Int64)

	order11 := &Order1{}
	json.FromJson(j, order11)
	t.Log(order11.CreateTime)
	t.Log(order11.Int64)

	jj := `{"OrderId":"123123","CreateTime":"2020-02-09 00:38:47","Int64":"64"}`
	ttt3 := &Order3{}
	json.FromJson(jj, ttt3)
	t.Log(ttt3.CreateTime)
	t.Log(ttt3.Int64)

	order12 := &Order1{}
	json.FromJson(jj, order12)
	t.Log(order12.CreateTime)
	t.Log(order12.Int64)

}
