package times

import (
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/types"
	"strconv"
	"time"
)

const (

	// MinuteSecond 一分钟的秒数
	MinuteSecond = 60
	// HourSecond 一小时的秒数
	HourSecond = MinuteSecond * 60
	// DaySecond 一天的秒数
	DaySecond = HourSecond * 24
)

func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetMillisecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func GetNowSecond() int64 {
	return time.Now().Unix()
}

func GetNowNanoSecond() int64 {
	return time.Now().UnixNano()
}

func GetBeiJingTime() time.Time {
	timelocal, _ := time.LoadLocation("Asia/Chongqing")
	time.Local = timelocal
	return time.Now().Local()
}

func Sleep(second int64) {
	time.Sleep(time.Duration(second) * time.Second)
}

func SleepMillisecond(millSecond int64) {
	time.Sleep(time.Duration(millSecond) * time.Millisecond)
}

func GetDateTimeStrBySecond(s int64) string {
	return time.Unix(s, 0).Format(consts.AppTimeFormat)
}

func GetDateTimeStrByMillisecond(ms int64) string {
	second := ms / 1000
	return time.Unix(second, 0).Format(consts.AppTimeFormat)
}

func GetDateTimeMillisecondStrByMillisecond(ms int64) string {
	return time.Unix(0, ms*int64(time.Millisecond)).Format(consts.AppTimeFormatMillisecond)
}

func GetUnixTime(t types.Time) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")                                   //设置时区
	tt, _ := time.ParseInLocation(consts.AppTimeFormat, FormatTimeByTypes(t), loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	return tt.Unix()
}

func GetWeeHours() string {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return FormatTime(t)
}

func FormatDate(t time.Time) string {
	return t.Format(consts.AppDateFormat)
}
func FormatDateByTypes(t types.Time) string {
	return time.Time(t).Format(consts.AppDateFormat)
}

func FormatTime(t time.Time) string {
	return t.Format(consts.AppTimeFormat)
}
func FormatTimeByTypes(t types.Time) string {
	return time.Time(t).Format(consts.AppTimeFormat)
}

func ParseTime(str string) (types.Time, errors.AppError) {
	t, err := time.Parse(consts.AppTimeFormat, str)
	if err != nil {
		return types.Time0(), errors.NewAppErrorByExistError(errors.ParseTimeFail, err)
	}
	return types.Time(t), nil
}

func Parse(str string) (time.Time, errors.AppError) {
	t, err := time.Parse(consts.AppTimeFormat, str)
	if err != nil {
		return time.Unix(0, 0), errors.NewAppErrorByExistError(errors.ParseTimeFail, err)
	}
	return t, nil
}

func ParseByFormat(format, str string) (time.Time, errors.AppError) {
	t, err := time.Parse(format, str)
	if err != nil {
		return time.Unix(0, 0), errors.NewAppErrorByExistError(errors.ParseTimeFail, err)
	}
	return t, nil
}

// GetWeekStart 获取某一天当前周的周一0点
func GetWeekStart(d time.Time) time.Time {
	offset := int(time.Monday - d.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

// GetMonthStart 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetMonthStart(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetDateInt(t types.Time) int {
	i, _ := strconv.Atoi(time.Time(t).Format(consts.AppDateFormat2))
	return i
}

func GetDateTimeLong(t types.Time) int64 {
	i, _ := strconv.ParseInt(time.Time(t).Format(consts.AppTimeFormat2), 10, 64)
	return i
}
func GetDateIntByTime(t time.Time) int {
	i, _ := strconv.Atoi(t.Format(consts.AppDateFormat2))
	return i
}

func GetDateTimeLongByTime(t time.Time) int64 {
	i, _ := strconv.ParseInt(t.Format(consts.AppTimeFormat2), 10, 64)
	return i
}

// 获取当天时间段：2019-08-12 00:00:00 - 2019-08-12 23:59:59
func GetTodayTimeQuantum() []time.Time {
	timeStr := time.Now().Format(consts.AppDateFormat)
	b, _ := time.Parse(consts.AppDateFormat, timeStr)
	a := b.Add(time.Duration(1000*60*60*24-1) * time.Millisecond)
	return []time.Time{b, a}
}

// 2019-09-03
func GetYesterdayDate() string {
	timeStr := time.Now().Format(consts.AppDateFormat)
	b, _ := time.Parse(consts.AppDateFormat, timeStr)
	a := b.Add(time.Duration(-1) * time.Millisecond)
	return a.Format(consts.AppDateFormat)
}

// GetLastMonth1Date 获取时间上个月1号0时0分0秒时间
func GetLastMonth1Date(t time.Time) time.Time {
	timestr := t.AddDate(0, -1, 0).Format(consts.AppMonthFormat) + "01000000"

	monthTime, _ := time.Parse(consts.AppTimeFormat2, timestr)
	return monthTime
}

// GetLastMonth1Date 获取时间当月1号0时0分0秒时间
func GetNowMonth1Date(t time.Time) time.Time {
	timestr := t.Format(consts.AppMonthFormat) + "01000000"

	monthTime, _ := time.Parse(consts.AppTimeFormat2, timestr)
	return monthTime
}

// GetLastMonthLastDate 获取时间上个月的最后一天0时0分0秒时间
func GetLastMonthLastDate(t time.Time) time.Time {
	timestr := t.Format(consts.AppMonthFormat) + "01000000"

	monthTime, _ := time.Parse(consts.AppTimeFormat2, timestr)
	return monthTime.AddDate(0, 0, -1)
}

// getLastMonthDayCount 获取时间上个月的天数
func GetLastMonthDayCount(t time.Time) int64 {
	begin := GetLastMonth1Date(t)
	end := GetLastMonthLastDate(t)
	return ((end.Unix() - begin.Unix()) / DaySecond) + 1
}

// GetMonthLastDay 获取传入日期月的最后一天
func GetMonthLastDay(t time.Time) time.Time {
	day1 := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	return day1.AddDate(0, 1, -1)
}
