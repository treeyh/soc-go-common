package types

import (
	"fmt"
	"github.com/treeyh/soc-go-common/core/utils/copyer"
	"testing"
	"time"

	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/utils/json"
)

type Order struct {
	Order_id    string `json:"OrderId"`
	Create_time Time   `json:"CreateTime"`
}

type Order1 struct {
	Order_id    string `json:"OrderId"`
	Create_time Time   `json:"CreateTime"`
}

type Order3 struct {
	Order_id    string    `json:"OrderId"`
	Create_time time.Time `json:"CreateTime"`
}

func TestUnixTime_MarshalJSON(t *testing.T) {
	order := Order{Order_id: "10001",
		Create_time: Time(time.Now())}

	orderBytes, err := json.ToJson(order)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(string(orderBytes))
	}

	order1 := &Order{}

	json.FromJson(orderBytes, order1)
	fmt.Println(order1.Create_time)

	order2 := &Order1{}
	copyer.Copy(&order, order2)
	//
	fmt.Println(order2.Create_time)
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

	tttt := &Order{
		Order_id:    "123123",
		Create_time: Time(time.Now()),
	}

	fmt.Println(tttt.Create_time)
	fmt.Println(time.Time(tttt.Create_time))

	j := `{"OrderId":"123123","CreateTime":"2020-02-09 00:38:47"}`

	tt := &Order{}
	json.FromJson(j, tt)
	fmt.Println(tt.Create_time)
	fmt.Println(json.ToJsonIgnoreError(tt))

	ttt := &Order3{
		Order_id:    tt.Order_id,
		Create_time: time.Time(tt.Create_time),
	}
	fmt.Println(ttt.Create_time)
	fmt.Println(json.ToJsonIgnoreError(ttt))
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
