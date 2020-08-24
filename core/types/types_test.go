package types

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
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

func TestUnixTime_MarshalJSON(t *testing.T) {

	convey.Convey("Test TestCopy", t, func() {
		order := Order{
			OrderId:    "10001",
			CreateTime: Time(time.Now()),
			Int64:      Int64(64),
		}

		orderBytes, err := json.ToJson(order)
		convey.So(err, convey.ShouldBeNil)

		t.Log(orderBytes)

		order1 := &Order{}
		json.FromJson(orderBytes, order1)
		t.Log(order1.CreateTime)
		convey.So(order1.Int64, convey.ShouldEqual, 64)

		t.Log(order1.Int64)

		order2 := &Order1{}
		copyer.Copy(&order, order2)
		//
		t.Log(order2.CreateTime)
		convey.So(order2.Int64, convey.ShouldEqual, 64)
		t.Log(order2.Int64)
	})

}

func TestTime(t *testing.T) {

	now, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", "2019-06-06T11:11:11+08:00", time.Local)
	t.Log(err)
	t.Log(now)
	ti := Time(now)
	t.Log(ti)
	t.Log(time.Time(ti))

}

func TestTime2(t *testing.T) {

	convey.Convey("Test TestCopy", t, func() {
		tttt := &Order{
			OrderId:    "123123",
			CreateTime: Time(time.Now()),
			Int64:      Int64(64),
		}

		t.Log(tttt.CreateTime)
		t.Log(time.Time(tttt.CreateTime))
		t.Log(tttt.Int64)
		t.Log(int64(tttt.Int64))

		j := `{"OrderId":"123123","CreateTime":"2020-02-09 00:38:47","Int64":64}`

		tt := &Order{}
		json.FromJson(j, tt)
		t.Log(tt.CreateTime)
		t.Log(tt.Int64)

		convey.So(tt.Int64, convey.ShouldEqual, 64)

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
		t.Log(tt3.Int64)

	})

}

func TestTime0(t *testing.T) {
	fmt.Println(time.Now().Format(consts.AppTimeFormat))
	fmt.Println(",,,", Time0().String())
	fmt.Println(time.Now())
	var s time.Time
	//var o Time = time.Now()
	fmt.Println(copyer.Copy(Time(time.Now()), s))
	fmt.Println(s)
	fmt.Println(time.Now().UTC())
	a, _ := time.Parse(consts.AppTimeFormat, time.Now().Format(consts.AppTimeFormat))
	fmt.Println(a, a.Add(-time.Hour*8))
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
